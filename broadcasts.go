package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func broadcastsListQuery(o *callOptions) values {
	return values{{query.AudienceID, o.value("audienceId")}}
}

func broadcastsRecipientQuery(o *callOptions) values {
	return values{{query.Filter, o.value("filter")}, {query.Q, o.value("q")}}
}

func broadcastsAnalyticsQuery(o *callOptions) values {
	return values{{query.BroadcastIDs, o.value("broadcastIds")}, {query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Since, o.value("since")}}
}

func (s *BroadcastsService) Preview(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.preview", opts)

	return s.t.object(ctx, o, call{
		path:       paths.BroadcastsPreview,
		method:     http.MethodPost,
		body:       pick(body, "audienceIds", "excludeAudienceIds", "segmentId"),
		repeatable: true,
	})
}

func (s *BroadcastsService) Send(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.send", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Broadcasts,
		method:     http.MethodPost,
		body:       bodyOf(body),
		idempotent: true,
		repeatable: true,
	})
}

func (s *BroadcastsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("broadcasts.list", opts)

	return s.t.page(ctx, o, call{path: paths.Broadcasts, query: broadcastsListQuery(o)}, cursorPages)
}

func (s *BroadcastsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("broadcasts.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Broadcasts, query: broadcastsListQuery(o)}, cursorPages)
}

func (s *BroadcastsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("broadcasts.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Broadcasts, query: broadcastsListQuery(o)}, cursorPages)
}

func (s *BroadcastsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.get", opts)

	return s.t.object(ctx, o, call{path: paths.Broadcast, params: []any{"id", id}})
}

func (s *BroadcastsService) Stats(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.stats", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastStats,
		params: []any{"id", id},
		query:  values{{query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Since, o.value("since")}},
	})
}

func (s *BroadcastsService) Analytics(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.analytics", opts)

	return s.t.object(ctx, o, call{path: paths.BroadcastsAnalytics, query: broadcastsAnalyticsQuery(o)})
}

func (s *BroadcastsService) ListRecipients(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("broadcasts.listRecipients", opts)

	return s.t.page(ctx, o, call{path: paths.BroadcastRecipients, params: []any{"id", id}, query: broadcastsRecipientQuery(o)}, cursorPages)
}

func (s *BroadcastsService) ListAllRecipients(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("broadcasts.listAllRecipients", opts)

	return s.t.all(ctx, o, call{path: paths.BroadcastRecipients, params: []any{"id", id}, query: broadcastsRecipientQuery(o)}, cursorPages)
}

func (s *BroadcastsService) IterateRecipients(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("broadcasts.iterateRecipients", opts)

	return s.t.iterate(ctx, o, call{path: paths.BroadcastRecipients, params: []any{"id", id}, query: broadcastsRecipientQuery(o)}, cursorPages)
}

func (s *BroadcastsService) GetRecipient(ctx context.Context, id string, emailID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.getRecipient", opts)

	return s.t.object(ctx, o, call{path: paths.BroadcastRecipient, params: []any{"id", id, "emailId", emailID}})
}

func (s *BroadcastsService) Cancel(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.cancel", opts)

	return s.t.object(ctx, o, call{
		path:       paths.BroadcastCancel,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *BroadcastsService) CreateDraft(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.createDraft", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastDrafts,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BroadcastsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Broadcast,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *BroadcastsService) SendDraft(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.sendDraft", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastSend,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BroadcastsService) Duplicate(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.duplicate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastDuplicate,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *BroadcastsService) Resend(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.resend", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastResend,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BroadcastsService) SendTest(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("broadcasts.sendTest", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BroadcastsTest,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
