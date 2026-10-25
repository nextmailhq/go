package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func formsSubmissionQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Status, o.value("status")}}
}

func formsWindowQuery(o *callOptions) values {
	return values{{query.Days, o.value("days")}, {query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.Since, o.value("since")}, {query.Until, o.value("until")}}
}

func (s *FormsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("forms.list", opts)

	return s.t.page(ctx, o, call{path: paths.Forms}, cursorPages)
}

func (s *FormsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("forms.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Forms}, cursorPages)
}

func (s *FormsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("forms.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Forms}, cursorPages)
}

func (s *FormsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Forms,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *FormsService) Design(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.design", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormsDesign,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *FormsService) Redesign(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.redesign", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormRedesign,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *FormsService) ListStarters(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("forms.listStarters", opts)

	return s.t.objects(ctx, o, call{path: paths.FormStarters})
}

func (s *FormsService) GetStarter(ctx context.Context, slug string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.getStarter", opts)

	return s.t.object(ctx, o, call{path: paths.FormStarter, params: []any{"slug", slug}})
}

func (s *FormsService) Get(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.get", opts)

	return s.t.object(ctx, o, call{path: paths.Form, params: []any{"id", id}})
}

func (s *FormsService) Update(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Form,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *FormsService) Delete(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Form,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *FormsService) Publish(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.publish", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FormPublish,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *FormsService) Pause(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.pause", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FormPause,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *FormsService) Resume(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.resume", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FormResume,
		params:     []any{"id", id},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *FormsService) Duplicate(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.duplicate", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormDuplicate,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}

func (s *FormsService) Analytics(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.analytics", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormAnalytics,
		params: []any{"id", id},
		query:  formsWindowQuery(o),
	})
}

func (s *FormsService) ListSubmissions(ctx context.Context, id string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("forms.listSubmissions", opts)

	return s.t.page(ctx, o, call{
		path:   paths.FormSubmissions,
		params: []any{"id", id},
		query:  formsSubmissionQuery(o),
	}, cursorPages)
}

func (s *FormsService) ListAllSubmissions(ctx context.Context, id string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("forms.listAllSubmissions", opts)

	return s.t.all(ctx, o, call{
		path:   paths.FormSubmissions,
		params: []any{"id", id},
		query:  formsSubmissionQuery(o),
	}, cursorPages)
}

func (s *FormsService) IterateSubmissions(ctx context.Context, id string, opts ...RequestOption) *Iterator {
	o := newCallOptions("forms.iterateSubmissions", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.FormSubmissions,
		params: []any{"id", id},
		query:  formsSubmissionQuery(o),
	}, cursorPages)
}

func (s *FormsService) GetSubmission(ctx context.Context, id string, submissionID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.getSubmission", opts)

	return s.t.object(ctx, o, call{path: paths.FormSubmission, params: []any{"id", id, "submissionId", submissionID}})
}

func (s *FormsService) DeleteSubmission(ctx context.Context, id string, submissionID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.deleteSubmission", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormSubmission,
		params: []any{"id", id, "submissionId", submissionID},
		method: http.MethodDelete,
	})
}

func (s *FormsService) DeleteSubmissions(ctx context.Context, id string, submissionIDs []string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.deleteSubmissions", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormSubmissionsBatchRemove,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   Body{"ids": submissionIDs},
	})
}

func (s *FormsService) ApproveSubmission(ctx context.Context, id string, submissionID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.approveSubmission", opts)

	return s.t.object(ctx, o, call{
		path:       paths.FormSubmissionApprove,
		params:     []any{"id", id, "submissionId", submissionID},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *FormsService) ResendConfirmation(ctx context.Context, id string, submissionID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.resendConfirmation", opts)

	return s.t.object(ctx, o, call{
		path:   paths.FormSubmissionResend,
		params: []any{"id", id, "submissionId", submissionID},
		method: http.MethodPost,
	})
}

func (s *FormsService) Subscribe(ctx context.Context, formID string, values Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("forms.subscribe", opts)

	return s.t.object(ctx, o, call{
		path:      paths.SubscribeForm,
		params:    []any{"formId", formID},
		method:    http.MethodPost,
		body:      bodyOf(values),
		anonymous: true,
	})
}
