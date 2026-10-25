package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func audiencesMemberQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Source, o.value("source")}, {query.Sort, o.value("sort")}, {query.Status, o.value("statuses")}, {query.InactiveDays, o.value("inactiveDays")}}
}

func audiencesGrowthQuery(o *callOptions) values {
	return values{{query.AudienceIDs, o.value("audienceIds")}, {query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func (s *AudiencesService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("audiences.list", opts)

	return s.t.page(ctx, o, call{path: paths.Audiences}, cursorPages)
}

func (s *AudiencesService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("audiences.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Audiences}, cursorPages)
}

func (s *AudiencesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("audiences.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Audiences}, cursorPages)
}

func (s *AudiencesService) Growth(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.growth", opts)

	return s.t.object(ctx, o, call{path: paths.AudiencesGrowth, query: audiencesGrowthQuery(o)})
}

func (s *AudiencesService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.get", opts)

	return s.t.object(ctx, o, call{path: paths.Audience, params: []any{"id", id}})
}

func (s *AudiencesService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Audiences,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *AudiencesService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Audience,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *AudiencesService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Audience,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *AudiencesService) Empty(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.empty", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AudienceEmpty,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *AudiencesService) ListContacts(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("audiences.listContacts", opts)

	return s.t.page(ctx, o, call{
		path:   paths.AudienceContacts,
		params: []any{"id", id},
		query:  audiencesMemberQuery(o),
	}, cursorPages)
}

func (s *AudiencesService) ListAllContacts(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("audiences.listAllContacts", opts)

	return s.t.all(ctx, o, call{
		path:   paths.AudienceContacts,
		params: []any{"id", id},
		query:  audiencesMemberQuery(o),
	}, cursorPages)
}

func (s *AudiencesService) IterateContacts(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("audiences.iterateContacts", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.AudienceContacts,
		params: []any{"id", id},
		query:  audiencesMemberQuery(o),
	}, cursorPages)
}

func (s *AudiencesService) AddContact(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.addContact", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AudienceContacts,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *AudiencesService) RemoveContact(ctx context.Context, id string, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.removeContact", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AudienceContact,
		params: []any{"id", id, "email", email},
		method: http.MethodDelete,
	})
}

func (s *AudiencesService) AddContacts(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.addContacts", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AudienceContactsBatch,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *AudiencesService) RemoveContacts(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.removeContacts", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AudienceContactsBatchRemove,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *AudiencesService) ImportContacts(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("audiences.importContacts", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AudienceImport,
		params:     []any{"id", id},
		method:     http.MethodPost,
		body:       bodyOf(body),
		repeatable: true,
	})
}
