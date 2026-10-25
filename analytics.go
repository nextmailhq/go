package nextmail

import (
	"context"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func analyticsMailboxQuery(o *callOptions) values {
	return values{{query.From, o.value("from")}, {query.To, o.value("to")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Grain, o.value("grain")}, {query.Address, o.value("address")}}
}

func analyticsSendingQuery(o *callOptions) values {
	return values{{query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Grain, o.value("grain")}, {query.Address, o.value("address")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}, {query.Recipient, o.value("recipient")}, {query.Tag, o.value("tag")}}
}

func (s *AnalyticsService) Mailbox(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("analytics.mailbox", opts)

	return s.t.object(ctx, o, call{path: paths.AnalyticsMailbox, query: analyticsMailboxQuery(o)})
}

func (s *AnalyticsService) Sending(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("analytics.sending", opts)

	return s.t.object(ctx, o, call{path: paths.AnalyticsSending, query: analyticsSendingQuery(o)})
}
