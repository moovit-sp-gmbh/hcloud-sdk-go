package entities

type High5Space struct {
	*Space
}

type EventCreation struct {
	Name string `json:"name"`
}

type Event struct {
	Id           string              `json:"_id,omitempty"`
	Name         string              `json:"name,omitempty"`
	SpaceId        string              `json:"spaceId,omitempty"`
	Organization ReducedOrganization `json:"organization,omitempty"`
	Creator      ReducedUser         `json:"creator,omitempty"`
	CreateDate   int                 `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

type StreamCreation struct {
	Name string `json:"name"`
}

type Stream struct {
	Id           string              `json:"_id,omitempty"`
	Name         string              `json:"name,omitempty"`
	EventId      string              `json:"eventId,omitempty"`
	SpaceId        string              `json:"spaceId,omitempty"`
	Order        int                 `json:"order,omitempty"`
	Organization ReducedOrganization `json:"organization,omitempty"`
	Creator      ReducedUser         `json:"creator,omitempty"`
	CreateDate   int                 `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

type StreamOrder struct {
	StreamId string `json:"streamId"`
	Order    int    `json:"order"`
}

type StreamExecutionRequest struct {
	Target        string                 `json:"target"`
	Payload       StreamExecutionPayload `json:"payload"`
	Timeout       int                    `json:"timeout"`
	WaitForResult bool                   `json:"waitForResult"`
}

type StreamPayloadType string

const (
	JSON    StreamPayloadType = "json"
	GENERIC StreamPayloadType = "generic"
)

type StreamExecutionPayload struct {
	Type StreamPayloadType `json:"type"`
	Data string            `json:"data"`
}

type StreamResult struct {
	Failed  bool   `json:"failed"`
	Message string `json:"message"`
}

type Webhook struct {
	Id              string                 `json:"_id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Url             string                 `json:"url,omitempty"`
	Target          string                 `json:"target"`
	SecurityHeaders WebhookSecurityHeaders `json:"securityHeaders,omitempty"`
	SpaceId           string                 `json:"spaceId,omitempty"`
	EventId         string                 `json:"eventId,omitempty"`
	Organization    ReducedOrganization    `json:"organization,omitempty"`
	Creator         ReducedUser            `json:"creator,omitempty"`
	CreateDate      int                    `json:"createDate,omitempty"` // UTC+0 unix timestamp
	ModifyDate      int                    `json:"modifyDate,omitempty"` // UTC+0 unix timestamp
}

type PatchWebhook struct {
	Name            string                 `json:"name"`
	Event           string                 `json:"event"`
	Target          string                 `json:"target"`
	SecurityHeaders WebhookSecurityHeaders `json:"securityHeaders,omitempty"`
}

type WebhookCreation struct {
	Name            string                 `json:"name"`
	Event           string                 `json:"event"`
	Target          string                 `json:"target"`
	SecurityHeaders WebhookSecurityHeaders `json:"securityHeaders,omitempty"`
}

type WebhookSecurityHeaders struct {
}

type WebhookLog struct {
	Id                 string      `json:"_id,omitempty"`
	WebhookId          string      `json:"webhookId,omitempty"`
	SpaceId              string      `json:"spaceId,omitempty"`
	EventId            string      `json:"eventId,omitempty"`
	OrganizationId     string      `json:"organizationId,omitempty"`
	SourceIP           string      `json:"sourceIp,omitempty"`
	CompleteHeader     interface{} `json:"completeHeader,omitempty"`
	RequestBody        interface{} `json:"requestBody,omitempty"`
	ResponseBody       interface{} `json:"responseBody,omitempty"`
	ResponseStatusCode int         `json:"responseStatusCode,omitempty"`
	Timestamp          int         `json:"timestamp,omitempty"`  // UTC+0 unix timestamp
	CreateDate         int         `json:"createDate,omitempty"` // UTC+0 unix timestamp
	ModifyDate         int         `json:"modifyDate,omitempty"` // UTC+0 unix timestamp
}
