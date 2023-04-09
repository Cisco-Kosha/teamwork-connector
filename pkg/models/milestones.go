package models

import "time"

type ReturnedMilestones struct {
	Milestones []Milestone `json:"milestones,omitempty"`
	Meta struct {
		Page struct {
			PageOffset int  `json:"pageOffset,omitempty"`
			PageSize   int  `json:"pageSize,omitempty"`
			Count      int  `json:"count,omitempty"`
			HasMore    bool `json:"hasMore,omitempty"`
		} `json:"page,omitempty"`
	} `json:"meta,omitempty"`
	Included Included `json:"included,omitempty"`
}

type ReturnedMilestone struct {
	Milestone Milestone `json:"milestones,omitempty"`
	Included Included `json:"included,omitempty"`
}

type Milestone struct {
	ID              int       `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	DescriptionHTML string    `json:"descriptionHTML,omitempty"`
	Deadline        time.Time `json:"deadline,omitempty"`
	OriginalDueDate time.Time `json:"originalDueDate,omitempty"`
	Completed       bool      `json:"completed,omitempty"`
	ProjectID       int       `json:"projectId,omitempty"`
	Project         struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"project"`
	CreatedOn             time.Time `json:"createdOn,omitempty"`
	LastChangedOn         time.Time `json:"lastChangedOn,omitempty"`
	UpdatedBy             int       `json:"updatedBy,omitempty"`
	CreatorUserID         int       `json:"creatorUserId,omitempty"`
	CreatedBy             int       `json:"createdBy,omitempty"`
	Reminder              bool      `json:"reminder,omitempty"`
	Private               bool      `json:"private,omitempty"`
	Status                string    `json:"status,omitempty"`
	UserFollowingComments bool      `json:"userFollowingComments,omitempty"`
	UserFollowingChanges  bool      `json:"userFollowingChanges,omitempty"`
	IsDeleted             bool      `json:"isDeleted,omitempty"`
	CanEdit               bool      `json:"canEdit,omitempty"`
	CanComplete           bool      `json:"canComplete,omitempty"`
	ResponsiblePartyIds   []int     `json:"responsiblePartyIds,omitempty"`
	ResponsibleParties    []struct {
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"responsibleParties,omitempty"`
	CommentsCount int   `json:"commentsCount,omitempty"`
	TagIds        []int `json:"tagIds,omitempty"`
	Tags          []struct {
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"tags,omitempty"`
	TasklistIds []int `json:"tasklistIds,omitempty"`
	Tasklists   []struct {
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"tasklists,omitempty"`
}

type Included struct {
	Tags struct {
		Num25488 struct {
			ID        int `json:"id,omitempty"`
			ProjectID int `json:"projectId,omitempty"`
			Project   struct {
				ID   int    `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"project,omitempty"`
			Name  string `json:"name,omitempty"`
			Color string `json:"color,omitempty"`
		} `json:"25488,omitempty"`
		Num30993 struct {
			ID        int `json:"id,omitempty"`
			ProjectID int `json:"projectId,omitempty"`
			Project   struct {
				ID   int    `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"project,omitempty"`
			Name  string `json:"name,omitempty"`
			Color string `json:"color,omitempty"`
		} `json:"30993,omitempty"`
	} `json:"tags,omitempty"`
}