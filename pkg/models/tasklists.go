package models

type ProjectsTasklists struct {
	Tasklists []Tasklist
}

type MultiTaskList struct {
	Status    string     `json:"STATUS,omitempty"`
	Tasklists []Tasklist `json:"tasklists,omitempty"`
}

type Tasklist struct {
	Name             string      `json:"name,omitempty"`
	Pinned           bool        `json:"pinned,omitempty"`
	MilestoneId      string      `json:"milestone-id,omitempty"`
	Description      string      `json:"description,omitempty"`
	UncompletedCount int         `json:"uncompleted-count,omitempty"`
	Id               string      `json:"id,omitempty"`
	Complete         bool        `json:"complete,omitempty"`
	Private          bool        `json:"private,omitempty"`
	isTemplate       bool        `json:"isTemplate,omitempty"`
	Position         int         `json:"position,omitempty"`
	Status           string      `json:"status,omitempty"`
	ProjectId        string      `json:"projectId,omitempty"`
	ProjectName      string      `json:"projectName,omitempty"`
	DLM              interface{} `json:"DLM,omitempty"`
}

type NewTaskList struct {
	ApplyDefaultsToExistingTasks bool     `json:"applyDefaultToExistingTasks,omitempty"`
	ToDoList                     ToDoList `json:"todo-list,omitempty"`
}

type Tasks struct {
	Status    string     `json:"STATUS,omitempty"`
	ToDoItems []ToDoItem `json:"todo-items,omitempty"`
}

type ToDoList struct {
	NewTaskDefaults struct {
		Description        string `json:"description,omitempty"`
		StartDateOffset    string `json:"start-date-offset,omitempty"`
		DueDateOffset      string `json:"due-date-offset,omitempty"`
		ResponsiblePartyId string `json:"responsible-party-id,omitempty"`
		Priority           int    `json:"priority,omitempty"`
		PriorityText       string `json:"priorityText,omitempty"`
		EstimatedMinutes   int    `json:"estimated-minutes,omitempty"`
		Tags               struct {
			Id        int    `json:"id,omitempty"`
			Color     string `json:"color,omitempty"`
			Name      string `json:"name,omitempty"`
			ProjectId int    `json:"projectId,omitempty"`
		} `json:"tags,omitempty"`
		ColumnId  int `json:"column-id,omitempty"`
		Reminders struct {
			UserId             int    `json:"user-id,omitempty"`
			Type               string `json:"type,omitempty"`
			Note               string `json:"note,omitempty"`
			PeopleAssigned     bool   `json:"people-assigned,omitempty"`
			IsRelative         bool   `json:"isRelative,omitempty"`
			RelativeNumberDays int    `json:"relative-number-days,omitempty"`
			UsingOffsetDueDate bool   `json:"usingOffSetDueDate,omitempty"`
			Time               string `json:"time,omitempty"`
		} `json:"reminders,omitempty"`
		RemoveAllReminders     bool        `json:"removeAllReminders,omitempty"`
		CommentFollowerIds     string      `json:"comment-follower-ids,omitempty"`
		ChangeFollowerIds      string      `json:"change-follower-ids,omitempty"`
		GrantAccessTo          string      `json:"grant-access-to,omitempty"`
		Private                bool        `json:"private,omitempty"`
		CustomFields           interface{} `json:"customFields,omitempty"`
		PendingFileAttachments interface{} `json:"pendingFileAttachments,omitempty"`
	} `json:"new-task-defaults,omitempty"`
}

type ToDoItem struct {
	Id                        int         `json:"id,omitempty"`
	CanComplete               bool        `json:"canComplete,omitempty"`
	CommentsCount             int         `json:"comments-count,omitempty"`
	Description               string      `json:"description,omitempty"`
	HasReminders              bool        `json:"has-reminders,omitempty"`
	HasUnreadComments         bool        `json:"has-unread-comments,omitempty"`
	Private                   int         `json:"private,omitempty"`
	Content                   string      `json:"content,omitempty"`
	Order                     int         `json:"order,omitempty"`
	ProjectId                 int         `json:"project-id,omitempty"`
	ProjectName               string      `json:"project-name,omitempty"`
	TodoListId                int         `json:"todo-list-id,omitempty"`
	ToDoListName              string      `json:"todo-list-name,omitempty"`
	TasklistPrivate           bool        `json:"tasklist-private,omitempty"`
	TasklistIsTemplate        bool        `json:"tasklist-isTemplate,omitempty"`
	Status                    string      `json:"status,omitempty"`
	CompanyName               string      `json:"company-name,omitempty"`
	CompanyId                 int         `json:"company-id,omitempty"`
	CreatorId                 int         `json:"creator-id,omitempty"`
	CreatorFirstName          string      `json:"creator-firstname,omitempty"`
	CreatorLastName           string      `json:"creator-lastname,omitempty"`
	Completed                 bool        `json:"completed,omitempty"`
	StartDate                 string      `json:"start-date,omitempty"`
	DueDateBase               string      `json:"due-date-base,omitempty"`
	DueDate                   string      `json:"due-date,omitempty"`
	CreatedOn                 string      `json:"created-on,omitempty"`
	LastChangedOn             string      `json:"last-changed-on,omitempty"`
	Position                  int         `json:"position,omitempty"`
	EstimatedMinutes          int         `json:"estimated-minutes,omitempty"`
	Priority                  string      `json:"priority,omitempty"`
	Progress                  int         `json:"progress,omitempty"`
	HarvestEnabled            bool        `json:"harvest-enabled,omitempty"`
	ParentTaskId              string      `json:"parentTaskId,omitempty"`
	LockdownId                string      `json:"lockdownId,omitempty"`
	TasklistLockDownId        string      `json:"tasklist-lockdownId,omitempty"`
	HasDependencies           int         `json:"has-dependencies,omitempty"`
	HasPredecessors           int         `json:"has-predecessors,omitempty"`
	HasTickets                bool        `json:"hasTickets,omitempty"`
	Tags                      interface{} `json:"tags,omitempty"`
	TimeIsLogged              string      `json:"timeIsLogged,omitempty"`
	AttachmentsCount          int         `json:"attachments-count,omitempty"`
	ResponsiblePartyIds       string      `json:"responsible-party-ids,omitempty"`
	ResponsiblePartyId        string      `json:"responsible-party-id,omitempty"`
	ResponsiblePartyNames     string      `json:"responsible-party-names,omitempty"`
	ResponsiblePartyType      string      `json:"responsible-party-type,omitempty"`
	ResponsiblePartyFirstName string      `json:"responsible-party-firstname,omitempty"`
	ResponsiblePartyLastName  string      `json:"responsible-party-lastname,omitempty"`
	ResponsiblePartySummary   string      `json:"responsible-party-summary,omitempty"`
	Predecessors              interface{} `json:"predecessors,omitempty"`
	CanEdit                   bool        `json:"canEdit,omitempty"`
	ViewEstimatedTime         bool        `json:"viewEstimatedTime,omitempty"`
	CreatorAvatarUrl          string      `json:"creator-avatar-url,omitempty"`
	CanLogTime                bool        `json:"canLogTime,omitempty"`
	UserFollowingComments     bool        `userFollowingComments:"type,omitempty"`
	UserFollowingChanges      bool        `json:"userFollowingChanges,omitempty"`
	DLM                       int         `json:"DLM,omitempty"`
}
