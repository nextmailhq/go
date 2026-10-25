package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func webhooksListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}}
}

func webhooksWindowQuery(o *callOptions) values {
	return values{{query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func webhooksDeliveryQuery(o *callOptions) values {
	return webhooksWindowQuery(o).merge(values{{query.Status, o.value("status")}})
}

func webhooksDeadLetterQuery(o *callOptions) values {
	return webhooksWindowQuery(o).merge(values{{query.Status, o.value("status")}})
}

func webhooksWorkspaceDeliveryQuery(o *callOptions) values {
	return webhooksDeliveryQuery(o).merge(values{{query.EndpointIDs, o.value("endpointIds")}})
}

func webhooksWorkspaceActivityQuery(o *callOptions) values {
	return webhooksWindowQuery(o).merge(values{{query.EndpointIDs, o.value("endpointIds")}})
}

func (s *WebhooksService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.list", opts)

	return s.t.page(ctx, o, call{path: paths.Webhooks, query: webhooksListQuery(o)}, cursorPages)
}

func (s *WebhooksService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Webhooks, query: webhooksListQuery(o)}, cursorPages)
}

func (s *WebhooksService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Webhooks, query: webhooksListQuery(o)}, cursorPages)
}

func (s *WebhooksService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.get", opts)

	return s.t.object(ctx, o, call{path: paths.Webhook, params: []any{"id", id}})
}

func (s *WebhooksService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Webhooks,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *WebhooksService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Webhook,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *WebhooksService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Webhook,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *WebhooksService) RotateSecret(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.rotateSecret", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WebhookRotateSecret,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *WebhooksService) Test(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.test", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WebhookTest,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *WebhooksService) ListDeliveries(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.listDeliveries", opts)

	return s.t.page(ctx, o, call{
		path:   paths.WebhookDeliveries,
		params: []any{"id", id},
		query:  webhooksDeliveryQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) ListAllDeliveries(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAllDeliveries", opts)

	return s.t.all(ctx, o, call{
		path:   paths.WebhookDeliveries,
		params: []any{"id", id},
		query:  webhooksDeliveryQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) IterateDeliveries(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterateDeliveries", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.WebhookDeliveries,
		params: []any{"id", id},
		query:  webhooksDeliveryQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) GetDelivery(ctx context.Context, id string, deliveryID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.getDelivery", opts)

	return s.t.object(ctx, o, call{path: paths.WebhookDelivery, params: []any{"id", id, "deliveryId", deliveryID}})
}

func (s *WebhooksService) ReplayDelivery(ctx context.Context, id string, deliveryID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.replayDelivery", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WebhookDeliveryReplay,
		params: []any{"id", id, "deliveryId", deliveryID},
		method: http.MethodPost,
	})
}

func (s *WebhooksService) ListDeadLetters(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.listDeadLetters", opts)

	return s.t.page(ctx, o, call{
		path:   paths.WebhookDeadLetters,
		params: []any{"id", id},
		query:  webhooksDeadLetterQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) ListAllDeadLetters(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAllDeadLetters", opts)

	return s.t.all(ctx, o, call{
		path:   paths.WebhookDeadLetters,
		params: []any{"id", id},
		query:  webhooksDeadLetterQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) IterateDeadLetters(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterateDeadLetters", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.WebhookDeadLetters,
		params: []any{"id", id},
		query:  webhooksDeadLetterQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) ReplayDeadLetters(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.replayDeadLetters", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WebhookDeadLettersReplay,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *WebhooksService) DiscardDeadLetter(ctx context.Context, id string, deadLetterID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.discardDeadLetter", opts)

	return s.t.object(ctx, o, call{
		path:       paths.WebhookDeadLetter,
		params:     []any{"id", id, "deadLetterId", deadLetterID},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *WebhooksService) ListWorkspaceDeliveries(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.listWorkspaceDeliveries", opts)

	return s.t.page(ctx, o, call{path: paths.WebhooksDeliveries, query: webhooksWorkspaceDeliveryQuery(o)}, cursorPages)
}

func (s *WebhooksService) ListAllWorkspaceDeliveries(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAllWorkspaceDeliveries", opts)

	return s.t.all(ctx, o, call{path: paths.WebhooksDeliveries, query: webhooksWorkspaceDeliveryQuery(o)}, cursorPages)
}

func (s *WebhooksService) IterateWorkspaceDeliveries(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterateWorkspaceDeliveries", opts)

	return s.t.iterate(ctx, o, call{path: paths.WebhooksDeliveries, query: webhooksWorkspaceDeliveryQuery(o)}, cursorPages)
}

func (s *WebhooksService) ListActivity(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.listActivity", opts)

	return s.t.page(ctx, o, call{
		path:   paths.WebhookActivity,
		params: []any{"id", id},
		query:  webhooksWindowQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) ListAllActivity(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAllActivity", opts)

	return s.t.all(ctx, o, call{
		path:   paths.WebhookActivity,
		params: []any{"id", id},
		query:  webhooksWindowQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) IterateActivity(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterateActivity", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.WebhookActivity,
		params: []any{"id", id},
		query:  webhooksWindowQuery(o),
	}, cursorPages)
}

func (s *WebhooksService) ListWorkspaceActivity(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("webhooks.listWorkspaceActivity", opts)

	return s.t.page(ctx, o, call{path: paths.WebhooksActivity, query: webhooksWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *WebhooksService) ListAllWorkspaceActivity(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("webhooks.listAllWorkspaceActivity", opts)

	return s.t.all(ctx, o, call{path: paths.WebhooksActivity, query: webhooksWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *WebhooksService) IterateWorkspaceActivity(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("webhooks.iterateWorkspaceActivity", opts)

	return s.t.iterate(ctx, o, call{path: paths.WebhooksActivity, query: webhooksWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *WebhooksService) ListEvents(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.listEvents", opts)

	return s.t.object(ctx, o, call{path: paths.WebhooksEvents})
}

func (s *WebhooksService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("webhooks.stats", opts)

	return s.t.object(ctx, o, call{
		path:  paths.WebhooksStats,
		query: webhooksWindowQuery(o).merge(values{{query.EndpointIDs, o.value("endpointIds")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}}),
	})
}
