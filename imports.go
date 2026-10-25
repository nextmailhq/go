package nextmail

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func importsListQuery(o *callOptions) values {
	return values{{"addressId", o.value("addressId")}, {query.Source, o.value("source")}, {query.Status, o.value("status")}, {query.Search, o.value("search")}}
}

func (s *ImportsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("imports.list", opts)

	return s.t.page(ctx, o, call{path: paths.Imports, query: importsListQuery(o)}, cursorPages)
}

func (s *ImportsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("imports.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Imports, query: importsListQuery(o)}, cursorPages)
}

func (s *ImportsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("imports.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Imports, query: importsListQuery(o)}, cursorPages)
}

func (s *ImportsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.get", opts)

	return s.t.object(ctx, o, call{path: paths.Import, params: []any{"id", id}})
}

func (s *ImportsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Imports,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ImportsService) UploadState(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.uploadState", opts)

	return s.t.object(ctx, o, call{path: paths.ImportUpload, params: []any{"id", id}})
}

func (s *ImportsService) UploadChunk(ctx context.Context, id string, file int, chunk int, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.uploadChunk", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ImportFileChunk,
		params:     []any{"id", id, "file", file, "chunk", chunk},
		method:     http.MethodPut,
		raw:        data,
		sendsRaw:   true,
		repeatable: true,
	})
}

func (s *ImportsService) Start(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.start", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ImportStart,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ImportsService) Cancel(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.cancel", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportCancel,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *ImportsService) ListFailures(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.listFailures", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportFailures,
		params: []any{"id", id},
		query:  values{{"after", o.value("after")}, {"limit", o.value("limit")}},
	})
}

func (s *ImportsService) DeleteUpload(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.deleteUpload", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportUpload,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *ImportsService) DiscoverMailbox(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.discoverMailbox", opts)

	return s.t.object(ctx, o, call{path: paths.ImportMailboxSettings, query: values{{query.Email, email}}})
}

func (s *ImportsService) CheckMailbox(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.checkMailbox", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportMailboxCheck,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ImportsService) ConnectMailbox(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.connectMailbox", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportMailbox,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ImportsService) Resume(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("imports.resume", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ImportResume,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ImportsService) ImportFiles(ctx context.Context, input Body, opts ...RequestOption) (Object, error) {
	sources, err := importSources(input["files"])

	if err != nil {
		return nil, err
	}

	files := make([]Body, 0, len(sources))

	for _, source := range sources {
		files = append(files, Body{"name": source.name, "bytes": len(source.data)})
	}

	body := pick(input, "addressId", "options")
	body["files"] = files

	created, err := s.Create(ctx, body, opts...)

	if err != nil || len(sources) == 0 {
		return created, err
	}

	total := created.Int("totalBytes")
	chunkBytes := created.Int("chunkBytes")
	planned, _ := created["files"].([]any)
	progress, _ := input["onProgress"].(func(uploaded, total int))
	uploaded := 0

	for index, source := range sources {
		chunks := 1

		if index < len(planned) && objectOf(planned[index]).Has("chunks") {
			chunks = objectOf(planned[index]).Int("chunks")
		}

		for chunk := 0; chunk < chunks; chunk++ {
			begin := min(len(source.data), chunk*chunkBytes)
			end := min(len(source.data), chunk*chunkBytes+chunkBytes)

			if _, err := s.UploadChunk(ctx, created.ID(), index, chunk, bytes.NewReader(source.data[begin:end]), opts...); err != nil {
				return nil, err
			}

			uploaded += end - begin

			if progress != nil {
				progress(uploaded, total)
			}
		}
	}

	return s.Start(ctx, created.ID(), opts...)
}

type importSource struct {
	name any
	data []byte
}

func importSources(value any) ([]importSource, error) {
	if value == nil {
		return nil, nil
	}

	items, ok := listOf(value)

	if !ok {
		return nil, argumentError(messageImportFilesShape)
	}

	sources := make([]importSource, 0, len(items))

	for _, item := range items {
		file, ok := mapOf(item)

		if !ok {
			return nil, argumentError(messageImportFileShape)
		}

		source := importSource{name: file["name"]}

		switch data := file["data"].(type) {
		case []byte:
			source.data = data
		case io.Reader:
			read, err := io.ReadAll(data)

			if err != nil {
				return nil, argumentFailure(err, messageRawBodyUnreadable, err.Error())
			}

			source.data = read
		default:
			return nil, argumentError(messageImportFileShape)
		}

		sources = append(sources, source)
	}

	return sources, nil
}
