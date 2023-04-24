package entities

type FuseApp struct {
	*App
}

type CronjobExpression struct {
	Expression string `json:"expression"`
}

type CronjobCreation struct {
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Expression       string            `json:"expression,omitempty"`
	Target           string            `json:"targetUrl,omitempty"`
	Method           CronjobHttpMethod `json:"httpMethod,omitempty"`
	AcceptInvalidSSL bool              `json:"acceptInvalidSSL,omitempty"`
	Timezone         string            `json:"timezone,omitempty"`
	Headers          []CronjobHeader   `json:"headers,omitempty"`
	Body             string            `json:"body,omitempty"`
	Enabled          bool              `json:"enabled,omitempty"`
}

type Cronjob struct {
	Id               string            `json:"_id,omitempty"`
	Enabled          bool              `json:"enabled,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Expression       string            `json:"expression,omitempty"`
	Target           string            `json:"targetUrl,omitempty"`
	Method           CronjobHttpMethod `json:"httpMethod,omitempty"`
	AcceptInvalidSSL bool              `json:"acceptInvalidSSL,omitempty"`
	Timezone         string            `json:"timezone,omitempty"`
	Creator          string            `json:"creatorId,omitempty"`
	CreateDate       int               `json:"createDate,omitempty"`
	ModifyDate       int               `json:"modifyDate,omitempty"`
	AppId            string            `json:"appId,omitempty"`
	OrganizationId   string            `json:"organizationId,omitempty"`
	Headers          []CronjobHeader   `json:"headers,omitempty"`
	Body             string            `json:"body,omitempty"`
}

type ReducedOrganization struct {
	Id   string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReducedUser struct {
	Id   string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CronjobLogCreation struct {
	StatusCode int             `json:"statusCode,omitempty"`
	Headers    []CronjobHeader `json:"headers,omitempty"`
	Body       string          `json:"body,omitempty"`
}

type CronjobLog struct {
	Id           string              `json:"_id,omitempty"`
	CronjobId    string              `json:"cronjobId,omitempty"`
	StatusCode   int                 `json:"statusCode,omitempty"`
	Organization ReducedOrganization `json:"organization,omitempty"`
	Creator      ReducedUser         `json:"creator,omitempty"`
	CreatorId    string              `json:"creatorId,omitempty"`
	App          string              `json:"app,omitempty"`
	Headers      []CronjobHeader     `json:"headers,omitempty"`
	Body         string              `json:"body,omitempty"`
	CreateDate   int                 `json:"createDate,omitempty"`
	ModifyDate   int                 `json:"modifyDate,omitempty"`
}

type CronjobHttpMethod string

const (
	GET    CronjobHttpMethod = "GET"
	POST   CronjobHttpMethod = "POST"
	PUT    CronjobHttpMethod = "PUT"
	PATCH  CronjobHttpMethod = "PATCH"
	DELETE CronjobHttpMethod = "DELETE"
)

type CronjobHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
