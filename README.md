<div align='center'>
   <a href='https://nextmail.sh'>
        <img
            src='https://nextmail.sh/logo.svg'
            alt='Nextmail Logo'
            width='180'
        />
   </a>

   <br />
</div>

<p align='center'>
    Email you can build on. Send mail, read the mailbox and automate a workspace from Go.
</p>

<p align='center'>
    <a href='https://nextmail.sh'><b>Website</b></a>
    •
    <a href='https://nextmail.sh/docs/go'><b>Documentation</b></a>
    •
    <a href='https://nextmail.sh/docs/go/reference/methods'><b>Every method</b></a>
    •
    <a href='https://nextmail.sh/docs/api/reference'><b>API Reference</b></a>
</p>

<br />

## The Go SDK

The official Go client for the Nextmail API. It has one method for every method of the TypeScript SDK, under the same name in Go's casing, and a parity check proves that each one sends the same request. It needs Go 1.24 or later and the standard library, and nothing else.

It carries a workspace API key or an OAuth access token, so it belongs on a server, in a job or in a tool that runs on your own machine.

### Installing
```bash
go get nextmail.sh
```

### Using
```go
package main

import (
	"context"
	"fmt"
	"log"

	"nextmail.sh"
)

func main() {
	client := nextmail.New()

	sent, err := client.Emails.Send(context.Background(), nextmail.Body{
		"from":    "Acme Billing <billing@acme.com>",
		"to":      "ada@example.com",
		"subject": "Your September invoice",
		"html":    "<p>Your invoice is attached.</p>",
		"attachments": []nextmail.Body{{
			"filename":    "invoice.pdf",
			"content":     []byte("%PDF-1.4"),
			"contentType": "application/pdf",
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sent.ID(), sent.String("status"))
}
```

`nextmail.New()` reads the key from `NEXTMAIL_API_KEY`. Create one in Nextmail under Settings, API keys. It is shown once, and it belongs in an environment variable or your secrets store rather than in code. Pass `nextmail.WithAPIKey` when your configuration lives somewhere else. `New` never fails: a missing or malformed credential is reported by `client.Err()` and by the first call.

A request body is `nextmail.Body`, a `map[string]any` whose keys keep the API's names, so `replyTo`, `scheduledAt` and `contentType` stay camelCase. Options are functions: `nextmail.WithLimit(25)`, `nextmail.WithCursor(next)`, `nextmail.WithIdempotencyKey(key)`. What comes back is `nextmail.Object`, the decoded JSON with typed readers such as `String`, `Int`, `Bool`, `Time` and `Object`. Every method takes a `context.Context` first, which cancels the call or gives it a deadline.

Every send carries an idempotency key, generated once per call and reused by its retries, so a retried request replays the original message rather than sending a second one. Pass your own with `nextmail.WithIdempotencyKey` to make that hold across processes and restarts.

### Reading the mailbox
```go
page, err := client.Threads.List(ctx, nextmail.WithFolder("inbox"), nextmail.WithLimit(25))
if err != nil {
	return err
}

for _, row := range page.Items {
	thread, err := client.Threads.Get(ctx, row.ID())
	if err != nil {
		return err
	}

	fmt.Println(thread.ID(), thread.Int("messageCount"), thread.Bool("hasUnread"))
}

for thread, err := range client.Threads.Iterate(ctx, nextmail.WithFolder("inbox"), nextmail.WithQuery("invoice")).All() {
	if err != nil {
		return err
	}

	fmt.Println(thread.ID())
}
```

A row of a thread list carries the thread's id, and `Threads.Get` returns the thread with its messages.

Every paginated resource has `List` for one page (`*nextmail.Page` with `Items`, `HasMore` and `NextCursor`), `ListAll` for every page at once as a slice, and `Iterate`, which returns `*nextmail.Iterator` that fetches the next page only when it gets there, so a `break` stops early. Range over `All()`, or call `Next`, `Item` and `Err` in a plain loop. `Addresses.ListAll` is the exception and returns the whole address book.

### Errors
```go
_, err := client.Templates.Send(ctx, "order-shipped", nextmail.Body{
	"from":  "dispatch@acme.com",
	"to":    "ada@example.com",
	"props": nextmail.Body{"orderId": "AC-4192"},
})

var failure *nextmail.Error

if errors.As(err, &failure) {
	fmt.Println(failure.Status, failure.Code, failure.Param, failure.RequestID)
}

if errors.Is(err, nextmail.ErrValidation) {
	fmt.Println("a field was refused")
}
```

Every failure is `*nextmail.Error`. An API refusal carries `Status`, `Type`, `Code`, `Param`, `DocURL`, `RequestID`, `RetryAfter`, `Fields` and `Body`, and matches a sentinel for its kind with `errors.Is`: `ErrAuthentication`, `ErrPermission`, `ErrScopeMissing`, `ErrStepUpRequired`, `ErrNotFound`, `ErrConflict`, `ErrValidation`, `ErrRateLimited`, `ErrInvalidRequest`, `ErrServer` and `ErrRetryable`. No response at all matches `ErrNetwork`, and `ErrTimeout` when the deadline passed. A mistake in the call itself, such as a malformed key or an empty id, matches `ErrInvalidArgument` and is returned before anything is sent. A cancelled context still matches `context.Canceled`.

### Webhooks
```go
http.HandleFunc("/webhooks/nextmail", func(writer http.ResponseWriter, request *http.Request) {
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	event, err := nextmail.VerifyWebhookSignature(payload, request.Header, os.Getenv("NEXTMAIL_WEBHOOK_SECRET"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	fmt.Println(event.String("type"), event.ID())
	writer.WriteHeader(http.StatusNoContent)
})
```

It checks the HMAC in constant time and rejects a delivery more than five minutes old, then returns the decoded event. Pass the raw body: a parsed or re-encoded body no longer matches its signature. `nextmail.WithTolerance` changes the five minutes.

### Disposable inboxes
```go
tempMail := nextmail.NewTempMail()

inbox, err := tempMail.Create(ctx, nextmail.Body{"ttlMinutes": 60})
if err != nil {
	return err
}

messages, err := tempMail.ListMessages(ctx, inbox.ID(), nextmail.WithInboxToken(inbox.String("token")))
if err != nil {
	return err
}

fmt.Println(len(messages.Items), messages.ExpiresAt)
```

`Create` needs no credential and is the only call that returns the inbox token, so keep it.

### OAuth access tokens
An app a person connected to Nextmail with OAuth, such as a command line tool or an agent, holds an access token rather than an API key. Pass it with `nextmail.WithAccessToken`, or pass a function that returns it:

```go
current := "token-from-your-oauth-flow"

client := nextmail.New(nextmail.WithAccessTokenFunc(func(ctx context.Context) (string, error) {
	return current, nil
}))

fmt.Println(client.Mode())
```

The function runs before every request, so renew the token there when it is close to expiring and the client never has to be rebuilt. Pass a key or a token, not both. A client reads `NEXTMAIL_ACCESS_TOKEN` when you pass neither and `NEXTMAIL_API_KEY` is not set.

A token acts for a person, so before a sensitive change, such as deleting a domain or changing a webhook, it is asked for the same verification code the web app asks for. The call fails with an error that matches `nextmail.ErrStepUpRequired`. Ask for a code, check it, then make the call again:

```go
domainID := "b3e1f0a4-6c2d-4e8a-9f17-2d5c8a0b4e6f"

_, err := client.Domains.Delete(ctx, domainID)

if errors.Is(err, nextmail.ErrStepUpRequired) {
	challenge, err := client.Security.BeginStepUp(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Enter the code sent by", challenge.String("method"))

	if _, err := client.Security.VerifyStepUp(ctx, nextmail.Body{"code": "123456"}); err != nil {
		return err
	}

	_, err = client.Domains.Delete(ctx, domainID)

	return err
}

return err
```

API keys are never asked for a code.

### Configuring
```go
client := nextmail.New(
	nextmail.WithAPIKey(os.Getenv("NEXTMAIL_API_KEY")),
	nextmail.WithBaseURL("https://api.nextmail.sh"),
	nextmail.WithTimeout(30*time.Second),
	nextmail.WithMaxRetries(2),
	nextmail.WithHeader("X-Team", "billing"),
	nextmail.WithHTTPClient(http.DefaultClient),
)

if err := client.Err(); err != nil {
	return err
}
```

Anything you leave out is read from the environment: `NEXTMAIL_API_KEY`, then `NEXTMAIL_ACCESS_TOKEN`, and `NEXTMAIL_BASE_URL`, where a bare host such as `localhost:2222` gets its scheme added.

Use an `https:` origin: the client refuses to send an API key, an access token or an inbox token over plain `http:` unless the server is on this machine, at `localhost`, a `127.x.x.x` address or `::1`. A base URL on `0.0.0.0` is refused when the client is built, since that is the address a server listens on: use `127.0.0.1` with the same port. `WithTimeout` bounds a whole attempt, zero turns it off, uploads wait at least ten minutes, and the chat calls that wait for the assistant to answer wait at least five. The context you pass bounds the whole call, retries included. Reads, sends and every write that is safe to repeat are retried on 408, 500, 502, 503 and 504 with backoff, and on a 429 only when it carries a `Retry-After` of a minute or less. Any other write is never retried. Every method takes `nextmail.WithAPIKey` to act with another key for one call, so one process can serve several workspaces with one client, and the `TempMail` methods take `nextmail.WithInboxToken` instead.

A client is safe for concurrent use. `WithHTTPClient` takes an `*http.Client` or anything with the same `Do` method, which is how a test drives the client without a network. The default client does not follow redirects, so a credential is never carried to another host.

An endpoint no method wraps yet is one `client.Request` away, with the client's credential, base URL, timeout and retry policy applied:

```go
label, err := client.Request(ctx, http.MethodPost, "/labels", nextmail.Body{"name": "Invoices"})
if err != nil {
	return err
}

fmt.Println(label)
```

The path must begin with a single `/`, and a path whose finished URL leaves the base URL's origin is refused before a request is sent, so the credential never reaches another host.

Printing a client, an option or an error never shows a credential: the key, the token and custom header values are kept out of `fmt`, `log/slog` and `encoding/json` output.

### Constants
The sets of values the API uses, such as webhook events, scopes and statuses, are in the `constants` package:

```go
webhook, err := client.Webhooks.Create(ctx, nextmail.Body{
	"url":    "https://example.com/webhooks/nextmail",
	"events": []string{constants.WebhookEventsEmailReceived, constants.WebhookEventsEmailBounced},
})
if err != nil {
	return err
}

fmt.Println(webhook.ID(), len(constants.WebhookEventsValues()))
```

Everything else, from templates, rules, tracking, calendar, contacts, audiences and broadcasts to roles, members, settings, mailbox imports and provider imports, plus the full reference for every method, lives in the [documentation](https://nextmail.sh/docs/go). What changed in each release is in the [changelog](https://nextmail.sh/docs/go/changelog).

### Licence
The package is proprietary. See the `LICENSE` file.
