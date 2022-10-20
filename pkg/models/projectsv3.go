package models

type ProjectResponseV3 struct {
	Included struct {
		Activities            interface{} `json:"activities,omitempty"`
		Companies             interface{} `json:"companies,omitempty"`
		Countries             interface{} `json:"countries,omitempty"`
		CustomfieldProjects   interface{} `json:"customfieldProjects,omitempty"`
		Customfields          interface{} `json:"customfields,omitempty"`
		Industries            interface{} `json:"industries,omitempty"`
		PortfolioBoards       interface{} `json:"portfolioBoards,omitempty"`
		PortfolioCards        interface{} `json:"portfolioCards,omitempty"`
		PortfolioColumns      interface{} `json:"portfolioColumns,omitempty"`
		ProjectBudgets        interface{} `json:"projectBudgets,omitempty"`
		ProjectCategories     interface{} `json:"projectCategories,omitempty"`
		ProjectEmailDropboxes interface{} `json:"projectEmailDropboxes,omitempty"`
		ProjectTaskStats      interface{} `json:"projectTaskStats,omitempty"`
		ProjectUpdates        interface{} `json:"projectUpdates,omitempty"`
		Tags                  interface{} `json:"tags,omitempty"`
		Users                 interface{} `json:"users,omitempty"`
	} `json:"included,omitempty"`
	Meta struct {
		Limit      int    `json:"limit,omitempty"`
		NextCursor string `json:"nextCursor,omitempty"`
		Page       struct {
			Count      int  `json:"count,omitempty"`
			HasMore    bool `json:"hasMore,omitempty"`
			PageOffset int  `json:"pageOffset,omitempty"`
			PageSize   int  `json:"pageSize,omitempty"`
		} `json:"page,omitempty"`
		PrevCursor    string `json:"prevCursor,omitempty"`
		TotalCapacity int    `json:"totalCapacity,omitempty"`
	} `json:"meta,omitempty"`
	Projects []ProjectV3 `json:"projects,omitempty"`
}

type ProjectV3 struct {
	ActivePages  *ActivePages `json:"activePages,omitempty"`
	Announcement string       `json:"announcement,omitempty"`
	Category     struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"category,omitempty"`
	CategoryID int `json:"categoryId,omitempty"`
	Company    struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"company,omitempty"`
	CompanyID           int    `json:"companyId,omitempty"`
	CreatedAt           string `json:"createdAt,omitempty"`
	CreatedBy           int    `json:"createdBy,omitempty"`
	CustomFieldValueIds []int  `json:"customFieldValueIds,omitempty"`
	CustomFieldValues   []struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"customFieldValues,omitempty"`
	CustomfieldValues []struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"customfieldValues,omitempty"`
	DefaultPrivacy           string `json:"defaultPrivacy,omitempty"`
	Description              string `json:"description,omitempty"`
	DirectFileUploadsEnabled bool   `json:"directFileUploadsEnabled,omitempty"`
	EndAt                    string `json:"endAt,omitempty"`
	EndDate                  string `json:"endDate,omitempty"`
	FinancialBudget          struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"financialBudget,omitempty"`
	FinancialBudgetID    int  `json:"financialBudgetId,omitempty"`
	HarvestTimersEnabled bool `json:"harvestTimersEnabled,omitempty"`
	ID                   int  `json:"id,omitempty"`
	Integrations         struct {
		OneDriveBusiness struct {
			Account    string `json:"account,omitempty"`
			Enabled    bool   `json:"enabled,omitempty"`
			Folder     string `json:"folder,omitempty"`
			FolderName string `json:"folderName,omitempty"`
		} `json:"oneDriveBusiness,omitempty"`
		Sharepoint struct {
			Account    string `json:"account,omitempty"`
			Enabled    bool   `json:"enabled,omitempty"`
			Folder     string `json:"folder,omitempty"`
			FolderName string `json:"folderName,omitempty"`
		} `json:"sharepoint,omitempty"`
		Xero struct {
			BaseCurrency string `json:"baseCurrency,omitempty"`
			Connected    bool   `json:"connected,omitempty"`
			CountryCode  string `json:"countryCode,omitempty"`
			Enabled      bool   `json:"enabled,omitempty"`
			Organisation string `json:"organisation,omitempty"`
		} `json:"xero,omitempty"`
	} `json:"integrations,omitempty"`
	IsBillable          bool   `json:"isBillable,omitempty"`
	IsOnBoardingProject bool   `json:"isOnBoardingProject,omitempty"`
	IsProjectAdmin      bool   `json:"isProjectAdmin,omitempty"`
	IsSampleProject     bool   `json:"isSampleProject,omitempty"`
	IsStarred           bool   `json:"isStarred,omitempty"`
	LastWorkedOn        string `json:"lastWorkedOn,omitempty"`
	LatestActivity      struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"latestActivity,omitempty"`
	Logo                 string `json:"logo,omitempty"`
	MinMaxAvailableDates struct {
		DeadlinesFound     bool   `json:"deadlinesFound,omitempty"`
		MaxEndDate         string `json:"maxEndDate,omitempty"`
		MinStartDate       string `json:"minStartDate,omitempty"`
		SuggestedEndDate   string `json:"suggestedEndDate,omitempty"`
		SuggestedStartDate string `json:"suggestedStartDate,omitempty"`
	} `json:"minMaxAvailableDates,omitempty"`
	Name                        string `json:"name,omitempty"`
	NotifyCommentIncludeCreator bool   `json:"notifyCommentIncludeCreator,omitempty"`
	NotifyEveryone              bool   `json:"notifyEveryone,omitempty"`
	NotifyTaskAssignee          bool   `json:"notifyTaskAssignee,omitempty"`
	OverviewStartPage           string `json:"overviewStartPage,omitempty"`
	OwnedBy                     int    `json:"ownedBy,omitempty"`
	OwnerID                     int    `json:"ownerId,omitempty"`
	PortfolioCards              []struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"portfolioCards,omitempty"`
	PrivacyEnabled bool `json:"privacyEnabled,omitempty"`
	ProjectOwner   struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"projectOwner,omitempty"`
	ProjectOwnerID      int    `json:"projectOwnerId,omitempty"`
	ReplyByEmailEnabled bool   `json:"replyByEmailEnabled,omitempty"`
	ShowAnnouncement    bool   `json:"showAnnouncement,omitempty"`
	SkipWeekends        bool   `json:"skipWeekends,omitempty"`
	StartAt             string `json:"startAt,omitempty"`
	StartDate           string `json:"startDate,omitempty"`
	StartPage           string `json:"startPage,omitempty"`
	Status              string `json:"status,omitempty"`
	SubStatus           string `json:"subStatus,omitempty"`
	TagIds              []int  `json:"tagIds,omitempty"`
	Tags                []struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"tags,omitempty"`
	TasksStartPage string `json:"tasksStartPage,omitempty"`
	TimeBudget     struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"timeBudget,omitempty"`
	TimeBudgetID int    `json:"timeBudgetId,omitempty"`
	Type         string `json:"type,omitempty"`
	Update       struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"update,omitempty"`
	UpdateID  int    `json:"updateId,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy int    `json:"updatedBy,omitempty"`
	Users     []struct {
		ID   int `json:"id,omitempty"`
		Meta struct {
		} `json:"meta,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"users,omitempty"`
}

type ActivePages struct {
	Billing      bool `json:"billing,omitempty"`
	Board        bool `json:"board,omitempty"`
	Comments     bool `json:"comments,omitempty"`
	Files        bool `json:"files,omitempty"`
	Finance      bool `json:"finance,omitempty"`
	Forms        bool `json:"forms,omitempty"`
	Gantt        bool `json:"gantt,omitempty"`
	Links        bool `json:"links,omitempty"`
	List         bool `json:"list,omitempty"`
	Messages     bool `json:"messages,omitempty"`
	Milestones   bool `json:"milestones,omitempty"`
	Notebooks    bool `json:"notebooks,omitempty"`
	RiskRegister bool `json:"riskRegister,omitempty"`
	Table        bool `json:"table,omitempty"`
	Tasks        bool `json:"tasks,omitempty"`
	Time         bool `json:"time,omitempty"`
}
