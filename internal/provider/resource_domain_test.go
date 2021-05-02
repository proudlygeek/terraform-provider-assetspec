package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceDomain(t *testing.T) {
	rnd := "example2"
	name := "assetspec_domain." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDomain(rnd, rnd+".com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						name, "fqdn", rnd+".com"),
					resource.TestCheckResourceAttr(name, "tcp_expect.#", "3"),
				),
			},
		},
	})
}

func testAccResourceDomain(resourceName, fqdn string) string {
	return fmt.Sprintf(`
	resource "assetspec_domain" "%[1]s" {
		fqdn = "%[2]s"
		tcp_expect = [80, 443, 9081]
	}`, resourceName, fqdn)
}
