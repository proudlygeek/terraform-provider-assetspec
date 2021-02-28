package repository

import (
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
