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

type App struct {
	Id            string       `json:"_id,omitempty"`
	Name          string       `json:"name,omitempty"`
	Permissions   []Permission `json:"permissions,omitempty"`
	OrgnizationId string       `json:"organizationId,omitempty"`
	CreatorId     string       `json:"creatorId,omitempty"`
	CreateDate    int          `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

func (a App) String() string {
	return fmt.Sprintf("Id:              %s\nName:            %s\nPermissions:     \n%s\nOrganizationId:  %s\nCreatorId:       %s\nCreateDate:      %s\n", a.Id, a.Name, a.Permissions, a.OrgnizationId, a.CreatorId, time.Unix(int64(a.CreateDate), 0))
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
	Id            string `json:"_id,omitempty"`
	Name          string `json:"name,omitempty"`
	AppId         string `json:"appId,omitempty"`
	OrgnizationId string `json:"organizationId,omitempty"`
	CreatorId     string `json:"creatorId,omitempty"`
	CreateDate    int    `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

func (e Event) String() string {
	return fmt.Sprintf("Id:              %s\nName:            %s\nAppId:           %s\nOrganizationId:  %s\nCreatorId:       %s\nCreateDate:      %s\n", e.Id, e.Name, e.AppId, e.OrgnizationId, e.CreatorId, time.Unix(int64(e.CreateDate), 0))
}

type Stream struct {
	Id            string `json:"_id,omitempty"`
	Name          string `json:"name,omitempty"`
	EventId       string `json:"eventId,omitempty"`
	AppId         string `json:"appId,omitempty"`
	OrgnizationId string `json:"organizationId,omitempty"`
	CreatorId     string `json:"creatorId,omitempty"`
	CreateDate    int    `json:"createDate,omitempty"` // UTC+0 unix timestamp
}

func (s Stream) String() string {
	return fmt.Sprintf("Id:              %s\nName:            %s\nEventId:         %s\nAppId:           %s\nOrganizationId:  %s\nCreatorId:       %s\nCreateDate:      %s\n", s.Id, s.Name, s.EventId, s.AppId, s.OrgnizationId, s.CreatorId, time.Unix(int64(s.CreateDate), 0))
}

type StreamExecutionRequest struct {
	StreamId      string `json:"streamId"`
	Target        string `json:"target"`
	Data          string `json:"data"`
	Timeout       int    `json:"timeout"`
	WaitForResult bool   `json:"waitForResult"`
}
