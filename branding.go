package nextmail

import (
	"context"
	"io"
	"net/http"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/query"
)

func (s *BrandingService) Get(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("branding.get", opts)

	return s.t.object(ctx, o, call{path: paths.Branding})
}

func (s *BrandingService) Update(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("branding.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Branding,
		method:     http.MethodPatch,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *BrandingService) UploadImage(ctx context.Context, variant string, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("branding.uploadImage", opts)

	return s.t.object(ctx, o, call{
		path:        paths.BrandingImage,
		params:      []any{"variant", variant},
		method:      http.MethodPut,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		repeatable:  true,
		upload:      true,
	})
}

func (s *BrandingService) RemoveImage(ctx context.Context, variant string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("branding.removeImage", opts)

	return s.t.object(ctx, o, call{
		path:       paths.BrandingImage,
		params:     []any{"variant", variant},
		method:     http.MethodDelete,
		repeatable: true,
	})
}
