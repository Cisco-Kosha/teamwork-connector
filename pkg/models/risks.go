package models

type CreateRisk struct {
	Risk Risk `json:"risk,omitempty"`
}

type Risk struct {
	Source            string `json:"source,omitempty"`
	ProbabilityValue  string `json:"probabilityValue,omitempty"`
	ImpactValue       string `json:"impactValue,omitempty"`
	ImpactCost        bool   `json:"impactCost,omitempty"`
	ImpactSchedule    bool   `json:"impactSchedule,omitempty"`
	ImpactPerformance bool   `json:"impactPerformance,omitempty"`
	Status            string `json:"status,omitempty"`
	MitigationPlan    string `json:"mitigationPlan,omitempty"`
}

type ReturnedRisks struct {
	Included interface{} `json:"included,omitempty"`
	Metadata Meta        `json:"meta,omitempty"`
	Risks    []Risks     `json:"risks,omitempty"`
}

type Risks struct {
	CanEdit             bool        `json:"canEdit,omitempty"`
	CreatedAt           string      `json:"createdAt,omitempty"`
	CreatedBy           int         `json:"createdBy,omitempty"`
	CreatedByUserId     int         `json:"createdByUserId,omitempty"`
	CreatedOn           string      `json:"createdOn,omitempty"`
	Deleted             bool        `json:"deleted,omitempty"`
	Id                  int         `json:"id,omitempty"`
	Impact              string      `json:"impact,omitempty"`
	ImpactCost          bool        `json:"impactCost,omitempty"`
	ImpactPerformance   bool        `json:"impactPerformance,omitempty"`
	ImpactSchedule      bool        `json:"impactSchedule,omitempty"`
	ImpactValue         int         `json:"impactValue,omitempty"`
	LastChangedByUserId int         `json:"lastChangedByUserId,omitempty"`
	LastChangedOn       string      `json:"lastChangedOn,omitempty"`
	MitigationPlan      string      `json:"mitigationPlan,omitempty"`
	Probability         string      `json:"probability,omitempty"`
	Probabilityvalue    int         `json:"probabilityValue,omitempty"`
	Project             interface{} `json:"project,omitempty"`
	ProjectId           int         `json:"projectId,omitempty"`
	Result              int         `json:"result,omitempty"`
	Source              string      `json:"source,omitempty"`
	Status              string      `json:"status,omitempty"`
	UpdatedAt           string      `json:"updatedAt,omitempty"`
	UpdatedBy           int         `json:"updatedBy,omitempty"`
}
