package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func automationsListQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}}
}

func automationsEnrollmentQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}, {query.StepKey, o.value("stepKey")}, {query.Q, o.value("q")}}
}

func automationsWindowQuery(o *callOptions) values {
	return values{{query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func automationsAnalyticsQuery(o *callOptions) values {
	return values{{query.AutomationIDs, o.value("automationIds")}}.merge(automationsWindowQuery(o))
}

func automationsRunQuery(o *callOptions) values {
	return values{{query.Event, o.value("event")}, {query.StepKey, o.value("stepKey")}, {query.Outcome, o.value("outcome")}}
}

func (s *AutomationsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("automations.list", opts)

	return s.t.page(ctx, o, call{path: paths.Automations, query: automationsListQuery(o)}, cursorPages)
}

func (s *AutomationsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("automations.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Automations, query: automationsListQuery(o)}, cursorPages)
}

func (s *AutomationsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("automations.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Automations, query: automationsListQuery(o)}, cursorPages)
}

func (s *AutomationsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Automations,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *AutomationsService) ListStarters(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("automations.listStarters", opts)

	return s.t.objects(ctx, o, call{path: paths.AutomationStarters})
}

func (s *AutomationsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.get", opts)

	return s.t.object(ctx, o, call{path: paths.Automation, params: []any{"id", id}})
}

func (s *AutomationsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Automation,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *AutomationsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Automation,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *AutomationsService) Publish(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.publish", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AutomationPublish,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AutomationsService) Pause(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.pause", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AutomationPause,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AutomationsService) Resume(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.resume", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AutomationResume,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AutomationsService) Archive(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.archive", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AutomationArchive,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AutomationsService) Duplicate(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.duplicate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AutomationDuplicate,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *AutomationsService) SendTest(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.sendTest", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AutomationTest,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *AutomationsService) ListVersions(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("automations.listVersions", opts)

	return s.t.objects(ctx, o, call{path: paths.AutomationVersions, params: []any{"id", id}})
}

func (s *AutomationsService) RestoreVersion(ctx context.Context, id string, version int, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.restoreVersion", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AutomationVersionRestore,
		params:     []any{"id", id, "version", version},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *AutomationsService) Stats(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.stats", opts)

	return s.t.object(ctx, o, call{path: paths.AutomationStats, params: []any{"id", id}, query: automationsWindowQuery(o)})
}

func (s *AutomationsService) Analytics(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.analytics", opts)

	return s.t.object(ctx, o, call{path: paths.AutomationsAnalytics, query: automationsAnalyticsQuery(o)})
}

func (s *AutomationsService) ListRuns(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("automations.listRuns", opts)

	return s.t.page(ctx, o, call{path: paths.AutomationRuns, params: []any{"id", id}, query: automationsRunQuery(o)}, cursorPages)
}

func (s *AutomationsService) ListAllRuns(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("automations.listAllRuns", opts)

	return s.t.all(ctx, o, call{path: paths.AutomationRuns, params: []any{"id", id}, query: automationsRunQuery(o)}, cursorPages)
}

func (s *AutomationsService) IterateRuns(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("automations.iterateRuns", opts)

	return s.t.iterate(ctx, o, call{path: paths.AutomationRuns, params: []any{"id", id}, query: automationsRunQuery(o)}, cursorPages)
}

func (s *AutomationsService) ListEnrollments(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("automations.listEnrollments", opts)

	return s.t.page(ctx, o, call{path: paths.AutomationEnrollments, params: []any{"id", id}, query: automationsEnrollmentQuery(o)}, cursorPages)
}

func (s *AutomationsService) ListAllEnrollments(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("automations.listAllEnrollments", opts)

	return s.t.all(ctx, o, call{path: paths.AutomationEnrollments, params: []any{"id", id}, query: automationsEnrollmentQuery(o)}, cursorPages)
}

func (s *AutomationsService) IterateEnrollments(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("automations.iterateEnrollments", opts)

	return s.t.iterate(ctx, o, call{path: paths.AutomationEnrollments, params: []any{"id", id}, query: automationsEnrollmentQuery(o)}, cursorPages)
}

func (s *AutomationsService) GetEnrollment(ctx context.Context, id string, enrollmentID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.getEnrollment", opts)

	return s.t.object(ctx, o, call{path: paths.AutomationEnrollment, params: []any{"id", id, "enrollmentId", enrollmentID}})
}

func (s *AutomationsService) Enroll(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.enroll", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AutomationEnrollments,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *AutomationsService) ExitEnrollment(ctx context.Context, id string, enrollmentID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("automations.exitEnrollment", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AutomationEnrollmentExit,
		params: []any{"id", id, "enrollmentId", enrollmentID},
		method: http.MethodPost,
	})
}
