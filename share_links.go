package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func shareLinksListQuery(o *callOptions) values {
	return values{{query.Kind, o.value("kind")}}
}

func (s *ShareLinksService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("shareLinks.list", opts)

	return s.t.page(ctx, o, call{path: paths.ShareLinks, query: shareLinksListQuery(o)}, cursorPages)
}

func (s *ShareLinksService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("shareLinks.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.ShareLinks, query: shareLinksListQuery(o)}, cursorPages)
}

func (s *ShareLinksService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("shareLinks.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.ShareLinks, query: shareLinksListQuery(o)}, cursorPages)
}

func (s *ShareLinksService) Get(ctx context.Context, kind string, itemID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("shareLinks.get", opts)

	return s.t.object(ctx, o, call{path: paths.ShareLink, params: []any{"kind", kind, "itemId", itemID}})
}

func (s *ShareLinksService) Set(ctx context.Context, kind string, itemID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("shareLinks.set", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ShareLink,
		params:     []any{"kind", kind, "itemId", itemID},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *ShareLinksService) Delete(ctx context.Context, kind string, itemID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("shareLinks.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ShareLink,
		params: []any{"kind", kind, "itemId", itemID},
		method: http.MethodDelete,
	})
}
