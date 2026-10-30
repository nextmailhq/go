package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SharedService) Get(ctx context.Context, token string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("shared.get", opts)

	return s.t.object(ctx, o, call{path: paths.SharedItem, params: []any{"token", token}})
}

func (s *SharedService) Copy(ctx context.Context, token string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("shared.copy", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SharedItemCopy,
		params: []any{"token", token},
		method: http.MethodPost,
	})
}
