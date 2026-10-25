package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *ToolsService) CheckDmarc(ctx context.Context, domain string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.checkDmarc", opts)

	return s.t.object(ctx, o, call{path: paths.ToolsDmarc, query: values{{query.Domain, domain}}})
}

func (s *ToolsService) CheckBimi(ctx context.Context, domain string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.checkBimi", opts)

	return s.t.object(ctx, o, call{path: paths.ToolsBimi, query: values{{query.Domain, domain}}})
}

func (s *ToolsService) CheckDeliverability(ctx context.Context, domain string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.checkDeliverability", opts)

	return s.t.object(ctx, o, call{path: paths.ToolsDeliverability, query: values{{query.Domain, domain}}})
}

func (s *ToolsService) ReadVerificationCode(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.readVerificationCode", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ToolsVerificationCode,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ToolsService) EstimateCost(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.estimateCost", opts)

	return s.t.object(ctx, o, call{
		path:  paths.ToolsCost,
		query: values{{query.Sends, o.value("sends")}, {query.Domains, o.value("domains")}, {query.People, o.value("people")}, {query.AiActions, o.value("aiActions")}, {query.Interval, o.value("interval")}},
	})
}

func (s *ToolsService) CheckAddress(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.checkAddress", opts)

	return s.t.object(ctx, o, call{
		path:  paths.ToolsAddress,
		query: values{{query.Email, email}, {query.Deep, o.value("deep")}},
	})
}

func (s *ToolsService) CheckAddresses(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tools.checkAddresses", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ToolsAddresses,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
