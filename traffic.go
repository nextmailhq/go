package nextmail

import (
	"context"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func trafficListQuery(o *callOptions) values {
	return values{
		{query.DomainID, o.value("domainId")},
		{query.Direction, o.value("direction")},
		{query.Outcome, o.value("outcomes")},
		{query.Address, o.value("address")},
		{query.Recipient, o.value("recipient")},
		{query.Since, o.value("since")},
		{query.Until, o.value("until")},
	}
}

func (s *TrafficService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("traffic.list", opts)

	return s.t.page(ctx, o, call{path: paths.Traffic, query: trafficListQuery(o)}, cursorPages)
}

func (s *TrafficService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("traffic.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Traffic, query: trafficListQuery(o)}, cursorPages)
}

func (s *TrafficService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("traffic.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Traffic, query: trafficListQuery(o)}, cursorPages)
}
