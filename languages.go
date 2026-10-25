package nextmail

import (
	"context"

	"nextmail.sh/internal/paths"
)

func (s *LanguagesService) List(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("languages.list", opts)

	return s.t.objects(ctx, o, call{path: paths.Languages})
}
