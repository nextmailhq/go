package nextmail

import (
	"context"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *SettingsService) Get(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("settings.get", opts)

	return s.t.object(ctx, o, call{path: paths.Settings, query: values{{query.Address, o.value("address")}}})
}

func (s *SettingsService) Update(ctx context.Context, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("settings.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Settings,
		method:     http.MethodPatch,
		query:      values{{query.Address, o.value("address")}},
		body:       bodyOf(patch),
		repeatable: true,
	})
}
