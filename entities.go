package hcloud

import (
	"fmt"
	"strings"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

func (t Token) String() string {
	return fmt.Sprintf("%s", t.Token)
}

type User struct {
	Id      string `json:"_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Company string `json:"company,omitempty"`
}

func (u User) String() string {
	return fmt.Sprintf("Id:      %s\nName:    %s\nEmail:   %s\nCompany: %s\n", u.Id, u.Name, u.Email, u.Company)
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
	Creator User   `json:"creator,omitempty"`
	Company string `json:"company,omitempty"`
}

func (o Organization) String() string {
	return fmt.Sprintf("Id:      %s\nName:    %s\nCreator:\n  Id:      %s\n  Name:    %s\n  Email:   %s\n  Company: %s\nCompany: %s\n", o.Id, o.Name, o.Creator.Id, o.Creator.Name, o.Creator.Email, o.Creator.Company, o.Company)
}

type OrganizationMember struct {
	Id             string   `json:"_id,omitempty"`
	OrganizationId string   `json:"organizationId,omitempty"`
	User           User     `json:"userDbRef,omitempty"`
	Roles          []string `json:"roles,omitempty"`
}

func (o OrganizationMember) String() string {
	return fmt.Sprintf("Id:             %s\nOrganizationId: %s\nCreator:\n  Id:      %s\n  Name:    %s\n  Email:   %s\n  Company: %s\nRoles: %s\n", o.Id, o.OrganizationId, o.User.Id, o.User.Name, o.User.Email, o.User.Company, strings.Join(o.Roles, ","))
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l Login) String() string {
	return fmt.Sprintf("Email: %s, Password: %s\n", l.Email, l.Password)
}

type Register struct {
	Name     string `json:"name"`
	Company  string `json:"company,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r Register) String() string {
	return fmt.Sprintf("Name: %s\nCompany: %s\nEmail: %s\nPassword: %s\n", r.Name, r.Company, r.Email, r.Password)
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

func (a AuditLog) String() string {
	return fmt.Sprintf("Origin:     %s\nLevel:      %s\nEvent:      %s\nType:       %s\nUser:       %s\nTime:       %s\nMessage:    %s\n", a.Origin, a.Level, a.Event, a.Type, a.User, time.Unix(a.Timestamp, 0), a.Message)
}
