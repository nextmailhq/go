package nextmail

import (
	"context"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func trackingHitQuery(o *callOptions) values {
	return values{{query.IncludeMachine, o.value("includeMachine")}, {query.Address, o.value("address")}}
}

func trackingListQuery(o *callOptions) values {
	return values{{query.Opened, o.value("opened")}, {query.Clicked, o.value("clicked")}, {query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Address, o.value("address")}}
}

func (s *TrackingService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tracking.list", opts)

	return s.t.page(ctx, o, call{path: paths.Tracking, query: trackingListQuery(o)}, cursorPages)
}

func (s *TrackingService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tracking.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Tracking, query: trackingListQuery(o)}, cursorPages)
}

func (s *TrackingService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("tracking.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Tracking, query: trackingListQuery(o)}, cursorPages)
}

func (s *TrackingService) GetStats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tracking.getStats", opts)

	return s.t.object(ctx, o, call{
		path:  paths.TrackingStats,
		query: values{{query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Address, o.value("address")}},
	})
}

func (s *TrackingService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tracking.get", opts)

	return s.t.object(ctx, o, call{path: paths.Tracked, params: []any{"id", id}})
}

func (s *TrackingService) ListOpens(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tracking.listOpens", opts)

	return s.t.page(ctx, o, call{
		path:   paths.TrackedOpens,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}

func (s *TrackingService) ListAllOpens(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tracking.listAllOpens", opts)

	return s.t.all(ctx, o, call{
		path:   paths.TrackedOpens,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}

func (s *TrackingService) IterateOpens(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("tracking.iterateOpens", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.TrackedOpens,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}

func (s *TrackingService) ListClicks(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tracking.listClicks", opts)

	return s.t.page(ctx, o, call{
		path:   paths.TrackedClicks,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}

func (s *TrackingService) ListAllClicks(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tracking.listAllClicks", opts)

	return s.t.all(ctx, o, call{
		path:   paths.TrackedClicks,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}

func (s *TrackingService) IterateClicks(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("tracking.iterateClicks", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.TrackedClicks,
		params: []any{"id", id},
		query:  trackingHitQuery(o),
	}, cursorPages)
}
