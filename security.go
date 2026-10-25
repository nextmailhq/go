package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SecurityService) ListEvents(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("security.listEvents", opts)

	return s.t.page(ctx, o, call{path: paths.SecurityEvents}, cursorPages)
}

func (s *SecurityService) ListAllEvents(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("security.listAllEvents", opts)

	return s.t.all(ctx, o, call{path: paths.SecurityEvents}, cursorPages)
}

func (s *SecurityService) IterateEvents(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("security.iterateEvents", opts)

	return s.t.iterate(ctx, o, call{path: paths.SecurityEvents}, cursorPages)
}

func (s *SecurityService) StepUpStatus(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("security.stepUpStatus", opts)

	return s.t.object(ctx, o, call{path: paths.SecurityStepUp})
}

func (s *SecurityService) BeginStepUp(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("security.beginStepUp", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SecurityStepUp,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SecurityService) VerifyStepUp(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("security.verifyStepUp", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SecurityStepUpVerify,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SecurityService) GetPolicy(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("security.getPolicy", opts)

	return s.t.object(ctx, o, call{path: paths.SecurityPolicy})
}

func (s *SecurityService) UpdatePolicy(ctx context.Context, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("security.updatePolicy", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SecurityPolicy,
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}
