package nextmail

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"nextmail.sh/internal/protocol"
)

type service struct {
	t *transport
}

type (
	MeService                 service
	SecurityService           service
	KeysService               service
	AddressesService          service
	LanguagesService          service
	EmailsService             service
	TemplatesService          service
	TrackingService           service
	ThreadsService            service
	FollowUpsService          service
	DraftsService             service
	LabelsService             service
	ContactsService           service
	AudiencesService          service
	BillingService            service
	FormsService              service
	AutomationsService        service
	EventsService             service
	BroadcastsService         service
	DomainsService            service
	AppHostService            service
	BrandingService           service
	RulesService              service
	WebhooksService           service
	ImportsService            service
	ProviderImportsService    service
	CalendarService           service
	SettingsService           service
	OutOfOfficeService        service
	RolesService              service
	MembersService            service
	SuppressionsService       service
	SendingService            service
	FilesService              service
	KnowledgeService          service
	TempMailService           service
	ExportsService            service
	ChatsService              service
	EncryptionService         service
	SupportService            service
	ToolsService              service
	SendersService            service
	AnalyticsService          service
	SubscriptionsService      service
	DNSConnectionsService     service
	WorkspacesService         service
	AccountService            service
	SignaturesService         service
	TasksService              service
	BookingsService           service
	ConnectorsService         service
	ConnectedMailboxesService service
	SegmentsService           service
)

type Client struct {
	Me                 *MeService
	Security           *SecurityService
	Keys               *KeysService
	Addresses          *AddressesService
	Languages          *LanguagesService
	Emails             *EmailsService
	Templates          *TemplatesService
	Tracking           *TrackingService
	Threads            *ThreadsService
	FollowUps          *FollowUpsService
	Drafts             *DraftsService
	Labels             *LabelsService
	Contacts           *ContactsService
	Audiences          *AudiencesService
	Billing            *BillingService
	Forms              *FormsService
	Automations        *AutomationsService
	Events             *EventsService
	Broadcasts         *BroadcastsService
	Domains            *DomainsService
	AppHost            *AppHostService
	Branding           *BrandingService
	Rules              *RulesService
	Webhooks           *WebhooksService
	Imports            *ImportsService
	ProviderImports    *ProviderImportsService
	Calendar           *CalendarService
	Settings           *SettingsService
	OutOfOffice        *OutOfOfficeService
	Roles              *RolesService
	Members            *MembersService
	Suppressions       *SuppressionsService
	Sending            *SendingService
	Files              *FilesService
	Knowledge          *KnowledgeService
	TempMail           *TempMailService
	Exports            *ExportsService
	Chats              *ChatsService
	Encryption         *EncryptionService
	Support            *SupportService
	Tools              *ToolsService
	Senders            *SendersService
	Analytics          *AnalyticsService
	Subscriptions      *SubscriptionsService
	DNSConnections     *DNSConnectionsService
	Workspaces         *WorkspacesService
	Account            *AccountService
	Signatures         *SignaturesService
	Tasks              *TasksService
	Bookings           *BookingsService
	Connectors         *ConnectorsService
	ConnectedMailboxes *ConnectedMailboxesService
	Segments           *SegmentsService

	common service
}

func New(opts ...ClientOption) *Client {
	config := configOf(opts)

	if config.inboxToken.present() {
		config.fail(argumentError(messageInboxTokenClient))
	}

	if !config.hasAPIKey && !config.hasToken {
		if apiKey := os.Getenv(protocol.EnvVarsAPIKey); apiKey != "" {
			config.apiKey = newSecret(apiKey)
		} else if accessToken := os.Getenv(protocol.EnvVarsAccessToken); accessToken != "" {
			config.accessToken = newSecret(accessToken)
		}
	}

	client := &Client{}
	client.common.t = newTransport(config, true)
	client.Me = (*MeService)(&client.common)
	client.Security = (*SecurityService)(&client.common)
	client.Keys = (*KeysService)(&client.common)
	client.Addresses = (*AddressesService)(&client.common)
	client.Languages = (*LanguagesService)(&client.common)
	client.Emails = (*EmailsService)(&client.common)
	client.Templates = (*TemplatesService)(&client.common)
	client.Tracking = (*TrackingService)(&client.common)
	client.Threads = (*ThreadsService)(&client.common)
	client.FollowUps = (*FollowUpsService)(&client.common)
	client.Drafts = (*DraftsService)(&client.common)
	client.Labels = (*LabelsService)(&client.common)
	client.Contacts = (*ContactsService)(&client.common)
	client.Audiences = (*AudiencesService)(&client.common)
	client.Billing = (*BillingService)(&client.common)
	client.Forms = (*FormsService)(&client.common)
	client.Automations = (*AutomationsService)(&client.common)
	client.Events = (*EventsService)(&client.common)
	client.Broadcasts = (*BroadcastsService)(&client.common)
	client.Domains = (*DomainsService)(&client.common)
	client.AppHost = (*AppHostService)(&client.common)
	client.Branding = (*BrandingService)(&client.common)
	client.Rules = (*RulesService)(&client.common)
	client.Webhooks = (*WebhooksService)(&client.common)
	client.Imports = (*ImportsService)(&client.common)
	client.ProviderImports = (*ProviderImportsService)(&client.common)
	client.Calendar = (*CalendarService)(&client.common)
	client.Settings = (*SettingsService)(&client.common)
	client.OutOfOffice = (*OutOfOfficeService)(&client.common)
	client.Roles = (*RolesService)(&client.common)
	client.Members = (*MembersService)(&client.common)
	client.Suppressions = (*SuppressionsService)(&client.common)
	client.Sending = (*SendingService)(&client.common)
	client.Files = (*FilesService)(&client.common)
	client.Knowledge = (*KnowledgeService)(&client.common)
	client.TempMail = (*TempMailService)(&client.common)
	client.Exports = (*ExportsService)(&client.common)
	client.Chats = (*ChatsService)(&client.common)
	client.Encryption = (*EncryptionService)(&client.common)
	client.Support = (*SupportService)(&client.common)
	client.Tools = (*ToolsService)(&client.common)
	client.Senders = (*SendersService)(&client.common)
	client.Analytics = (*AnalyticsService)(&client.common)
	client.Subscriptions = (*SubscriptionsService)(&client.common)
	client.DNSConnections = (*DNSConnectionsService)(&client.common)
	client.Workspaces = (*WorkspacesService)(&client.common)
	client.Account = (*AccountService)(&client.common)
	client.Signatures = (*SignaturesService)(&client.common)
	client.Tasks = (*TasksService)(&client.common)
	client.Bookings = (*BookingsService)(&client.common)
	client.Connectors = (*ConnectorsService)(&client.common)
	client.ConnectedMailboxes = (*ConnectedMailboxesService)(&client.common)
	client.Segments = (*SegmentsService)(&client.common)

	return client
}

func NewTempMail(opts ...ClientOption) *TempMailService {
	config := configOf(opts)

	if config.hasAPIKey || config.hasToken {
		config.fail(argumentError(messageTempMailCredential))
	}

	return (*TempMailService)(&service{t: newTransport(config, false)})
}

func configOf(opts []ClientOption) *clientConfig {
	config := &clientConfig{}

	for _, option := range opts {
		if option != nil {
			option.applyClient(config)
		}
	}

	if !config.hasBaseURL {
		config.baseURL = os.Getenv(protocol.EnvVarsBaseURL)
	}

	config.baseURL = normaliseEndpoint(config.baseURL)

	return config
}

func (c *Client) Err() error {
	return c.common.t.err
}

func (c *Client) Mode() string {
	return c.common.t.mode
}

func (c *Client) BaseURL() string {
	return c.common.t.baseURL
}

func (c *Client) Request(ctx context.Context, method, path string, body any, opts ...RequestOption) (any, error) {
	options := newCallOptions(rawRequestAccessor, opts)
	decoded, _, err := c.common.t.json(ctx, options, call{
		method:     method,
		path:       path,
		body:       body,
		idempotent: options.has(optionIdempotencyKey),
	})

	return decoded, err
}

func (c *Client) String() string {
	if c == nil {
		return "nextmail.Client(nil)"
	}

	return "nextmail.Client{mode: " + c.Mode() + ", baseURL: " + c.BaseURL() + ", credential: " + redacted + "}"
}

func (c *Client) GoString() string {
	return c.String()
}

func (c *Client) Format(state fmt.State, _ rune) {
	_, _ = fmt.Fprint(state, c.String())
}

func (c *Client) LogValue() slog.Value {
	return slog.StringValue(c.String())
}
