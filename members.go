package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func membersListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}}
}

func (s *MembersService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("members.list", opts)

	return s.t.page(ctx, o, call{path: paths.Members, query: membersListQuery(o)}, cursorPages)
}

func (s *MembersService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("members.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Members, query: membersListQuery(o)}, cursorPages)
}

func (s *MembersService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("members.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Members, query: membersListQuery(o)}, cursorPages)
}

func (s *MembersService) Get(ctx context.Context, userID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.get", opts)

	return s.t.object(ctx, o, call{path: paths.Member, params: []any{"userId", userID}})
}

func (s *MembersService) Add(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.add", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Members,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *MembersService) Update(ctx context.Context, userID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Member,
		params:     []any{"userId", userID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *MembersService) Remove(ctx context.Context, userID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.remove", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Member,
		params: []any{"userId", userID},
		method: http.MethodDelete,
	})
}

func (s *MembersService) GrantAddress(ctx context.Context, userID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.grantAddress", opts)

	return s.t.object(ctx, o, call{
		path:       paths.MemberAddresses,
		params:     []any{"userId", userID},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *MembersService) RevokeAddress(ctx context.Context, userID string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.revokeAddress", opts)

	return s.t.object(ctx, o, call{
		path:   paths.MemberAddress,
		params: []any{"userId", userID, "addressId", addressID},
		method: http.MethodDelete,
	})
}

func (s *MembersService) GrantDomain(ctx context.Context, userID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.grantDomain", opts)

	return s.t.object(ctx, o, call{
		path:       paths.MemberDomains,
		params:     []any{"userId", userID},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *MembersService) RevokeDomain(ctx context.Context, userID string, domainID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.revokeDomain", opts)

	return s.t.object(ctx, o, call{
		path:   paths.MemberDomain,
		params: []any{"userId", userID, "domainId", domainID},
		method: http.MethodDelete,
	})
}

func (s *MembersService) ListInvitations(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("members.listInvitations", opts)

	return s.t.page(ctx, o, call{path: paths.MemberInvitations}, cursorPages)
}

func (s *MembersService) ListAllInvitations(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("members.listAllInvitations", opts)

	return s.t.all(ctx, o, call{path: paths.MemberInvitations}, cursorPages)
}

func (s *MembersService) IterateInvitations(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("members.iterateInvitations", opts)

	return s.t.iterate(ctx, o, call{path: paths.MemberInvitations}, cursorPages)
}

func (s *MembersService) RevokeInvitation(ctx context.Context, invitationID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.revokeInvitation", opts)

	return s.t.object(ctx, o, call{
		path:   paths.MemberInvitation,
		params: []any{"invitationId", invitationID},
		method: http.MethodDelete,
	})
}

func (s *MembersService) ResendInvitation(ctx context.Context, invitationID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.resendInvitation", opts)

	return s.t.object(ctx, o, call{
		path:   paths.MemberInvitationResend,
		params: []any{"invitationId", invitationID},
		method: http.MethodPost,
	})
}

func (s *MembersService) GetDirectorySync(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.getDirectorySync", opts)

	return s.t.object(ctx, o, call{path: paths.DirectorySync})
}

func (s *MembersService) ConnectDirectorySync(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.connectDirectorySync", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DirectorySync,
		method: http.MethodPost,
	})
}

func (s *MembersService) DisconnectDirectorySync(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.disconnectDirectorySync", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DirectorySync,
		method: http.MethodDelete,
	})
}

func (s *MembersService) ListDirectoryPeople(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.listDirectoryPeople", opts)

	return s.t.object(ctx, o, call{
		path:  paths.DirectorySyncPeople,
		query: values{{query.StartIndex, o.value("startIndex")}, {query.Count, o.value("count")}},
	})
}

func (s *MembersService) SetDirectoryDefaultRole(ctx context.Context, roleID *string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.setDirectoryDefaultRole", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DirectorySyncDefaultRole,
		method:     http.MethodPut,
		body:       Body{"roleId": nullable(roleID)},
		repeatable: true,
	})
}

func (s *MembersService) SetDirectoryGroupRole(ctx context.Context, groupID string, roleID *string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("members.setDirectoryGroupRole", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DirectorySyncGroupRole,
		params:     []any{"groupId", groupID},
		method:     http.MethodPut,
		body:       Body{"roleId": nullable(roleID)},
		repeatable: true,
	})
}
