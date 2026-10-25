package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
)

func (s *LabelsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("labels.list", opts)

	return s.t.page(ctx, o, call{path: paths.Labels}, cursorPages)
}

func (s *LabelsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("labels.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Labels}, cursorPages)
}

func (s *LabelsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("labels.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Labels}, cursorPages)
}

func (s *LabelsService) ListColors(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("labels.listColors", opts)

	return s.t.objects(ctx, o, call{path: paths.LabelColors})
}

func (s *LabelsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("labels.get", opts)

	return s.t.object(ctx, o, call{path: paths.Label, params: []any{"id", id}})
}

func (s *LabelsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("labels.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Labels,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *LabelsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("labels.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Label,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *LabelsService) Test(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("labels.test", opts)

	return s.t.object(ctx, o, call{
		path:   paths.LabelTest,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *LabelsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("labels.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Label,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}
