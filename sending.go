package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SendingService) ListStreams(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("sending.listStreams", opts)

	return s.t.object(ctx, o, call{path: paths.SendingStreams})
}

func (s *SendingService) ResumeStream(ctx context.Context, stream string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("sending.resumeStream", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SendingStreamResume,
		params: []any{"stream", stream},
		method: http.MethodPost,
	})
}

func (s *SendingService) ListDedicatedIPs(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("sending.listDedicatedIps", opts)

	return s.t.object(ctx, o, call{path: paths.SendingDedicatedIPs})
}

func (s *SendingService) GetDedicatedIP(ctx context.Context, ip string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("sending.getDedicatedIp", opts)

	return s.t.object(ctx, o, call{path: paths.SendingDedicatedIP, params: []any{"ip", ip}})
}
