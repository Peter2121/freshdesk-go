package freshdesk

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/ratelimit"
)

type Client interface {
	// GetAPIStatus() (*interface{}, error)

	GetTicket(ID uint64) (*Ticket, error)
	GetAllTickets() ([]Ticket, error)
	GetTicketsByCompanyID(companyID, pageSize, page int) ([]Ticket, error, bool)
	CreateTicket(payload TicketCreatePayload) (*Ticket, error)
	CreateTicketWithAttachments(payload TicketCreatePayload, files []Attachment) (*Ticket, error)
	UpdateTicket(ID uint64, payload TicketUpdatePayload) (*Ticket, error)
	DeleteTicket(ID uint64) (*interface{}, error)

	FindContactByEmail(email string) (Contact, error)
	GetContact(ID uint64) (*Contact, error)
	GetAllContacts() ([]Contact, error)
	CreateContact(payload ContactCreatePayload) (*Contact, error)
	UpdateContact(ID uint64, payload ContactUpdatePayload) (*Contact, error)
	SoftDeleteContact(ID uint64) (*interface{}, error)
	PermanentlyDeleteContact(ID uint64) (*interface{}, error)
	AddOtherCompanyForContact(fd_contact *Contact, id_client uint64, view_all bool) (bool, error)

	GetCompany(ID uint64) (*Company, error)
	GetAllCompanies() ([]Company, error)
	SearchCompanies(mask string) ([]CompanyName, error)
	CreateCompany(payload CompanyCreatePayload) (*Company, error)
	UpdateCompany(ID uint64, payload CompanyUpdatePayload) (*Company, error)
	DeleteCompany(ID uint64) (*interface{}, error)
}

type freshDeskService struct {
	restyClient *resty.Client
	rateLimiter ratelimit.Limiter
}

func NewClient(baseUrl string, user string, password string, maxRequestPerMinute int) Client {
	_freshDeskService := freshDeskService{
		restyClient: resty.New(),
		rateLimiter: ratelimit.New(maxRequestPerMinute, ratelimit.Per(time.Second*60), ratelimit.WithSlack(100)),
	}

	auth := user + ":" + password

	_freshDeskService.restyClient.SetBaseURL(baseUrl)
	_freshDeskService.restyClient.SetHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	return &_freshDeskService
}

// Ticket
func (service *freshDeskService) GetTicket(ID uint64) (*Ticket, error) {

	var responseSchema Ticket
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("%v%v", "/api/v2/tickets/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) GetAllTickets() ([]Ticket, error) {

	var responseSchema []Ticket
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get("/api/v2/tickets")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return responseSchema, nil
}

func (service *freshDeskService) CreateTicket(payload TicketCreatePayload) (*Ticket, error) {

	var responseSchema Ticket
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Post("/api/v2/tickets")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) CreateTicketWithAttachments(payload TicketCreatePayload, files []Attachment) (*Ticket, error) {

	var responseSchema Ticket
	new_ticket, err := service.CreateTicket(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, att := range files {
		req := service.restyClient.R()
		//req = req.SetFileReader("attachments[]", att.FileName, att.FileData) /**** Does not work like this ****/
		req = req.SetFile("attachments[]", att.FileData.Name())
		req = req.SetResult(&responseSchema)
		resp, err1 := req.Put(fmt.Sprintf("/api/v2/tickets/%v", new_ticket.ID))
		if err1 != nil {
			log.Println(err1)
		}
		if resp.StatusCode() != http.StatusOK {
			log.Println(string(resp.Body()))
		}
	}

	new_ticket, err = service.GetTicket(new_ticket.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return new_ticket, nil
}

func (service *freshDeskService) UpdateTicket(ID uint64, payload TicketUpdatePayload) (*Ticket, error) {
	var responseSchema Ticket
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Put(fmt.Sprintf("/api/v2/tickets/%v", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) DeleteTicket(ID uint64) (*interface{}, error) {
	var responseSchema interface{}
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Delete(fmt.Sprintf("%v%v", "/api/v2/tickets/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusNoContent {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

// Contact
func (service *freshDeskService) GetContact(ID uint64) (*Contact, error) {

	var responseSchema Contact
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("%v%v", "/api/v2/contacts/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}
func (service *freshDeskService) FindContactByEmail(email string) (Contact, error) {

	var responseSchema SrchContactResp
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("/api/v2/search/contacts?query=\"email:%s\"", url.QueryEscape("'"+email+"'")))

	if err != nil {
		log.Println(err)
		return Contact{}, err
	}

	if resp.StatusCode() != http.StatusOK {
		return Contact{}, errors.New(string(resp.Body()))
	}

	if responseSchema.Total == 0 {
		return Contact{}, errors.New("Contact not found")
	}

	return responseSchema.Results[0], nil
}

func (service *freshDeskService) GetAllContacts() ([]Contact, error) {

	var responseSchema []Contact
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get("/api/v2/contacts")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return responseSchema, nil
}

func (service *freshDeskService) CreateContact(payload ContactCreatePayload) (*Contact, error) {
	var responseSchema Contact
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Post("/api/v2/contacts")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) UpdateContact(ID uint64, payload ContactUpdatePayload) (*Contact, error) {
	var responseSchema Contact
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Put(fmt.Sprintf("/api/v2/contacts/%v", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) SoftDeleteContact(ID uint64) (*interface{}, error) {
	var responseSchema interface{}
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Delete(fmt.Sprintf("%v%v", "/api/v2/contacts/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusNoContent {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) PermanentlyDeleteContact(ID uint64) (*interface{}, error) {
	var responseSchema interface{}
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Delete(fmt.Sprintf("%v%v%v", "/api/v2/contacts/", ID, "/hard_delete?force=true"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusNoContent {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

// Company
func (service *freshDeskService) GetCompany(ID uint64) (*Company, error) {
	var responseSchema Company
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("%v%v", "/api/v2/companies/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) GetAllCompanies() ([]Company, error) {

	var responseSchema []Company
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get("/api/v2/companies")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return responseSchema, nil
}

func (service *freshDeskService) SearchCompanies(mask string) ([]CompanyName, error) {

	var responseSchema SrchCompanyResp
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("%v%v", "/api/v2/companies/autocomplete?name=", mask))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return responseSchema.CompanyNames, nil
}

func (service *freshDeskService) CreateCompany(payload CompanyCreatePayload) (*Company, error) {

	var responseSchema Company
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Post("/api/v2/companies")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) UpdateCompany(ID uint64, payload CompanyUpdatePayload) (*Company, error) {
	var responseSchema Company
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).SetResult(&responseSchema).
		Put(fmt.Sprintf("/api/v2/companies/%v", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) DeleteCompany(ID uint64) (*interface{}, error) {
	var responseSchema interface{}
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Delete(fmt.Sprintf("%v%v", "/api/v2/companies/", ID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusNoContent {
		return nil, errors.New(string(resp.Body()))
	}

	return &responseSchema, nil
}

func (service *freshDeskService) GetTicketsByCompanyID(companyID, pageSize, page int) ([]Ticket, error, bool) {
	service.rateLimiter.Take()

	var responseSchema []Ticket
	resp, err := service.restyClient.R().
		SetHeader("Content-Type", "application/json").SetResult(&responseSchema).
		Get(fmt.Sprintf("/api/v2/tickets?company_id=%v&per_page=%v&page=%v", companyID, pageSize, page))

	if err != nil {
		return nil, err, false
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body())), false
	}

	return responseSchema, nil, resp.Header().Get("Link") != ""
}

func (service *freshDeskService) AddOtherCompanyForContact(fd_contact *Contact, id_client uint64, view_all bool) (bool, error) {
	other_company := CompanyContactOtherUpdatePayload{
		ID:             id_client,
		ViewAllTickets: view_all,
	}
	var other_companies []CompanyContactOtherUpdatePayload
	for _, c := range fd_contact.OtherCompanies {
		other_companies = append(other_companies, CompanyContactOtherUpdatePayload{ID: c.ID, ViewAllTickets: c.ViewAllTickets})
	}
	other_companies = append(other_companies, other_company)
	update_contact := ContactUpdatePayload{
		Name:             fd_contact.Name,
		Email:            fd_contact.Email,
		Phone:            fd_contact.Phone,
		Mobile:           fd_contact.Mobile,
		TwitterID:        fd_contact.TwitterID,
		UniqueExternalID: fd_contact.UniqueExternalID,
		OtherEmails:      fd_contact.OtherEmails,
		CompanyID:        fd_contact.CompanyID,
		ViewAllTickets:   fd_contact.ViewAllTickets,
		OtherCompanies:   other_companies,
		Address:          fd_contact.Address,
		Avatar:           fd_contact.Avatar,
		CustomFields:     fd_contact.CustomFields,
		Description:      fd_contact.Description,
		JobTitle:         fd_contact.JobTitle,
		Languages:        fd_contact.Language,
		Tags:             fd_contact.Tags,
		TimeZone:         fd_contact.TimeZone,
	}
	_, err3 := service.UpdateContact(fd_contact.ID, update_contact)
	if err3 != nil {
		return false, err3
	}
	return true, nil
}
