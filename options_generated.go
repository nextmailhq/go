// Code generated from the TypeScript SDK. DO NOT EDIT.

package nextmail

import (
	"time"
)

func WithActivate(activate bool) RequestOption {
	return callOption{function: "WithActivate", name: "activate", value: activate}
}

func WithAddress(address string) RequestOption {
	return callOption{function: "WithAddress", name: "address", value: address}
}

func WithAddressID(addressID string) RequestOption {
	return callOption{function: "WithAddressID", name: "addressId", value: addressID}
}

func WithAfter(after int) RequestOption {
	return callOption{function: "WithAfter", name: "after", value: after}
}

func WithAiActions(aiActions int) RequestOption {
	return callOption{function: "WithAiActions", name: "aiActions", value: aiActions}
}

func WithApp(app string) RequestOption {
	return callOption{function: "WithApp", name: "app", value: app}
}

func WithAudienceID(audienceID string) RequestOption {
	return callOption{function: "WithAudienceID", name: "audienceId", value: audienceID}
}

func WithAudienceIDs(audienceIDs ...string) RequestOption {
	return callOption{function: "WithAudienceIDs", name: "audienceIds", value: audienceIDs}
}

func WithAutomationIDs(automationIDs ...string) RequestOption {
	return callOption{function: "WithAutomationIDs", name: "automationIds", value: automationIDs}
}

func WithBlocked(blocked bool) RequestOption {
	return callOption{function: "WithBlocked", name: "blocked", value: blocked}
}

func WithBookingID(bookingID string) RequestOption {
	return callOption{function: "WithBookingID", name: "bookingId", value: bookingID}
}

func WithBroadcastID(broadcastID string) RequestOption {
	return callOption{function: "WithBroadcastID", name: "broadcastId", value: broadcastID}
}

func WithBroadcastIDs(broadcastIDs ...string) RequestOption {
	return callOption{function: "WithBroadcastIDs", name: "broadcastIds", value: broadcastIDs}
}

func WithCategory(category string) RequestOption {
	return callOption{function: "WithCategory", name: "category", value: category}
}

func WithClicked(clicked bool) RequestOption {
	return callOption{function: "WithClicked", name: "clicked", value: clicked}
}

func WithConnectedApps(connectedApps bool) RequestOption {
	return callOption{function: "WithConnectedApps", name: "connectedApps", value: connectedApps}
}

func WithConnectorID(connectorID string) RequestOption {
	return callOption{function: "WithConnectorID", name: "connectorId", value: connectorID}
}

func WithContentType(contentType string) RequestOption {
	return callOption{function: "WithContentType", name: "contentType", value: contentType}
}

func WithCount(count int) RequestOption {
	return callOption{function: "WithCount", name: "count", value: count}
}

func WithCrop(crop ImageCrop) RequestOption {
	return callOption{function: "WithCrop", name: "crop", value: crop}
}

func WithCursor(cursor string) RequestOption {
	return callOption{function: "WithCursor", name: "cursor", value: cursor}
}

func WithDateFrom(dateFrom time.Time) RequestOption {
	return callOption{function: "WithDateFrom", name: "dateFrom", value: dateFrom}
}

func WithDateTo(dateTo time.Time) RequestOption {
	return callOption{function: "WithDateTo", name: "dateTo", value: dateTo}
}

func WithDays(days int) RequestOption {
	return callOption{function: "WithDays", name: "days", value: days}
}

func WithDaysWindow(days string) RequestOption {
	return callOption{function: "WithDaysWindow", name: "days", value: days}
}

func WithDeep(deep bool) RequestOption {
	return callOption{function: "WithDeep", name: "deep", value: deep}
}

func WithDirection(direction string) RequestOption {
	return callOption{function: "WithDirection", name: "direction", value: direction}
}

func WithDomain(domain string) RequestOption {
	return callOption{function: "WithDomain", name: "domain", value: domain}
}

func WithDomains(domains int) RequestOption {
	return callOption{function: "WithDomains", name: "domains", value: domains}
}

func WithDurationMinutes(durationMinutes int) RequestOption {
	return callOption{function: "WithDurationMinutes", name: "durationMinutes", value: durationMinutes}
}

func WithEmail(email string) RequestOption {
	return callOption{function: "WithEmail", name: "email", value: email}
}

func WithEnabled(enabled bool) RequestOption {
	return callOption{function: "WithEnabled", name: "enabled", value: enabled}
}

func WithEndpointIDs(endpointIDs ...string) RequestOption {
	return callOption{function: "WithEndpointIDs", name: "endpointIds", value: endpointIDs}
}

func WithEvent(event string) RequestOption {
	return callOption{function: "WithEvent", name: "event", value: event}
}

func WithExpectedUpdatedAt(expectedUpdatedAt string) RequestOption {
	return callOption{function: "WithExpectedUpdatedAt", name: "expectedUpdatedAt", value: expectedUpdatedAt}
}

func WithExpectedVersion(expectedVersion int) RequestOption {
	return callOption{function: "WithExpectedVersion", name: "expectedVersion", value: expectedVersion}
}

func WithFailedOnly(failedOnly bool) RequestOption {
	return callOption{function: "WithFailedOnly", name: "failedOnly", value: failedOnly}
}

func WithFilename(filename string) RequestOption {
	return callOption{function: "WithFilename", name: "filename", value: filename}
}

func WithFilter(filter string) RequestOption {
	return callOption{function: "WithFilter", name: "filter", value: filter}
}

func WithFolder(folder string) RequestOption {
	return callOption{function: "WithFolder", name: "folder", value: folder}
}

func WithFromAddress(from string) RequestOption {
	return callOption{function: "WithFromAddress", name: "from", value: from}
}

func WithFrom(from time.Time) RequestOption {
	return callOption{function: "WithFrom", name: "from", value: from}
}

func WithFromContacts(fromContacts bool) RequestOption {
	return callOption{function: "WithFromContacts", name: "fromContacts", value: fromContacts}
}

func WithGrain(grain string) RequestOption {
	return callOption{function: "WithGrain", name: "grain", value: grain}
}

func WithGroupID(groupID string) RequestOption {
	return callOption{function: "WithGroupID", name: "groupId", value: groupID}
}

func WithHideLabelled(hideLabelled bool) RequestOption {
	return callOption{function: "WithHideLabelled", name: "hideLabelled", value: hideLabelled}
}

func WithInactiveDays(inactiveDays int) RequestOption {
	return callOption{function: "WithInactiveDays", name: "inactiveDays", value: inactiveDays}
}

func WithIncludeMachine(includeMachine bool) RequestOption {
	return callOption{function: "WithIncludeMachine", name: "includeMachine", value: includeMachine}
}

func WithIncludeRevoked(includeRevoked bool) RequestOption {
	return callOption{function: "WithIncludeRevoked", name: "includeRevoked", value: includeRevoked}
}

func WithIncludeTeam(includeTeam bool) RequestOption {
	return callOption{function: "WithIncludeTeam", name: "includeTeam", value: includeTeam}
}

func WithInterval(interval string) RequestOption {
	return callOption{function: "WithInterval", name: "interval", value: interval}
}

func WithItemID(itemID string) RequestOption {
	return callOption{function: "WithItemID", name: "itemId", value: itemID}
}

func WithKeyIDs(keyIDs ...string) RequestOption {
	return callOption{function: "WithKeyIDs", name: "keyIds", value: keyIDs}
}

func WithKind(kind string) RequestOption {
	return callOption{function: "WithKind", name: "kind", value: kind}
}

func WithLabelIDs(labelIDs ...string) RequestOption {
	return callOption{function: "WithLabelIDs", name: "labelIds", value: labelIDs}
}

func WithLevel(level string) RequestOption {
	return callOption{function: "WithLevel", name: "level", value: level}
}

func WithLimit(limit int) RequestOption {
	return callOption{function: "WithLimit", name: "limit", value: limit}
}

func WithMessageID(messageID string) RequestOption {
	return callOption{function: "WithMessageID", name: "messageId", value: messageID}
}

func WithMinutes(minutes int) RequestOption {
	return callOption{function: "WithMinutes", name: "minutes", value: minutes}
}

func WithName(name string) RequestOption {
	return callOption{function: "WithName", name: "name", value: name}
}

func WithOccurrenceStart(occurrenceStart time.Time) RequestOption {
	return callOption{function: "WithOccurrenceStart", name: "occurrenceStart", value: occurrenceStart}
}

func WithOffset(offset int) RequestOption {
	return callOption{function: "WithOffset", name: "offset", value: offset}
}

func WithOffsetMinutes(offsetMinutes int) RequestOption {
	return callOption{function: "WithOffsetMinutes", name: "offsetMinutes", value: offsetMinutes}
}

func WithOpened(opened bool) RequestOption {
	return callOption{function: "WithOpened", name: "opened", value: opened}
}

func WithOrder(order string) RequestOption {
	return callOption{function: "WithOrder", name: "order", value: order}
}

func WithOutcome(outcome string) RequestOption {
	return callOption{function: "WithOutcome", name: "outcome", value: outcome}
}

func WithPage(page int) RequestOption {
	return callOption{function: "WithPage", name: "page", value: page}
}

func WithPageID(pageID string) RequestOption {
	return callOption{function: "WithPageID", name: "pageId", value: pageID}
}

func WithPageIDs(pageIDs ...string) RequestOption {
	return callOption{function: "WithPageIDs", name: "pageIds", value: pageIDs}
}

func WithPageSize(pageSize int) RequestOption {
	return callOption{function: "WithPageSize", name: "pageSize", value: pageSize}
}

func WithPath(path string) RequestOption {
	return callOption{function: "WithPath", name: "path", value: path}
}

func WithPeople(people int) RequestOption {
	return callOption{function: "WithPeople", name: "people", value: people}
}

func WithPinned(pinned bool) RequestOption {
	return callOption{function: "WithPinned", name: "pinned", value: pinned}
}

func WithProvider(provider string) RequestOption {
	return callOption{function: "WithProvider", name: "provider", value: provider}
}

func WithQ(q string) RequestOption {
	return callOption{function: "WithQ", name: "q", value: q}
}

func WithQuery(query string) RequestOption {
	return callOption{function: "WithQuery", name: "query", value: query}
}

func WithReason(reason string) RequestOption {
	return callOption{function: "WithReason", name: "reason", value: reason}
}

func WithReassignTo(reassignTo string) RequestOption {
	return callOption{function: "WithReassignTo", name: "reassignTo", value: reassignTo}
}

func WithRecipient(recipient string) RequestOption {
	return callOption{function: "WithRecipient", name: "recipient", value: recipient}
}

func WithRefresh(refresh bool) RequestOption {
	return callOption{function: "WithRefresh", name: "refresh", value: refresh}
}

func WithReuse(reuse bool) RequestOption {
	return callOption{function: "WithReuse", name: "reuse", value: reuse}
}

func WithRuleID(ruleID string) RequestOption {
	return callOption{function: "WithRuleID", name: "ruleId", value: ruleID}
}

func WithRuleIDs(ruleIDs ...string) RequestOption {
	return callOption{function: "WithRuleIDs", name: "ruleIds", value: ruleIDs}
}

func WithRunID(runID string) RequestOption {
	return callOption{function: "WithRunID", name: "runId", value: runID}
}

func WithScheduledFrom(scheduledFrom time.Time) RequestOption {
	return callOption{function: "WithScheduledFrom", name: "scheduledFrom", value: scheduledFrom}
}

func WithScheduledTo(scheduledTo time.Time) RequestOption {
	return callOption{function: "WithScheduledTo", name: "scheduledTo", value: scheduledTo}
}

func WithScope(scope string) RequestOption {
	return callOption{function: "WithScope", name: "scope", value: scope}
}

func WithSearch(search string) RequestOption {
	return callOption{function: "WithSearch", name: "search", value: search}
}

func WithSemantic(semantic bool) RequestOption {
	return callOption{function: "WithSemantic", name: "semantic", value: semantic}
}

func WithSends(sends int) RequestOption {
	return callOption{function: "WithSends", name: "sends", value: sends}
}

func WithSince(since time.Time) RequestOption {
	return callOption{function: "WithSince", name: "since", value: since}
}

func WithSort(sort string) RequestOption {
	return callOption{function: "WithSort", name: "sort", value: sort}
}

func WithSource(source string) RequestOption {
	return callOption{function: "WithSource", name: "source", value: source}
}

func WithStartIndex(startIndex int) RequestOption {
	return callOption{function: "WithStartIndex", name: "startIndex", value: startIndex}
}

func WithStatus(status ...string) RequestOption {
	return callOption{function: "WithStatus", name: "status", value: status}
}

func WithStatuses(statuses ...string) RequestOption {
	return callOption{function: "WithStatuses", name: "statuses", value: statuses}
}

func WithStatusCodes(statuses ...int) RequestOption {
	return callOption{function: "WithStatusCodes", name: "statuses", value: statuses}
}

func WithStepKey(stepKey string) RequestOption {
	return callOption{function: "WithStepKey", name: "stepKey", value: stepKey}
}

func WithSuggested(suggested bool) RequestOption {
	return callOption{function: "WithSuggested", name: "suggested", value: suggested}
}

func WithTag(tag string) RequestOption {
	return callOption{function: "WithTag", name: "tag", value: tag}
}

func WithTaskID(taskID string) RequestOption {
	return callOption{function: "WithTaskID", name: "taskId", value: taskID}
}

func WithThreadID(threadID string) RequestOption {
	return callOption{function: "WithThreadID", name: "threadId", value: threadID}
}

func WithTimezone(timezone string) RequestOption {
	return callOption{function: "WithTimezone", name: "timezone", value: timezone}
}

func WithTitle(title string) RequestOption {
	return callOption{function: "WithTitle", name: "title", value: title}
}

func WithTo(to time.Time) RequestOption {
	return callOption{function: "WithTo", name: "to", value: to}
}

func WithTracked(tracked bool) RequestOption {
	return callOption{function: "WithTracked", name: "tracked", value: tracked}
}

func WithUntil(until time.Time) RequestOption {
	return callOption{function: "WithUntil", name: "until", value: until}
}

func WithVersion(version int) RequestOption {
	return callOption{function: "WithVersion", name: "version", value: version}
}

func WithView(view string) RequestOption {
	return callOption{function: "WithView", name: "view", value: view}
}

func WithWithoutEmail(withoutEmail bool) RequestOption {
	return callOption{function: "WithWithoutEmail", name: "withoutEmail", value: withoutEmail}
}

var methodOptions = map[string][]string{
	"me.get":                              {"WithAPIKey"},
	"me.ping":                             {"WithAPIKey"},
	"me.rotate":                           {"WithAPIKey"},
	"me.listRequests":                     {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.listAllRequests":                  {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.iterateRequests":                  {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.listActivity":                     {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.listAllActivity":                  {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.iterateActivity":                  {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"me.stats":                            {"WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"security.stepUpStatus":               {"WithAPIKey"},
	"security.beginStepUp":                {"WithAPIKey"},
	"security.verifyStepUp":               {"WithAPIKey"},
	"security.listEvents":                 {"WithLimit", "WithCursor", "WithAPIKey"},
	"security.listAllEvents":              {"WithLimit", "WithCursor", "WithAPIKey"},
	"security.iterateEvents":              {"WithLimit", "WithCursor", "WithAPIKey"},
	"security.getPolicy":                  {"WithAPIKey"},
	"security.updatePolicy":               {"WithAPIKey"},
	"keys.list":                           {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"keys.listAll":                        {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"keys.iterate":                        {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"keys.get":                            {"WithAPIKey"},
	"keys.create":                         {"WithAPIKey"},
	"keys.update":                         {"WithAPIKey"},
	"keys.delete":                         {"WithAPIKey"},
	"keys.rotate":                         {"WithAPIKey"},
	"keys.revoke":                         {"WithAPIKey"},
	"keys.listRequests":                   {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listAllRequests":                {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.iterateRequests":                {"WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listActivity":                   {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listAllActivity":                {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.iterateActivity":                {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listWorkspaceRequests":          {"WithKeyIDs", "WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listAllWorkspaceRequests":       {"WithKeyIDs", "WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.iterateWorkspaceRequests":       {"WithKeyIDs", "WithFailedOnly", "WithStatusCodes", "WithPath", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listWorkspaceActivity":          {"WithKeyIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.listAllWorkspaceActivity":       {"WithKeyIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.iterateWorkspaceActivity":       {"WithKeyIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"keys.stats":                          {"WithKeyIDs", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"addresses.list":                      {"WithLimit", "WithCursor", "WithAPIKey"},
	"addresses.listAll":                   {"WithLimit", "WithCursor", "WithAPIKey"},
	"addresses.iterate":                   {"WithLimit", "WithCursor", "WithAPIKey"},
	"languages.list":                      {"WithAPIKey"},
	"emails.send":                         {"WithIdempotencyKey", "WithAPIKey"},
	"emails.sendBatch":                    {"WithIdempotencyKey", "WithAPIKey"},
	"emails.translate":                    {"WithAPIKey"},
	"emails.check":                        {"WithAPIKey"},
	"emails.list":                         {"WithStatus", "WithFromAddress", "WithBroadcastID", "WithScheduledFrom", "WithScheduledTo", "WithSince", "WithUntil", "WithRecipient", "WithTag", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.listAll":                      {"WithStatus", "WithFromAddress", "WithBroadcastID", "WithScheduledFrom", "WithScheduledTo", "WithSince", "WithUntil", "WithRecipient", "WithTag", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.iterate":                      {"WithStatus", "WithFromAddress", "WithBroadcastID", "WithScheduledFrom", "WithScheduledTo", "WithSince", "WithUntil", "WithRecipient", "WithTag", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.get":                          {"WithAPIKey"},
	"emails.listEvents":                   {"WithAddress", "WithMessageID", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.listAllEvents":                {"WithAddress", "WithMessageID", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.iterateEvents":                {"WithAddress", "WithMessageID", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.listEventsByMessageId":        {"WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.listAllEventsByMessageId":     {"WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.iterateEventsByMessageId":     {"WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"emails.getTracking":                  {"WithAddress", "WithAPIKey"},
	"emails.cancel":                       {"WithAPIKey"},
	"emails.reschedule":                   {"WithAPIKey"},
	"emails.update":                       {"WithAPIKey"},
	"emails.compose":                      {"WithAPIKey"},
	"emails.rewrite":                      {"WithAPIKey"},
	"emails.suggestSubject":               {"WithAPIKey"},
	"templates.list":                      {"WithStatus", "WithSearch", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"templates.listAll":                   {"WithStatus", "WithSearch", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"templates.iterate":                   {"WithStatus", "WithSearch", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"templates.get":                       {"WithAPIKey"},
	"templates.create":                    {"WithAPIKey"},
	"templates.update":                    {"WithAPIKey"},
	"templates.duplicate":                 {"WithAPIKey"},
	"templates.replaceContent":            {"WithAPIKey"},
	"templates.delete":                    {"WithAPIKey"},
	"templates.listVersions":              {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.listAllVersions":           {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.iterateVersions":           {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.getVersion":                {"WithAPIKey"},
	"templates.publish":                   {"WithExpectedVersion", "WithExpectedUpdatedAt", "WithAPIKey"},
	"templates.restoreVersion":            {"WithAPIKey"},
	"templates.deleteVersion":             {"WithAPIKey"},
	"templates.listStarters":              {"WithAPIKey"},
	"templates.getStarter":                {"WithAPIKey"},
	"templates.listFonts":                 {"WithAPIKey"},
	"templates.render":                    {"WithAPIKey"},
	"templates.preview":                   {"WithAPIKey"},
	"templates.getAnalytics":              {"WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"templates.listSends":                 {"WithPage", "WithPageSize", "WithSearch", "WithSource", "WithVersion", "WithOpened", "WithClicked", "WithTracked", "WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"templates.send":                      {"WithIdempotencyKey", "WithAPIKey"},
	"templates.listImages":                {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.listAllImages":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.iterateImages":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"templates.uploadImage":               {"WithContentType", "WithCrop", "WithAPIKey"},
	"templates.design":                    {"WithAPIKey"},
	"templates.redesign":                  {"WithAPIKey"},
	"tracking.list":                       {"WithOpened", "WithClicked", "WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.listAll":                    {"WithOpened", "WithClicked", "WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.iterate":                    {"WithOpened", "WithClicked", "WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.getStats":                   {"WithDays", "WithMinutes", "WithGrain", "WithOffsetMinutes", "WithAddress", "WithAPIKey"},
	"tracking.get":                        {"WithAPIKey"},
	"tracking.listOpens":                  {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.listAllOpens":               {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.iterateOpens":               {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.listClicks":                 {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.listAllClicks":              {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"tracking.iterateClicks":              {"WithIncludeMachine", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.list":                        {"WithFolder", "WithQuery", "WithLabelIDs", "WithSort", "WithDateFrom", "WithDateTo", "WithFromContacts", "WithSemantic", "WithCategory", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.listAll":                     {"WithFolder", "WithQuery", "WithLabelIDs", "WithSort", "WithDateFrom", "WithDateTo", "WithFromContacts", "WithSemantic", "WithCategory", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.iterate":                     {"WithFolder", "WithQuery", "WithLabelIDs", "WithSort", "WithDateFrom", "WithDateTo", "WithFromContacts", "WithSemantic", "WithCategory", "WithAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.get":                         {"WithAPIKey"},
	"threads.update":                      {"WithAPIKey"},
	"threads.trash":                       {"WithAPIKey"},
	"threads.snooze":                      {"WithAPIKey"},
	"threads.unsnooze":                    {"WithAPIKey"},
	"threads.mute":                        {"WithAPIKey"},
	"threads.unmute":                      {"WithAPIKey"},
	"threads.listMessages":                {"WithAPIKey"},
	"threads.getMessage":                  {"WithAPIKey"},
	"threads.getRawMessage":               {"WithAPIKey"},
	"threads.listAttachments":             {"WithAPIKey"},
	"threads.listFileLinks":               {"WithIncludeRevoked", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.listAllFileLinks":            {"WithIncludeRevoked", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.iterateFileLinks":            {"WithIncludeRevoked", "WithLimit", "WithCursor", "WithAPIKey"},
	"threads.listTracking":                {"WithAddress", "WithAPIKey"},
	"threads.getDelivery":                 {"WithAddress", "WithAPIKey"},
	"threads.listNotes":                   {"WithAPIKey"},
	"threads.createNote":                  {"WithAPIKey"},
	"threads.updateNote":                  {"WithAPIKey"},
	"threads.deleteNote":                  {"WithAPIKey"},
	"threads.reorderNotes":                {"WithAPIKey"},
	"threads.counts":                      {"WithAddress", "WithHideLabelled", "WithAPIKey"},
	"threads.summary":                     {"WithAPIKey"},
	"threads.replySuggestions":            {"WithAPIKey"},
	"threads.restore":                     {"WithAPIKey"},
	"threads.delete":                      {"WithAPIKey"},
	"threads.unsubscribe":                 {"WithAPIKey"},
	"threads.getEvent":                    {"WithAPIKey"},
	"threads.listSenderCategories":        {"WithAPIKey"},
	"threads.setSenderCategory":           {"WithAPIKey"},
	"threads.clearSenderCategory":         {"WithAPIKey"},
	"threads.getTeam":                     {"WithAPIKey"},
	"threads.assign":                      {"WithAPIKey"},
	"threads.unassign":                    {"WithAPIKey"},
	"threads.setReplying":                 {"WithAPIKey"},
	"threads.clearReplying":               {"WithAPIKey"},
	"threads.markDone":                    {"WithAPIKey"},
	"threads.reopen":                      {"WithAPIKey"},
	"threads.listComments":                {"WithAPIKey"},
	"threads.createComment":               {"WithAPIKey"},
	"threads.updateComment":               {"WithAPIKey"},
	"threads.deleteComment":               {"WithAPIKey"},
	"followUps.list":                      {"WithStatus", "WithThreadID", "WithLimit", "WithAPIKey"},
	"followUps.set":                       {"WithAPIKey"},
	"followUps.cancel":                    {"WithAPIKey"},
	"drafts.list":                         {"WithQuery", "WithLimit", "WithCursor", "WithAPIKey"},
	"drafts.listAll":                      {"WithQuery", "WithLimit", "WithCursor", "WithAPIKey"},
	"drafts.iterate":                      {"WithQuery", "WithLimit", "WithCursor", "WithAPIKey"},
	"drafts.get":                          {"WithAPIKey"},
	"drafts.create":                       {"WithAPIKey"},
	"drafts.update":                       {"WithAPIKey"},
	"drafts.delete":                       {"WithAPIKey"},
	"labels.list":                         {"WithLimit", "WithCursor", "WithAPIKey"},
	"labels.listAll":                      {"WithLimit", "WithCursor", "WithAPIKey"},
	"labels.iterate":                      {"WithLimit", "WithCursor", "WithAPIKey"},
	"labels.listColors":                   {"WithAPIKey"},
	"labels.get":                          {"WithAPIKey"},
	"labels.create":                       {"WithAPIKey"},
	"labels.update":                       {"WithAPIKey"},
	"labels.test":                         {"WithAPIKey"},
	"labels.delete":                       {"WithAPIKey"},
	"contacts.list":                       {"WithLimit", "WithCursor", "WithSource", "WithAudienceID", "WithQ", "WithAPIKey"},
	"contacts.listAll":                    {"WithSource", "WithAudienceID", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iterate":                    {"WithSource", "WithAudienceID", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.get":                        {"WithAPIKey"},
	"contacts.create":                     {"WithAPIKey"},
	"contacts.update":                     {"WithAPIKey"},
	"contacts.delete":                     {"WithAPIKey"},
	"contacts.setAudiences":               {"WithAPIKey"},
	"contacts.listPeople":                 {"WithQ", "WithEmail", "WithSort", "WithBlocked", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.listAllPeople":              {"WithQ", "WithEmail", "WithSort", "WithBlocked", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iteratePeople":              {"WithQ", "WithEmail", "WithSort", "WithBlocked", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.save":                       {"WithAPIKey"},
	"contacts.deleteMany":                 {"WithAPIKey"},
	"contacts.setPhoto":                   {"WithContentType", "WithCrop", "WithAPIKey"},
	"contacts.removePhoto":                {"WithAPIKey"},
	"contacts.block":                      {"WithAPIKey"},
	"contacts.unblock":                    {"WithAPIKey"},
	"contacts.listThreads":                {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.listAllThreads":             {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iterateThreads":             {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.activity":                   {"WithMinutes", "WithFrom", "WithTo", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"contacts.listEvents":                 {"WithName", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.listAllEvents":              {"WithName", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iterateEvents":              {"WithName", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.listCards":                  {"WithLimit", "WithCursor", "WithEmail", "WithWithoutEmail", "WithKind", "WithGroupID", "WithAPIKey"},
	"contacts.listAllCards":               {"WithEmail", "WithWithoutEmail", "WithKind", "WithGroupID", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iterateCards":               {"WithEmail", "WithWithoutEmail", "WithKind", "WithGroupID", "WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.getCard":                    {"WithAPIKey"},
	"contacts.createCard":                 {"WithAPIKey"},
	"contacts.updateCard":                 {"WithAPIKey"},
	"contacts.deleteCard":                 {"WithAPIKey"},
	"contacts.findDuplicateCards":         {"WithAPIKey"},
	"contacts.mergeCards":                 {"WithAPIKey"},
	"contacts.listBirthdays":              {"WithTimezone", "WithAPIKey"},
	"contacts.importVcf":                  {"WithFilename", "WithAPIKey"},
	"contacts.listImports":                {"WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.listAllImports":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.iterateImports":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"contacts.getImport":                  {"WithAPIKey"},
	"contacts.undoImport":                 {"WithAPIKey"},
	"audiences.list":                      {"WithLimit", "WithCursor", "WithAPIKey"},
	"audiences.listAll":                   {"WithLimit", "WithCursor", "WithAPIKey"},
	"audiences.iterate":                   {"WithLimit", "WithCursor", "WithAPIKey"},
	"audiences.growth":                    {"WithAudienceIDs", "WithDays", "WithMinutes", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"audiences.get":                       {"WithAPIKey"},
	"audiences.create":                    {"WithAPIKey"},
	"audiences.update":                    {"WithAPIKey"},
	"audiences.delete":                    {"WithAPIKey"},
	"audiences.empty":                     {"WithAPIKey"},
	"audiences.listContacts":              {"WithLimit", "WithCursor", "WithQ", "WithSource", "WithSort", "WithStatuses", "WithInactiveDays", "WithAPIKey"},
	"audiences.listAllContacts":           {"WithLimit", "WithCursor", "WithQ", "WithSource", "WithSort", "WithStatuses", "WithInactiveDays", "WithAPIKey"},
	"audiences.iterateContacts":           {"WithLimit", "WithCursor", "WithQ", "WithSource", "WithSort", "WithStatuses", "WithInactiveDays", "WithAPIKey"},
	"audiences.addContact":                {"WithAPIKey"},
	"audiences.removeContact":             {"WithAPIKey"},
	"audiences.addContacts":               {"WithAPIKey"},
	"audiences.removeContacts":            {"WithAPIKey"},
	"audiences.importContacts":            {"WithAPIKey"},
	"segments.list":                       {"WithAPIKey"},
	"segments.get":                        {"WithAPIKey"},
	"segments.create":                     {"WithAPIKey"},
	"segments.update":                     {"WithAPIKey"},
	"segments.delete":                     {"WithAPIKey"},
	"segments.preview":                    {"WithAPIKey"},
	"forms.list":                          {"WithLimit", "WithCursor", "WithAPIKey"},
	"forms.listAll":                       {"WithLimit", "WithCursor", "WithAPIKey"},
	"forms.iterate":                       {"WithLimit", "WithCursor", "WithAPIKey"},
	"forms.create":                        {"WithAPIKey"},
	"forms.design":                        {"WithAPIKey"},
	"forms.redesign":                      {"WithAPIKey"},
	"forms.listStarters":                  {"WithAPIKey"},
	"forms.getStarter":                    {"WithAPIKey"},
	"forms.get":                           {"WithAPIKey"},
	"forms.update":                        {"WithAPIKey"},
	"forms.delete":                        {"WithAPIKey"},
	"forms.publish":                       {"WithAPIKey"},
	"forms.pause":                         {"WithAPIKey"},
	"forms.resume":                        {"WithAPIKey"},
	"forms.duplicate":                     {"WithAPIKey"},
	"forms.analytics":                     {"WithDays", "WithMinutes", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"forms.listSubmissions":               {"WithLimit", "WithCursor", "WithQ", "WithStatus", "WithAPIKey"},
	"forms.listAllSubmissions":            {"WithLimit", "WithCursor", "WithQ", "WithStatus", "WithAPIKey"},
	"forms.iterateSubmissions":            {"WithLimit", "WithCursor", "WithQ", "WithStatus", "WithAPIKey"},
	"forms.getSubmission":                 {"WithAPIKey"},
	"forms.deleteSubmission":              {"WithAPIKey"},
	"forms.deleteSubmissions":             {"WithAPIKey"},
	"forms.approveSubmission":             {"WithAPIKey"},
	"forms.resendConfirmation":            {"WithAPIKey"},
	"forms.subscribe":                     {"WithAPIKey"},
	"automations.list":                    {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.listAll":                 {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.iterate":                 {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.create":                  {"WithAPIKey"},
	"automations.listStarters":            {"WithAPIKey"},
	"automations.get":                     {"WithAPIKey"},
	"automations.update":                  {"WithAPIKey"},
	"automations.delete":                  {"WithAPIKey"},
	"automations.publish":                 {"WithAPIKey"},
	"automations.pause":                   {"WithAPIKey"},
	"automations.resume":                  {"WithAPIKey"},
	"automations.archive":                 {"WithAPIKey"},
	"automations.duplicate":               {"WithAPIKey"},
	"automations.sendTest":                {"WithAPIKey"},
	"automations.listVersions":            {"WithAPIKey"},
	"automations.restoreVersion":          {"WithAPIKey"},
	"automations.stats":                   {"WithSince", "WithUntil", "WithAPIKey"},
	"automations.analytics":               {"WithAutomationIDs", "WithSince", "WithUntil", "WithAPIKey"},
	"automations.listRuns":                {"WithEvent", "WithStepKey", "WithOutcome", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.listAllRuns":             {"WithEvent", "WithStepKey", "WithOutcome", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.iterateRuns":             {"WithEvent", "WithStepKey", "WithOutcome", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.listEnrollments":         {"WithStatus", "WithStepKey", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.listAllEnrollments":      {"WithStatus", "WithStepKey", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.iterateEnrollments":      {"WithStatus", "WithStepKey", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"automations.getEnrollment":           {"WithAPIKey"},
	"automations.enroll":                  {"WithAPIKey"},
	"automations.exitEnrollment":          {"WithAPIKey"},
	"events.send":                         {"WithIdempotencyKey", "WithAPIKey"},
	"events.sendBatch":                    {"WithIdempotencyKey", "WithAPIKey"},
	"events.listNames":                    {"WithAPIKey"},
	"broadcasts.preview":                  {"WithAPIKey"},
	"broadcasts.send":                     {"WithIdempotencyKey", "WithAPIKey"},
	"broadcasts.list":                     {"WithAudienceID", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.listAll":                  {"WithAudienceID", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.iterate":                  {"WithAudienceID", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.get":                      {"WithAPIKey"},
	"broadcasts.stats":                    {"WithGrain", "WithOffsetMinutes", "WithDays", "WithMinutes", "WithSince", "WithAPIKey"},
	"broadcasts.analytics":                {"WithBroadcastIDs", "WithDays", "WithMinutes", "WithSince", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"broadcasts.listRecipients":           {"WithFilter", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.listAllRecipients":        {"WithFilter", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.iterateRecipients":        {"WithFilter", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"broadcasts.getRecipient":             {"WithAPIKey"},
	"broadcasts.cancel":                   {"WithAPIKey"},
	"broadcasts.createDraft":              {"WithAPIKey"},
	"broadcasts.update":                   {"WithAPIKey"},
	"broadcasts.sendDraft":                {"WithAPIKey"},
	"broadcasts.duplicate":                {"WithAPIKey"},
	"broadcasts.resend":                   {"WithAPIKey"},
	"broadcasts.sendTest":                 {"WithAPIKey"},
	"domains.list":                        {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"domains.listAll":                     {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"domains.iterate":                     {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"domains.get":                         {"WithAPIKey"},
	"domains.create":                      {"WithAPIKey"},
	"domains.verify":                      {"WithAPIKey"},
	"domains.update":                      {"WithAPIKey"},
	"domains.delete":                      {"WithAPIKey"},
	"domains.getLogo":                     {"WithAPIKey"},
	"domains.setLogo":                     {"WithAPIKey"},
	"domains.removeLogo":                  {"WithAPIKey"},
	"domains.setLogoCertificate":          {"WithAPIKey"},
	"domains.removeLogoCertificate":       {"WithAPIKey"},
	"domains.listAddresses":               {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.listAllAddresses":            {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.iterateAddresses":            {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.createAddress":               {"WithAPIKey"},
	"domains.getAddress":                  {"WithAPIKey"},
	"domains.updateAddress":               {"WithAPIKey"},
	"domains.setAddressPhoto":             {"WithContentType", "WithCrop", "WithAPIKey"},
	"domains.removeAddressPhoto":          {"WithAPIKey"},
	"domains.startGravatarPublish":        {"WithAPIKey"},
	"domains.deleteAddress":               {"WithAPIKey"},
	"domains.listAddressForwards":         {"WithAPIKey"},
	"domains.addAddressForwards":          {"WithAPIKey"},
	"domains.updateAddressForward":        {"WithAPIKey"},
	"domains.deleteAddressForward":        {"WithAPIKey"},
	"domains.resendAddressForwardConsent": {"WithAPIKey"},
	"domains.listAddressMembers":          {"WithAPIKey"},
	"domains.getAddressLogin":             {"WithAPIKey"},
	"domains.setAddressLogin":             {"WithAPIKey"},
	"domains.deleteAddressLogin":          {"WithAPIKey"},
	"domains.getMailAppSettings":          {"WithAPIKey"},
	"domains.updateMailAppSettings":       {"WithAPIKey"},
	"domains.getAddressLimits":            {"WithAPIKey"},
	"domains.updateAddressLimits":         {"WithAPIKey"},
	"domains.listAppPasswords":            {"WithAPIKey"},
	"domains.createAppPassword":           {"WithAPIKey"},
	"domains.updateAppPassword":           {"WithAPIKey"},
	"domains.deleteAppPassword":           {"WithAPIKey"},
	"domains.sendMailAppSetup":            {"WithAPIKey"},
	"domains.createMailAppProfileLink":    {"WithAPIKey"},
	"domains.getDns":                      {"WithRefresh", "WithAPIKey"},
	"domains.setDnsZone":                  {"WithAPIKey"},
	"domains.syncDns":                     {"WithAPIKey"},
	"domains.listRoutes":                  {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.listAllRoutes":               {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.iterateRoutes":               {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.createRoute":                 {"WithAPIKey"},
	"domains.getRoute":                    {"WithAPIKey"},
	"domains.updateRoute":                 {"WithAPIKey"},
	"domains.deleteRoute":                 {"WithAPIKey"},
	"domains.testRoute":                   {"WithAPIKey"},
	"domains.listGroupMembers":            {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.listAllGroupMembers":         {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.iterateGroupMembers":         {"WithLimit", "WithCursor", "WithAPIKey"},
	"domains.addGroupMembers":             {"WithAPIKey"},
	"domains.removeGroupMember":           {"WithAPIKey"},
	"appHost.get":                         {"WithAPIKey"},
	"appHost.set":                         {"WithAPIKey"},
	"appHost.verify":                      {"WithAPIKey"},
	"appHost.delete":                      {"WithAPIKey"},
	"branding.get":                        {"WithAPIKey"},
	"branding.update":                     {"WithAPIKey"},
	"branding.uploadImage":                {"WithContentType", "WithCrop", "WithAPIKey"},
	"branding.removeImage":                {"WithAPIKey"},
	"rules.list":                          {"WithEnabled", "WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"rules.listAll":                       {"WithEnabled", "WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"rules.iterate":                       {"WithEnabled", "WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"rules.get":                           {"WithAPIKey"},
	"rules.create":                        {"WithAPIKey"},
	"rules.update":                        {"WithAPIKey"},
	"rules.delete":                        {"WithAPIKey"},
	"rules.reorder":                       {"WithAPIKey"},
	"rules.test":                          {"WithAPIKey"},
	"rules.listRuns":                      {"WithRuleID", "WithRuleIDs", "WithThreadID", "WithLimit", "WithCursor", "WithAPIKey"},
	"rules.listAllRuns":                   {"WithRuleID", "WithRuleIDs", "WithThreadID", "WithLimit", "WithCursor", "WithAPIKey"},
	"rules.iterateRuns":                   {"WithRuleID", "WithRuleIDs", "WithThreadID", "WithLimit", "WithCursor", "WithAPIKey"},
	"rules.stats":                         {"WithRuleIDs", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"webhooks.list":                       {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"webhooks.listAll":                    {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"webhooks.iterate":                    {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"webhooks.get":                        {"WithAPIKey"},
	"webhooks.create":                     {"WithAPIKey"},
	"webhooks.update":                     {"WithAPIKey"},
	"webhooks.delete":                     {"WithAPIKey"},
	"webhooks.rotateSecret":               {"WithAPIKey"},
	"webhooks.test":                       {"WithAPIKey"},
	"webhooks.listDeliveries":             {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listAllDeliveries":          {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.iterateDeliveries":          {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.getDelivery":                {"WithAPIKey"},
	"webhooks.replayDelivery":             {"WithAPIKey"},
	"webhooks.listDeadLetters":            {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listAllDeadLetters":         {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.iterateDeadLetters":         {"WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.replayDeadLetters":          {"WithAPIKey"},
	"webhooks.discardDeadLetter":          {"WithAPIKey"},
	"webhooks.listWorkspaceDeliveries":    {"WithEndpointIDs", "WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listAllWorkspaceDeliveries": {"WithEndpointIDs", "WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.iterateWorkspaceDeliveries": {"WithEndpointIDs", "WithStatus", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listActivity":               {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listAllActivity":            {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.iterateActivity":            {"WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listWorkspaceActivity":      {"WithEndpointIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listAllWorkspaceActivity":   {"WithEndpointIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.iterateWorkspaceActivity":   {"WithEndpointIDs", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"webhooks.listEvents":                 {"WithAPIKey"},
	"webhooks.stats":                      {"WithEndpointIDs", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"imports.list":                        {"WithAddressID", "WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"imports.listAll":                     {"WithAddressID", "WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"imports.iterate":                     {"WithAddressID", "WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"imports.get":                         {"WithAPIKey"},
	"imports.create":                      {"WithAPIKey"},
	"imports.uploadState":                 {"WithAPIKey"},
	"imports.uploadChunk":                 {"WithAPIKey"},
	"imports.start":                       {"WithAPIKey"},
	"imports.cancel":                      {"WithAPIKey"},
	"imports.listFailures":                {"WithAfter", "WithLimit", "WithAPIKey"},
	"imports.deleteUpload":                {"WithAPIKey"},
	"imports.importFiles":                 {"WithAPIKey"},
	"imports.discoverMailbox":             {"WithAPIKey"},
	"imports.checkMailbox":                {"WithAPIKey"},
	"imports.connectMailbox":              {"WithAPIKey"},
	"imports.resume":                      {"WithAPIKey"},
	"providerImports.inspect":             {"WithProvider", "WithAPIKey"},
	"providerImports.create":              {"WithAPIKey"},
	"providerImports.list":                {"WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"providerImports.listAll":             {"WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"providerImports.iterate":             {"WithSource", "WithStatus", "WithSearch", "WithLimit", "WithCursor", "WithAPIKey"},
	"providerImports.get":                 {"WithAPIKey"},
	"providerImports.cancel":              {"WithAPIKey"},
	"calendar.listEvents":                 {"WithTimezone", "WithIncludeTeam", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.listAllEvents":              {"WithTimezone", "WithIncludeTeam", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.iterateEvents":              {"WithTimezone", "WithIncludeTeam", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.getEvent":                   {"WithAPIKey"},
	"calendar.getEventIcs":                {"WithAPIKey"},
	"calendar.createEvent":                {"WithAPIKey"},
	"calendar.updateEvent":                {"WithAPIKey"},
	"calendar.deleteEvent":                {"WithScope", "WithOccurrenceStart", "WithAPIKey"},
	"calendar.cancelEvent":                {"WithAPIKey"},
	"calendar.respondToEvent":             {"WithAPIKey"},
	"calendar.declineProposal":            {"WithAPIKey"},
	"calendar.hideEvent":                  {"WithAPIKey"},
	"calendar.unhideEvent":                {"WithAPIKey"},
	"calendar.freeBusy":                   {"WithTimezone", "WithDurationMinutes", "WithAPIKey"},
	"calendar.listFeeds":                  {"WithAPIKey"},
	"calendar.createFeed":                 {"WithAPIKey"},
	"calendar.resetFeed":                  {"WithAPIKey"},
	"calendar.deleteFeed":                 {"WithAPIKey"},
	"calendar.listSubscriptions":          {"WithAPIKey"},
	"calendar.createSubscription":         {"WithAPIKey"},
	"calendar.updateSubscription":         {"WithAPIKey"},
	"calendar.refreshSubscription":        {"WithAPIKey"},
	"calendar.deleteSubscription":         {"WithAPIKey"},
	"calendar.importIcs":                  {"WithFilename", "WithAPIKey"},
	"calendar.listImports":                {"WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.listAllImports":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.iterateImports":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"calendar.getImport":                  {"WithAPIKey"},
	"calendar.undoImport":                 {"WithAPIKey"},
	"settings.get":                        {"WithAddress", "WithAPIKey"},
	"settings.update":                     {"WithAddress", "WithAPIKey"},
	"outOfOffice.list":                    {"WithAddress", "WithAPIKey"},
	"outOfOffice.set":                     {"WithAPIKey"},
	"outOfOffice.clear":                   {"WithAPIKey"},
	"roles.list":                          {"WithLimit", "WithCursor", "WithAPIKey"},
	"roles.listAll":                       {"WithLimit", "WithCursor", "WithAPIKey"},
	"roles.iterate":                       {"WithLimit", "WithCursor", "WithAPIKey"},
	"roles.get":                           {"WithAPIKey"},
	"roles.create":                        {"WithAPIKey"},
	"roles.update":                        {"WithAPIKey"},
	"roles.delete":                        {"WithReassignTo", "WithAPIKey"},
	"roles.listPermissions":               {"WithAPIKey"},
	"members.list":                        {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"members.listAll":                     {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"members.iterate":                     {"WithLimit", "WithCursor", "WithQ", "WithAPIKey"},
	"members.get":                         {"WithAPIKey"},
	"members.add":                         {"WithAPIKey"},
	"members.update":                      {"WithAPIKey"},
	"members.remove":                      {"WithAPIKey"},
	"members.grantAddress":                {"WithAPIKey"},
	"members.revokeAddress":               {"WithAPIKey"},
	"members.grantDomain":                 {"WithAPIKey"},
	"members.revokeDomain":                {"WithAPIKey"},
	"members.listInvitations":             {"WithLimit", "WithCursor", "WithAPIKey"},
	"members.listAllInvitations":          {"WithLimit", "WithCursor", "WithAPIKey"},
	"members.iterateInvitations":          {"WithLimit", "WithCursor", "WithAPIKey"},
	"members.revokeInvitation":            {"WithAPIKey"},
	"members.resendInvitation":            {"WithAPIKey"},
	"members.getDirectorySync":            {"WithAPIKey"},
	"members.connectDirectorySync":        {"WithAPIKey"},
	"members.disconnectDirectorySync":     {"WithAPIKey"},
	"members.listDirectoryPeople":         {"WithStartIndex", "WithCount", "WithAPIKey"},
	"members.setDirectoryDefaultRole":     {"WithAPIKey"},
	"members.setDirectoryGroupRole":       {"WithAPIKey"},
	"suppressions.list":                   {"WithQ", "WithReason", "WithLimit", "WithCursor", "WithAPIKey"},
	"suppressions.listAll":                {"WithQ", "WithReason", "WithLimit", "WithCursor", "WithAPIKey"},
	"suppressions.iterate":                {"WithQ", "WithReason", "WithLimit", "WithCursor", "WithAPIKey"},
	"suppressions.get":                    {"WithAPIKey"},
	"suppressions.add":                    {"WithAPIKey"},
	"suppressions.remove":                 {"WithAPIKey"},
	"suppressions.listUnsubscribes":       {"WithQ", "WithSource", "WithLimit", "WithCursor", "WithAPIKey"},
	"suppressions.addUnsubscribe":         {"WithAPIKey"},
	"suppressions.removeUnsubscribe":      {"WithAPIKey"},
	"sending.listStreams":                 {"WithAPIKey"},
	"sending.resumeStream":                {"WithAPIKey"},
	"sending.listDedicatedIps":            {"WithAPIKey"},
	"sending.getDedicatedIp":              {"WithAPIKey"},
	"files.list":                          {"WithQ", "WithKind", "WithDirection", "WithAddress", "WithSince", "WithUntil", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"files.listAll":                       {"WithQ", "WithKind", "WithDirection", "WithAddress", "WithSince", "WithUntil", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"files.iterate":                       {"WithQ", "WithKind", "WithDirection", "WithAddress", "WithSince", "WithUntil", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"files.get":                           {"WithAPIKey"},
	"files.stats":                         {"WithAPIKey"},
	"files.download":                      {"WithAPIKey"},
	"files.listLinks":                     {"WithLimit", "WithCursor", "WithAPIKey"},
	"files.listAllLinks":                  {"WithLimit", "WithCursor", "WithAPIKey"},
	"files.iterateLinks":                  {"WithLimit", "WithCursor", "WithAPIKey"},
	"files.createLink":                    {"WithDomain", "WithReuse", "WithAPIKey"},
	"files.revokeLink":                    {"WithAPIKey"},
	"files.revokeAllLinks":                {"WithAPIKey"},
	"files.upload":                        {"WithContentType", "WithTimeout", "WithAPIKey"},
	"files.delete":                        {"WithAPIKey"},
	"files.deleteMany":                    {"WithAPIKey"},
	"knowledge.list":                      {"WithScope", "WithLevel", "WithKind", "WithStatus", "WithPinned", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.listAll":                   {"WithScope", "WithLevel", "WithKind", "WithStatus", "WithPinned", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.iterate":                   {"WithScope", "WithLevel", "WithKind", "WithStatus", "WithPinned", "WithQ", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.levels":                    {"WithAPIKey"},
	"knowledge.usage":                     {"WithAPIKey"},
	"knowledge.search":                    {"WithAPIKey"},
	"knowledge.createNote":                {"WithAPIKey"},
	"knowledge.addLink":                   {"WithAPIKey"},
	"knowledge.uploadFile":                {"WithContentType", "WithScope", "WithTitle", "WithTimeout", "WithAPIKey"},
	"knowledge.get":                       {"WithAPIKey"},
	"knowledge.update":                    {"WithAPIKey"},
	"knowledge.delete":                    {"WithAPIKey"},
	"knowledge.refresh":                   {"WithAPIKey"},
	"knowledge.stats":                     {"WithDays", "WithAPIKey"},
	"knowledge.listSuggestions":           {"WithKind", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.listAllSuggestions":        {"WithKind", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.iterateSuggestions":        {"WithKind", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"knowledge.acceptSuggestion":          {"WithAPIKey"},
	"knowledge.dismissSuggestion":         {"WithAPIKey"},
	"knowledge.listFlags":                 {"WithItemID", "WithLimit", "WithAPIKey"},
	"knowledge.dismissFlag":               {"WithAPIKey"},
	"knowledge.listConnectors":            {"WithAPIKey"},
	"knowledge.addConnector":              {"WithAPIKey"},
	"knowledge.getConnector":              {"WithAPIKey"},
	"knowledge.updateConnector":           {"WithAPIKey"},
	"knowledge.deleteConnector":           {"WithAPIKey"},
	"knowledge.syncConnector":             {"WithAPIKey"},
	"knowledge.draftFromThread":           {"WithAPIKey"},
	"tempMail.listDomains":                {"WithInboxToken"},
	"tempMail.create":                     {"WithInboxToken"},
	"tempMail.get":                        {"WithInboxToken"},
	"tempMail.extend":                     {"WithInboxToken"},
	"tempMail.delete":                     {"WithInboxToken"},
	"tempMail.listMessages":               {"WithLimit", "WithCursor", "WithInboxToken"},
	"tempMail.listAllMessages":            {"WithLimit", "WithCursor", "WithInboxToken"},
	"tempMail.iterateMessages":            {"WithLimit", "WithCursor", "WithInboxToken"},
	"tempMail.getMessage":                 {"WithInboxToken"},
	"tempMail.deleteMessage":              {"WithInboxToken"},
	"tempMail.listAttachments":            {"WithInboxToken"},
	"tools.checkDmarc":                    {"WithAPIKey"},
	"tools.checkBimi":                     {"WithAPIKey"},
	"tools.checkDeliverability":           {"WithAPIKey"},
	"tools.readVerificationCode":          {"WithAPIKey"},
	"tools.estimateCost":                  {"WithSends", "WithDomains", "WithPeople", "WithAiActions", "WithInterval", "WithAPIKey"},
	"tools.checkAddress":                  {"WithDeep", "WithAPIKey"},
	"tools.checkAddresses":                {"WithAPIKey"},
	"support.send":                        {"WithAPIKey"},
	"encryption.listKeys":                 {"WithAPIKey"},
	"encryption.publishKey":               {"WithAPIKey"},
	"encryption.lookupKeys":               {"WithAPIKey"},
	"senders.get":                         {"WithAPIKey"},
	"senders.getAvatar":                   {"WithAPIKey"},
	"senders.research":                    {"WithAPIKey"},
	"chats.list":                          {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.listAll":                       {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.iterate":                       {"WithQ", "WithSort", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.get":                           {"WithAPIKey"},
	"chats.rename":                        {"WithAPIKey"},
	"chats.delete":                        {"WithAPIKey"},
	"chats.listMessages":                  {"WithOrder", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.listAllMessages":               {"WithOrder", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.iterateMessages":               {"WithOrder", "WithLimit", "WithCursor", "WithAPIKey"},
	"chats.sendMessage":                   {"WithTimeout", "WithAPIKey"},
	"chats.stopTurn":                      {"WithAPIKey"},
	"chats.approveToolCall":               {"WithTimeout", "WithAPIKey"},
	"chats.denyToolCall":                  {"WithTimeout", "WithAPIKey"},
	"analytics.mailbox":                   {"WithFrom", "WithTo", "WithOffsetMinutes", "WithGrain", "WithAddress", "WithAPIKey"},
	"analytics.sending":                   {"WithDays", "WithMinutes", "WithOffsetMinutes", "WithGrain", "WithAddress", "WithSince", "WithUntil", "WithRecipient", "WithTag", "WithAPIKey"},
	"subscriptions.list":                  {"WithStatus", "WithQ", "WithSort", "WithSuggested", "WithLimit", "WithOffset", "WithAddress", "WithAPIKey"},
	"subscriptions.listDomains":           {"WithStatus", "WithQ", "WithSort", "WithSuggested", "WithLimit", "WithOffset", "WithAddress", "WithAPIKey"},
	"subscriptions.keep":                  {"WithAPIKey"},
	"subscriptions.unsubscribe":           {"WithAPIKey"},
	"subscriptions.unsubscribeMany":       {"WithAPIKey"},
	"subscriptions.unsubscribeDomain":     {"WithAPIKey"},
	"subscriptions.move":                  {"WithAPIKey"},
	"subscriptions.moveMany":              {"WithAPIKey"},
	"exports.preview":                     {"WithAPIKey"},
	"exports.list":                        {"WithLimit", "WithOffset", "WithAPIKey"},
	"exports.get":                         {"WithAPIKey"},
	"exports.start":                       {"WithAPIKey"},
	"exports.download":                    {"WithAPIKey"},
	"dnsConnections.list":                 {"WithAPIKey"},
	"dnsConnections.start":                {"WithAPIKey"},
	"dnsConnections.get":                  {"WithAPIKey"},
	"dnsConnections.delete":               {"WithAPIKey"},
	"billing.get":                         {"WithAPIKey"},
	"billing.listPlans":                   {"WithAPIKey"},
	"billing.getUsage":                    {"WithDaysWindow", "WithDays", "WithAPIKey"},
	"billing.listAlerts":                  {"WithAPIKey"},
	"billing.listCountries":               {"WithAPIKey"},
	"billing.getPayAsYouGo":               {"WithAPIKey"},
	"billing.setPayAsYouGo":               {"WithAPIKey"},
	"billing.setPayAsYouGoLimit":          {"WithAPIKey"},
	"billing.confirmPayAsYouGo":           {"WithAPIKey"},
	"billing.startCheckout":               {"WithAPIKey"},
	"billing.openPortal":                  {"WithAPIKey"},
	"billing.listInvoices":                {"WithPage", "WithLimit", "WithSort", "WithSearch", "WithAPIKey"},
	"billing.getInvoice":                  {"WithAPIKey"},
	"billing.saveInvoiceDetails":          {"WithAPIKey"},
	"billing.listAddOns":                  {"WithAPIKey"},
	"billing.startAddOnCheckout":          {"WithAPIKey"},
	"billing.cancelAddOn":                 {"WithAPIKey"},
	"workspaces.list":                     {"WithAPIKey"},
	"workspaces.create":                   {"WithAPIKey"},
	"workspaces.getActive":                {"WithAPIKey"},
	"workspaces.setActive":                {"WithAPIKey"},
	"workspaces.delete":                   {"WithAPIKey"},
	"workspaces.getOwnershipOffer":        {"WithAPIKey"},
	"workspaces.offerOwnership":           {"WithAPIKey"},
	"workspaces.withdrawOwnershipOffer":   {"WithAPIKey"},
	"account.setName":                     {"WithAPIKey"},
	"account.getNotifications":            {"WithAPIKey"},
	"account.setEmailNotification":        {"WithAPIKey"},
	"account.setPushMuted":                {"WithAPIKey"},
	"account.listPushDevices":             {"WithAPIKey"},
	"account.removePushDevice":            {"WithAPIKey"},
	"account.sendTestPush":                {"WithAPIKey"},
	"account.setPhoto":                    {"WithContentType", "WithCrop", "WithAPIKey"},
	"account.removePhoto":                 {"WithAPIKey"},
	"account.getUsername":                 {"WithAPIKey"},
	"account.checkUsername":               {"WithAPIKey"},
	"account.setUsername":                 {"WithAPIKey"},
	"account.listConnectedApps":           {"WithAPIKey"},
	"account.listConnectedAppRequests":    {"WithFailedOnly", "WithStatusCodes", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"account.listAllConnectedAppRequests": {"WithFailedOnly", "WithStatusCodes", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"account.iterateConnectedAppRequests": {"WithFailedOnly", "WithStatusCodes", "WithSince", "WithUntil", "WithLimit", "WithCursor", "WithAPIKey"},
	"account.revokeConnectedApp":          {"WithAPIKey"},
	"account.endConnectedAppElevation":    {"WithAPIKey"},
	"account.listSessions":                {"WithAPIKey"},
	"account.revokeSession":               {"WithAPIKey"},
	"account.revokeSessions":              {"WithConnectedApps", "WithAPIKey"},
	"account.listSecurityEvents":          {"WithLimit", "WithCursor", "WithAPIKey"},
	"account.listAllSecurityEvents":       {"WithLimit", "WithCursor", "WithAPIKey"},
	"account.iterateSecurityEvents":       {"WithLimit", "WithCursor", "WithAPIKey"},
	"account.listPasskeys":                {"WithAPIKey"},
	"account.renamePasskey":               {"WithAPIKey"},
	"account.removePasskey":               {"WithAPIKey"},
	"account.listInvitations":             {"WithAPIKey"},
	"account.acceptInvitation":            {"WithActivate", "WithAPIKey"},
	"account.declineInvitation":           {"WithAPIKey"},
	"account.listOwnershipOffers":         {"WithAPIKey"},
	"account.acceptOwnershipOffer":        {"WithAPIKey"},
	"account.declineOwnershipOffer":       {"WithAPIKey"},
	"signatures.list":                     {"WithAddress", "WithFromAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"signatures.listAll":                  {"WithAddress", "WithFromAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"signatures.iterate":                  {"WithAddress", "WithFromAddress", "WithLimit", "WithCursor", "WithAPIKey"},
	"signatures.get":                      {"WithAPIKey"},
	"signatures.create":                   {"WithAPIKey"},
	"signatures.update":                   {"WithAPIKey"},
	"signatures.delete":                   {"WithAPIKey"},
	"signatures.copy":                     {"WithAPIKey"},
	"tasks.list":                          {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.listAll":                       {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.iterate":                       {"WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.get":                           {"WithAPIKey"},
	"tasks.create":                        {"WithAPIKey"},
	"tasks.update":                        {"WithAPIKey"},
	"tasks.delete":                        {"WithAPIKey"},
	"tasks.pause":                         {"WithAPIKey"},
	"tasks.resume":                        {"WithAPIKey"},
	"tasks.run":                           {"WithAPIKey"},
	"tasks.listRuns":                      {"WithTaskID", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.listAllRuns":                   {"WithTaskID", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.iterateRuns":                   {"WithTaskID", "WithStatus", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.getRun":                        {"WithAPIKey"},
	"tasks.startRun":                      {"WithAPIKey"},
	"tasks.cancelRun":                     {"WithAPIKey"},
	"tasks.listApprovals":                 {"WithStatus", "WithRunID", "WithTaskID", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.listAllApprovals":              {"WithStatus", "WithRunID", "WithTaskID", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.iterateApprovals":              {"WithStatus", "WithRunID", "WithTaskID", "WithLimit", "WithCursor", "WithAPIKey"},
	"tasks.getApproval":                   {"WithAPIKey"},
	"tasks.approve":                       {"WithAPIKey"},
	"tasks.deny":                          {"WithAPIKey"},
	"bookings.listPages":                  {"WithQ", "WithStatus", "WithSince", "WithUntil", "WithSort", "WithAPIKey"},
	"bookings.getPage":                    {"WithAPIKey"},
	"bookings.createPage":                 {"WithAPIKey"},
	"bookings.updatePage":                 {"WithAPIKey"},
	"bookings.uploadBackground":           {"WithContentType", "WithCrop", "WithAPIKey"},
	"bookings.removeBackground":           {"WithAPIKey"},
	"bookings.deletePage":                 {"WithAPIKey"},
	"bookings.listTimes":                  {"WithFrom", "WithTo", "WithDurationMinutes", "WithBookingID", "WithAPIKey"},
	"bookings.list":                       {"WithView", "WithPageID", "WithPageIDs", "WithQ", "WithFrom", "WithTo", "WithLimit", "WithCursor", "WithAPIKey"},
	"bookings.listAll":                    {"WithView", "WithPageID", "WithPageIDs", "WithQ", "WithFrom", "WithTo", "WithLimit", "WithCursor", "WithAPIKey"},
	"bookings.iterate":                    {"WithView", "WithPageID", "WithPageIDs", "WithQ", "WithFrom", "WithTo", "WithLimit", "WithCursor", "WithAPIKey"},
	"bookings.get":                        {"WithAPIKey"},
	"bookings.approve":                    {"WithAPIKey"},
	"bookings.decline":                    {"WithAPIKey"},
	"bookings.cancel":                     {"WithAPIKey"},
	"bookings.resendEmail":                {"WithAPIKey"},
	"bookings.analytics":                  {"WithPageIDs", "WithDays", "WithMinutes", "WithSince", "WithUntil", "WithGrain", "WithOffsetMinutes", "WithAPIKey"},
	"bookings.reschedule":                 {"WithAPIKey"},
	"connectors.listApps":                 {"WithSearch", "WithCategory", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectors.listAllApps":              {"WithSearch", "WithCategory", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectors.iterateApps":              {"WithSearch", "WithCategory", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectors.getApp":                   {"WithAPIKey"},
	"connectors.listCategories":           {"WithAPIKey"},
	"connectors.list":                     {"WithScope", "WithStatus", "WithApp", "WithAPIKey"},
	"connectors.get":                      {"WithAPIKey"},
	"connectors.connect":                  {"WithAPIKey"},
	"connectors.update":                   {"WithAPIKey"},
	"connectors.disconnect":               {"WithAPIKey"},
	"connectors.reconnect":                {"WithAPIKey"},
	"connectors.listActions":              {"WithSearch", "WithLimit", "WithAPIKey"},
	"connectors.getAction":                {"WithAPIKey"},
	"connectors.runAction":                {"WithAPIKey"},
	"connectors.listCalls":                {"WithConnectorID", "WithScope", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectors.listAllCalls":             {"WithConnectorID", "WithScope", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectors.iterateCalls":             {"WithConnectorID", "WithScope", "WithLimit", "WithCursor", "WithAPIKey"},
	"connectedMailboxes.list":             {"WithQ", "WithAPIKey"},
	"connectedMailboxes.get":              {"WithAPIKey"},
	"connectedMailboxes.discover":         {"WithAPIKey"},
	"connectedMailboxes.connect":          {"WithAPIKey"},
	"connectedMailboxes.update":           {"WithAPIKey"},
	"connectedMailboxes.sync":             {"WithAPIKey"},
	"connectedMailboxes.disconnect":       {"WithAPIKey"},
}
