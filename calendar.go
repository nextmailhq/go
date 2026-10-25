package nextmail

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

func calendarRangeQuery(from, to time.Time, o *callOptions) values {
	return values{
		{query.From, instant(from)},
		{query.To, instant(to)},
		{query.Timezone, o.value("timezone")},
		{query.IncludeTeam, o.value("includeTeam")},
		{query.Q, o.value("q")},
	}
}

func (s *CalendarService) ListEvents(ctx context.Context, from time.Time, to time.Time, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("calendar.listEvents", opts)

	return s.t.page(ctx, o, call{path: paths.CalendarEvents, query: calendarRangeQuery(from, to, o)}, cursorPages)
}

func (s *CalendarService) ListAllEvents(ctx context.Context, from time.Time, to time.Time, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("calendar.listAllEvents", opts)

	return s.t.all(ctx, o, call{path: paths.CalendarEvents, query: calendarRangeQuery(from, to, o)}, cursorPages)
}

func (s *CalendarService) IterateEvents(ctx context.Context, from time.Time, to time.Time, opts ...RequestOption) *Iterator {
	o := newCallOptions("calendar.iterateEvents", opts)

	return s.t.iterate(ctx, o, call{path: paths.CalendarEvents, query: calendarRangeQuery(from, to, o)}, cursorPages)
}

func (s *CalendarService) GetEvent(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.getEvent", opts)

	return s.t.object(ctx, o, call{path: paths.CalendarEvent, params: []any{"id", id}})
}

func (s *CalendarService) GetEventIcs(ctx context.Context, id string, opts ...RequestOption) (string, error) {
	o := newCallOptions("calendar.getEventIcs", opts)

	return s.t.text(ctx, o, call{
		path:   paths.CalendarEventIcs,
		params: []any{"id", id},
		accept: protocol.ContentTypesCalendar,
	})
}

func (s *CalendarService) CreateEvent(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.createEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEvents,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) UpdateEvent(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.updateEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEvent,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *CalendarService) DeleteEvent(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.deleteEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEvent,
		params: []any{"id", id},
		method: http.MethodDelete,
		query:  values{{query.Scope, o.value("scope")}, {query.OccurrenceStart, o.value("occurrenceStart")}},
	})
}

func (s *CalendarService) HideEvent(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.hideEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEventHide,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *CalendarService) UnhideEvent(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.unhideEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEventUnhide,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *CalendarService) FreeBusy(ctx context.Context, emails []string, from time.Time, to time.Time, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.freeBusy", opts)

	return s.t.object(ctx, o, call{
		path: paths.CalendarFreeBusy,
		query: values{
			{query.From, instant(from)},
			{query.To, instant(to)},
			{query.Timezone, o.value("timezone")},
			{query.Emails, strings.Join(emails, protocol.ListSeparator)},
			{query.DurationMinutes, o.value("durationMinutes")},
		},
	})
}

func (s *CalendarService) CancelEvent(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.cancelEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEventCancel,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) RespondToEvent(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.respondToEvent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEventRespond,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) DeclineProposal(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.declineProposal", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarEventDeclineProposal,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) ListFeeds(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("calendar.listFeeds", opts)

	return s.t.objects(ctx, o, call{path: paths.CalendarFeeds})
}

func (s *CalendarService) CreateFeed(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.createFeed", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarFeeds,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) ResetFeed(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.resetFeed", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarFeedReset,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *CalendarService) DeleteFeed(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.deleteFeed", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarFeed,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *CalendarService) ListSubscriptions(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("calendar.listSubscriptions", opts)

	return s.t.objects(ctx, o, call{path: paths.CalendarSubscriptions})
}

func (s *CalendarService) CreateSubscription(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.createSubscription", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarSubscriptions,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *CalendarService) UpdateSubscription(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.updateSubscription", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarSubscription,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *CalendarService) RefreshSubscription(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.refreshSubscription", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarSubscriptionRefresh,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *CalendarService) DeleteSubscription(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.deleteSubscription", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarSubscription,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *CalendarService) ImportIcs(ctx context.Context, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.importIcs", opts)

	return s.t.object(ctx, o, call{
		path:        paths.CalendarImports,
		method:      http.MethodPost,
		query:       values{{query.Filename, o.value("filename")}},
		raw:         data,
		sendsRaw:    true,
		contentType: protocol.ContentTypesCalendar,
	})
}

func (s *CalendarService) ListImports(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("calendar.listImports", opts)

	return s.t.page(ctx, o, call{path: paths.CalendarImports}, cursorPages)
}

func (s *CalendarService) ListAllImports(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("calendar.listAllImports", opts)

	return s.t.all(ctx, o, call{path: paths.CalendarImports}, cursorPages)
}

func (s *CalendarService) IterateImports(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("calendar.iterateImports", opts)

	return s.t.iterate(ctx, o, call{path: paths.CalendarImports}, cursorPages)
}

func (s *CalendarService) GetImport(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.getImport", opts)

	return s.t.object(ctx, o, call{path: paths.CalendarImport, params: []any{"id", id}})
}

func (s *CalendarService) UndoImport(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("calendar.undoImport", opts)

	return s.t.object(ctx, o, call{
		path:   paths.CalendarImportUndo,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}
