package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func providerImportsListQuery(o *callOptions) values {
	return values{{query.Source, o.value("source")}, {query.Status, o.value("status")}, {query.Search, o.value("search")}}
}

func (s *ProviderImportsService) Inspect(ctx context.Context, apiKey string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("providerImports.inspect", opts)
	body := optionsBody(o, "provider")
	body["apiKey"] = apiKey

	return s.t.object(ctx, o, call{
		path:   paths.ProviderImportsInspect,
		method: http.MethodPost,
		body:   body,
	})
}

func (s *ProviderImportsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("providerImports.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ProviderImports,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ProviderImportsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("providerImports.list", opts)

	return s.t.page(ctx, o, call{path: paths.ProviderImports, query: providerImportsListQuery(o)}, cursorPages)
}

func (s *ProviderImportsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("providerImports.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.ProviderImports, query: providerImportsListQuery(o)}, cursorPages)
}

func (s *ProviderImportsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("providerImports.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.ProviderImports, query: providerImportsListQuery(o)}, cursorPages)
}

func (s *ProviderImportsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("providerImports.get", opts)

	return s.t.object(ctx, o, call{path: paths.ProviderImport, params: []any{"id", id}})
}

func (s *ProviderImportsService) Cancel(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("providerImports.cancel", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ProviderImportCancel,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}
