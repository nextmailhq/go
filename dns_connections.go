package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *DNSConnectionsService) List(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("dnsConnections.list", opts)

	return s.t.object(ctx, o, call{path: paths.DNSConnections})
}

func (s *DNSConnectionsService) Start(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("dnsConnections.start", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DNSConnections,
		method: http.MethodPost,
	})
}

func (s *DNSConnectionsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("dnsConnections.get", opts)

	return s.t.object(ctx, o, call{path: paths.DNSConnection, params: []any{"id", id}})
}

func (s *DNSConnectionsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("dnsConnections.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.DNSConnection,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}
