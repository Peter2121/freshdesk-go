package freshdesk

import "time"

type Ticket struct {
	Attachments     []interface{} `json:"attachments"`
	CcEmails        []string      `json:"cc_emails"`
	CompanyID       uint64        `json:"company_id"`
	CustomFields    interface{}   `json:"custom_fields"`
	Deleted         bool          `json:"deleted"`
	Description     string        `json:"description"`
	DescriptionText string        `json:"description_text"`
	DueBy           *time.Time    `json:"due_by"`
	Email           string        `json:"email"`
	EmailConfigID   int64         `json:"email_config_id"`
	FacebookID      string        `json:"facebook_id"`
	FrDueBy         *time.Time    `json:"fr_due_by"`
	FrEscalated     bool          `json:"fr_escalated"`
	FwdEmails       []string      `json:"fwd_emails"`
	GroupID         int64         `json:"group_id"`
	ID              uint64        `json:"id"`
	IsEscalated     bool          `json:"is_escalated"`
	Name            string        `json:"name"`
	Phone           string        `json:"phone"`
	Priority        Priority      `json:"priority"`
	ProductID       int64         `json:"product_id"`
	ReplyCcEmails   []string      `json:"reply_cc_emails"`
	RequesterID     int64         `json:"requester_id"` // UserID of the requester
	ResponderID     int64         `json:"responder_id"`
	Source          int64         `json:"source"`
	Spam            bool          `json:"spam"`
	Status          Status        `json:"status"`
	Subject         string        `json:"subject"`
	Tags            []string      `json:"tags"`
	ToEmails        []string      `json:"to_emails"`
	TwitterID       string        `json:"twitter_id"`
	Type            string        `json:"type"`
	CreatedAt       *time.Time    `json:"created_at"`
	UpdatedAt       *time.Time    `json:"updated_at"`
}

type Priority int64
type Status int64

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
	PriorityUrgent Priority = 4
)
const (
	StatusOpen     Status = 2
	StatusPending  Status = 3
	StatusResolved Status = 4
	StatusClosed   Status = 5
)

type TicketCreatePayload struct {
	Name             string        `json:"name,omitempty"`
	RequesterID      int64         `json:"requester_id,omitempty"` // UserID of the requester
	Email            string        `json:"email,omitempty"`
	FacebookID       string        `json:"facebook_id,omitempty"`
	Phone            string        `json:"phone,omitempty"`
	TwitterID        string        `json:"twitter_id,omitempty"`
	UniqueExternalID string        `json:"unique_external_id,omitempty"`
	Subject          string        `json:"subject,omitempty"`
	Type             string        `json:"type,omitempty"`
	Status           int64         `json:"status,omitempty"`
	Priority         int64         `json:"priority,omitempty"`
	Description      string        `json:"description,omitempty"`
	ResponderID      int64         `json:"responder_id,omitempty"`
	Attachments      []interface{} `json:"attachments,omitempty"`
	CcEmails         []string      `json:"cc_emails,omitempty"`
	CustomFields     interface{}   `json:"custom_fields,omitempty"`
	DueBy            *time.Time    `json:"due_by,omitempty"`
	EmailConfigID    int64         `json:"email_config_id,omitempty"`
	FrDueBy          *time.Time    `json:"fr_due_by,omitempty"`
	GroupID          int64         `json:"group_id,omitempty"`
	ProductID        int64         `json:"product_id,omitempty"`
	Source           int64         `json:"source,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	CompanyID        uint64        `json:"company_id,omitempty"`
	InternalAgentID  int64         `json:"internal_agent_id,omitempty"`
	InternalGroupID  int64         `json:"internal_group_id,omitempty"`
}

type TicketUpdatePayload struct {
	Name             string        `json:"name,omitempty"`
	RequesterID      int64         `json:"requester_id,omitempty"` // UserID of the requester
	Email            string        `json:"email,omitempty"`
	FacebookID       string        `json:"facebook_id,omitempty"`
	Phone            string        `json:"phone,omitempty"`
	TwitterID        string        `json:"twitter_id,omitempty"`
	UniqueExternalID string        `json:"unique_external_id,omitempty"`
	Subject          string        `json:"subject,omitempty"`
	Type             string        `json:"type,omitempty"`
	Status           int64         `json:"status,omitempty"`
	Priority         int64         `json:"priority,omitempty"`
	Description      string        `json:"description,omitempty"`
	ResponderID      int64         `json:"responder_id,omitempty"`
	Attachments      []interface{} `json:"attachments,omitempty"`
	CustomFields     interface{}   `json:"custom_fields,omitempty"`
	DueBy            *time.Time    `json:"due_by,omitempty"`
	EmailConfigID    int64         `json:"email_config_id,omitempty"`
	FrDueBy          *time.Time    `json:"fr_due_by,omitempty"`
	GroupID          int64         `json:"group_id,omitempty"`
	ProductID        int64         `json:"product_id,omitempty"`
	Source           int64         `json:"source,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	CompanyID        uint64        `json:"company_id,omitempty"`
	InternalAgentID  int64         `json:"internal_agent_id,omitempty"`
	InternalGroupID  int64         `json:"internal_group_id,omitempty"`
}

type Contact struct {
	Active            bool          `json:"active"`
	Address           string        `json:"address"`
	Avatar            interface{}   `json:"avatar"`
	CompanyID         uint64        `json:"company_id"`
	CreatedAt         *time.Time    `json:"created_at"`
	CsatRating        interface{}   `json:"csat_rating"`
	CustomFields      interface{}   `json:"custom_fields"`
	Deleted           bool          `json:"deleted"`
	Description       string        `json:"description"`
	Email             string        `json:"email"`
	FacebookID        interface{}   `json:"facebook_id"`
	FirstName         string        `json:"first_name"`
	ID                uint64        `json:"id"`
	JobTitle          string        `json:"job_title"`
	Language          string        `json:"language"`
	LastName          string        `json:"last_name"`
	Mobile            string        `json:"mobile"`
	Name              string        `json:"name"`
	OrgContactID      int64         `json:"org_contact_id"`
	OtherCompanies    []interface{} `json:"other_companies"`
	OtherEmails       []string      `json:"other_emails"`
	OtherPhoneNumbers []interface{} `json:"other_phone_numbers"`
	Phone             string        `json:"phone"`
	PreferredSource   string        `json:"preferred_source"`
	Tags              []string      `json:"tags"`
	TimeZone          string        `json:"time_zone"`
	TwitterID         string        `json:"twitter_id"`
	UniqueExternalID  string        `json:"unique_external_id"`
	UpdatedAt         *time.Time    `json:"updated_at"`
	ViewAllTickets    bool          `json:"view_all_tickets"`
	VisitorID         string        `json:"visitor_id"`
}

type ContactCreatePayload struct {
	Name             string        `json:"name,omitempty"`
	Email            string        `json:"email,omitempty"`
	Phone            string        `json:"phone,omitempty"`
	Mobile           string        `json:"mobile,omitempty"`
	TwitterID        string        `json:"twitter_id,omitempty"`
	UniqueExternalID string        `json:"unique_external_id,omitempty"`
	OtherEmails      []string      `json:"other_emails,omitempty"`
	CompanyID        uint64        `json:"company_id,omitempty"`
	ViewAllTickets   bool          `json:"view_all_tickets,omitempty"`
	OtherCompanies   []interface{} `json:"other_companies,omitempty"`
	Address          string        `json:"address,omitempty"`
	Avatar           interface{}   `json:"avatar,omitempty"`
	CustomFields     interface{}   `json:"custom_fields,omitempty"`
	Description      string        `json:"description,omitempty"`
	JobTitle         string        `json:"job_title,omitempty"`
	Languages        string        `json:"language,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	TimeZone         string        `json:"time_zone,omitempty"`
}

type ContactUpdatePayload struct {
	Name             string        `json:"name,omitempty"`
	Email            string        `json:"email,omitempty"`
	Phone            string        `json:"phone,omitempty"`
	Mobile           string        `json:"mobile,omitempty"`
	TwitterID        string        `json:"twitter_id,omitempty"`
	UniqueExternalID string        `json:"unique_external_id,omitempty"`
	OtherEmails      []string      `json:"other_emails,omitempty"`
	CompanyID        uint64        `json:"company_id,omitempty"`
	ViewAllTickets   bool          `json:"view_all_tickets,omitempty"`
	OtherCompanies   []interface{} `json:"other_companies,omitempty"`
	Address          string        `json:"address,omitempty"`
	Avatar           interface{}   `json:"avatar,omitempty"`
	CustomFields     interface{}   `json:"custom_fields,omitempty"`
	Description      string        `json:"description,omitempty"`
	JobTitle         string        `json:"job_title,omitempty"`
	Languages        string        `json:"language,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	TimeZone         string        `json:"time_zone,omitempty"`
}

type Company struct {
	CustomFields interface{} `json:"custom_fields"`
	Description  string      `json:"description"`
	Domains      []string    `json:"domains"`
	ID           uint64      `json:"id"`
	Name         string      `json:"name"`
	Note         string      `json:"note"`
	HealthScore  string      `json:"health_score"`
	AccountTier  string      `json:"account_tier"`
	RenewalDate  *time.Time  `json:"renewal_date"`
	Industry     string      `json:"industry"`
	CreatedAt    *time.Time  `json:"created_at"`
	UpdatedAt    *time.Time  `json:"updated_at"`
}

type CompanyName struct {
	ID           uint64      `json:"id"`
	Name         string      `json:"name"`
}

type SrchCompanyResp struct {
	CompanyNames []CompanyName `json:"companies"`
}

type CompanyCreatePayload struct {
	CustomFields interface{} `json:"custom_fields,omitempty"`
	Description  string      `json:"description,omitempty"`
	Domains      []string    `json:"domains,omitempty"`
	Name         string      `json:"name,omitempty"`
	Note         string      `json:"note,omitempty"`
	HealthScore  string      `json:"health_score,omitempty"`
	AccountTier  string      `json:"account_tier,omitempty"`
	RenewalDate  string      `json:"renewal_date,omitempty"`
	Industry     string      `json:"industry,omitempty"`
}

type CompanyUpdatePayload struct {
	CustomFields interface{} `json:"custom_fields,omitempty"`
	Description  string      `json:"description,omitempty"`
	Domains      []string    `json:"domains,omitempty"`
	Name         string      `json:"name,omitempty"`
	Note         string      `json:"note,omitempty"`
	HealthScore  string      `json:"health_score,omitempty"`
	AccountTier  string      `json:"account_tier,omitempty"`
	RenewalDate  string      `json:"renewal_date,omitempty"`
	Industry     string      `json:"industry,omitempty"`
}
