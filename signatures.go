package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func signaturesListQuery(o *callOptions) values {
	return values{{query.Address, o.value("address")}, {query.From, o.value("from")}}
}

func (s *SignaturesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("signatures.list", opts)

	return s.t.page(ctx, o, call{path: paths.Signatures, query: signaturesListQuery(o)}, cursorPages)
}

func (s *SignaturesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("signatures.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Signatures, query: signaturesListQuery(o)}, cursorPages)
}

func (s *SignaturesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("signatures.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Signatures, query: signaturesListQuery(o)}, cursorPages)
}

func (s *SignaturesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signatures.get", opts)

	return s.t.object(ctx, o, call{path: paths.Signature, params: []any{"id", id}})
}

func (s *SignaturesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signatures.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Signatures,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SignaturesService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signatures.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Signature,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *SignaturesService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signatures.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Signature,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *SignaturesService) Copy(ctx context.Context, id string, body Body, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("signatures.copy", opts)

	return s.t.objects(ctx, o, call{
		path:       paths.SignatureCopy,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}
