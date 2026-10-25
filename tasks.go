package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func tasksListQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}}
}

func tasksRunQuery(o *callOptions) values {
	return values{{query.TaskID, o.value("taskId")}, {query.Status, o.value("status")}}
}

func tasksApprovalQuery(o *callOptions) values {
	return values{{query.Status, o.value("status")}, {query.RunID, o.value("runId")}, {query.TaskID, o.value("taskId")}}
}

func (s *TasksService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tasks.list", opts)

	return s.t.page(ctx, o, call{path: paths.Tasks, query: tasksListQuery(o)}, cursorPages)
}

func (s *TasksService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tasks.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Tasks, query: tasksListQuery(o)}, cursorPages)
}

func (s *TasksService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("tasks.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Tasks, query: tasksListQuery(o)}, cursorPages)
}

func (s *TasksService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.get", opts)

	return s.t.object(ctx, o, call{path: paths.Task, params: []any{"id", id}})
}

func (s *TasksService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Tasks,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TasksService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Task,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *TasksService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Task,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *TasksService) Pause(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.pause", opts)

	return s.t.object(ctx, o, call{
		path:       paths.TaskPause,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *TasksService) Resume(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.resume", opts)

	return s.t.object(ctx, o, call{
		path:       paths.TaskResume,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *TasksService) Run(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.run", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TaskRunNow,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *TasksService) ListRuns(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tasks.listRuns", opts)

	return s.t.page(ctx, o, call{path: paths.TaskRuns, query: tasksRunQuery(o)}, cursorPages)
}

func (s *TasksService) ListAllRuns(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tasks.listAllRuns", opts)

	return s.t.all(ctx, o, call{path: paths.TaskRuns, query: tasksRunQuery(o)}, cursorPages)
}

func (s *TasksService) IterateRuns(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("tasks.iterateRuns", opts)

	return s.t.iterate(ctx, o, call{path: paths.TaskRuns, query: tasksRunQuery(o)}, cursorPages)
}

func (s *TasksService) GetRun(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.getRun", opts)

	return s.t.object(ctx, o, call{path: paths.TaskRun, params: []any{"id", id}})
}

func (s *TasksService) StartRun(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.startRun", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TaskRuns,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TasksService) CancelRun(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.cancelRun", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TaskRunCancel,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *TasksService) ListApprovals(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("tasks.listApprovals", opts)

	return s.t.page(ctx, o, call{path: paths.TaskApprovals, query: tasksApprovalQuery(o)}, cursorPages)
}

func (s *TasksService) ListAllApprovals(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("tasks.listAllApprovals", opts)

	return s.t.all(ctx, o, call{path: paths.TaskApprovals, query: tasksApprovalQuery(o)}, cursorPages)
}

func (s *TasksService) IterateApprovals(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("tasks.iterateApprovals", opts)

	return s.t.iterate(ctx, o, call{path: paths.TaskApprovals, query: tasksApprovalQuery(o)}, cursorPages)
}

func (s *TasksService) GetApproval(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.getApproval", opts)

	return s.t.object(ctx, o, call{path: paths.TaskApproval, params: []any{"id", id}})
}

func (s *TasksService) Approve(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.approve", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TaskApprovalApprove,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *TasksService) Deny(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("tasks.deny", opts)

	return s.t.object(ctx, o, call{
		path:   paths.TaskApprovalDeny,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}
