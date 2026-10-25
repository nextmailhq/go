package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *MeService) Get(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("me.get", opts)

	return s.t.object(ctx, o, call{path: paths.KeySelf})
}

func (s *MeService) Ping(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("me.ping", opts)

	return s.t.object(ctx, o, call{path: paths.Ping})
}

func (s *MeService) Rotate(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("me.rotate", opts)

	return s.t.object(ctx, o, call{path: paths.KeySelfRotate, method: http.MethodPost})
}

func (s *MeService) ListRequests(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("me.listRequests", opts)

	return s.t.page(ctx, o, call{path: paths.KeySelfRequests, query: keysRequestQuery(o)}, cursorPages)
}

func (s *MeService) ListAllRequests(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("me.listAllRequests", opts)

	return s.t.all(ctx, o, call{path: paths.KeySelfRequests, query: keysRequestQuery(o)}, cursorPages)
}

func (s *MeService) IterateRequests(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("me.iterateRequests", opts)

	return s.t.iterate(ctx, o, call{path: paths.KeySelfRequests, query: keysRequestQuery(o)}, cursorPages)
}

func (s *MeService) ListActivity(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("me.listActivity", opts)

	return s.t.page(ctx, o, call{path: paths.KeySelfActivity, query: keysWindowQuery(o)}, cursorPages)
}

func (s *MeService) ListAllActivity(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("me.listAllActivity", opts)

	return s.t.all(ctx, o, call{path: paths.KeySelfActivity, query: keysWindowQuery(o)}, cursorPages)
}

func (s *MeService) IterateActivity(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("me.iterateActivity", opts)

	return s.t.iterate(ctx, o, call{path: paths.KeySelfActivity, query: keysWindowQuery(o)}, cursorPages)
}

func (s *MeService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("me.stats", opts)

	return s.t.object(ctx, o, call{
		path:  paths.KeySelfStats,
		query: keysWindowQuery(o).merge(values{{query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}}),
	})
}
