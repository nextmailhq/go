package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *EventsService) Send(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("events.send", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Events,
		method:     http.MethodPost,
		body:       bodyOf(body),
		idempotent: true,
		repeatable: true,
	})
}

func (s *EventsService) SendBatch(ctx context.Context, events []Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("events.sendBatch", opts)
	wire := make([]Body, 0, len(events))

	for _, event := range events {
		wire = append(wire, bodyOf(event))
	}

	return s.t.object(ctx, o, call{
		path:       paths.EventsBatch,
		method:     http.MethodPost,
		body:       Body{"events": wire},
		idempotent: true,
		repeatable: true,
	})
}

func (s *EventsService) ListNames(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("events.listNames", opts)

	return s.t.objects(ctx, o, call{path: paths.EventNames})
}
