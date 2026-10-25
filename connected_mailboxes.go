package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *ConnectedMailboxesService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectedMailboxes.list", opts)

	return s.t.objects(ctx, o, call{path: paths.ConnectedMailboxes, query: values{{query.Q, o.value("q")}}})
}

func (s *ConnectedMailboxesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.get", opts)

	return s.t.object(ctx, o, call{path: paths.ConnectedMailbox, params: []any{"id", id}})
}

func (s *ConnectedMailboxesService) Discover(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.discover", opts)

	return s.t.object(ctx, o, call{path: paths.ConnectedMailboxSettings, query: values{{query.Email, email}}})
}

func (s *ConnectedMailboxesService) Connect(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.connect", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ConnectedMailboxes,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ConnectedMailboxesService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.update", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ConnectedMailbox,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *ConnectedMailboxesService) Sync(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.sync", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ConnectedMailboxSync,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ConnectedMailboxesService) Disconnect(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectedMailboxes.disconnect", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ConnectedMailbox,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}
