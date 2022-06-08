package entities

type TeamPatchUserOperation string

const (
	ADD    TeamPatchUserOperation = "ADD"
	REMOVE TeamPatchUserOperation = "REMOVE"
	SET    TeamPatchUserOperation = "SET"
)

type TeamPatch struct {
	Name          string                 `json:"name"`
	Users         []string               `json:"users"`
	UserOperation TeamPatchUserOperation `json:"userOperation"`
}
type TeamCreation struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
type Team struct {
	Id             string   `json:"_id"`
	Name           string   `json:"name"`
	OrganizationId string   `json:"organizationId"`
	Users          []string `json:"users"`
	CreateDate     int      `json:"createDate"`
	ModifyDate     int      `json:"modifyDate"`
}

type UserEmail struct {
	Email string `json:"email"`
}
type User struct {
	Id                 string `json:"_id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Company            string `json:"company,omitempty"`
	ActiveOrganization string `json:"activeOrganization"`
}

type Token struct {
	Token string `json:"token"`
}

type PatchUser struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Company string `json:"company,omitempty"`
	// ActiveOrganization holds the name for the active organization, mainly being used within the web frontend
	// Changing this value makes all web frontend reload and act in the context of the new active organization
	ActiveOrganization string `json:"activeOrganization,omitempty"`
}

type SearchOrganizationFilter struct {
	Name               string           `json:"name,omitempty"`
	CreatorId          string           `json:"creatorId,omitempty"`
	Company            string           `json:"company,omitempty"`
	IsUserOrganization bool             `json:"isUserOrganization,omitempty"`
	CreateDateFilter   SearchDateFilter `json:"createDateFilter,omitempty"`
	ModifyDateFilter   SearchDateFilter `json:"modifyDateFilter,omitempty"`
}

type Organization struct {
	Id                 string                 `json:"_id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Creator            *User                  `json:"creator,omitempty"`
	Company            string                 `json:"company,omitempty"`
	IsUserOrganization bool                   `json:"isUserOrganization"`
	CreateDate         int                    `json:"createDate"`
	ModifyDate         int                    `json:"modifyDate"`
	Permission         OrganizationPermission `json:"permission,omitempty"`
}

type OrganizationMember struct {
	Id             string                 `json:"_id"`
	OrganizationId string                 `json:"organizationId"`
	User           *User                  `json:"user"`
	Permission     OrganizationPermission `json:"permission"`
	CreateDate     int                    `json:"createDate"`
	ModifyDate     int                    `json:"modifyDate"`
}

type OrganizationMemberCreation struct {
	Email      string                 `json:"email"`
	Permission OrganizationPermission `json:"permission"`
}

type PatchOrganizationMemberPermission struct {
	Permission OrganizationPermission `json:"permission"`
}

type OrganizationPermission string

const (
	ORGANIZATION_MEMBER OrganizationPermission = "MEMBER"
	ORGANIZATION_MANAGE OrganizationPermission = "MANAGE"
	ORGANIZATION_ADMIN  OrganizationPermission = "ADMIN"
	ORGANIZATION_OWNER  OrganizationPermission = "OWNER"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Totp     string `json:"totp,omitempty"`
}

type Register struct {
	Name     string `json:"name"`
	Company  string `json:"company,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type RegisterVerify struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verificationCode"`
}

type UserPatPatch struct {
	Name   string  `json:"name"`
	Scopes []Scope `json:"scopes"`
}

type UserPatCreation struct {
	Name       string  `json:"name"`
	Expiration int     `json:"expiration"`
	Scopes     []Scope `json:"scopes"`
}

type UserPat struct {
	Id         string      `json:"_id"`
	Name       string      `json:"name"`
	Expiration int         `json:"expiration"`
	ModifyDate int         `json:"modifyDate"`
	Scopes     []Scope     `json:"scopes"`
	Type       UserPatType `json:"type"`
	UserId     string      `json:"userId"`
}

type UserPatType string

const (
	PAT   UserPatType = "PAT"
	OAUTH UserPatType = "OAUTH"
)

type Scope string

const (
	HCLOUD_FULL             Scope = "hcloud:full"
	IDP_EMAIL_READ          Scope = "idp:email:read"
	IDP_USER_READ           Scope = "idp:user:read"
	IDP_USER_WRITE          Scope = "idp:user:write"
	IDP_USER_DELETE         Scope = "idp:user:delete"
	IDP_ORGANIZATION_READ   Scope = "idp:organization:read"
	IDP_ORGANIZATION_WRITE  Scope = "idp:organization:write"
	IDP_ORGANIZATION_DELETE Scope = "idp:organization:delete"
	HIGH5_APP_READ          Scope = "high5:app:read"
	HIGH5_APP_EXECUTE       Scope = "high5:app:execute"
	HIGH5_APP_WRITE         Scope = "high5:app:write"
	HIGH5_APP_DELETE        Scope = "high5:app:delete"
	FUSE_APP_READ           Scope = "fuse:app:read"
	FUSE_APP_EXECUTE        Scope = "fuse:app:execute"
	FUSE_APP_WRITE          Scope = "fuse:app:write"
	FUSE_APP_DELETE         Scope = "fuse:app:delete"
)

type TOTP struct {
	Activated  bool   `json:"activated"`
	OTPAuthUrl string `json:"otpAuthUrl"`
	QRCode     string `json:"qrcode"`
	CreateDate int    `json:"createDate"`
	ModifyDate int    `json:"modifyDate"`
}

type TOTPActivated struct {
	Activated    bool     `json:"activated"`
	RecoverCodes []string `json:"recoverCode"`
	CreateDate   int      `json:"createDate"`
	ModifyDate   int      `json:"modifyDate"`
}

type TOTPToken struct {
	Token string `json:"token"`
}

type OAuthApplicationClientSecretCreation struct {
	SecretName string `json:"secretName"`
}
type OAuthApplicationClientSecret struct {
	Id         string `json:"_id"`
	Name       string `json:"name"`
	Secret     string `json:"secret"`
	Uuid       string `json:"uuid"`
	CreatorId  string `json:"creatorId"`
	Used       bool   `json:"used"`
	CreateDate int    `json:"createDate"`
	ModifyDate int    `json:"modifyDate"`
}

type OAuthApplication struct {
	Id          string   `json:"_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Avatar      string   `json:"avatar"`
	Homepage    string   `json:"homepage"`
	Callback    []string `json:"callback"`

	ClientId     string                         `json:"clientId"`
	ClientSecret []OAuthApplicationClientSecret `json:"clientSecret"`

	OrganizationId string `json:"organizationId"`
	CreatorId      string `json:"creatorId"`
}

type OAuthApplicationCreation struct {
	*OAuthApplicationUpdate
	*OAuthApplicationClientSecretCreation
}

type OAuthApplicationUpdate struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Avatar      string   `json:"avatar"`
	Homepage    string   `json:"homepage"`
	Callback    []string `json:"callback"`
}
