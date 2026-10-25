package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *SegmentsService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("segments.list", opts)

	return s.t.objects(ctx, o, call{path: paths.Segments})
}

func (s *SegmentsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("segments.get", opts)

	return s.t.object(ctx, o, call{path: paths.Segment, params: []any{"id", id}})
}

func (s *SegmentsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("segments.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Segments,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *SegmentsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("segments.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Segment,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *SegmentsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("segments.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Segment,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *SegmentsService) Preview(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("segments.preview", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SegmentsPreview,
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}
