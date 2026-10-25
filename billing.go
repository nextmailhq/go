package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *BillingService) Get(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.get", opts)

	return s.t.object(ctx, o, call{path: paths.Billing})
}

func (s *BillingService) ListPlans(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.listPlans", opts)

	return s.t.object(ctx, o, call{path: paths.BillingPlans})
}

func (s *BillingService) GetUsage(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.getUsage", opts)

	return s.t.object(ctx, o, call{path: paths.BillingUsage, query: values{{query.Days, o.value("days")}}})
}

func (s *BillingService) ListAlerts(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.listAlerts", opts)

	return s.t.object(ctx, o, call{path: paths.BillingAlerts})
}

func (s *BillingService) ListCountries(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.listCountries", opts)

	return s.t.object(ctx, o, call{path: paths.BillingCountries})
}

func (s *BillingService) GetPayAsYouGo(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.getPayAsYouGo", opts)

	return s.t.object(ctx, o, call{path: paths.BillingPayAsYouGo})
}

func (s *BillingService) SetPayAsYouGo(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.setPayAsYouGo", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingPayAsYouGo,
		method: http.MethodPatch,
		body:   bodyOf(body),
	})
}

func (s *BillingService) SetPayAsYouGoLimit(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.setPayAsYouGoLimit", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingPayAsYouGoLimit,
		method: http.MethodPut,
		body:   bodyOf(body),
	})
}

func (s *BillingService) ConfirmPayAsYouGo(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.confirmPayAsYouGo", opts)

	return s.t.object(ctx, o, call{path: paths.BillingPayAsYouGoConfirm, method: http.MethodPost})
}

func (s *BillingService) StartCheckout(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.startCheckout", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingCheckout,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BillingService) OpenPortal(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.openPortal", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingPortal,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BillingService) ListInvoices(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.listInvoices", opts)

	return s.t.object(ctx, o, call{
		path:  paths.BillingInvoices,
		query: values{{query.Page, o.value("page")}, {query.Limit, o.value("limit")}, {query.Sort, o.value("sort")}, {query.Search, o.value("search")}},
	})
}

func (s *BillingService) GetInvoice(ctx context.Context, orderID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.getInvoice", opts)

	return s.t.object(ctx, o, call{path: paths.BillingInvoice, params: []any{"orderId", orderID}})
}

func (s *BillingService) SaveInvoiceDetails(ctx context.Context, orderID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.saveInvoiceDetails", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingInvoiceDetails,
		params: []any{"orderId", orderID},
		method: http.MethodPut,
		body:   bodyOf(body),
	})
}

func (s *BillingService) ListAddOns(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.listAddOns", opts)

	return s.t.object(ctx, o, call{path: paths.BillingAddOns})
}

func (s *BillingService) StartAddOnCheckout(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.startAddOnCheckout", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingAddOnCheckout,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *BillingService) CancelAddOn(ctx context.Context, addOn string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("billing.cancelAddOn", opts)

	return s.t.object(ctx, o, call{
		path:   paths.BillingAddOn,
		params: []any{"addOn", addOn},
		method: http.MethodDelete,
	})
}
