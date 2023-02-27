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
