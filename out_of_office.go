package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *OutOfOfficeService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("outOfOffice.list", opts)

	return s.t.objects(ctx, o, call{path: paths.OutOfOffice, query: values{{query.Address, o.value("address")}}})
}

func (s *OutOfOfficeService) Set(ctx context.Context, address string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("outOfOffice.set", opts)

	return s.t.object(ctx, o, call{
		path:       paths.OutOfOffice,
		method:     http.MethodPut,
		query:      values{{query.Address, address}},
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *OutOfOfficeService) Clear(ctx context.Context, address string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("outOfOffice.clear", opts)

	return s.t.object(ctx, o, call{
		path:       paths.OutOfOffice,
		method:     http.MethodDelete,
		query:      values{{query.Address, address}},
		repeatable: true,
	})
}
