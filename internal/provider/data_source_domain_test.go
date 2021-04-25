package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDomain(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDomain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.assetspec_domain.example", "id", regexp.MustCompile("^example.com")),
				),
			},
		},
	})
}

const testAccDataSourceDomain = `
data "assetspec_domain" "example" {
  id = "example.com"
}
`
