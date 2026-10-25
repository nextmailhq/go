package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func keysListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}}
}

func keysWindowQuery(o *callOptions) values {
	return values{{query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func keysRequestQuery(o *callOptions) values {
	return keysWindowQuery(o).merge(values{{query.FailedOnly, o.value("failedOnly")}, {query.Status, o.value("statuses")}, {query.Path, o.value("path")}})
}

func keysWorkspaceRequestQuery(o *callOptions) values {
	return keysRequestQuery(o).merge(values{{query.KeyIDs, o.value("keyIds")}})
}

func keysWorkspaceActivityQuery(o *callOptions) values {
	return keysWindowQuery(o).merge(values{{query.KeyIDs, o.value("keyIds")}})
}

func (s *KeysService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("keys.list", opts)

	return s.t.page(ctx, o, call{path: paths.Keys, query: keysListQuery(o)}, cursorPages)
}

func (s *KeysService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("keys.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Keys, query: keysListQuery(o)}, cursorPages)
}

func (s *KeysService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("keys.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Keys, query: keysListQuery(o)}, cursorPages)
}

func (s *KeysService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.get", opts)

	return s.t.object(ctx, o, call{path: paths.Key, params: []any{"id", id}})
}

func (s *KeysService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Keys,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *KeysService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Key,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *KeysService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Key,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *KeysService) Rotate(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.rotate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KeyRotate,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *KeysService) Revoke(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.revoke", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KeyRevoke,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *KeysService) ListRequests(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("keys.listRequests", opts)

	return s.t.page(ctx, o, call{
		path:   paths.KeyRequests,
		params: []any{"id", id},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *KeysService) ListAllRequests(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("keys.listAllRequests", opts)

	return s.t.all(ctx, o, call{
		path:   paths.KeyRequests,
		params: []any{"id", id},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *KeysService) IterateRequests(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("keys.iterateRequests", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.KeyRequests,
		params: []any{"id", id},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *KeysService) ListActivity(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("keys.listActivity", opts)

	return s.t.page(ctx, o, call{
		path:   paths.KeyActivity,
		params: []any{"id", id},
		query:  keysWindowQuery(o),
	}, cursorPages)
}

func (s *KeysService) ListAllActivity(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("keys.listAllActivity", opts)

	return s.t.all(ctx, o, call{
		path:   paths.KeyActivity,
		params: []any{"id", id},
		query:  keysWindowQuery(o),
	}, cursorPages)
}

func (s *KeysService) IterateActivity(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("keys.iterateActivity", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.KeyActivity,
		params: []any{"id", id},
		query:  keysWindowQuery(o),
	}, cursorPages)
}

func (s *KeysService) ListWorkspaceRequests(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("keys.listWorkspaceRequests", opts)

	return s.t.page(ctx, o, call{path: paths.KeysRequests, query: keysWorkspaceRequestQuery(o)}, cursorPages)
}

func (s *KeysService) ListAllWorkspaceRequests(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("keys.listAllWorkspaceRequests", opts)

	return s.t.all(ctx, o, call{path: paths.KeysRequests, query: keysWorkspaceRequestQuery(o)}, cursorPages)
}

func (s *KeysService) IterateWorkspaceRequests(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("keys.iterateWorkspaceRequests", opts)

	return s.t.iterate(ctx, o, call{path: paths.KeysRequests, query: keysWorkspaceRequestQuery(o)}, cursorPages)
}

func (s *KeysService) ListWorkspaceActivity(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("keys.listWorkspaceActivity", opts)

	return s.t.page(ctx, o, call{path: paths.KeysActivity, query: keysWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *KeysService) ListAllWorkspaceActivity(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("keys.listAllWorkspaceActivity", opts)

	return s.t.all(ctx, o, call{path: paths.KeysActivity, query: keysWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *KeysService) IterateWorkspaceActivity(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("keys.iterateWorkspaceActivity", opts)

	return s.t.iterate(ctx, o, call{path: paths.KeysActivity, query: keysWorkspaceActivityQuery(o)}, cursorPages)
}

func (s *KeysService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("keys.stats", opts)

	return s.t.object(ctx, o, call{
		path:  paths.KeysStats,
		query: keysWindowQuery(o).merge(values{{query.KeyIDs, o.value("keyIds")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}}),
	})
}
