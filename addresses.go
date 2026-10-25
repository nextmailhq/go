package nextmail

import (
	"context"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

func (s *AddressesService) List(ctx context.Context, opts ...RequestOption) (*AddressBookPage, error) {
	o := newCallOptions("addresses.list", opts)

	return s.page(ctx, o, o.text(optionCursor))
}

func (s *AddressesService) page(ctx context.Context, o *callOptions, cursor string) (*AddressBookPage, error) {
	envelope, err := s.t.object(ctx, o, call{
		path:  paths.Addresses,
		query: values{{query.Limit, o.value("limit")}, {protocol.CursorStylesCursorParam, cursor}},
	})

	if err != nil {
		return nil, err
	}

	next := cursorOf(envelope[fieldNextCursor])

	return &AddressBookPage{
		Unrestricted: envelope.Bool("unrestricted"),
		Addresses:    objectsOf(envelope[fieldData]),
		Domains:      objectsOf(envelope["domains"]),
		HasMore:      hasMoreOf(envelope, next),
		NextCursor:   next,
	}, nil
}

func (s *AddressesService) ListAll(ctx context.Context, opts ...RequestOption) (*AddressBook, error) {
	o := newCallOptions("addresses.listAll", opts)
	cursor := o.text(optionCursor)
	current, err := s.page(ctx, o, cursor)

	if err != nil {
		return nil, err
	}

	addresses := current.Addresses
	followed := map[string]struct{}{cursor: {}}

	for current.HasMore && current.NextCursor != "" {
		if _, seen := followed[current.NextCursor]; seen {
			break
		}

		followed[current.NextCursor] = struct{}{}

		if current, err = s.page(ctx, o, current.NextCursor); err != nil {
			return nil, err
		}

		addresses = append(addresses, current.Addresses...)
	}

	return &AddressBook{Unrestricted: current.Unrestricted, Addresses: addresses, Domains: current.Domains}, nil
}

func (s *AddressesService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("addresses.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Addresses}, cursorPages)
}
