package main

import (
	"fmt"
	"net/url"
	"encoding/json"
)

type CreditorService service

type SchemeIdentifier struct {
	Name string
}

type SchemeIdentifierList struct {
	list []SchemeIdentifier
}

type Creditor struct {
	Id                 string   `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	Region             string   `json:"region,omitempty"`
	PostalCode         string   `json:"postal_code,omitempty"`
	LogoUrl            string   `json:"logo_url,omitempty"`
	City               string   `json:"city,omitempty"`
	AddressLine1       string   `json:"address_line1,omitempty"`
	AddressLine2       string   `json:"address_line2,omitempty"`
	AddressLine3       string   `json:"address_line3,omitempty"`
	CountryCode        string   `json:"country_code,omitempty"`
	CreatedAt          string   `json:"created_at,omitempty"`
	VerificationStatus []string `json:"verification_status,omitempty"`
	Links              []string `json:"links,omitempty"`
	SchemeIdentifiers  []string `json:"scheme_identifiers,omitempty"`
	ResponseUrl        string   `json:"responseurl,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}

type CreditorCreateRequest struct {
	Name         string   `json:"name,omitempty"`
	Region       string   `json:"region,omitempty"`
	PostalCode   string   `json:"postal_code,omitempty"`
	City         string   `json:"city,omitempty"`
	AddressLine1 string   `json:"address_line1,omitempty"`
	AddressLine2 string   `json:"address_line2,omitempty"`
	AddressLine3 string   `json:"address_line3,omitempty"`
	CountryCode  string   `json:"country_code,omitempty"`
	Links        []string `json:"links,omitempty"`
}

type CreatedAt struct {
	Gt  string `json:"gt,omitempty"`
	Gte string `json:"gte,omitempty"`
	Lt  string `json:"lt,omitempty"`
	Lte string `json:"lte,omitempty"`
}

type CreditorListRequest struct {
	CreatedAt CreatedAt `json:"created_at,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	Before    string    `json:"before,omitempty"`
	After     string    `json:"after,omitempty"`
}

type CreditorList struct {
	Meta   ListMeta
	Values []Creditor `json:"data"`
}

type CreditorUpdateRequest struct {
	Name         string   `json:"name,omitempty"`
	Region       string   `json:"region,omitempty"`
	PostalCode   string   `json:"postal_code,omitempty"`
	City         string   `json:"city,omitempty"`
	AddressLine1 string   `json:"address_line1,omitempty"`
	AddressLine2 string   `json:"address_line2,omitempty"`
	AddressLine3 string   `json:"address_line3,omitempty"`
	CountryCode  string   `json:"country_code,omitempty"`
	Identity     string   `json:"identity,omitempty"`
	Links        []string `json:"links,omitempty"`
}

// creates a new creditor
// https://developer.gocardless.com/api-reference/#creditors-create-a-creditor
func (c *CreditorService) CreateCreditor(creditor *Creditor) (*Creditor, error) {
	u := fmt.Sprintf("/creditors")
	crd := &Creditor{}
	rel := map[string]interface{}{
		"creditors": creditor,
	}

	custJson, _ := json.Marshal(rel)
	creditorObject := string(custJson[:])
	fmt.Println(creditorObject)

	err := c.client.Call("POST", u, rel, crd)

	return crd, err
}

func (s *CreditorService) ListCreditors(req *CreditorListRequest) (*CreditorList, error) {
	return s.ListNCreditors(100, 10, req)
}

func (s *CreditorService) ListNCreditors(count, offset int, req *CreditorListRequest) (*CreditorList, error) {
	params := url.Values{}
	params.Add("after", req.After)
	params.Add("before", req.Before)
	params.Add("created_at[gt]", req.CreatedAt.Gt)
	params.Add("created_at[gte]", req.CreatedAt.Gte)
	params.Add("created_at[lt]", req.CreatedAt.Lt)
	params.Add("created_at[lte]", req.CreatedAt.Lte)
	params.Add("limit", string(req.Limit))
	u := paginateURL("/creditors", count, offset)
	sub := &CreditorList{}
	err := s.client.Call("GET", u, params, sub)

	return sub, err
}

// Get:: returns the details of an existing creditor.
// https://developer.gocardless.com/api-reference/#creditors-get-a-single-creditor
func (s *CreditorService) GetCreditor(id string) (*Creditor, error) {
	u := fmt.Sprintf("/creditors/%s", id)
	sub := &Creditor{}
	err := s.client.Call("GET", u, nil, sub)

	return sub, err
}

// Update updates a creditor's properties.
func (s *CreditorService) UpdateCreditor(creditor *Creditor) (*Creditor, error) {
	u := fmt.Sprintf("/creditors/%s", creditor.Id)
	sub := &Creditor{}
	rel := map[string]interface{}{
		"creditors": creditor,
	}
	err := s.client.Call("PUT", u, rel, sub)

	return sub, err
}
