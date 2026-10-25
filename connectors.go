package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func connectorsAppQuery(o *callOptions) values {
	return values{{query.Search, o.value("search")}, {query.Category, o.value("category")}}
}

func connectorsListQuery(o *callOptions) values {
	return values{{query.Scope, o.value("scope")}, {query.Status, o.value("status")}, {query.App, o.value("app")}}
}

func connectorsActionQuery(o *callOptions) values {
	return values{{query.Search, o.value("search")}, {query.Limit, o.value("limit")}}
}

func connectorsCallQuery(o *callOptions) values {
	return values{{query.ConnectorID, o.value("connectorId")}, {query.Scope, o.value("scope")}}
}

func (s *ConnectorsService) ListApps(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("connectors.listApps", opts)

	return s.t.page(ctx, o, call{path: paths.ConnectorApps, query: connectorsAppQuery(o)}, cursorPages)
}

func (s *ConnectorsService) ListAllApps(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectors.listAllApps", opts)

	return s.t.all(ctx, o, call{path: paths.ConnectorApps, query: connectorsAppQuery(o)}, cursorPages)
}

func (s *ConnectorsService) IterateApps(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("connectors.iterateApps", opts)

	return s.t.iterate(ctx, o, call{path: paths.ConnectorApps, query: connectorsAppQuery(o)}, cursorPages)
}

func (s *ConnectorsService) GetApp(ctx context.Context, slug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.getApp", opts)

	return s.t.object(ctx, o, call{path: paths.ConnectorApp, params: []any{"slug", slug}})
}

func (s *ConnectorsService) ListCategories(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectors.listCategories", opts)

	return s.t.objects(ctx, o, call{path: paths.ConnectorCategories})
}

func (s *ConnectorsService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectors.list", opts)

	return s.t.objects(ctx, o, call{path: paths.Connectors, query: connectorsListQuery(o)})
}

func (s *ConnectorsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.get", opts)

	return s.t.object(ctx, o, call{path: paths.Connector, params: []any{"id", id}})
}

func (s *ConnectorsService) Connect(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.connect", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Connectors,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ConnectorsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.update", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Connector,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *ConnectorsService) Disconnect(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.disconnect", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Connector,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *ConnectorsService) Reconnect(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.reconnect", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ConnectorReconnect,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *ConnectorsService) ListActions(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectors.listActions", opts)

	return s.t.objects(ctx, o, call{
		path:   paths.ConnectorActions,
		params: []any{"id", id},
		query:  connectorsActionQuery(o),
	})
}

func (s *ConnectorsService) GetAction(ctx context.Context, id string, action string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.getAction", opts)

	return s.t.object(ctx, o, call{path: paths.ConnectorAction, params: []any{"id", id, "action", action}})
}

func (s *ConnectorsService) RunAction(ctx context.Context, id string, action string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("connectors.runAction", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ConnectorActionRun,
		params: []any{"id", id, "action", action},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ConnectorsService) ListCalls(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("connectors.listCalls", opts)

	return s.t.page(ctx, o, call{path: paths.ConnectorCalls, query: connectorsCallQuery(o)}, cursorPages)
}

func (s *ConnectorsService) ListAllCalls(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("connectors.listAllCalls", opts)

	return s.t.all(ctx, o, call{path: paths.ConnectorCalls, query: connectorsCallQuery(o)}, cursorPages)
}

func (s *ConnectorsService) IterateCalls(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("connectors.iterateCalls", opts)

	return s.t.iterate(ctx, o, call{path: paths.ConnectorCalls, query: connectorsCallQuery(o)}, cursorPages)
}
