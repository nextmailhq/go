package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *AccountService) SetName(ctx context.Context, name string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.setName", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Account,
		method:     http.MethodPatch,
		body:       Body{"name": name},
		repeatable: true,
	})
}

func (s *AccountService) GetNotifications(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.getNotifications", opts)

	return s.t.object(ctx, o, call{path: paths.AccountNotifications})
}

func (s *AccountService) SetEmailNotification(ctx context.Context, category string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.setEmailNotification", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountNotificationEmail,
		params:     []any{"category", category},
		method:     http.MethodPut,
		body:       pick(body, "enabled"),
		repeatable: true,
	})
}

func (s *AccountService) SetPushMuted(ctx context.Context, workspaceID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.setPushMuted", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountNotificationPush,
		params:     []any{"workspaceId", workspaceID},
		method:     http.MethodPut,
		body:       pick(body, "muted"),
		repeatable: true,
	})
}

func (s *AccountService) ListPushDevices(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listPushDevices", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountPushDevices})
}

func (s *AccountService) RemovePushDevice(ctx context.Context, deviceID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.removePushDevice", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountPushDevice,
		params: []any{"deviceId", deviceID},
		method: http.MethodDelete,
	})
}

func (s *AccountService) SendTestPush(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.sendTestPush", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountPushDevicesTest,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *AccountService) SetPhoto(ctx context.Context, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.setPhoto", opts)

	return s.t.object(ctx, o, call{
		path:        paths.AccountPhoto,
		method:      http.MethodPut,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		repeatable:  true,
		upload:      true,
	})
}

func (s *AccountService) RemovePhoto(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.removePhoto", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountPhoto,
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *AccountService) GetUsername(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.getUsername", opts)

	return s.t.object(ctx, o, call{path: paths.AccountUsername})
}

func (s *AccountService) CheckUsername(ctx context.Context, username string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.checkUsername", opts)

	return s.t.object(ctx, o, call{path: paths.AccountUsernameAvailability, query: values{{query.Username, username}}})
}

func (s *AccountService) SetUsername(ctx context.Context, username string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.setUsername", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountUsername,
		method: http.MethodPut,
		body:   Body{"username": username},
	})
}

func (s *AccountService) ListConnectedApps(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listConnectedApps", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountConnectedApps})
}

func (s *AccountService) ListConnectedAppRequests(ctx context.Context, clientID string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("account.listConnectedAppRequests", opts)

	return s.t.page(ctx, o, call{
		path:   paths.AccountConnectedAppRequests,
		params: []any{"clientId", clientID},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *AccountService) ListAllConnectedAppRequests(ctx context.Context, clientID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listAllConnectedAppRequests", opts)

	return s.t.all(ctx, o, call{
		path:   paths.AccountConnectedAppRequests,
		params: []any{"clientId", clientID},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *AccountService) IterateConnectedAppRequests(ctx context.Context, clientID string, opts ...RequestOption) *Iterator {
	o := newCallOptions("account.iterateConnectedAppRequests", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.AccountConnectedAppRequests,
		params: []any{"clientId", clientID},
		query:  keysRequestQuery(o),
	}, cursorPages)
}

func (s *AccountService) RevokeConnectedApp(ctx context.Context, clientID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.revokeConnectedApp", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountConnectedApp,
		params: []any{"clientId", clientID},
		method: http.MethodDelete,
	})
}

func (s *AccountService) EndConnectedAppElevation(ctx context.Context, clientID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.endConnectedAppElevation", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountConnectedAppElevation,
		params:     []any{"clientId", clientID},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *AccountService) ListSessions(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listSessions", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountSessions})
}

func (s *AccountService) RevokeSession(ctx context.Context, sessionID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.revokeSession", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountSession,
		params: []any{"sessionId", sessionID},
		method: http.MethodDelete,
	})
}

func (s *AccountService) RevokeSessions(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.revokeSessions", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountSessions,
		method: http.MethodDelete,
		query:  values{{query.ConnectedApps, o.value("connectedApps")}},
	})
}

func (s *AccountService) ListSecurityEvents(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("account.listSecurityEvents", opts)

	return s.t.page(ctx, o, call{path: paths.AccountSecurityEvents}, cursorPages)
}

func (s *AccountService) ListAllSecurityEvents(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listAllSecurityEvents", opts)

	return s.t.all(ctx, o, call{path: paths.AccountSecurityEvents}, cursorPages)
}

func (s *AccountService) IterateSecurityEvents(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("account.iterateSecurityEvents", opts)

	return s.t.iterate(ctx, o, call{path: paths.AccountSecurityEvents}, cursorPages)
}

func (s *AccountService) ListPasskeys(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listPasskeys", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountPasskeys})
}

func (s *AccountService) RenamePasskey(ctx context.Context, passkeyID string, name string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.renamePasskey", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountPasskey,
		params:     []any{"passkeyId", passkeyID},
		method:     http.MethodPatch,
		body:       Body{"name": name},
		repeatable: true,
	})
}

func (s *AccountService) RemovePasskey(ctx context.Context, passkeyID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.removePasskey", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountPasskey,
		params: []any{"passkeyId", passkeyID},
		method: http.MethodDelete,
	})
}

func (s *AccountService) ListInvitations(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listInvitations", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountInvitations})
}

func (s *AccountService) AcceptInvitation(ctx context.Context, invitationID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.acceptInvitation", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountInvitationAccept,
		params:     []any{"invitationId", invitationID},
		method:     http.MethodPost,
		body:       optionsBody(o, "activate"),
		repeatable: true,
	})
}

func (s *AccountService) DeclineInvitation(ctx context.Context, invitationID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.declineInvitation", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AccountInvitationDecline,
		params:     []any{"invitationId", invitationID},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AccountService) ListOwnershipOffers(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("account.listOwnershipOffers", opts)

	return s.t.objects(ctx, o, call{path: paths.AccountOwnershipOffers})
}

func (s *AccountService) AcceptOwnershipOffer(ctx context.Context, workspaceID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.acceptOwnershipOffer", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountOwnershipOfferAccept,
		params: []any{"workspaceId", workspaceID},
		method: http.MethodPost,
	})
}

func (s *AccountService) DeclineOwnershipOffer(ctx context.Context, workspaceID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("account.declineOwnershipOffer", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AccountOwnershipOfferDecline,
		params: []any{"workspaceId", workspaceID},
		method: http.MethodPost,
	})
}
