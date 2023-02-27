package hcloud

import (
	"fmt"
)

type Token struct {
	Token string `json:"token"`
}

type User struct {
	Id                   string `json:"_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
	Company              string `json:"company,omitempty"`
	ActiveOrganizationId string `json:"activeOrganizationId,omitempty"`
}

type PatchUser struct {
	Name                 string `json:"name,omitempty"`
	Email                string `json:"email,omitempty"`
	Company              string `json:"company,omitempty"`
	ActiveOrganizationId string `json:"activeOrganizationId,omitempty"`
}

type Version struct {
	Version string `json:"version"`
}

func (v Version) String() string {
	return fmt.Sprintf("%s\n", v.Version)
}

type Organization struct {
	Id      string `json:"_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Creator *User  `json:"creator,omitempty"`
	Company string `json:"company,omitempty"`
}

type OrganizationMember struct {
	Id             string                 `json:"_id,omitempty"`
	OrganizationId string                 `json:"organizationId,omitempty"`
	User           *User                  `json:"user,omitempty"`
	Permission     OrganizationPermission `json:"permission,omitempty"`
}

type AddOrganizationMember struct {
	Email      string                 `json:"email,omitempty"`
	Permission OrganizationPermission `json:"permission,omitempty"`
}

type OrganizationPermission string

const (
	ORGANIZATION_READ   OrganizationPermission = "READ"
	ORGANIZATION_MANAGE OrganizationPermission = "MANAGE"
	ORGANIZATION_ADMIN  OrganizationPermission = "ADMIN"
	ORGANIZATION_OWNER  OrganizationPermission = "OWNER"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Totp     string `json:"totp"`
}

type Register struct {
	Name     string `json:"name"`
	Company  string `json:"company,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuditLog struct {
	Origin    string      `json:"origin"`
	Level     string      `json:"level"`
	Event     string      `json:"event"`
	Type      string      `json:"type"`
	User      string      `json:"user"`
	Timestamp int64       `json:"timestamp"`
	Message   interface{} `json:"message"`
}

type App struct {
	Id             string       `json:"_id,omitempty"`
	Name           string       `json:"name,omitempty"`
	Permissions    []Permission `json:"permissions,omitempty"`
	OrganizationId string       `json:"organizationId,omitempty"`
	CreatorId      string       `json:"creatorId,omitempty"`
	CreateDate     int          `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

type AppPermission string

const (
	NONE    AppPermission = "NONE"
	READ    AppPermission = "READ"
	EXECUTE AppPermission = "EXECUTE"
	WRITE   AppPermission = "WRITE"
	MANAGE  AppPermission = "MANAGE"
	OWNER   AppPermission = "OWNER"
)

type Permission struct {
	UserId     string        `json:"userId,omitempty"`
	Permission AppPermission `json:"permission,omitempty"`
}

func (p Permission) String() string {
	return fmt.Sprintf("     UserId:         %s\n      Permissions:    %s\n", p.UserId, p.Permission)
}

type Event struct {
	Id             string `json:"_id,omitempty"`
	Name           string `json:"name,omitempty"`
	AppId          string `json:"appId,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
	CreatorId      string `json:"creatorId,omitempty"`
	CreateDate     int    `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

type Stream struct {
	Id             string `json:"_id,omitempty"`
	Name           string `json:"name,omitempty"`
	EventId        string `json:"eventId,omitempty"`
	AppId          string `json:"appId,omitempty"`
	Order          int    `json:"order,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
	CreatorId      string `json:"creatorId,omitempty"`
	CreateDate     int    `json:"createDate,omitempty"` // UTC+0 unix timestamp
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

type StreamExecutionPayload struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type EventExecutionRequest struct {
	EventName     string `json:"eventName"`
	Target        string `json:"target"`
	Payload       StreamExecutionPayload
	Timeout       int  `json:"timeout"`
	WaitForResult bool `json:"waitForResult"`
}

type Webhook struct {
	Id              string                 `json:"_id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Url             string                 `json:"url,omitempty"`
	Token           string                 `json:"token,omitempty"`
	Target          string                 `json:"target"`
	SecurityHeaders WebhookSecurityHeaders `json:"securityHeaders,omitempty"`
	AppId           string                 `json:"appId,omitempty"`
	EventId         string                 `json:"eventId,omitempty"`
	OrganizationId  string                 `json:"organizationId,omitempty"`
	CreatorId       string                 `json:"creatorId,omitempty"`
	CreateDate      int                    `json:"createDate,omitempty"` // UTC+0 unix timestamp
	ModifyDate      int                    `json:"modifyDate,omitempty"` // UTC+0 unix timestamp
}

type WebhookCreation struct {
	Name            string                 `json:"name"`
	Token           string                 `json:"token"`
	EventId         string                 `json:"eventId"`
	AppId           string                 `json:"appId"`
	Target          string                 `json:"target"`
	SecurityHeaders WebhookSecurityHeaders `json:"securityHeaders,omitempty"`
}

type WebhookSecurityHeaders map[string]string

type WebhookLog struct {
	Id                 string      `json:"_id,omitempty"`
	WebhookId          string      `json:"webhookId,omitempty"`
	AppId              string      `json:"appId,omitempty"`
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

type Cronjob struct {
	Id               string          `json:"_id,omitempty"`
	Enabled          bool            `json:"enabled,omitempty"`
	Name             string          `json:"name,omitempty"`
	Expression       string          `json:"expression,omitempty"`
	Target           string          `json:"targetUrl,omitempty"`
	Method           string          `json:"httpMethod,omitempty"`
	AcceptInvalidSSL bool            `json:"acceptInvalidSSL,omitempty"`
	Timezone         string          `json:"timezone,omitempty"`
	Creator          string          `json:"creatorId,omitempty"`
	CreateDate       int             `json:"createDate,omitempty"`
	ModifyDate       int             `json:"modifyDate,omitempty"`
	AppId            string          `json:"appId,omitempty"`
	OrganizationId   string          `json:"organizationId,omitempty"`
	Headers          []CronjobHeader `json:"headers,omitempty"`
	Body             string          `json:"body,omitempty"`
}

type CronjobHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
