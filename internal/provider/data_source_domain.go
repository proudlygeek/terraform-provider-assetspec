package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-scaffolding/internal/httpclient"
	"github.com/hashicorp/terraform-provider-scaffolding/internal/repository"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDomainRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"datetime_lastcheck": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_expect": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"webhook_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*httpclient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)

	if id == "" {
		return diag.Errorf("id must be specified")
	}

	repo := &repository.DomainRepository{Client: client}

	domain, err := repo.GetDomain(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)
	d.Set("check_interval", domain.Detail.CheckInterval)
	d.Set("datetime_lastcheck", domain.Detail.DatetimeLastcheck)
	d.Set("labels", domain.Detail.Labels)
	d.Set("last_result", domain.Detail.LastResult)
	d.Set("tcp_expect", domain.Detail.TCPExpect)
	d.Set("webhook_target", domain.Detail.WebhookTarget)

	return diags
}
