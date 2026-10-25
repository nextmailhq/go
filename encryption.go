package nextmail

import (
	"context"
	"net/http"
	"strings"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

func (s *EncryptionService) ListKeys(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("encryption.listKeys", opts)

	return s.t.objects(ctx, o, call{path: paths.EncryptionKeys})
}

func (s *EncryptionService) PublishKey(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("encryption.publishKey", opts)

	return s.t.object(ctx, o, call{
		path:   paths.EncryptionKeys,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *EncryptionService) LookupKeys(ctx context.Context, addresses []string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("encryption.lookupKeys", opts)

	return s.t.objects(ctx, o, call{
		path:  paths.EncryptionKeysLookup,
		query: values{{query.Addresses, strings.Join(addresses, protocol.ListSeparator)}},
	})
}
