package entities

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
