package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func suppressionsListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Reason, o.value("reason")}}
}

func (s *SuppressionsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("suppressions.list", opts)

	return s.t.page(ctx, o, call{path: paths.Suppressions, query: suppressionsListQuery(o)}, cursorPages)
}

func (s *SuppressionsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("suppressions.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Suppressions, query: suppressionsListQuery(o)}, cursorPages)
}

func (s *SuppressionsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("suppressions.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Suppressions, query: suppressionsListQuery(o)}, cursorPages)
}

func (s *SuppressionsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("suppressions.get", opts)

	return s.t.object(ctx, o, call{path: paths.Suppression, params: []any{"id", id}})
}

func (s *SuppressionsService) Add(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("suppressions.add", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Suppressions,
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *SuppressionsService) Remove(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("suppressions.remove", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Suppression,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func suppressionsUnsubscribeQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Source, o.value("source")}}
}

func (s *SuppressionsService) ListUnsubscribes(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("suppressions.listUnsubscribes", opts)

	return s.t.page(ctx, o, call{path: paths.SuppressionUnsubscribes, query: suppressionsUnsubscribeQuery(o)}, cursorPages)
}

func (s *SuppressionsService) AddUnsubscribe(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("suppressions.addUnsubscribe", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SuppressionUnsubscribes,
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *SuppressionsService) RemoveUnsubscribe(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("suppressions.removeUnsubscribe", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SuppressionUnsubscribe,
		params: []any{"email", email},
		method: http.MethodDelete,
	})
}
