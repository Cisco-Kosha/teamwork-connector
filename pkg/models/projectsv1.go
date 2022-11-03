package models

import "time"

type ProjectResponseV1 struct {
	Status   string      `json:"STATUS,omitempty"`
	Message  string      `json:"MESSAGE,omitempty"`
	Projects []ProjectV1 `json:"projects,omitempty"`
}

type SingleProjectResponseV1 struct {
	Status  string    `json:"STATUS,omitempty"`
	Message string    `json:"MESSAGE,omitempty"`
	Project ProjectV1 `json:"project,omitempty"`
}

type ProjectV1 struct {
	StartDate      string    `json:"startDate,omitempty"`
	LastChangedOn  time.Time `json:"last-changed-on,omitempty"`
	Logo           string    `json:"logo,omitempty"`
	CreatedOn      time.Time `json:"created-on,omitempty"`
	PrivacyEnabled bool      `json:"privacyEnabled,omitempty"`
	Status         string    `json:"status,omitempty"`
	BoardData      struct {
	} `json:"boardData,omitempty"`
	ReplyByEmailEnabled  bool   `json:"replyByEmailEnabled,omitempty"`
	HarvestTimersEnabled bool   `json:"harvest-timers-enabled,omitempty"`
	Description          string `json:"description,omitempty"`
	Category             struct {
		Color string `json:"color,omitempty"`
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"category,omitempty"`
	ID                string       `json:"id,omitempty"`
	OverviewStartPage string       `json:"overview-start-page,omitempty"`
	StartPage         string       `json:"start-page,omitempty"`
	Integrations      Integrations `json:"integrations,omitempty"`
	Defaults          struct {
		Privacy string `json:"privacy,omitempty"`
	} `json:"defaults,omitempty"`
	Notifyeveryone      bool   `json:"notifyeveryone,omitempty"`
	FilesAutoNewVersion bool   `json:"filesAutoNewVersion,omitempty"`
	DefaultPrivacy      string `json:"defaultPrivacy,omitempty"`
	TasksStartPage      string `json:"tasks-start-page,omitempty"`
	Starred             bool   `json:"starred,omitempty"`
	AnnouncementHTML    string `json:"announcementHTML,omitempty"`
	IsProjectAdmin      bool   `json:"isProjectAdmin,omitempty"`
	Name                string `json:"name,omitempty"`
	Company             struct {
		IsOwner string `json:"is-owner,omitempty"`
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"company,omitempty"`
	EndDate          string        `json:"endDate,omitempty"`
	ActivePages      ActivePagesV1 `json:"active-pages,omitempty"`
	Announcement     string        `json:"announcement,omitempty"`
	ShowAnnouncement bool          `json:"show-announcement,omitempty"`
	SubStatus        string        `json:"subStatus,omitempty"`
	Tags             []interface{} `json:"tags,omitempty"`
	Owner            Owner         `json:"owner,omitempty"`
}

type Owner struct {
	AvatarURL string `json:"avatarUrl,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	FullName  string `json:"fullName,omitempty"`
	Id        string `json:"id,omitempty"`
	LastName  string `json:"basecurrency,omitempty"`
}

type Integrations struct {
	Xero struct {
		Basecurrency string `json:"basecurrency,omitempty"`
		Countrycode  string `json:"countrycode,omitempty"`
		Enabled      bool   `json:"enabled,omitempty"`
		Connected    string `json:"connected,omitempty"`
		Organisation string `json:"organisation,omitempty"`
	} `json:"xero,omitempty"`
	Sharepoint struct {
		Account    string `json:"account,omitempty"`
		Foldername string `json:"foldername,omitempty"`
		Enabled    bool   `json:"enabled,omitempty"`
		Folder     string `json:"folder,omitempty"`
	} `json:"sharepoint,omitempty"`
	MicrosoftConnectors struct {
		Enabled bool `json:"enabled,omitempty"`
	} `json:"microsoftConnectors,omitempty"`
	Onedrivebusiness struct {
		Account    string `json:"account,omitempty"`
		Foldername string `json:"foldername,omitempty"`
		Enabled    bool   `json:"enabled,omitempty"`
		Folder     string `json:"folder,omitempty"`
	} `json:"onedrivebusiness,omitempty"`
}

type ActivePagesV1 struct {
	Links        string `json:"links,omitempty"`
	Tasks        string `json:"tasks,omitempty"`
	Time         string `json:"time,omitempty"`
	Billing      string `json:"billing,omitempty"`
	Notebooks    string `json:"notebooks,omitempty"`
	Files        string `json:"files,omitempty"`
	Comments     string `json:"comments,omitempty"`
	RiskRegister string `json:"riskRegister,omitempty"`
	Milestones   string `json:"milestones,omitempty"`
	Messages     string `json:"messages,omitempty"`
}

type ProjectUpdateResponseV1 struct {
	Update []ProjectUpdateV1 `json:"update"`
	Status string            `json:"STATUS"`
}

type ProjectUpdateV1 struct {
	DeletedDate string    `json:"deletedDate"`
	Text        string    `json:"text"`
	Health      string    `json:"health"`
	ID          string    `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	ProjectID   string    `json:"projectId"`
	Color       string    `json:"color"`
	Deleted     bool      `json:"deleted"`
	User        struct {
		AvatarURL string `json:"avatarUrl"`
		FirstName string `json:"firstName"`
		FullName  string `json:"fullName"`
		ID        string `json:"id"`
		LastName  string `json:"lastName"`
	} `json:"user"`
	ProjectStatus string `json:"projectStatus"`
}
