package entities

import "fmt"

// Total represents the total amount of entities that exist in a paginated request
type Total int

type Version struct {
	Version string `json:"version"`
}

func (v Version) String() string {
	return fmt.Sprintf("%s\n", v.Version)
}

type SearchDateOperator string

const (
	EQ  SearchDateOperator = "$eq"
	NE  SearchDateOperator = "$ne"
	GTE SearchDateOperator = "$gte"
	GT  SearchDateOperator = "$gt"
	LTE SearchDateOperator = "$lte"
	LT  SearchDateOperator = "$lt"
)

type SearchDateFilter struct {
	Date     int                `json:"date"`
	Operator SearchDateOperator `json:"searchDateOperator"`
}

type AppPermissionPatch struct {
	UserId string `json:"userId"`
	// Use AppPermission.NONE to remove access
	Permission AppPermission `json:"permission"`
}

type AppCreation struct {
	Name string `json:"name"`
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
