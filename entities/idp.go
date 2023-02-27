package entities

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
