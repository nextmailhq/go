package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *NotificationsService) List(ctx context.Context, opts ...RequestOption) (*NotificationPage, error) {
	o := newCallOptions("notifications.list", opts)
	envelope, err := s.t.object(ctx, o, call{
		path:  paths.Notifications,
		query: values{{query.Limit, o.value("limit")}, {query.Cursor, o.value("cursor")}},
	})

	if err != nil {
		return nil, err
	}

	next := cursorOf(envelope[fieldNextCursor])

	return &NotificationPage{
		Items:       objectsOf(envelope[fieldData]),
		UnreadCount: envelope.Int("unreadCount"),
		HasMore:     hasMoreOf(envelope, next),
		NextCursor:  next,
	}, nil
}

func (s *NotificationsService) UnreadCount(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("notifications.unreadCount", opts)

	return s.t.object(ctx, o, call{path: paths.NotificationsUnreadCount})
}

func (s *NotificationsService) MarkRead(ctx context.Context, ids []string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("notifications.markRead", opts)

	return s.t.object(ctx, o, call{
		path:       paths.NotificationsRead,
		method:     http.MethodPost,
		body:       Body{"ids": ids},
		repeatable: true,
	})
}

func (s *NotificationsService) MarkAllRead(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("notifications.markAllRead", opts)

	return s.t.object(ctx, o, call{
		path:       paths.NotificationsReadAll,
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *NotificationsService) Clear(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("notifications.clear", opts)

	return s.t.object(ctx, o, call{
		path:       paths.NotificationsClear,
		method:     http.MethodPost,
		body:       optionsBody(o, "declineInvitations"),
		repeatable: true,
	})
}
