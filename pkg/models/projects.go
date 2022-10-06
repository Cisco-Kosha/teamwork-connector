package models

type MultiProject struct {
	Projects []Project       `json:"projects,omitempty"`
	Metadata Meta            `json:"meta,omitempty"`
	Included ProjectIncluded `json:"included,omitempty"`
}

type SingleProject struct {
	Project  Project         `json:"project,omitempty"`
	Metadata Meta            `json:"meta,omitempty"`
	Included ProjectIncluded `json:"included,omitempty"`
}

type MultiMinimalProject struct {
	Projects []MinimalProject `json:"projects,omitempty"`
	Metadata Meta             `json:"meta,omitempty"`
	Included ProjectIncluded  `json:"included,omitempty"`
}

type MinimalProject struct {
	Name            string       `json:"name,omitempty"`
	Description     string       `json:"description,omitempty"`
	ID              int          `json:"id,omitempty"`
	Type            string       `json:"type,omitempty"`
	UpdatedAt       string       `json:"updatedAt,omitempty"`
	OwnedBy         int          `json:"ownedBy,omitempty"`
	FinancialBudget Relationship `json:"financialBudgetRelationship,omitempty"`
}

type Project struct {
	ActivePages struct {
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
	} `json:"activePages,omitempty"`
	Announcement             string         `json:"announcement,omitempty"`
	Category                 Relationship   `json:"category,omitempty"`
	Company                  Relationship   `json:"company,omitempty"`
	CreatedAt                string         `json:"createdAt,omitempty"`
	CreatedBy                int            `json:"createdBy,omitempty"`
	CustomFieldValues        []Relationship `json:"customFieldValues,omitempty"`
	DefaultPrivacy           string         `json:"defaultPrivacy,omitempty"`
	Description              string         `json:"description,omitempty"`
	DirectFileUploadsEnabled bool           `json:"directFileUploadsEnabled,omitempty"`
	EndAt                    string         `json:"endAt,omitempty "`
	FinancialBudget          Relationship   `json:"financialBudgetRelationship,omitempty"`
	HarvestTimersEnabled     bool           `json:"harvestTimersEnabled,omitempty"`
	ID                       int            `json:"id,omitempty"`
	Integrations             struct {
		OneDriveBusiness IntegrationsMeta `json:"oneDriveBusiness,omitempty"`
		Sharepoint       IntegrationsMeta `json:"sharepoint,omitempty"`
		Xero             struct {
			BaseCurrency string `json:"baseCurrency,omitempty"`
			Connected    bool   `json:"connected,omitempty"`
			CountryCode  string `json:"countryCode,omitempty"`
			Enabled      bool   `json:"enabled,omitempty"`
			Organisation string `json:"organisation,omitempty"`
		} `json:"xero,omitempty"`
	} `json:"integrations,omitempty"`
	IsBillable           bool         `json:"isBillable,omitempty"`
	IsOnBoardingProject  bool         `json:"isOnBoardingProject,omitempty"`
	IsProjectAdmin       bool         `json:"isProjectAdmin,omitempty"`
	IsSampleProject      bool         `json:"isSampleProject,omitempty"`
	IsStarred            bool         `json:"isStarred,omitempty"`
	LastWorkedOn         string       `json:"lastWorkedOn,omitempty"`
	LatestActivity       Relationship `json:"latestActivity,omitempty"`
	Logo                 string       `json:"logo,omitempty"`
	MinMaxAvailableDates struct {
		DeadlinesFound     bool   `json:"deadlinesFound,omitempty"`
		MaxEndDate         string `json:"maxEndDate,omitempty"`
		MinStartDate       string `json:"minStartDate,omitempty"`
		SuggestedEndDate   string `json:"suggestedEndDate,omitempty"`
		SuggestedStartDate string `json:"suggestedStartDate,omitempty"`
	} `json:"minMaxAvailableDates,omitempty"`
	Name                        string         `json:"name,omitempty"`
	NotifyCommentIncludeCreater bool           `json:"notifyCommentIncludeCreator,omitempty"`
	NotifyEveryone              bool           `json:"notifyEveryone,omitempty"`
	NotifyTaskAssignee          bool           `json:"notifyTaskAssignee,omitempty"`
	OverviewStartPage           string         `json:"overviewStartPage,omitempty"`
	OwnedBy                     int            `json:"ownedBy,omitempty"`
	PortfolioCards              []Relationship `json:"portfolioCards,omitempty"`
	PrivacyEnabled              bool           `json:"privacyEnabled,omitempty"`
	ProjectOwner                Relationship   `json:"projectOwner,omitempty"`
	ReplyByEmailEnabled         bool           `json:"replyByEmailEnabled,omitempty"`
	ShowAnnouncement            bool           `json:"showAnnouncement,omitempty"`
	SkipWeekends                bool           `json:"skipWeekends,omitempty"`
	StartAt                     string         `json:"startAt,omitempty"`
	StartPage                   string         `json:"startPage,omitempty"`
	Status                      string         `json:"status,omitempty"`
	SubStatus                   string         `json:"subStatus,omitempty"`
	Tags                        []Relationship `json:"tags,omitempty"`
	TasksStartPage              string         `json:"tasksStartPage,omitempty"`
	TimeBudget                  Relationship   `json:"timeBudget,omitempty"`
	Type                        string         `json:"type,omitempty"`
	Update                      Relationship   `json:"update,omitempty"`
	UpdatedAt                   string         `json:"updatedAt,omitempty"`
	UpdatedBy                   int            `json:"updatedBy,omitempty"`
	Users                       []Relationship `json:"users,omitempty"`
}

type IntegrationsMeta struct {
	Account    string `json:"account,omitempty"`
	Enabled    bool   `json:"enabled,omitempty"`
	Folder     string `json:"folder,omitempty"`
	FolderName string `json:"folderName,omitempty"`
}

type ProjectIncluded struct {
	Activities          interface{}
	Companies           interface{}
	Countries           interface{}
	CustomFieldProjects interface{}
	CustomFields        interface{}
	Industries          interface{}
	PortfolioBoards     interface{}
	PortfolioCards      interface{}
	PortfolioColumns    interface{}
	ProjectBudgets      interface{}
	ProjectCategories   interface{}
	ProjectTaskStats    interface{}
	ProjectUpdates      interface{}
	Tags                interface{}
	Users               interface{}
}

type ProjectUpdate struct {
	Update    Update `json:"update,omitempty"`
	NotifyIds string `json:"notifyIds,omitempty"`
}

type Update struct {
	Text   string `json:"text,omitempty"`
	Health string `json:"health,omitempty"`
}

// type ProjectUpdateResponse struct {
// 	Included       interface{}        `json:"included,omitempty"`
// 	Metadata       interface{}        `json:"meta,omitempty"`
// 	ProjectUpdates MultiProjectUpdate `json:"projectUpdates,omitempty"`
// }

type MultiProjectUpdate struct {
	Color           string      `json:"color,omitempty"`
	CreatedAt       string      `json:"createdAt,omitempty"`
	CreatedBy       int         `json:"createdBy,omitempty"`
	Deleted         bool        `json:"deleted,omitempty"`
	DeletedAt       string      `json:"deletedAt,omitempty"`
	DeletedBy       int         `json:"deletedBy,omitempty"`
	Health          int         `json:"health,omitempty"`
	HealthLabel     string      `json:"healthLabel,omitempty"`
	Id              int         `json:"id,omitempty"`
	LikeFromUserIDs interface{} `json:"likeFromUserIDs,omitempty"`
	LikeFromUsers   interface{} `json:"likeFromUsers,omitempty"`
	Project         interface{} `json:"project,omitempty"`
	ProjectId       int         `json:"projectId,omitempty"`
	Reactions       interface{} `json:"reactions,omitempty"`
	Text            string      `json:"text,omitempty"`
	UpdatedAt       string      `json:"updatedAt,omitempty"`
}

type ProjectUpdateResponse struct {
	Included struct {
		Projects struct {
			Property1 struct {
				ActivePages struct {
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
				} `json:"activePages,omitempty"`
				Announcement string `json:"announcement,omitempty"`
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
			} `json:"property1,omitempty"`
			Property2 struct {
				ActivePages struct {
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
				} `json:"activePages,omitempty"`
				Announcement string `json:"announcement,omitempty"`
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
			} `json:"property2,omitempty"`
		} `json:"projects,omitempty"`
		Users struct {
			Property1 struct {
				AvatarURL      string `json:"avatarUrl,omitempty"`
				CanAddProjects bool   `json:"canAddProjects,omitempty"`
				Company        struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"company,omitempty"`
				CompanyID        int    `json:"companyId,omitempty"`
				CompanyRoleID    int    `json:"companyRoleId,omitempty"`
				Deleted          bool   `json:"deleted,omitempty"`
				Email            string `json:"email,omitempty"`
				FirstName        string `json:"firstName,omitempty"`
				ID               int    `json:"id,omitempty"`
				IsAdmin          bool   `json:"isAdmin,omitempty"`
				IsClientUser     bool   `json:"isClientUser,omitempty"`
				IsServiceAccount bool   `json:"isServiceAccount,omitempty"`
				LastName         string `json:"lastName,omitempty"`
				LengthOfDay      int    `json:"lengthOfDay,omitempty"`
				Teams            []struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"teams,omitempty"`
				Title       string `json:"title,omitempty"`
				Type        string `json:"type,omitempty"`
				UserCost    int    `json:"userCost,omitempty"`
				UserRate    int    `json:"userRate,omitempty"`
				WorkingHour struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"workingHour,omitempty"`
				WorkingHoursID int `json:"workingHoursId,omitempty"`
			} `json:"property1,omitempty"`
			Property2 struct {
				AvatarURL      string `json:"avatarUrl,omitempty"`
				CanAddProjects bool   `json:"canAddProjects,omitempty"`
				Company        struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"company,omitempty"`
				CompanyID        int    `json:"companyId,omitempty"`
				CompanyRoleID    int    `json:"companyRoleId,omitempty"`
				Deleted          bool   `json:"deleted,omitempty"`
				Email            string `json:"email,omitempty"`
				FirstName        string `json:"firstName,omitempty"`
				ID               int    `json:"id,omitempty"`
				IsAdmin          bool   `json:"isAdmin,omitempty"`
				IsClientUser     bool   `json:"isClientUser,omitempty"`
				IsServiceAccount bool   `json:"isServiceAccount,omitempty"`
				LastName         string `json:"lastName,omitempty"`
				LengthOfDay      int    `json:"lengthOfDay,omitempty"`
				Teams            []struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"teams,omitempty"`
				Title       string `json:"title,omitempty"`
				Type        string `json:"type,omitempty"`
				UserCost    int    `json:"userCost,omitempty"`
				UserRate    int    `json:"userRate,omitempty"`
				WorkingHour struct {
					ID   int `json:"id,omitempty"`
					Meta struct {
					} `json:"meta,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"workingHour,omitempty"`
				WorkingHoursID int `json:"workingHoursId,omitempty"`
			} `json:"property2,omitempty"`
		} `json:"users,omitempty"`
	} `json:"included,omitempty"`
	Metadata struct {
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
	ProjectUpdates []struct {
		Color           string `json:"color,omitempty"`
		CreatedAt       string `json:"createdAt,omitempty"`
		CreatedBy       int    `json:"createdBy,omitempty"`
		Deleted         bool   `json:"deleted,omitempty"`
		DeletedAt       string `json:"deletedAt,omitempty"`
		DeletedBy       int    `json:"deletedBy,omitempty"`
		Health          int    `json:"health,omitempty"`
		HealthLabel     string `json:"healthLabel,omitempty"`
		ID              int    `json:"id,omitempty"`
		LikeFromUserIDs []int  `json:"likeFromUserIDs,omitempty"`
		LikeFromUsers   []struct {
			ID   int `json:"id,omitempty"`
			Meta struct {
			} `json:"meta,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"likeFromUsers,omitempty"`
		Project struct {
			ID   int `json:"id,omitempty"`
			Meta struct {
			} `json:"meta,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"project,omitempty"`
		ProjectID int `json:"projectId,omitempty"`
		Reactions struct {
			Counts struct {
				Dislike int `json:"dislike,omitempty"`
				Frown   int `json:"frown,omitempty"`
				Heart   int `json:"heart,omitempty"`
				Joy     int `json:"joy,omitempty"`
				Like    int `json:"like,omitempty"`
			} `json:"counts,omitempty"`
			Mine []string `json:"mine,omitempty"`
		} `json:"reactions,omitempty"`
		Text      string `json:"text,omitempty"`
		UpdatedAt string `json:"updatedAt,omitempty"`
	} `json:"projectUpdates,omitempty"`
}
