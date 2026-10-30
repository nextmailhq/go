package nextmail

import (
	"context"
	"io"
	"net/http"
	"time"

	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

func contactsListQuery(o *callOptions) values {
	return values{{query.Source, o.value("source")}, {query.Q, o.value("q")}, {query.AudienceID, o.value("audienceId")}}
}

func contactsCardQuery(o *callOptions) values {
	return values{
		{query.Email, o.value("email")},
		{query.WithoutEmail, o.value("withoutEmail")},
		{query.Kind, o.value("kind")},
		{query.GroupID, o.value("groupId")},
		{query.AddressBook, o.value("addressBook")},
	}
}

func contactsPeopleQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Email, o.value("email")}, {query.Sort, o.value("sort")}, {query.Blocked, o.value("blocked")}}
}

func contactsThreadQuery(o *callOptions) values {
	return values{{query.Q, o.value("q")}, {query.Sort, o.value("sort")}}
}

func contactsActivityQuery(o *callOptions) values {
	return values{{query.Minutes, o.value("minutes")}, {query.Grain, o.value("grain")}, {query.OffsetMinutes, o.value("offsetMinutes")}, {query.From, o.value("from")}, {query.To, o.value("to")}}
}

func contactsEventQuery(o *callOptions) values {
	return values{{query.Name, o.value("name")}}
}

func (s *ContactsService) List(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("contacts.list", opts)

	return s.t.page(ctx, o, call{path: paths.Contacts, query: contactsListQuery(o)}, cursorPages)
}

func (s *ContactsService) ListAll(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAll", opts)

	return s.t.all(ctx, o, call{path: paths.Contacts, query: contactsListQuery(o)}, cursorPages)
}

func (s *ContactsService) Iterate(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iterate", opts)

	return s.t.iterate(ctx, o, call{path: paths.Contacts, query: contactsListQuery(o)}, cursorPages)
}

func (s *ContactsService) Get(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.get", opts)

	return s.t.object(ctx, o, call{path: paths.Contact, params: []any{"email", email}})
}

func (s *ContactsService) Create(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.create", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Contacts,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ContactsService) Update(ctx context.Context, email string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.update", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Contact,
		params:     []any{"email", email},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *ContactsService) Delete(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.delete", opts)

	return s.t.object(ctx, o, call{
		path:   paths.Contact,
		params: []any{"email", email},
		method: http.MethodDelete,
	})
}

func (s *ContactsService) SetAudiences(ctx context.Context, email string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.setAudiences", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactAudiences,
		params:     []any{"email", email},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *ContactsService) ListPeople(ctx context.Context, opts ...RequestOption) (*PeoplePage, error) {
	o := newCallOptions("contacts.listPeople", opts)
	payload, err := s.t.object(ctx, o, call{
		path: paths.ContactsPeople,
		query: contactsPeopleQuery(o).merge(values{
			{query.Limit, o.value("limit")},
			{protocol.CursorStylesCursorParam, o.value("cursor")},
		}),
	})

	if err != nil {
		return nil, err
	}

	next := cursorOf(payload[protocol.CursorStylesCursorField])

	return &PeoplePage{Items: objectsOf(payload[fieldData]), HasMore: hasMoreOf(payload, next), NextCursor: next, Seen: payload.Bool("seen")}, nil
}

func (s *ContactsService) ListAllPeople(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAllPeople", opts)

	return s.t.all(ctx, o, call{path: paths.ContactsPeople, query: contactsPeopleQuery(o)}, cursorPages)
}

func (s *ContactsService) IteratePeople(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iteratePeople", opts)

	return s.t.iterate(ctx, o, call{path: paths.ContactsPeople, query: contactsPeopleQuery(o)}, cursorPages)
}

func (s *ContactsService) Save(ctx context.Context, email string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.save", opts)

	return s.t.object(ctx, o, call{
		path:       paths.Contact,
		params:     []any{"email", email},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *ContactsService) DeleteMany(ctx context.Context, emails []string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.deleteMany", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactsBatchDelete,
		method:     http.MethodPost,
		body:       Body{"emails": emails},
		repeatable: true,
	})
}

func (s *ContactsService) SetPhoto(ctx context.Context, email string, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.setPhoto", opts)

	return s.t.object(ctx, o, call{
		path:        paths.ContactPhoto,
		params:      []any{"email", email},
		method:      http.MethodPut,
		query:       values{{query.Crop, o.value("crop")}},
		raw:         data,
		sendsRaw:    true,
		contentType: o.text("contentType"),
		repeatable:  true,
		upload:      true,
	})
}

func (s *ContactsService) RemovePhoto(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.removePhoto", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactPhoto,
		params:     []any{"email", email},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *ContactsService) Block(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.block", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactBlock,
		params:     []any{"email", email},
		method:     http.MethodPost,
		repeatable: true,
	})
}

func (s *ContactsService) Unblock(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.unblock", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactBlock,
		params:     []any{"email", email},
		method:     http.MethodDelete,
		repeatable: true,
	})
}

func (s *ContactsService) ListThreads(ctx context.Context, email string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("contacts.listThreads", opts)

	return s.t.page(ctx, o, call{
		path:   paths.ContactThreads,
		params: []any{"email", email},
		query:  contactsThreadQuery(o),
	}, cursorPages)
}

func (s *ContactsService) ListAllThreads(ctx context.Context, email string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAllThreads", opts)

	return s.t.all(ctx, o, call{
		path:   paths.ContactThreads,
		params: []any{"email", email},
		query:  contactsThreadQuery(o),
	}, cursorPages)
}

func (s *ContactsService) IterateThreads(ctx context.Context, email string, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iterateThreads", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.ContactThreads,
		params: []any{"email", email},
		query:  contactsThreadQuery(o),
	}, cursorPages)
}

func (s *ContactsService) Activity(ctx context.Context, email string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.activity", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ContactActivity,
		params: []any{"email", email},
		query:  contactsActivityQuery(o),
	})
}

func (s *ContactsService) ListEvents(ctx context.Context, email string, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("contacts.listEvents", opts)

	return s.t.page(ctx, o, call{
		path:   paths.ContactEvents,
		params: []any{"email", email},
		query:  contactsEventQuery(o),
	}, cursorPages)
}

func (s *ContactsService) ListAllEvents(ctx context.Context, email string, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAllEvents", opts)

	return s.t.all(ctx, o, call{
		path:   paths.ContactEvents,
		params: []any{"email", email},
		query:  contactsEventQuery(o),
	}, cursorPages)
}

func (s *ContactsService) IterateEvents(ctx context.Context, email string, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iterateEvents", opts)

	return s.t.iterate(ctx, o, call{
		path:   paths.ContactEvents,
		params: []any{"email", email},
		query:  contactsEventQuery(o),
	}, cursorPages)
}

func (s *ContactsService) ListCards(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("contacts.listCards", opts)

	return s.t.page(ctx, o, call{path: paths.ContactCards, query: contactsCardQuery(o)}, cursorPages)
}

func (s *ContactsService) ListAllCards(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAllCards", opts)

	return s.t.all(ctx, o, call{path: paths.ContactCards, query: contactsCardQuery(o)}, cursorPages)
}

func (s *ContactsService) IterateCards(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iterateCards", opts)

	return s.t.iterate(ctx, o, call{path: paths.ContactCards, query: contactsCardQuery(o)}, cursorPages)
}

func (s *ContactsService) GetCard(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.getCard", opts)

	return s.t.object(ctx, o, call{path: paths.ContactCard, params: []any{"id", id}})
}

func (s *ContactsService) CreateCard(ctx context.Context, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.createCard", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ContactCards,
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ContactsService) UpdateCard(ctx context.Context, id string, patch Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.updateCard", opts)

	return s.t.object(ctx, o, call{
		path:       paths.ContactCard,
		params:     []any{"id", id},
		method:     http.MethodPatch,
		body:       bodyOf(patch),
		repeatable: true,
	})
}

func (s *ContactsService) DeleteCard(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.deleteCard", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ContactCard,
		params: []any{"id", id},
		method: http.MethodDelete,
	})
}

func (s *ContactsService) FindDuplicateCards(ctx context.Context, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.findDuplicateCards", opts)

	return s.t.object(ctx, o, call{path: paths.ContactCardDuplicates})
}

func (s *ContactsService) MergeCards(ctx context.Context, id string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.mergeCards", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ContactCardMerge,
		params: []any{"id", id},
		method: http.MethodPost,
		body:   bodyOf(body),
	})
}

func (s *ContactsService) ListAddressBooks(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAddressBooks", opts)

	return s.t.objects(ctx, o, call{path: paths.AddressBooks})
}

func (s *ContactsService) SetAddressBookMember(ctx context.Context, id string, memberID string, body Body, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.setAddressBookMember", opts)

	return s.t.object(ctx, o, call{
		path:       paths.AddressBookMember,
		params:     []any{"id", id, "memberId", memberID},
		method:     http.MethodPut,
		body:       bodyOf(body),
		repeatable: true,
	})
}

func (s *ContactsService) RemoveAddressBookMember(ctx context.Context, id string, memberID string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.removeAddressBookMember", opts)

	return s.t.object(ctx, o, call{
		path:   paths.AddressBookMember,
		params: []any{"id", id, "memberId", memberID},
		method: http.MethodDelete,
	})
}

func (s *ContactsService) ListBirthdays(ctx context.Context, from time.Time, to time.Time, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listBirthdays", opts)

	return s.t.objects(ctx, o, call{
		path: paths.ContactBirthdays,
		query: values{
			{query.From, instant(from)},
			{query.To, instant(to)},
			{query.Timezone, o.value("timezone")},
		},
	})
}

func (s *ContactsService) ImportVcf(ctx context.Context, data io.Reader, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.importVcf", opts)

	return s.t.object(ctx, o, call{
		path:        paths.ContactImports,
		method:      http.MethodPost,
		query:       values{{query.Filename, o.value("filename")}},
		raw:         data,
		sendsRaw:    true,
		contentType: protocol.ContentTypesVcard,
	})
}

func (s *ContactsService) ListImports(ctx context.Context, opts ...RequestOption) (*Page, error) {
	o := newCallOptions("contacts.listImports", opts)

	return s.t.page(ctx, o, call{path: paths.ContactImports}, cursorPages)
}

func (s *ContactsService) ListAllImports(ctx context.Context, opts ...RequestOption) ([]Object, error) {
	o := newCallOptions("contacts.listAllImports", opts)

	return s.t.all(ctx, o, call{path: paths.ContactImports}, cursorPages)
}

func (s *ContactsService) IterateImports(ctx context.Context, opts ...RequestOption) *Iterator {
	o := newCallOptions("contacts.iterateImports", opts)

	return s.t.iterate(ctx, o, call{path: paths.ContactImports}, cursorPages)
}

func (s *ContactsService) GetImport(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.getImport", opts)

	return s.t.object(ctx, o, call{path: paths.ContactImport, params: []any{"id", id}})
}

func (s *ContactsService) UndoImport(ctx context.Context, id string, opts ...RequestOption) (Object, error) {
	o := newCallOptions("contacts.undoImport", opts)

	return s.t.object(ctx, o, call{
		path:   paths.ContactImportUndo,
		params: []any{"id", id},
		method: http.MethodPost,
	})
}
