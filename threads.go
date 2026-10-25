package nextmail

import (
	"context"
	"net/http"
	"time"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func threadsListQuery(o *callOptions) values {
	return values{{query.Folder, o.value("folder")}, {query.Query, o.value("query")}, {query.LabelIDs, o.value("labelIds")}, {query.Sort, o.value("sort")}, {query.DateFrom, o.value("dateFrom")}, {query.DateTo, o.value("dateTo")}, {query.FromContacts, o.value("fromContacts")}, {query.Semantic, o.value("semantic")}, {query.Category, o.value("category")}, {query.Address, o.value("address")}}
}

func (s *ThreadsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("threads.list", opts)

	return s.t.page(ctx, o, call{path: paths.Threads, query: threadsListQuery(o)}, pageTokenPages)
}

func (s *ThreadsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Threads, query: threadsListQuery(o)}, pageTokenPages)
}

func (s *ThreadsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("threads.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Threads, query: threadsListQuery(o)}, pageTokenPages)
}

func (s *ThreadsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.get", opts)

	return s.t.object(ctx, o, call{path: paths.Thread, params: []any{"id", id}})
}

func (s *ThreadsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Thread,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *ThreadsService) Trash(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.trash", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadTrash,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) Snooze(ctx context.Context, id string, wakeAt time.Time, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.snooze", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadSnooze,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       Body{"wakeAt": instant(wakeAt)},
		repeatable: true,
	})
}

func (s *ThreadsService) Unsnooze(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.unsnooze", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadUnsnooze,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) Mute(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.mute", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadMute,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) Unmute(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.unmute", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadUnmute,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) ListMessages(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listMessages", opts)

	return s.t.objects(ctx, o, call{path: paths.ThreadMessages, params: []any{"id", id}})
}

func (s *ThreadsService) GetMessage(ctx context.Context, id string, messageID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.getMessage", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadMessage, params: []any{"id", id, "messageId", messageID}})
}

func (s *ThreadsService) GetRawMessage(ctx context.Context, id string, messageID string, opts ...RequestOption) ([]byte, error) {
	o := newCallOptions("threads.getRawMessage", opts)

	return s.t.bytes(ctx, o, call{path: paths.ThreadMessageRaw, params: []any{"id", id, "messageId", messageID}})
}

func (s *ThreadsService) ListAttachments(ctx context.Context, id string, messageID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listAttachments", opts)

	return s.t.objects(ctx, o, call{path: paths.ThreadMessageAttachments, params: []any{"id", id, "messageId", messageID}})
}

func threadsFileLinksQuery(o *callOptions) values {
	return values{{query.IncludeRevoked, o.value("includeRevoked")}}
}

func (s *ThreadsService) ListFileLinks(ctx context.Context, id string, messageID string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("threads.listFileLinks", opts)

	return s.t.page(ctx, o, call{
		path:   paths.ThreadMessageFileLinks,
		params: []any{"id", id, "messageId", messageID},
		query:  threadsFileLinksQuery(o),
	}, cursorPages)
}

func (s *ThreadsService) ListAllFileLinks(ctx context.Context, id string, messageID string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listAllFileLinks", opts)

	return s.t.all(ctx, o, call{
		path:   paths.ThreadMessageFileLinks,
		params: []any{"id", id, "messageId", messageID},
		query:  threadsFileLinksQuery(o),
	}, cursorPages)
}

func (s *ThreadsService) IterateFileLinks(ctx context.Context, id string, messageID string, opts ...RequestOption) *Iterator {
	o := newCallOptions("threads.iterateFileLinks", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.ThreadMessageFileLinks,
		params: []any{"id", id, "messageId", messageID},
		query:  threadsFileLinksQuery(o),
	}, cursorPages)
}

func (s *ThreadsService) ListTracking(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listTracking", opts)

	return s.t.objects(ctx, o, call{
		path:   paths.ThreadTracking,
		params: []any{"id", id},
		query:  values{{query.Address, o.value("address")}},
	})
}

func (s *ThreadsService) GetDelivery(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.getDelivery", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadDelivery,
		params: []any{"id", id},
		query:  values{{query.Address, o.value("address")}},
	})
}

func (s *ThreadsService) ListNotes(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listNotes", opts)

	return s.t.objects(ctx, o, call{path: paths.ThreadNotes, params: []any{"id", id}})
}

func (s *ThreadsService) CreateNote(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.createNote", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadNotes,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ThreadsService) UpdateNote(ctx context.Context, id string, noteID string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.updateNote", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadNote,
		params:     []any{"id", id, "noteId", noteID},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *ThreadsService) DeleteNote(ctx context.Context, id string, noteID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.deleteNote", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadNote,
		params: []any{"id", id, "noteId", noteID},
		method: http.MethodDelete,
	})
}

func (s *ThreadsService) ReorderNotes(ctx context.Context, id string, ids []string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.reorderNotes", opts)

	return s.t.objects(ctx, o, call{
		path:       paths.ThreadNotesReorder,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       Body{"ids": ids},
		repeatable: true,
	})
}

func (s *ThreadsService) Counts(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.counts", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadsCounts, query: values{{query.Address, o.value("address")}, {query.HideLabelled, o.value("hideLabelled")}}})
}

func (s *ThreadsService) Summary(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.summary", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadSummary, params: []any{"id", id}})
}

func (s *ThreadsService) ReplySuggestions(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.replySuggestions", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadReplySuggestions, params: []any{"id", id}})
}

func (s *ThreadsService) Restore(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.restore", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadRestore,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Thread,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *ThreadsService) Unsubscribe(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.unsubscribe", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadUnsubscribe,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *ThreadsService) GetEvent(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.getEvent", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadEvent, params: []any{"id", id}})
}

func (s *ThreadsService) ListSenderCategories(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listSenderCategories", opts)

	return s.t.objects(ctx, o, call{path: paths.SenderCategories})
}

func (s *ThreadsService) SetSenderCategory(ctx context.Context, email string, category string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.setSenderCategory", opts)

	return s.t.object(ctx, o, call{
		path:       paths.SenderCategory,
		params:     []any{"email", email},
		method:     http.MethodPut,
		body:       Body{"category": category},
		repeatable: true,
	})
}

func (s *ThreadsService) ClearSenderCategory(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.clearSenderCategory", opts)

	return s.t.object(ctx, o, call{
		path:   paths.SenderCategory,
		params: []any{"email", email},
		method: http.MethodDelete,
	})
}

func (s *ThreadsService) GetTeam(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.getTeam", opts)

	return s.t.object(ctx, o, call{path: paths.ThreadTeam, params: []any{"id", id}})
}

func (s *ThreadsService) Assign(ctx context.Context, id string, assigneeID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.assign", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadAssignee,
		params:     []any{"id", id},
		method:     http.MethodPut,
		body:       Body{"assigneeId": assigneeID},
		repeatable: true,
	})
}

func (s *ThreadsService) Unassign(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.unassign", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadAssignee,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *ThreadsService) SetReplying(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.setReplying", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadReplying,
		params:     []any{"id", id},
		method:     http.MethodPut,
		repeatable: true,
	})
}

func (s *ThreadsService) ClearReplying(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.clearReplying", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadReplying,
		params:     []any{"id", id},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *ThreadsService) MarkDone(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.markDone", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadDone,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) Reopen(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.reopen", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadReopen,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ThreadsService) ListComments(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("threads.listComments", opts)

	return s.t.objects(ctx, o, call{path: paths.ThreadComments, params: []any{"id", id}})
}

func (s *ThreadsService) CreateComment(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.createComment", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadComments,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ThreadsService) UpdateComment(ctx context.Context, id string, commentID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.updateComment", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ThreadComment,
		params:     []any{"id", id, "commentId", commentID},
		method:     http.MethodPatch,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *ThreadsService) DeleteComment(ctx context.Context, id string, commentID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("threads.deleteComment", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ThreadComment,
		params: []any{"id", id, "commentId", commentID},
		method: http.MethodDelete,
	})
}
