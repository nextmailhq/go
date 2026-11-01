package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SignupsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signups.create", opts)

	return s.t.object(ctx, o, call{
		path:      paths.Signups,
		method:    http.MethodPost,
		body:      bodyOf(body),
		anonymous: true,
	})
}

func (s *SignupsService) Verify(ctx context.Context, signUpID string, code string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("signups.verify", opts)

	return s.t.object(ctx, o, call{
		path:      paths.SignupVerify,
		params:    []any{"id", signUpID},
		method:    http.MethodPost,
		body:      Body{"code": code},
		anonymous: true,
	})
}
