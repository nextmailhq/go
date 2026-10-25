package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func filesListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Kind, o.value("kind")}, {query.Direction, o.value("direction")}, {query.Address, o.value("address")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}, {query.Sort, o.value("sort")}}
}

func (s *FilesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("files.list", opts)

	return s.t.page(ctx, o, call{path: paths.Files, query: filesListQuery(o)}, cursorPages)
}

func (s *FilesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("files.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Files, query: filesListQuery(o)}, cursorPages)
}

func (s *FilesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("files.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Files, query: filesListQuery(o)}, cursorPages)
}

func (s *FilesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.get", opts)

	return s.t.object(ctx, o, call{path: paths.File, params: []any{"id", id}})
}

func (s *FilesService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.stats", opts)

	return s.t.object(ctx, o, call{path: paths.FilesStats})
}

func (s *FilesService) Download(ctx context.Context, id string, opts ...RequestOption) ([]byte, error) {
	o := newCallOptions("files.download", opts)

	return s.t.bytes(ctx, o, call{path: paths.FileContent, params: []any{"id", id}})
}

func (s *FilesService) ListLinks(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("files.listLinks", opts)

	return s.t.page(ctx, o, call{path: paths.FileLinks, params: []any{"id", id}}, cursorPages)
}

func (s *FilesService) ListAllLinks(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("files.listAllLinks", opts)

	return s.t.all(ctx, o, call{path: paths.FileLinks, params: []any{"id", id}}, cursorPages)
}

func (s *FilesService) IterateLinks(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("files.iterateLinks", opts)

	return s.t.iterate(ctx, o, call{path: paths.FileLinks, params: []any{"id", id}}, cursorPages)
}

func (s *FilesService) CreateLink(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.createLink", opts)
	body := optionsBody(o, "reuse")

	if domain := o.text("domain"); domain != "" {
		body["domain"] = domain
	}

	return s.t.object(ctx, o, call{
		path:   paths.FileLinks,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   body,
	})
}

func (s *FilesService) RevokeLink(ctx context.Context, id string, linkID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.revokeLink", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FileLink,
		params:     []any{"id", id, "linkId", linkID},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *FilesService) RevokeAllLinks(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.revokeAllLinks", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FileLinks,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *FilesService) Upload(ctx context.Context, data io.Reader, filename string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.upload", opts)
	name, err := uploadName(filename)

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:        paths.Files,
		method:      http.MethodPost,
		query:       values{{query.Filename, name}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		upload:      true,
	})
}

func (s *FilesService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.File,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *FilesService) DeleteMany(ctx context.Context, ids []string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("files.deleteMany", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FilesBatchDelete,
		method: http.MethodPost,
		body:   Body{"ids": ids},
	})
}
