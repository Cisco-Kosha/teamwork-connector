package models

import "time"

type MultiActivity struct {
	Activities []Activity       `json:"activities,omitempty"`
	Metadata   Meta             `json:"meta,omitempty"`
	Included   ActivityIncluded `json:"included,omitempty"`
}

type Activity struct {
	ActivityType       string       `json:"activityType,omitempty"`
	Company            Relationship `json:"company,omitempty"`
	CompanyId          int          `json:"companyId,omitempty"`
	DateTime           string       `json:"dateTime,omitempty"`
	Description        string       `json:"description,omitempty"`
	DueDate            string       `json:"dueDate,omitempty"`
	ExtraDescription   string       `json:"extraDescription,omitempty"`
	ForUser            Relationship `json:"forUser,omitempty"`
	ForUserId          int          `json:"forUserId,omitempty"`
	ForUserName        string       `json:"forUserName,omitempty"`
	Id                 int          `json:"id,omitempty"`
	IsPrivate          int          `json:"isPrivate,omitempty"`
	Item               Relationship `json:"item,omitempty"`
	ItemId             int          `json:"itemId,omitempty"`
	ItemLink           string       `json:"itemLink,omitempty"`
	LatestActivityType string       `json:"latestActivityType,omitempty"`
	Link               string       `json:"link,omitempty"`
	Lockdown           Relationship `json:"lockdown,omitempty"`
	LockdownId         int          `json:"lockdownId,omitempty"`
	Project            Relationship `json:"project,omitempty"`
	ProjectId          int          `json:"projectId,omitempty"`
	PublicInfo         string       `json:"publicInfo,omitempty"`
	Type               string       `json:"type,omitempty"`
	User               Relationship `json:"user,omitempty"`
	UserId             int          `json:"userId,omitempty"`
}

type ActivityIncluded struct {
	Companies interface{} `json:"companies,omitempty"`
	Projects  interface{} `json:"projects,omitempty"`
	Users     interface{} `json:"users,omitempty"`
}

type ActivityResponseV1 struct {
	Activity []ActivityV1 `json:"activity"`
	Message  string       `json:"MESSAGE,omitempty"`
	Status   string       `json:"STATUS"`
}

type ActivityV1 struct {
	ProjectID         string    `json:"project-id"`
	Itemid            string    `json:"itemid"`
	TodoListName      string    `json:"todo-list-name"`
	FromUserAvatarURL string    `json:"from-user-avatar-url"`
	Description       string    `json:"description"`
	Forusername       string    `json:"forusername"`
	Publicinfo        string    `json:"publicinfo"`
	Foruserid         string    `json:"foruserid"`
	Itemlink          string    `json:"itemlink"`
	Datetime          time.Time `json:"datetime"`
	Activitytype      string    `json:"activitytype"`
	ProjectName       string    `json:"project-name"`
	Link              string    `json:"link"`
	Extradescription  string    `json:"extradescription"`
	Isprivate         string    `json:"isprivate"`
	ID                string    `json:"id"`
	DueDate           string    `json:"due-date"`
	Fromusername      string    `json:"fromusername"`
	Type              string    `json:"type"`
	ForUserAvatarURL  string    `json:"for-user-avatar-url"`
	Userid            string    `json:"userid"`
}
