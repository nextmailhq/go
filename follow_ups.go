package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *FollowUpsService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("followUps.list", opts)

	return s.t.objects(ctx, o, call{
		path:  paths.FollowUps,
		query: values{{query.Status, o.value("status")}, {query.ThreadID, o.value("threadId")}, {query.Limit, o.value("limit")}},
	})
}

func (s *FollowUpsService) Set(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("followUps.set", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FollowUps,
		method:     http.MethodPost,
		body:       pick(body, "threadId", "remindAt"),
		repeatable: true,
	})
}

func (s *FollowUpsService) Cancel(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("followUps.cancel", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FollowUp,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}
