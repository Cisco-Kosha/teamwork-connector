package models

type Relationship struct {
	ID int `json:"id,omitempty"`
	Metadata Meta `json:"meta,omitempty"`
	Type string `json:"type,omitempty"`
}

type Meta struct {
	Limit int `json:"limit,omitempty"`
	NextCursor string `json:"nextCursor,omitempty"`
	Page MetaPage `json:"page,omitempty"`
	PrevCursor string `json:"prevCursor,omitempty"`
}

type MetaPage struct {
	Count int `json:"count,omitempty"`
	HasMore bool `json:"hasMore,omitempty"`
	PageOffset int `json:"pageOffset,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type Account struct {
	realTimeNotifications bool `json:"count,omitempty"`
	DefaultUserType string `json:"defaultUserType,omitempty"`
	Logo string `json:"logo,omitempty"`
	AllowProjectAdminsCreateUsers bool `json:"allowProjectAdminsCreateUsers,omitempty"`
	PricePlan string `json:"pricePlan,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	TrialExpiryDate string `json:"trialExpiryDate,omitempty"`
	PriceplanType string `json:"priceplanType,omitempty"`
	PaidUntilDateTime string `json:"paidUntilDateTime,omitempty"`
	EmailNotificationEnabled  bool `json:"email-notification-enabled,omitempty"`
	CompanyId string `json:"companyid,omitempty"`
	PasswordPolicyIsOn bool `json:"passwordPolicyIsOn,omitempty"`
	SiteOwnerEmail string `json:"siteOwnerEmail,omitempty"`
	URL string `json:"URL,omitempty"`
	TrialDaysRemaining string `json:"trialDaysRemaining,omitempty"`
	PdfServerLink string `json:"pdfServerLink,omitempty"`
	IndustryCategoryId string `json:"industryCategoryId,omitempty"`
	DashboardProjectsList string `json:"dashboardProjectsList,omitempty"`
	SSO interface{} `json:"SSO,omitempty"`
	DashboardMessageHTML string `json:"dashboardMessageHTML,omitempty"`
	PaymentMethod string `json:"paymentMethod,omitempty"`
	BillingAmount string `json:"billingAmount,omitempty"`
	Currency interface{} `json:"currency,omitempty"`
	Lang string `json:"lang,omitempty"`
	Datesignedup string `json:"datesignedup,omitempty"`
	IsStaging bool `json:"isStaging,omitempty"`
	Id string `json:"id,omitempty"`
	AwsRegion string `json:"awsRegion,omitempty"`
	Companyname string `json:"companyname,omitempty"`
	SiteOwnerName string `json:"siteOwnerName,omitempty"`
}

type SingleAccount struct {
	Status string `json:"STATUS,omitempty"`
	Account Account `json:"account,omitempty"`
}


type Specification struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	DomainName   string `json:"domain_name,omitempty"`
}