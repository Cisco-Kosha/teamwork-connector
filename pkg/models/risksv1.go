package models

import "time"

type ReturnedRisksV1 struct {
	Status  string    `json:"STATUS,omitempty"`
	Message string    `json:"MESSAGE,omitempty"`
	Risks   []RisksV1 `json:"risks,omitempty"`
}

type RisksV1 struct {
	LastChangedOn              time.Time `json:"lastChangedOn,omitempty"`
	CreatedByUserID            string    `json:"createdByUserId,omitempty"`
	Impact                     string    `json:"impact,omitempty"`
	ImpactSchedule             string    `json:"impactSchedule,omitempty"`
	ImpactValue                string    `json:"impactValue,omitempty"`
	ProjectID                  string    `json:"projectId,omitempty"`
	Status                     string    `json:"status,omitempty"`
	ProbabilityValue           string    `json:"probabilityValue,omitempty"`
	Source                     string    `json:"source,omitempty"`
	Result                     string    `json:"result,omitempty"`
	CreatedOn                  time.Time `json:"createdOn,omitempty"`
	ID                         string    `json:"id,omitempty"`
	LastChangedByUserID        string    `json:"lastChangedByUserId,omitempty"`
	Deleted                    bool      `json:"deleted,omitempty"`
	CreatedByUserLastName      string    `json:"createdByUserLastName,omitempty"`
	CompanyName                string    `json:"companyName,omitempty"`
	CreatedByUserFirstName     string    `json:"createdByUserFirstName,omitempty"`
	LastChangedByUserFirstName string    `json:"lastChangedByUserFirstName,omitempty"`
	ProjectName                string    `json:"projectName,omitempty"`
	ProjectIsActive            bool      `json:"projectIsActive,omitempty"`
	ImpactPerformance          string    `json:"impactPerformance,omitempty"`
	MitigationPlan             string    `json:"mitigationPlan,omitempty"`
	Probability                string    `json:"probability,omitempty"`
	LastChangedByUserLastName  string    `json:"lastChangedByUserLastName,omitempty"`
	ImpactCost                 string    `json:"impactCost,omitempty"`
	CompanyID                  string    `json:"companyId,omitempty"`
}
