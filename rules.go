package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func rulesListQuery(o *callOptions) values {
	return values{{query.Enabled, o.value("enabled")}, {query.Q, o.value("q")}}
}

func rulesRunsQuery(o *callOptions) values {
	return values{{query.RuleID, o.value("ruleId")}, {query.RuleIDs, o.value("ruleIds")}, {query.ThreadID, o.value("threadId")}}
}

func (s *RulesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("rules.list", opts)

	return s.t.page(ctx, o, call{path: paths.Rules, query: rulesListQuery(o)}, cursorPages)
}

func (s *RulesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("rules.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Rules, query: rulesListQuery(o)}, cursorPages)
}

func (s *RulesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("rules.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Rules, query: rulesListQuery(o)}, cursorPages)
}

func (s *RulesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.get", opts)

	return s.t.object(ctx, o, call{path: paths.Rule, params: []any{"id", id}})
}

func (s *RulesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Rules,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *RulesService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.update", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Rule,
		params: []any{"id", id},
		method: http.MethodPatch,
		body:   bodyOf(patch),
	})
}

func (s *RulesService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Rule,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *RulesService) Reorder(ctx context.Context, ruleIDs []string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("rules.reorder", opts)

	return s.t.objects(ctx, o, call{
		path:       paths.RulesReorder,
		method:     http.MethodPost,
		body:       Body{"ruleIds": ruleIDs},
		repeatable: true,
	})
}

func (s *RulesService) Test(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.test", opts)

	return s.t.object(ctx, o, call{
		path:       paths.RuleTest,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *RulesService) ListRuns(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("rules.listRuns", opts)

	return s.t.page(ctx, o, call{path: paths.RuleRuns, query: rulesRunsQuery(o)}, cursorPages)
}

func (s *RulesService) ListAllRuns(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("rules.listAllRuns", opts)

	return s.t.all(ctx, o, call{path: paths.RuleRuns, query: rulesRunsQuery(o)}, cursorPages)
}

func (s *RulesService) IterateRuns(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("rules.iterateRuns", opts)

	return s.t.iterate(ctx, o, call{path: paths.RuleRuns, query: rulesRunsQuery(o)}, cursorPages)
}

func (s *RulesService) Stats(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("rules.stats", opts)

	return s.t.object(ctx, o, call{
		path:  paths.RulesStats,
		query: values{{query.Since, o.value("since")}, {query.Until, o.value("until")}, {query.RuleIDs, o.value("ruleIds")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}},
	})
}
