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

type SpacePermissionPatch struct {
	UserId string `json:"userId"`
	// Use AppPermission.NONE to remove access
	Permission SpacePermission `json:"permission"`
}

type SpaceCreation struct {
	Name string `json:"name"`
}
type Space struct {
	Id           string              `json:"_id,omitempty"`
	Name         string              `json:"name,omitempty"`
	Permissions  []Permission        `json:"permissions,omitempty"`
	Organization ReducedOrganization `json:"organization,omitempty"`
	Creator      ReducedUser         `json:"creator,omitempty"`
	CreateDate   int                 `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

type SpacePermission string

const (
	NONE    SpacePermission = "NONE"
	READ    SpacePermission = "READ"
	EXECUTE SpacePermission = "EXECUTE"
	WRITE   SpacePermission = "WRITE"
	MANAGE  SpacePermission = "MANAGE"
	OWNER   SpacePermission = "OWNER"
)

type Permission struct {
	UserId     string          `json:"userId,omitempty"`
	Permission SpacePermission `json:"permission,omitempty"`
}

func (p Permission) String() string {
	return fmt.Sprintf("     UserId:         %s\n      Permissions:    %s\n", p.UserId, p.Permission)
}

type Search struct {
	Filters []SearchFilter `json:"filters,omitempty"`
	Sorting *Sorting       `json:"sorting,omitempty"`
}

type Comparator string

const (
	IS                    Comparator = "is"
	IS_NOT                Comparator = "is not"
	STARTS_WITH           Comparator = "starts with"
	ENDS_WITH             Comparator = "ends with"
	CONTAINS              Comparator = "contains"
	GREATER_THAN          Comparator = "greater than"
	LESS_THAN             Comparator = "less than"
	GREATER_THAN_OR_EQUAL Comparator = "greater than or equal"
	LESS_THAN_OR_EQUAL    Comparator = "less than or equal"
)

type SearchFilter struct {
	Type       SearchFilterType
	Key        string
	Comparator Comparator
	Value      any
}

type SearchFilterType string

const (
	STRING      SearchFilterType = "STRING"
	SELECT      SearchFilterType = "SELECT"
	MULTISELECT SearchFilterType = "MULTISELECT"
	TYPEAHEAD   SearchFilterType = "TYPEAHEAD"
	NUMBER      SearchFilterType = "NUMBER"
	DATE        SearchFilterType = "DATE"
	BOOLEAN     SearchFilterType = "BOOLEAN"
)

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type Sorting struct {
	Direction SortDirection `json:"direction,omitempty"`
	Fields    []string      `json:"fields,omitempty"`
}
