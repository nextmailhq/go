package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func subscriptionsListQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}, {query.Q, o.value("q")}, {query.Sort, o.value("sort")}, {query.Suggested, o.value("suggested")}, {query.Limit, o.value("limit")}, {query.Offset, o.value("offset")}, {query.Address, o.value("address")}}
}

func subscriptionsBatchBody(ids []string, body Body) Body {
	batch := Body{"ids": ids}

	for key, value := range body {
		batch[key] = value
	}

	return batch
}

func (s *SubscriptionsService) List(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.list", opts)

	return s.t.object(ctx, o, call{path: paths.Subscriptions, query: subscriptionsListQuery(o)})
}

func (s *SubscriptionsService) ListDomains(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.listDomains", opts)

	return s.t.object(ctx, o, call{path: paths.SubscriptionDomains, query: subscriptionsListQuery(o)})
}

func (s *SubscriptionsService) Unsubscribe(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.unsubscribe", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SubscriptionUnsubscribe,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SubscriptionsService) UnsubscribeDomain(ctx context.Context, domain string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.unsubscribeDomain", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SubscriptionDomainUnsubscribe,
		params: []any{"domain", domain},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SubscriptionsService) Move(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.move", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SubscriptionMove,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *SubscriptionsService) Keep(ctx context.Context, ids []string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.keep", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SubscriptionsBatchKeep,
		method:     http.MethodPost,
		body:       subscriptionsBatchBody(ids, body),
		repeatable: true,
	})
}

func (s *SubscriptionsService) UnsubscribeMany(ctx context.Context, ids []string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.unsubscribeMany", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SubscriptionsBatchUnsubscribe,
		method: http.MethodPost,
		body:   subscriptionsBatchBody(ids, body),
	})
}

func (s *SubscriptionsService) MoveMany(ctx context.Context, ids []string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("subscriptions.moveMany", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SubscriptionsBatchMove,
		method:     http.MethodPost,
		body:       subscriptionsBatchBody(ids, body),
		repeatable: true,
	})
}
