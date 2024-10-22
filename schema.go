package freshdesk

import (
	"os"
	"time"
)

type TicketMessage struct {
	Body              string      `json:"body"`
	BodyText          string      `json:"body_text"`
	ID                uint64      `json:"id"`
	IsIncoming        bool        `json:"incoming"`
	IsPrivate         bool        `json:"private"`
	UserId            uint64      `json:"user_id"`
	SupportEmail      string      `json:"support_email"`
	Source            uint64      `json:"source"`
	Category          uint64      `json:"category"`
	EmailsTo          []string    `json:"to_emails"`
	EmailFrom         string      `json:"from_email"`
	EmailsCc          []string    `json:"cc_emails"`
	EmailsBcc         []string    `json:"bcc_emails"`
	EmailFailureCount uint64      `json:"email_failure_count"`
	OutgoingFailures  uint64      `json:"outgoing_failures"`
	ThreadId          uint64      `json:"thread_id"`
	ThreadMessageId   uint64      `json:"thread_message_id"`
	CreatedAt         *time.Time  `json:"created_at"`
	UpdatedAt         *time.Time  `json:"updated_at"`
	EditedAt          *time.Time  `json:"last_edited_at"`
	EditedByUserId    uint64      `json:"last_edited_user_id"`
	Attachments       interface{} `json:"attachments"`
	AutomationId      uint64      `json:"automation_id"`
	AutomationTypeId  uint64      `json:"automation_type_id"`
	IsAutoResponse    bool        `json:"auto_response"`
	TicketId          uint64      `json:"ticket_id"`
	SrcAdditionalInfo interface{} `json:"source_additional_info"`
}

type TicketMessageCreatePayload struct {
	BodyHtml    string        `json:"body"`
	Attachments []interface{} `json:"attachments,omitempty"`
	EmailFrom   string        `json:"from_email,omitempty"`
	UserId      uint64        `json:"user_id,omitempty"`
	CcEmails    []string      `json:"cc_emails,omitempty"`
	BccEmails   []string      `json:"bcc_emails,omitempty"`
}

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
	Conversations   []TicketMessage `json:"conversations"`
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

const (
	ERR_CONTACT_NOT_FOUND string = "Contact not found"
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

type TicketStatusUpdatePayload struct {
	Status           int64         `json:"status"`
}

type Contact struct {
	Active            bool                  `json:"active"`
	Address           string                `json:"address"`
	Avatar            interface{}           `json:"avatar"`
	CompanyID         uint64                `json:"company_id"`
	CreatedAt         *time.Time            `json:"created_at"`
	CsatRating        interface{}           `json:"csat_rating"`
	CustomFields      interface{}           `json:"custom_fields"`
	Deleted           bool                  `json:"deleted"`
	Description       string                `json:"description"`
	Email             string                `json:"email"`
	FacebookID        interface{}           `json:"facebook_id"`
	FirstName         string                `json:"first_name"`
	ID                uint64                `json:"id"`
	JobTitle          string                `json:"job_title"`
	Language          string                `json:"language"`
	LastName          string                `json:"last_name"`
	Mobile            string                `json:"mobile"`
	Name              string                `json:"name"`
	OrgContactID      int64                 `json:"org_contact_id"`
	OtherCompanies    []CompanyContactOther `json:"other_companies"`
	OtherEmails       []string              `json:"other_emails"`
	OtherPhoneNumbers []interface{}         `json:"other_phone_numbers"`
	Phone             string                `json:"phone"`
	PreferredSource   string                `json:"preferred_source"`
	Tags              []string              `json:"tags"`
	TimeZone          string                `json:"time_zone"`
	TwitterID         string                `json:"twitter_id"`
	UniqueExternalID  string                `json:"unique_external_id"`
	UpdatedAt         *time.Time            `json:"updated_at"`
	ViewAllTickets    bool                  `json:"view_all_tickets"`
	VisitorID         string                `json:"visitor_id"`
}

type ContactShort struct {
	Active         bool        `json:"active"`
	Address        string      `json:"address"`
	CompanyID      uint64      `json:"company_id"`
	CreatedAt      *time.Time  `json:"created_at"`
	CustomFields   interface{} `json:"custom_fields"`
	Deleted        bool        `json:"deleted"`
	Description    string      `json:"description"`
	Email          string      `json:"email"`
	FacebookID     interface{} `json:"facebook_id"`
	ID             uint64      `json:"id"`
	JobTitle       string      `json:"job_title"`
	Language       string      `json:"language"`
	Mobile         string      `json:"mobile"`
	Name           string      `json:"name"`
	OtherCompanies []int64     `json:"other_companies"`
	Phone          string      `json:"phone"`
	Tags           []string    `json:"tags"`
	TwitterID      string      `json:"twitter_id"`
	UpdatedAt      *time.Time  `json:"updated_at"`
}

type ContactCreatePayload struct {
	Name             string                             `json:"name,omitempty"`
	Email            string                             `json:"email,omitempty"`
	Phone            string                             `json:"phone,omitempty"`
	Mobile           string                             `json:"mobile,omitempty"`
	TwitterID        string                             `json:"twitter_id,omitempty"`
	UniqueExternalID string                             `json:"unique_external_id,omitempty"`
	OtherEmails      []string                           `json:"other_emails,omitempty"`
	CompanyID        uint64                             `json:"company_id,omitempty"`
	ViewAllTickets   bool                               `json:"view_all_tickets,omitempty"`
	OtherCompanies   []CompanyContactOtherUpdatePayload `json:"other_companies,omitempty"`
	Address          string                             `json:"address,omitempty"`
	Avatar           interface{}                        `json:"avatar,omitempty"`
	CustomFields     interface{}                        `json:"custom_fields,omitempty"`
	Description      string                             `json:"description,omitempty"`
	JobTitle         string                             `json:"job_title,omitempty"`
	Languages        string                             `json:"language,omitempty"`
	Tags             []string                           `json:"tags,omitempty"`
	TimeZone         string                             `json:"time_zone,omitempty"`
}

type ContactUpdatePayload struct {
	Name             string                             `json:"name,omitempty"`
	Email            string                             `json:"email,omitempty"`
	Phone            string                             `json:"phone,omitempty"`
	Mobile           string                             `json:"mobile,omitempty"`
	TwitterID        string                             `json:"twitter_id,omitempty"`
	UniqueExternalID string                             `json:"unique_external_id,omitempty"`
	OtherEmails      []string                           `json:"other_emails,omitempty"`
	CompanyID        uint64                             `json:"company_id,omitempty"`
	ViewAllTickets   bool                               `json:"view_all_tickets,omitempty"`
	OtherCompanies   []CompanyContactOtherUpdatePayload `json:"other_companies,omitempty"`
	Address          string                             `json:"address,omitempty"`
	Avatar           interface{}                        `json:"avatar,omitempty"`
	CustomFields     interface{}                        `json:"custom_fields,omitempty"`
	Description      string                             `json:"description,omitempty"`
	JobTitle         string                             `json:"job_title,omitempty"`
	Languages        string                             `json:"language,omitempty"`
	Tags             []string                           `json:"tags,omitempty"`
	TimeZone         string                             `json:"time_zone,omitempty"`
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
	OrgCompanyID uint64      `json:"org_company_id"`
}

type CompanyName struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type SrchCompanyResp struct {
	CompanyNames []CompanyName `json:"companies"`
}

type SrchContactResp struct {
	Total   uint64    `json:"total"`
	Results []Contact `json:"results"`
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

type CompanyContactOther struct {
	ID             uint64      `json:"id"`
	ViewAllTickets bool        `json:"view_all_tickets,omitempty"`
	Name           string      `json:"name,omitempty"`
	Avatar         interface{} `json:"avatar,omitempty"`
}

type CompanyContactOtherUpdatePayload struct {
	ID             uint64 `json:"company_id"`
	ViewAllTickets bool   `json:"view_all_tickets,omitempty"`
}

type Attachment struct {
	FileName string
	FileType string
	FileData *os.File
}

type Group struct {
	ID                     int64       `json:"id"`
	Name                   string      `json:"name,omitempty"`
	Description            string      `json:"description,omitempty"`
	EscalateTo             uint64      `json:"escalate_to,omitempty"`
	UnassignedFor          string      `json:"unassigned_for,omitempty"`
	Agents                 []uint64    `json:"agent_ids,omitempty"`
	CreatedAt              *time.Time  `json:"created_at"`
	UpdatedAt              *time.Time  `json:"updated_at"`
	AllowAgentsChangeAvail bool        `json:"allow_agents_to_change_availability,omitempty"`
	BusinessCalendar       uint64      `json:"business_calendar_id,omitempty"`
	Type                   string      `json:"type,omitempty"`
	AutoAgentAssign        interface{} `json:"automatic_agent_assignment,omitempty"`
}

type CustomObject struct {
	DisplayID    string                  `json:"display_id"`
	CreatedTime  uint64                  `json:"created_time"`
	UpdatedTime  uint64                  `json:"updated_time"`
	Data         map[string]interface{}  `json:"data"`
	Version      uint64                  `json:"version"`
	Metadata     map[string]interface{}  `json:"metadata"`
	Links        map[string]interface{}  `json:"_links"`
}

type CustomObjectSearchResp struct {
	Records  []CustomObject          `json:"records"`
	Links    map[string]interface{}  `json:"_links"`	
}

type CustomObjectUpdatePayload struct {
	DisplayID    string                  `json:"display_id"`
	Version      uint64                  `json:"version"`
	Data         map[string]interface{}  `json:"data"`
}

type CustomObjectUpdateResult struct {
	DisplayID    string                  `json:"display_id"`
	CreatedTime  uint64                  `json:"created_time"`
	UpdatedTime  uint64                  `json:"updated_time"`
	Data         map[string]interface{}  `json:"data"`
}
