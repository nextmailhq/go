package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func knowledgeListQuery(o *callOptions) values {
	return values{{query.Scope, o.value("scope")}, {query.Level, o.value("level")}, {query.Kind, o.value("kind")}, {query.Status, o.value("status")}, {query.Pinned, o.value("pinned")}, {query.Q, o.value("q")}}
}

func knowledgeSuggestionQuery(o *callOptions) values {
	return values{{query.Kind, o.value("kind")}, {query.Status, o.value("status")}}
}

func (s *KnowledgeService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("knowledge.list", opts)

	return s.t.page(ctx, o, call{path: paths.Knowledge, query: knowledgeListQuery(o)}, cursorPages)
}

func (s *KnowledgeService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("knowledge.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Knowledge, query: knowledgeListQuery(o)}, cursorPages)
}

func (s *KnowledgeService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("knowledge.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Knowledge, query: knowledgeListQuery(o)}, cursorPages)
}

func (s *KnowledgeService) Levels(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("knowledge.levels", opts)

	return s.t.objects(ctx, o, call{path: paths.KnowledgeLevels})
}

func (s *KnowledgeService) Usage(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.usage", opts)

	return s.t.object(ctx, o, call{path: paths.KnowledgeUsage})
}

func (s *KnowledgeService) Search(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.search", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeSearch,
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *KnowledgeService) CreateNote(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.createNote", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeNotes,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *KnowledgeService) AddLink(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.addLink", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeLinks,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *KnowledgeService) UploadFile(ctx context.Context, data io.Reader, filename string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.uploadFile", opts)
	name, err := uploadName(filename)

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:        paths.KnowledgeFiles,
		method:      http.MethodPost,
		query:       values{{query.Filename, name}, {query.Scope, o.value("scope")}, {query.Title, o.value("title")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		upload:      true,
	})
}

func (s *KnowledgeService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.get", opts)

	return s.t.object(ctx, o, call{path: paths.KnowledgeItem, params: []any{"id", id}})
}

func (s *KnowledgeService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeItem,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *KnowledgeService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeItem,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *KnowledgeService) Refresh(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.refresh", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeItemRefresh,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *KnowledgeService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.stats", opts)

	return s.t.object(ctx, o, call{path: paths.KnowledgeStats, query: values{{query.Days, o.value("days")}}})
}

func (s *KnowledgeService) ListSuggestions(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("knowledge.listSuggestions", opts)

	return s.t.page(ctx, o, call{path: paths.KnowledgeSuggestions, query: knowledgeSuggestionQuery(o)}, cursorPages)
}

func (s *KnowledgeService) ListAllSuggestions(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("knowledge.listAllSuggestions", opts)

	return s.t.all(ctx, o, call{path: paths.KnowledgeSuggestions, query: knowledgeSuggestionQuery(o)}, cursorPages)
}

func (s *KnowledgeService) IterateSuggestions(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("knowledge.iterateSuggestions", opts)

	return s.t.iterate(ctx, o, call{path: paths.KnowledgeSuggestions, query: knowledgeSuggestionQuery(o)}, cursorPages)
}

func (s *KnowledgeService) AcceptSuggestion(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.acceptSuggestion", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeSuggestionAccept,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *KnowledgeService) DismissSuggestion(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.dismissSuggestion", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeSuggestionDismiss,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *KnowledgeService) ListFlags(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("knowledge.listFlags", opts)

	return s.t.objects(ctx, o, call{
		path:  paths.KnowledgeFlags,
		query: values{{query.ItemID, o.value("itemId")}, {query.Limit, o.value("limit")}},
	})
}

func (s *KnowledgeService) DismissFlag(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.dismissFlag", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeFlagDismiss,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *KnowledgeService) ListConnectors(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("knowledge.listConnectors", opts)

	return s.t.objects(ctx, o, call{path: paths.KnowledgeConnectors})
}

func (s *KnowledgeService) AddConnector(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.addConnector", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeConnectors,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *KnowledgeService) GetConnector(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.getConnector", opts)

	return s.t.object(ctx, o, call{path: paths.KnowledgeConnector, params: []any{"id", id}})
}

func (s *KnowledgeService) UpdateConnector(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.updateConnector", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeConnector,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *KnowledgeService) DeleteConnector(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.deleteConnector", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeConnector,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *KnowledgeService) SyncConnector(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.syncConnector", opts)

	return s.t.object(ctx, o, call{
		path:       paths.KnowledgeConnectorSync,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *KnowledgeService) DraftFromThread(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("knowledge.draftFromThread", opts)

	return s.t.object(ctx, o, call{
		path:   paths.KnowledgeDrafts,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}
