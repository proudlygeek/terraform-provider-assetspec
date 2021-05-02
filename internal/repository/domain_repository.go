package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-provider-scaffolding/internal/httpclient"
)

type Domain struct {
	Detail *Detail `json:"detail"`
}

type Detail struct {
	CheckInterval     int      `json:"check_interval"`
	DatetimeLastcheck string   `json:"datetime_lastcheck"`
	LastResult        string   `json:"last_result"`
	TCPExpect         []int    `json:"tcp_expect"`
	Labels            []string `json:"labels"`
	WebhookTarget     *string  `json:"webhook_target,omitempty"`
}

type DomainRepository struct {
	Client *httpclient.Client
}

type CreateDomainBody struct {
	FQDN          string   `json:"fqdn"`
	TCPExpect     []int    `json:"tcp_expect"`
	WebhookTarget *string  `json:"webhook_target,omitempty"`
	Labels        []string `json:"labels"`
	CheckInterval int      `json:"check_interval"`
}

type UpdateDomainBody struct {
	TCPExpect     []int    `json:"tcp_expect"`
	WebhookTarget *string  `json:"webhook_target,omitempty"`
	Labels        []string `json:"labels"`
	CheckInterval int      `json:"check_interval"`
}

type DomainResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *DomainRepository) GetDomain(id string) (*Domain, error) {
	path := fmt.Sprintf("%s/domain/%s", httpclient.BaseURL, id)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	body, err := r.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}

	domain := &Domain{}
	err = json.Unmarshal(body, domain)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *DomainRepository) CreateDomain(params *CreateDomainBody) (*DomainResponse, error) {
	path := fmt.Sprintf("%s/domain", httpclient.BaseURL)
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(&params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", path, &buffer)
	if err != nil {
		return nil, err
	}

	body, err := r.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}

	res := &DomainResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DomainRepository) UpdateDomain(id string, params *UpdateDomainBody) (*DomainResponse, error) {
	path := fmt.Sprintf("%s/domain/%s", httpclient.BaseURL, id)
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(&params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", path, &buffer)
	if err != nil {
		return nil, err
	}

	body, err := r.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}

	res := &DomainResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DomainRepository) DeleteDomain(id string) (*DomainResponse, error) {
	path := fmt.Sprintf("%s/domain/%s", httpclient.BaseURL, id)

	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	body, err := r.Client.DoRequest(req)
	if err != nil {
		return nil, err
	}

	res := &DomainResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
