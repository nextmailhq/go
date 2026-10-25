package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func templatesListQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}, {query.Search, o.value("search")}, {query.Sort, o.value("sort")}}
}

func templatesWindowQuery(o *callOptions) values {
	return values{{query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}}
}

func (s *TemplatesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("templates.list", opts)

	return s.t.page(ctx, o, call{path: paths.Templates, query: templatesListQuery(o)}, cursorPages)
}

func (s *TemplatesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("templates.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Templates, query: templatesListQuery(o)}, cursorPages)
}

func (s *TemplatesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("templates.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Templates, query: templatesListQuery(o)}, cursorPages)
}

func (s *TemplatesService) Get(ctx context.Context, idOrSlug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.get", opts)

	return s.t.object(ctx, o, call{path: paths.Template, params: []any{"id", idOrSlug}})
}

func (s *TemplatesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Templates,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TemplatesService) Update(ctx context.Context, idOrSlug string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.update", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Template,
		params: []any{"id", idOrSlug},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *TemplatesService) Duplicate(ctx context.Context, idOrSlug string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.duplicate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateDuplicate,
		params: []any{"id", idOrSlug},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TemplatesService) ReplaceContent(ctx context.Context, idOrSlug string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.replaceContent", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateContent,
		params: []any{"id", idOrSlug},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TemplatesService) Delete(ctx context.Context, idOrSlug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Template,
		params: []any{"id", idOrSlug},
		method: http.MethodDelete,
	})
}

func (s *TemplatesService) ListVersions(ctx context.Context, idOrSlug string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("templates.listVersions", opts)

	return s.t.page(ctx, o, call{path: paths.TemplateVersions, params: []any{"id", idOrSlug}}, cursorPages)
}

func (s *TemplatesService) ListAllVersions(ctx context.Context, idOrSlug string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("templates.listAllVersions", opts)

	return s.t.all(ctx, o, call{path: paths.TemplateVersions, params: []any{"id", idOrSlug}}, cursorPages)
}

func (s *TemplatesService) IterateVersions(ctx context.Context, idOrSlug string, opts ...RequestOption) *Iterator {
	o := newCallOptions("templates.iterateVersions", opts)

	return s.t.iterate(ctx, o, call{path: paths.TemplateVersions, params: []any{"id", idOrSlug}}, cursorPages)
}

func (s *TemplatesService) GetVersion(ctx context.Context, idOrSlug string, version int, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.getVersion", opts)

	return s.t.object(ctx, o, call{path: paths.TemplateVersion, params: []any{"id", idOrSlug, "version", version}})
}

func (s *TemplatesService) Publish(ctx context.Context, idOrSlug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.publish", opts)
	request := call{
		path:       paths.TemplateVersions,
		params:     []any{"id", idOrSlug},
		method:     http.MethodPost,
		repeatable: true,
	}

	if o.has("expectedVersion") || o.has("expectedUpdatedAt") {
		request.body = optionsBody(o, "expectedVersion", "expectedUpdatedAt")
	}

	return s.t.object(ctx, o, request)
}

func (s *TemplatesService) RestoreVersion(ctx context.Context, idOrSlug string, version int, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.restoreVersion", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateVersionRestore,
		params: []any{"id", idOrSlug, "version", version},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TemplatesService) DeleteVersion(ctx context.Context, idOrSlug string, version int, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.deleteVersion", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateVersion,
		params: []any{"id", idOrSlug, "version", version},
		method: http.MethodDelete,
	})
}

func (s *TemplatesService) ListStarters(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("templates.listStarters", opts)

	return s.t.objects(ctx, o, call{path: paths.TemplateStarters})
}

func (s *TemplatesService) GetStarter(ctx context.Context, slug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.getStarter", opts)

	return s.t.object(ctx, o, call{path: paths.TemplateStarter, params: []any{"slug", slug}})
}

func (s *TemplatesService) ListFonts(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("templates.listFonts", opts)

	return s.t.objects(ctx, o, call{path: paths.TemplateFonts})
}

func (s *TemplatesService) Render(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.render", opts)

	return s.t.object(ctx, o, call{
		path:       paths.TemplateRender,
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *TemplatesService) Preview(ctx context.Context, idOrSlug string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.preview", opts)

	return s.t.object(ctx, o, call{
		path:       paths.TemplatePreview,
		params:     []any{"id", idOrSlug},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *TemplatesService) GetAnalytics(ctx context.Context, idOrSlug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.getAnalytics", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateAnalytics,
		params: []any{"id", idOrSlug},
		query:  templatesWindowQuery(o),
	})
}

func (s *TemplatesService) ListSends(ctx context.Context, idOrSlug string, opts ...RequestOption) (*TemplateSends, error) {
	o := newCallOptions("templates.listSends", opts)
	envelope, err := s.t.object(ctx, o, call{
		path:   paths.TemplateSends,
		params: []any{"id", idOrSlug},
		query: templatesWindowQuery(o).merge(values{
			{query.Page, o.value("page")},
			{query.PageSize, o.value("pageSize")},
			{query.Search, o.value("search")},
			{query.Source, o.value("source")},
			{query.Version, o.value("version")},
			{query.Opened, o.value("opened")},
			{query.Clicked, o.value("clicked")},
			{query.Tracked, o.value("tracked")},
		}),
	})

	if err != nil {
		return nil, err
	}

	sends := &TemplateSends{Items: objectsOf(envelope[fieldData]), Total: envelope.Int("total"), Page: 1, PageSize: envelope.Int("pageSize")}

	if envelope.Has("page") {
		sends.Page = envelope.Int("page")
	}

	return sends, nil
}

func (s *TemplatesService) Send(ctx context.Context, idOrSlug string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.send", opts)

	return s.t.object(ctx, o, call{
		path:       paths.TemplateSend,
		params:     []any{"id", idOrSlug},
		method:     http.MethodPost,
		body:       recipientsToWire(body),
		idempotent: true,
		repeatable: true,
	})
}

func (s *TemplatesService) ListImages(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("templates.listImages", opts)

	return s.t.page(ctx, o, call{path: paths.TemplateImages}, cursorPages)
}

func (s *TemplatesService) ListAllImages(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("templates.listAllImages", opts)

	return s.t.all(ctx, o, call{path: paths.TemplateImages}, cursorPages)
}

func (s *TemplatesService) IterateImages(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("templates.iterateImages", opts)

	return s.t.iterate(ctx, o, call{path: paths.TemplateImages}, cursorPages)
}

func (s *TemplatesService) UploadImage(ctx context.Context, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.uploadImage", opts)

	return s.t.object(ctx, o, call{
		path:        paths.TemplateImages,
		method:      http.MethodPost,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		upload:      true,
	})
}

func (s *TemplatesService) Design(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.design", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplatesDesign,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TemplatesService) Redesign(ctx context.Context, idOrSlug string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("templates.redesign", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TemplateRedesign,
		params: []any{"id", idOrSlug},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
