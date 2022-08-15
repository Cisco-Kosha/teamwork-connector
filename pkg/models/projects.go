package models

type MultiProject struct {
	Projects []Project `json:"projects,omitempty"`
	Metadata Meta `json:"meta,omitempty"`
	Included ProjectIncluded `json:"included,omitempty"`
}

type SingleProject struct {
	Project Project `json:"project,omitempty"`
	Metadata Meta `json:"meta,omitempty"`
	Included ProjectIncluded `json:"included,omitempty"`
}

type Project struct {
	ActivePages struct {
		Billing bool `json:"billing,omitempty"`
		Board bool `json:"board,omitempty"`
		Comments bool `json:"comments,omitempty"`
		Files bool `json:"files,omitempty"`
		Finance bool `json:"finance,omitempty"`
		Forms bool `json:"forms,omitempty"`
		Gantt bool `json:"gantt,omitempty"`
		Links bool `json:"links,omitempty"`
		List bool `json:"list,omitempty"`
		Messages bool `json:"messages,omitempty"`
		Milestones bool `json:"milestones,omitempty"`
		Notebooks bool `json:"notebooks,omitempty"`
		RiskRegister bool `json:"riskRegister,omitempty"`
		Table bool `json:"table,omitempty"`
		Tasks bool `json:"tasks,omitempty"`
		Time bool `json:"time,omitempty"`
	} `json:"activePages,omitempty"`
	Announcement string `json:"announcement,omitempty"`
	Category Relationship `json:"category,omitempty"`
	Company Relationship `json:"company,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy int `json:"createdBy,omitempty"`
	CustomFieldValues []Relationship `json:"customFieldValues,omitempty"`
	DefaultPrivacy string `json:"defaultPrivacy,omitempty"`
	Description string `json:"description,omitempty"`
	DirectFileUploadsEnabled bool `json:"directFileUploadsEnabled,omitempty"`
	EndAt string `json:"endAt,omitempty "`
	FinancialBudget Relationship `json:"financialBudgetRelationship,omitempty"`
	HarvestTimersEnabled bool `json:"harvestTimersEnabled,omitempty"`
	ID int `json:"id,omitempty"`
	Integrations struct {
		OneDriveBusiness IntegrationsMeta `json:"oneDriveBusiness,omitempty"`
		Sharepoint IntegrationsMeta `json:"sharepoint,omitempty"`
		Xero struct {
			BaseCurrency string `json:"baseCurrency,omitempty"`
			Connected bool `json:"connected,omitempty"`
			CountryCode string `json:"countryCode,omitempty"`
			Enabled bool `json:"enabled,omitempty"`
			Organisation string `json:"organisation,omitempty"`
		} `json:"xero,omitempty"`
	} `json:"integrations,omitempty"`
	IsBillable bool `json:"isBillable,omitempty"`
	IsOnBoardingProject bool `json:"isOnBoardingProject,omitempty"`
	IsProjectAdmin bool `json:"isProjectAdmin,omitempty"`
	IsSampleProject bool `json:"isSampleProject,omitempty"`
	IsStarred bool `json:"isStarred,omitempty"`
	LastWorkedOn string `json:"lastWorkedOn,omitempty"`
	LatestActivity Relationship `json:"latestActivity,omitempty"`
	Logo string `json:"logo,omitempty"`
	MinMaxAvailableDates struct {
		DeadlinesFound bool `json:"deadlinesFound,omitempty"`
		MaxEndDate string `json:"maxEndDate,omitempty"`
		MinStartDate string `json:"minStartDate,omitempty"`
		SuggestedEndDate string `json:"suggestedEndDate,omitempty"`
		SuggestedStartDate string `json:"suggestedStartDate,omitempty"`
	} `json:"minMaxAvailableDates,omitempty"`
	Name string `json:"name,omitempty"`
	NotifyCommentIncludeCreater bool `json:"notifyCommentIncludeCreator,omitempty"`
	NotifyEveryone bool `json:"notifyEveryone,omitempty"`
	NotifyTaskAssignee bool `json:"notifyTaskAssignee,omitempty"`
	OverviewStartPage string `json:"overviewStartPage,omitempty"`
	OwnedBy int `json:"ownedBy,omitempty"`
	PortfolioCards []Relationship `json:"portfolioCards,omitempty"`
	PrivacyEnabled bool `json:"privacyEnabled,omitempty"`
	ProjectOwner Relationship `json:"projectOwner,omitempty"`
	ReplyByEmailEnabled bool `json:"replyByEmailEnabled,omitempty"`
	ShowAnnouncement bool `json:"showAnnouncement,omitempty"`
	SkipWeekends bool `json:"skipWeekends,omitempty"`
	StartAt string `json:"startAt,omitempty"`
	StartPage string `json:"startPage,omitempty"`
	Status string `json:"status,omitempty"`
	SubStatus string `json:"subStatus,omitempty"`
	Tags []Relationship `json:"tags,omitempty"`
	TasksStartPage string `json:"tasksStartPage,omitempty"`
	TimeBudget Relationship `json:"timeBudget,omitempty"`
	Type string `json:"type,omitempty"`
	Update Relationship `json:"update,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy int `json:"updatedBy,omitempty"`
	Users []Relationship `json:"users,omitempty"`
}	

type IntegrationsMeta struct {
	Account string `json:"account,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Folder string `json:"folder,omitempty"`
	FolderName string `json:"folderName,omitempty"`
}

type ProjectIncluded struct {
	Activities interface{}
	Companies interface{}
	Countries interface{}
	CustomFieldProjects interface{}
	CustomFields interface{}
	Industries interface{}
	PortfolioBoards interface{}
	PortfolioCards interface{}
	PortfolioColumns interface{}
	ProjectBudgets interface{}
	ProjectCategories interface{}
	ProjectTaskStats interface{}
	ProjectUpdates interface{}
	Tags interface{}
	Users interface{}
}