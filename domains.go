package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func domainsListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}}
}

func (s *DomainsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("domains.list", opts)

	return s.t.page(ctx, o, call{path: paths.Domains, query: domainsListQuery(o)}, cursorPages)
}

func (s *DomainsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("domains.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Domains, query: domainsListQuery(o)}, cursorPages)
}

func (s *DomainsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("domains.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Domains, query: domainsListQuery(o)}, cursorPages)
}

func (s *DomainsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.get", opts)

	return s.t.object(ctx, o, call{path: paths.Domain, params: []any{"id", id}})
}

func (s *DomainsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Domains,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) Verify(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.verify", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainVerify,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *DomainsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Domain,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Domain,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) GetLogo(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getLogo", opts)

	return s.t.object(ctx, o, call{path: paths.DomainLogo, params: []any{"id", id}})
}

func (s *DomainsService) SetLogo(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.setLogo", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainLogo,
		params:     []any{"id", id},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *DomainsService) RemoveLogo(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.removeLogo", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainLogo,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *DomainsService) SetLogoCertificate(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.setLogoCertificate", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainLogoCertificate,
		params:     []any{"id", id},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *DomainsService) RemoveLogoCertificate(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.removeLogoCertificate", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainLogoCertificate,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *DomainsService) ListAddresses(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("domains.listAddresses", opts)

	return s.t.page(ctx, o, call{path: paths.DomainAddresses, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) ListAllAddresses(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("domains.listAllAddresses", opts)

	return s.t.all(ctx, o, call{path: paths.DomainAddresses, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) IterateAddresses(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("domains.iterateAddresses", opts)

	return s.t.iterate(ctx, o, call{path: paths.DomainAddresses, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) CreateAddress(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.createAddress", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddresses,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *DomainsService) GetAddress(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getAddress", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddress, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) UpdateAddress(ctx context.Context, id string, addressID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateAddress", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddress,
		params:     []any{"id", id, "addressId", addressID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) SetAddressPhoto(ctx context.Context, id string, addressID string, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.setAddressPhoto", opts)

	return s.t.object(ctx, o, call{
		path:        paths.DomainAddressPhoto,
		params:      []any{"id", id, "addressId", addressID},
		method:      http.MethodPut,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		repeatable:  true,
		upload:      true,
	})
}

func (s *DomainsService) RemoveAddressPhoto(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.removeAddressPhoto", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddressPhoto,
		params:     []any{"id", id, "addressId", addressID},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *DomainsService) StartGravatarPublish(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.startGravatarPublish", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressPhotoGravatar,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
	})
}

func (s *DomainsService) DeleteAddress(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.deleteAddress", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddress,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) ListAddressForwards(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.listAddressForwards", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressForwards, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) AddAddressForwards(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.addAddressForwards", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressForwards,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) UpdateAddressForward(ctx context.Context, id string, addressID string, forwardID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateAddressForward", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddressForward,
		params:     []any{"id", id, "addressId", addressID, "forwardId", forwardID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) DeleteAddressForward(ctx context.Context, id string, addressID string, forwardID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.deleteAddressForward", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressForward,
		params: []any{"id", id, "addressId", addressID, "forwardId", forwardID},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) ResendAddressForwardConsent(ctx context.Context, id string, addressID string, forwardID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.resendAddressForwardConsent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressForwardResend,
		params: []any{"id", id, "addressId", addressID, "forwardId", forwardID},
		method: http.MethodPost,
	})
}

func (s *DomainsService) ListAddressMembers(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.listAddressMembers", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressMembers, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) GetAddressLogin(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getAddressLogin", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressLogin, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) SetAddressLogin(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.setAddressLogin", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressLogin,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPut,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) DeleteAddressLogin(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.deleteAddressLogin", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressLogin,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) GetMailAppSettings(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getMailAppSettings", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressMailApps, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) UpdateMailAppSettings(ctx context.Context, id string, addressID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateMailAppSettings", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddressMailApps,
		params:     []any{"id", id, "addressId", addressID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) GetAddressLimits(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getAddressLimits", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressLimits, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) UpdateAddressLimits(ctx context.Context, id string, addressID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateAddressLimits", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddressLimits,
		params:     []any{"id", id, "addressId", addressID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) ListAppPasswords(ctx context.Context, id string, addressID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.listAppPasswords", opts)

	return s.t.object(ctx, o, call{path: paths.DomainAddressAppPasswords, params: []any{"id", id, "addressId", addressID}})
}

func (s *DomainsService) CreateAppPassword(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.createAppPassword", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressAppPasswords,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) UpdateAppPassword(ctx context.Context, id string, addressID string, passwordID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateAppPassword", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainAddressAppPassword,
		params:     []any{"id", id, "addressId", addressID, "passwordId", passwordID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *DomainsService) DeleteAppPassword(ctx context.Context, id string, addressID string, passwordID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.deleteAppPassword", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressAppPassword,
		params: []any{"id", id, "addressId", addressID, "passwordId", passwordID},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) SendMailAppSetup(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.sendMailAppSetup", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressMailAppSetup,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) CreateMailAppProfileLink(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.createMailAppProfileLink", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressMailAppProfileLink,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) GetDNS(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getDns", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainDNS,
		params: []any{"id", id},
		query:  values{{query.Refresh, o.value("refresh")}},
	})
}

func (s *DomainsService) SetDNSZone(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.setDnsZone", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainDNS,
		params: []any{"id", id},
		method: http.MethodPut,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) SyncDNS(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.syncDns", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainDNSSync,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) ListRoutes(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("domains.listRoutes", opts)

	return s.t.page(ctx, o, call{path: paths.DomainRoutes, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) ListAllRoutes(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("domains.listAllRoutes", opts)

	return s.t.all(ctx, o, call{path: paths.DomainRoutes, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) IterateRoutes(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("domains.iterateRoutes", opts)

	return s.t.iterate(ctx, o, call{path: paths.DomainRoutes, params: []any{"id", id}}, cursorPages)
}

func (s *DomainsService) CreateRoute(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.createRoute", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainRoutes,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) GetRoute(ctx context.Context, id string, routeID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.getRoute", opts)

	return s.t.object(ctx, o, call{path: paths.DomainRoute, params: []any{"id", id, "routeId", routeID}})
}

func (s *DomainsService) UpdateRoute(ctx context.Context, id string, routeID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.updateRoute", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainRoute,
		params: []any{"id", id, "routeId", routeID},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *DomainsService) TestRoute(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.testRoute", opts)

	return s.t.object(ctx, o, call{
		path:       paths.DomainRouteTest,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *DomainsService) DeleteRoute(ctx context.Context, id string, routeID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.deleteRoute", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainRoute,
		params: []any{"id", id, "routeId", routeID},
		method: http.MethodDelete,
	})
}

func (s *DomainsService) ListGroupMembers(ctx context.Context, id string, addressID string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("domains.listGroupMembers", opts)

	return s.t.page(ctx, o, call{path: paths.DomainAddressGroupMembers, params: []any{"id", id, "addressId", addressID}}, cursorPages)
}

func (s *DomainsService) ListAllGroupMembers(ctx context.Context, id string, addressID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("domains.listAllGroupMembers", opts)

	return s.t.all(ctx, o, call{path: paths.DomainAddressGroupMembers, params: []any{"id", id, "addressId", addressID}}, cursorPages)
}

func (s *DomainsService) IterateGroupMembers(ctx context.Context, id string, addressID string, opts ...RequestOption) *Iterator {
	o := newCallOptions("domains.iterateGroupMembers", opts)

	return s.t.iterate(ctx, o, call{path: paths.DomainAddressGroupMembers, params: []any{"id", id, "addressId", addressID}}, cursorPages)
}

func (s *DomainsService) AddGroupMembers(ctx context.Context, id string, addressID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.addGroupMembers", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressGroupMembers,
		params: []any{"id", id, "addressId", addressID},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *DomainsService) RemoveGroupMember(ctx context.Context, id string, addressID string, memberID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("domains.removeGroupMember", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DomainAddressGroupMember,
		params: []any{"id", id, "addressId", addressID, "memberId", memberID},
		method: http.MethodDelete,
	})
}
