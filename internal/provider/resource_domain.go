package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-scaffolding/internal/httpclient"
	"github.com/hashicorp/terraform-provider-scaffolding/internal/repository"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Domain maps to an internet domain.",

		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,

		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1440,
			},
			"datetime_lastcheck": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
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
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"webhook_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*httpclient.Client)

	var diags diag.Diagnostics

	repo := &repository.DomainRepository{Client: client}

	id := d.Get("fqdn").(string)

	domain := &repository.CreateDomainBody{
		FQDN:          id,
		Labels:        formatLabels(d),
		TCPExpect:     formatTCPExpect(d),
		CheckInterval: d.Get("check_interval").(int),
		WebhookTarget: formatWebhookTarget(d),
	}

	_, err := repo.CreateDomain(domain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	resourceDomainRead(ctx, d, meta)

	return diags
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*httpclient.Client)

	var diags diag.Diagnostics

	id := d.Get("fqdn").(string)

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

func resourceDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*httpclient.Client)

	var diags diag.Diagnostics

	repo := &repository.DomainRepository{Client: client}

	id := d.Get("fqdn").(string)

	domain := &repository.UpdateDomainBody{
		Labels:        formatLabels(d),
		TCPExpect:     formatTCPExpect(d),
		CheckInterval: d.Get("check_interval").(int),
		WebhookTarget: formatWebhookTarget(d),
	}

	_, err := repo.UpdateDomain(id, domain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	resourceDomainRead(ctx, d, meta)

	return diags
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*httpclient.Client)

	var diags diag.Diagnostics

	repo := &repository.DomainRepository{Client: client}

	id := d.Get("fqdn").(string)

	_, err := repo.DeleteDomain(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func formatLabels(d *schema.ResourceData) []string {
	items := d.Get("labels").([]interface{})

	if len(items) == 0 {
		return make([]string, 0)
	}

	labels := []string{}
	for _, item := range items {
		labels = append(labels, item.(string))
	}

	return labels
}

func formatTCPExpect(d *schema.ResourceData) []int {
	items := d.Get("tcp_expect").([]interface{})

	if len(items) == 0 {
		return []int{80}
	}

	tcpExpect := []int{}
	for _, item := range items {
		tcpExpect = append(tcpExpect, item.(int))
	}

	return tcpExpect
}

func formatWebhookTarget(d *schema.ResourceData) *string {
	webhookTarget := d.Get("webhook_target").(string)
	if webhookTarget == "" {
		webhookTarget = "-"
	}

	return &webhookTarget
}
