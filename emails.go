package nextmail

import (
	"context"
	"net/http"
	"time"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func emailsListQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}, {query.From, o.value("from")}, {query.BroadcastID, o.value("broadcastId")}, {query.ScheduledFrom, o.value("scheduledFrom")}, {query.ScheduledTo, o.value("scheduledTo")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}, {query.Recipient, o.value("recipient")}, {query.Tag, o.value("tag")}}
}

func emailsAddressQuery(o *callOptions) values {
	return values{{query.Address, o.value("address")}}
}

func emailsEventsQuery(o *callOptions) values {
	return append(emailsAddressQuery(o), pair{query.MessageID, o.value("messageId")})
}

func emailsMessageIDEventsQuery(o *callOptions, messageID string) values {
	return append(emailsAddressQuery(o), pair{query.MessageID, messageID})
}

func (s *EmailsService) Send(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.send", opts)
	wire, err := emailToWire(body)

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:       paths.Emails,
		method:     http.MethodPost,
		body:       wire,
		idempotent: true,
		repeatable: true,
	})
}

func (s *EmailsService) SendRaw(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.sendRaw", opts)
	wire, err := emailRawToWire(body)

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:       paths.EmailsRaw,
		method:     http.MethodPost,
		body:       wire,
		idempotent: true,
		repeatable: true,
	})
}

func (s *EmailsService) SendBatch(ctx context.Context, emails []Body, opts ...RequestOption) (*BatchResult, error) {
	o := newCallOptions("emails.sendBatch", opts)
	wire := make([]Body, 0, len(emails))

	for _, email := range emails {
		prepared, err := emailToWire(email)

		if err != nil {
			return nil, err
		}

		wire = append(wire, prepared)
	}

	result, err := s.t.object(ctx, o, call{
		path:       paths.EmailsBatch,
		method:     http.MethodPost,
		body:       Body{"emails": wire},
		idempotent: true,
		repeatable: true,
	})

	if err != nil {
		return nil, err
	}

	return &BatchResult{Items: objectsOf(result[fieldData]), Sent: result.Int("sent"), Failed: result.Int("failed")}, nil
}

func (s *EmailsService) Translate(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.translate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EmailsTranslate,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *EmailsService) Check(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.check", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EmailsCheck,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *EmailsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("emails.list", opts)

	return s.t.page(ctx, o, call{path: paths.Emails, query: emailsListQuery(o)}, cursorPages)
}

func (s *EmailsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("emails.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Emails, query: emailsListQuery(o)}, cursorPages)
}

func (s *EmailsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("emails.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Emails, query: emailsListQuery(o)}, cursorPages)
}

func (s *EmailsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.get", opts)

	return s.t.object(ctx, o, call{path: paths.Email, params: []any{"id", id}})
}

func (s *EmailsService) ListEvents(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("emails.listEvents", opts)

	return s.t.page(ctx, o, call{path: paths.EmailEvents, params: []any{"id", id}, query: emailsEventsQuery(o)}, cursorPages)
}

func (s *EmailsService) ListAllEvents(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("emails.listAllEvents", opts)

	return s.t.all(ctx, o, call{path: paths.EmailEvents, params: []any{"id", id}, query: emailsEventsQuery(o)}, cursorPages)
}

func (s *EmailsService) IterateEvents(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("emails.iterateEvents", opts)

	return s.t.iterate(ctx, o, call{path: paths.EmailEvents, params: []any{"id", id}, query: emailsEventsQuery(o)}, cursorPages)
}

func (s *EmailsService) ListEventsByMessageID(ctx context.Context, messageID string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("emails.listEventsByMessageId", opts)

	return s.t.page(ctx, o, call{path: paths.EmailsEvents, query: emailsMessageIDEventsQuery(o, messageID)}, cursorPages)
}

func (s *EmailsService) ListAllEventsByMessageID(ctx context.Context, messageID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("emails.listAllEventsByMessageId", opts)

	return s.t.all(ctx, o, call{path: paths.EmailsEvents, query: emailsMessageIDEventsQuery(o, messageID)}, cursorPages)
}

func (s *EmailsService) IterateEventsByMessageID(ctx context.Context, messageID string, opts ...RequestOption) *Iterator {
	o := newCallOptions("emails.iterateEventsByMessageId", opts)

	return s.t.iterate(ctx, o, call{path: paths.EmailsEvents, query: emailsMessageIDEventsQuery(o, messageID)}, cursorPages)
}

func (s *EmailsService) GetTracking(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.getTracking", opts)

	return s.t.object(ctx, o, call{path: paths.EmailTracking, params: []any{"id", id}, query: emailsAddressQuery(o)})
}

func (s *EmailsService) Cancel(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.cancel", opts)

	return s.t.object(ctx, o, call{
		path:       paths.EmailCancel,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *EmailsService) Reschedule(ctx context.Context, id string, scheduledAt time.Time, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.reschedule", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Email,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       Body{"scheduledAt": instant(scheduledAt)},
		repeatable: true,
	})
}

func (s *EmailsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Email,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       recipientsToWire(patch),
		repeatable: true,
	})
}

func (s *EmailsService) Compose(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.compose", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EmailsCompose,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *EmailsService) Rewrite(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.rewrite", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EmailsRewrite,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *EmailsService) SuggestSubject(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("emails.suggestSubject", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EmailsSubject,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
