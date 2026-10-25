package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func bookingsPagesQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Status, o.value("status")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}, {query.Sort, o.value("sort")}}
}

func bookingsListQuery(o *callOptions) values {
	return values{{query.View, o.value("view")}, {query.PageID, o.value("pageId")}, {query.PageIDs, o.value("pageIds")}, {query.Q, o.value("q")}, {query.From, o.value("from")}, {query.To, o.value("to")}}
}

func bookingsTimesQuery(o *callOptions) values {
	return values{{query.From, o.value("from")}, {query.To, o.value("to")}, {query.DurationMinutes, o.value("durationMinutes")}, {query.BookingID, o.value("bookingId")}}
}

func bookingsAnalyticsQuery(o *callOptions) values {
	return values{{query.PageIDs, o.value("pageIds")}, {query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func (s *BookingsService) ListPages(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("bookings.listPages", opts)

	return s.t.objects(ctx, o, call{path: paths.BookingPages, query: bookingsPagesQuery(o)})
}

func (s *BookingsService) GetPage(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.getPage", opts)

	return s.t.object(ctx, o, call{path: paths.BookingPage, params: []any{"id", id}})
}

func (s *BookingsService) CreatePage(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.createPage", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingPages,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BookingsService) UpdatePage(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.updatePage", opts)

	return s.t.object(ctx, o, call{
		path:       paths.BookingPage,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *BookingsService) UploadBackground(ctx context.Context, id string, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.uploadBackground", opts)

	return s.t.object(ctx, o, call{
		path:        paths.BookingPageBackground,
		params:      []any{"id", id},
		method:      http.MethodPut,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		repeatable:  true,
		upload:      true,
	})
}

func (s *BookingsService) RemoveBackground(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.removeBackground", opts)

	return s.t.object(ctx, o, call{
		path:       paths.BookingPageBackground,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *BookingsService) DeletePage(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.deletePage", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingPage,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *BookingsService) ListTimes(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.listTimes", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingPageTimes,
		params: []any{"id", id},
		query:  bookingsTimesQuery(o),
	})
}

func (s *BookingsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("bookings.list", opts)

	return s.t.page(ctx, o, call{path: paths.Bookings, query: bookingsListQuery(o)}, cursorPages)
}

func (s *BookingsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("bookings.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Bookings, query: bookingsListQuery(o)}, cursorPages)
}

func (s *BookingsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("bookings.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Bookings, query: bookingsListQuery(o)}, cursorPages)
}

func (s *BookingsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.get", opts)

	return s.t.object(ctx, o, call{path: paths.Booking, params: []any{"id", id}})
}

func (s *BookingsService) Approve(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.approve", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingApprove,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *BookingsService) Decline(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.decline", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingDecline,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BookingsService) Cancel(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.cancel", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingCancel,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BookingsService) Reschedule(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.reschedule", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingReschedule,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   pick(body, "start", "reason"),
	})
}

func (s *BookingsService) ResendEmail(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.resendEmail", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BookingResend,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *BookingsService) Analytics(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("bookings.analytics", opts)

	return s.t.object(ctx, o, call{path: paths.BookingsAnalytics, query: bookingsAnalyticsQuery(o)})
}
