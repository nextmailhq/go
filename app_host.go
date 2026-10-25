package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *AppHostService) Get(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("appHost.get", opts)

	return s.t.object(ctx, o, call{path: paths.AppHost})
}

func (s *AppHostService) Set(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("appHost.set", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AppHost,
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *AppHostService) Verify(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("appHost.verify", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AppHostVerify,
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AppHostService) Delete(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("appHost.delete", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AppHost,
		method:     http.MethodDelete,
		repeatable: true,
	})
}
