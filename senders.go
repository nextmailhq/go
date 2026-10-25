package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SendersService) Get(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("senders.get", opts)

	return s.t.object(ctx, o, call{path: paths.Sender, params: []any{"email", email}})
}

func (s *SendersService) GetAvatar(ctx context.Context, email string, opts ...RequestOption) ([]byte, error) {
	o := newCallOptions("senders.getAvatar", opts)

	return s.t.bytes(ctx, o, call{path: paths.SenderAvatar, params: []any{"email", email}})
}

func (s *SendersService) Research(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("senders.research", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SenderResearch,
		params:     []any{"email", email},
		method:     http.MethodPost,
		repeatable: true,
	})
}
