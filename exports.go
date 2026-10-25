package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *ExportsService) Preview(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("exports.preview", opts)

	return s.t.object(ctx, o, call{path: paths.MailboxExportsPreview})
}

func (s *ExportsService) List(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("exports.list", opts)

	return s.t.object(ctx, o, call{
		path:  paths.MailboxExports,
		query: values{{query.Limit, o.value("limit")}, {query.Offset, o.value("offset")}},
	})
}

func (s *ExportsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("exports.get", opts)

	return s.t.object(ctx, o, call{path: paths.MailboxExport, params: []any{"id", id}})
}

func (s *ExportsService) Start(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("exports.start", opts)

	return s.t.object(ctx, o, call{path: paths.MailboxExports, method: http.MethodPost})
}

func (s *ExportsService) Download(ctx context.Context, id string, opts ...RequestOption) ([]byte, error) {
	o := newCallOptions("exports.download", opts)

	return s.t.bytes(ctx, o, call{path: paths.MailboxExportContent, params: []any{"id", id}})
}
