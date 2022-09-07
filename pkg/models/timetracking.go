package models

type ReturnedTimeEntries struct {
	TimeEntries []ReturnedTimeEntryObject `json:"time-entries,omitempty"`
	Status string `json:"STATUS,omitempty"`
}

type CreateTimeEntry struct {
	TimeEntry TimeEntry `json:"time-entry, omitempty"`
}

type TimeEntry struct {
	Description string `json:"description,omitempty"`
	PersonId int `json:"person-id,omitempty"`
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
	Hours int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`
	IsBillable bool `json:"isbillable,omitempty"`
	Tags string `json:"tags,omitempty"`
}

type CreatedTimeEntry struct {
	timeLogId string
	STATUS string
}

type ReturnedTimeEntryObject struct {
	ProjectId string `json:"project-id,omitempty"`
	IsBillable string `json:"isbillable,omitempty"`
	// TaskListId string `json:"tasklistId,omitempty"`
	TodoListName string `json:"todo-list-name,omitempty"`
	TodoItemName string `json:"todo-item-name,omitempty"`
	// IsBilled string `json:"isbilled,omitempty"`
	// UpdatedDate string `json:"updated-date,omitempty"`
	TodoListId string `json:"todo-list-id,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	// CanEdit bool `json:"canEdit,omitempty"`
	// TaskEstimatedTime string `json:"taskEstimatedTime,omitempty"`
	CompanyName string `json:"company-name,omitempty"`
	Id string `json:"id,omitempty"`
	InvoiceNo string `json:"invoiceNo,omitempty"`
	PersonLastName string `json:"person-last-name,omitempty"`
	// ParentTaskName string `json:"parentTaskName,omitempty"`
	// DateUserPerspective string `json:"dateUserPerspective,omitempty"`
	Minutes string `json:"minutes,omitempty"`
	PersonFirstName string `json:"person-first-name,omitempty"`
	Description string `json:"description,omitempty"`
	// TicketId string `json:"ticket-id,omitempty"`
	// CreatedAt string `json:"createdAt,omitempty"`
	// TaskIsPrivate string `json:"taskIsPrivate,omitempty"`
	// ParentTaskId string `json:"parentTaskId,omitempty"`
	CompanyId string `json:"company-id,omitempty"`
	ProjectStatus string `json:"project-status,omitempty"`
	PersonId string `json:"person-id,omitempty"`
	ProjectName string `json:"project-name,omitempty"`
	// TaskTags interface{} `json:"task-tags,omitempty"`
	// TaskIsSubTask string `json:"taskIsSubTask,omitempty"`
	TodoItemId string `json:"todo-item-id,omitempty"`
	Date string `json:"date,omitempty"`
	HasStartTime string `json:"has-start-time,omitempty"`
	Hours string `json:"hours,omitempty"`
}