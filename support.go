package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SupportService) Send(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("support.send", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Support,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
