package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func chatsListQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Sort, o.value("sort")}}
}

func (s *ChatsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("chats.list", opts)

	return s.t.page(ctx, o, call{path: paths.Chats, query: chatsListQuery(o)}, cursorPages)
}

func (s *ChatsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("chats.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Chats, query: chatsListQuery(o)}, cursorPages)
}

func (s *ChatsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("chats.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Chats, query: chatsListQuery(o)}, cursorPages)
}

func (s *ChatsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.get", opts)

	return s.t.object(ctx, o, call{path: paths.Chat, params: []any{"id", id}})
}

func (s *ChatsService) Rename(ctx context.Context, id string, title string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.rename", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Chat,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       Body{"title": title},
		repeatable: true,
	})
}

func (s *ChatsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Chat,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func chatsMessagesQuery(o *callOptions) values {
	return values{{query.Order, o.value("order")}}
}

func (s *ChatsService) ListMessages(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("chats.listMessages", opts)

	return s.t.page(ctx, o, call{path: paths.ChatMessages, params: []any{"id", id}, query: chatsMessagesQuery(o)}, cursorPages)
}

func (s *ChatsService) ListAllMessages(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("chats.listAllMessages", opts)

	return s.t.all(ctx, o, call{path: paths.ChatMessages, params: []any{"id", id}, query: chatsMessagesQuery(o)}, cursorPages)
}

func (s *ChatsService) IterateMessages(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("chats.iterateMessages", opts)

	return s.t.iterate(ctx, o, call{path: paths.ChatMessages, params: []any{"id", id}, query: chatsMessagesQuery(o)}, cursorPages)
}

func (s *ChatsService) SendMessage(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.sendMessage", opts)
	wire, err := attachmentsToWire(bodyOf(body))

	if err != nil {
		return nil, err
	}

	return s.t.object(ctx, o, call{
		path:   paths.ChatMessages,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   wire,
		turn:   true,
	})
}

func (s *ChatsService) StopTurn(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.stopTurn", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ChatStop,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ChatsService) ApproveToolCall(ctx context.Context, id string, toolCallID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.approveToolCall", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ChatToolCallApprove,
		params: []any{"id", id, "toolCallId", toolCallID},
		method: http.MethodPost,
		body:   bodyOf(body),
		turn:   true,
	})
}

func (s *ChatsService) DenyToolCall(ctx context.Context, id string, toolCallID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("chats.denyToolCall", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ChatToolCallDeny,
		params: []any{"id", id, "toolCallId", toolCallID},
		method: http.MethodPost,
		turn:   true,
	})
}
