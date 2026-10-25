package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *TempMailService) ListDomains(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tempMail.listDomains", opts)

	return s.t.objects(ctx, o, call{path: paths.TempMailDomains, anonymous: true})
}

func (s *TempMailService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.create", opts)

	return s.t.object(ctx, o, call{
		path:      paths.TempMailInboxes,
		method:    http.MethodPost,
		body:      bodyOf(body),
		anonymous: true,
	})
}

func (s *TempMailService) Get(ctx context.Context, inboxID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.get", opts)

	return s.t.object(ctx, o, call{path: paths.TempMailInbox, params: []any{"id", inboxID}})
}

func (s *TempMailService) Extend(ctx context.Context, inboxID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.extend", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TempMailInboxExtend,
		params: []any{"id", inboxID},
		method: http.MethodPost,
	})
}

func (s *TempMailService) Delete(ctx context.Context, inboxID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TempMailInbox,
		params: []any{"id", inboxID},
		method: http.MethodDelete,
	})
}

func (s *TempMailService) ListMessages(ctx context.Context, inboxID string, opts ...RequestOption) (*TempMessagesPage, error) {
	o := newCallOptions("tempMail.listMessages", opts)

	return s.messagePage(ctx, o, inboxID, o.text(optionCursor))
}

func (s *TempMailService) messagePage(ctx context.Context, o *callOptions, inboxID, cursor string) (*TempMessagesPage, error) {
	envelope, err := s.t.object(ctx, o, call{
		path:   paths.TempMailMessages,
		params: []any{"id", inboxID},
		query:  values{{query.Limit, o.value("limit")}, {query.Cursor, cursor}},
	})

	if err != nil {
		return nil, err
	}

	next := cursorOf(envelope[fieldNextCursor])

	return &TempMessagesPage{
		Items:      objectsOf(envelope[fieldData]),
		HasMore:    hasMoreOf(envelope, next),
		NextCursor: next,
		ExpiresAt:  envelope.String("expiresAt"),
	}, nil
}

func (s *TempMailService) messages(ctx context.Context, o *callOptions, inboxID string) *Iterator {
	return newIterator(ctx, o.text(optionCursor), func(ctx context.Context, cursor string) ([]Object, bool, string, error) {
		page, err := s.messagePage(ctx, o, inboxID, cursor)

		if err != nil {
			return nil, false, "", err
		}

		return page.Items, page.HasMore, page.NextCursor, nil
	})
}

func (s *TempMailService) ListAllMessages(ctx context.Context, inboxID string, opts ...RequestOption) ([]Object, error) {
	return s.messages(ctx, newCallOptions("tempMail.listAllMessages", opts), inboxID).collect()
}

func (s *TempMailService) IterateMessages(ctx context.Context, inboxID string, opts ...RequestOption) *Iterator {
	return s.messages(ctx, newCallOptions("tempMail.iterateMessages", opts), inboxID)
}

func (s *TempMailService) GetMessage(ctx context.Context, inboxID string, messageID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.getMessage", opts)

	return s.t.object(ctx, o, call{path: paths.TempMailMessage, params: []any{"id", inboxID, "messageId", messageID}})
}

func (s *TempMailService) DeleteMessage(ctx context.Context, inboxID string, messageID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tempMail.deleteMessage", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TempMailMessage,
		params: []any{"id", inboxID, "messageId", messageID},
		method: http.MethodDelete,
	})
}

func (s *TempMailService) ListAttachments(ctx context.Context, inboxID string, messageID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tempMail.listAttachments", opts)

	return s.t.objects(ctx, o, call{
		path:   paths.TempMailMessageAttachments,
		params: []any{"id", inboxID, "messageId", messageID},
	})
}
