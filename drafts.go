package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func draftsListQuery(o *callOptions) values {
	return values{{query.Query, o.value("query")}}
}

func (s *DraftsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("drafts.list", opts)

	return s.t.page(ctx, o, call{path: paths.Drafts, query: draftsListQuery(o)}, pageTokenPages)
}

func (s *DraftsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("drafts.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Drafts, query: draftsListQuery(o)}, pageTokenPages)
}

func (s *DraftsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("drafts.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Drafts, query: draftsListQuery(o)}, pageTokenPages)
}

func (s *DraftsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("drafts.get", opts)

	return s.t.object(ctx, o, call{path: paths.Draft, params: []any{"id", id}})
}

func (s *DraftsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("drafts.create", opts)
	wire, err := attachmentsToWire(bodyOf(body))

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:   paths.Drafts,
		method: http.MethodPost,
		body:   wire,
	})
}

func (s *DraftsService) Update(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("drafts.update", opts)
	wire, err := attachmentsToWire(bodyOf(body))

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:   paths.Draft,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   wire,
	})
}

func (s *DraftsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("drafts.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Draft,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}
