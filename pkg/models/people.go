package models

type OverallPerson struct {
	Person Person
	Projects []Project
	Permissions Permissions
}

type SinglePerson struct {
	Person Person `json:"person,omitempty"`
	Status string `json:"STATUS,omitempty"`
}

type People struct {
	People []MinimalPerson `json:"people,omitempty"`
	Status string `json:"STATUS,omitempty"`
}

type MinimalPerson struct {
	Id string `json:"id,omitempty"`
	CompanyId string `json:"companyId,omitempty"`
	Title string `json:"title,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
}


type Person struct {
	UserUUID string `json:"userUUID,omitempty"`
	MarketoID string `json:"marketoId,omitempty"`
	Username string `json:"user-name,omitempty"`
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	CompanyId string `json:"companyId,omitempty"`
	CompanyName string `json:"company-name,omitempty"`
	Pid string `json:"pid,omitempty"`
	Profile string `json:"profile,omitempty"`
	ProfileText string `json:"profile-text,omitempty"`
	PhoneNumberOffice string `json:"phone-number-office,omitempty"`
	PhoneNumberOfficeExt string `json:"phone-number-office-ext,omitempty"`
	PhoneNumberMobile string `json:"phone-number-mobile,omitempty"`
	PhoneNumberHome string `json:"phone-number-home,omitempty"`
	EmailAddress string `json:"email-address,omitempty"`
	LengthOfDay string `json:"lengthOfDay,omitempty"`
	UserType string `json:"user-type,omitempty"`
	FirstName string `json:"first-name,omitempty"`
	LastName string `json:"last-name,omitempty"`
	IMService string `json:"im-service,omitempty"`
	IMHandle string `json:"im-handle,omitempty"`
	PrivateNotesText string `json:"private-notes-text,omitempty"`
	PrivateNotes string `json:"private-notes,omitempty"`
	LastLogin string `json:"last-login,omitempty"`
	SiteOwner bool `json:"site-owner,omitempty"`
	Administrator bool `json:"administrator,omitempty"`
	TwoFactorAuthEnabled bool `json:"twoFactorAuthEnabled,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
	CreatedAt string `json:"created-at,omitempty"`
}

type Permissions struct {
	CanManagePeople bool `json:"can-manage-people,omitempty"`
	CanAddProjects bool `json:"can-add-projects,omitempty"`
	CanAccessTemplates bool `json:"can-access-templates,omitempty"`
	CanAccessCalendar bool `json:"canAccessCalendar,omitempty"`
	CanViewProjectTemplates bool `json:"canViewProjectTemplates,omitempty"`
	CanManageCustomFields bool `json:"canManageCustomFields,omitempty"`
	CanManageProjectTemplates bool `json:"canManageProjectTemplates,omitempty"`
}