package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *WorkspacesService) List(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.list", opts)

	return s.t.object(ctx, o, call{path: paths.Workspaces})
}

func (s *WorkspacesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Workspaces,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *WorkspacesService) GetActive(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.getActive", opts)

	return s.t.object(ctx, o, call{path: paths.WorkspacesActive})
}

func (s *WorkspacesService) SetActive(ctx context.Context, workspaceID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.setActive", opts)

	return s.t.object(ctx, o, call{
		path:       paths.WorkspacesActive,
		method:     http.MethodPut,
		body:       Body{"workspaceId": workspaceID},
		repeatable: true,
	})
}

func (s *WorkspacesService) GetOwnershipOffer(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.getOwnershipOffer", opts)

	return s.t.object(ctx, o, call{path: paths.WorkspaceOwnershipOffer, params: []any{"id", id}})
}

func (s *WorkspacesService) OfferOwnership(ctx context.Context, id string, userID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.offerOwnership", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WorkspaceOwnershipOffer,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   Body{"userId": userID},
	})
}

func (s *WorkspacesService) WithdrawOwnershipOffer(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.withdrawOwnershipOffer", opts)

	return s.t.object(ctx, o, call{
		path:   paths.WorkspaceOwnershipOffer,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *WorkspacesService) Delete(ctx context.Context, id string, confirm string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("workspaces.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Workspace,
		params: []any{"id", id},
		method: http.MethodDelete,
		query:  values{{query.Confirm, confirm}},
	})
}
