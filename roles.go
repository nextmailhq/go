package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *RolesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("roles.list", opts)

	return s.t.page(ctx, o, call{path: paths.Roles}, cursorPages)
}

func (s *RolesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("roles.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Roles}, cursorPages)
}

func (s *RolesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("roles.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Roles}, cursorPages)
}

func (s *RolesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("roles.get", opts)

	return s.t.object(ctx, o, call{path: paths.Role, params: []any{"id", id}})
}

func (s *RolesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("roles.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Roles,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *RolesService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("roles.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Role,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *RolesService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("roles.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Role,
		params: []any{"id", id},
		method: http.MethodDelete,
		query:  values{{query.ReassignTo, o.value("reassignTo")}},
	})
}

func (s *RolesService) ListPermissions(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("roles.listPermissions", opts)

	return s.t.objects(ctx, o, call{path: paths.RolePermissions})
}
