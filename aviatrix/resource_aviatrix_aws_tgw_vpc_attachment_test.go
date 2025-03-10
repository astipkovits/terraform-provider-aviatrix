package aviatrix

import (
	"fmt"
	"os"
	"testing"

	"github.com/AviatrixSystems/terraform-provider-aviatrix/v2/goaviatrix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAviatrixAwsTgwVpcAttachment_basic(t *testing.T) {
	var awsTgwVpcAttachment goaviatrix.AwsTgwVpcAttachment

	rName := acctest.RandString(5)
	resourceName := "aviatrix_aws_tgw_vpc_attachment.test"

	skipAcc := os.Getenv("SKIP_AWS_TGW_VPC_ATTACHMENT")
	if skipAcc == "yes" {
		t.Skip("Skipping AWS TGW VPC ATTACH test as SKIP_AWS_TGW_VPC_ATTACHMENT is set")
	}
	msg := ". Set SKIP_AWS_TGW_VPC_ATTACHMENT to yes to skip AWS TGW VPC ATTACH tests"

	awsSideAsNumber := "64512"
	nDm := "myNdn"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			preAccountCheck(t, msg)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsTgwVpcAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsTgwVpcAttachmentConfigBasic(rName, awsSideAsNumber, nDm),
				Check: resource.ComposeTestCheckFunc(
					tesAccCheckAwsTgwVpcAttachmentExists(resourceName, &awsTgwVpcAttachment),
					resource.TestCheckResourceAttr(resourceName, "tgw_name", fmt.Sprintf("tft-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "region", os.Getenv("AWS_REGION")),
					resource.TestCheckResourceAttr(resourceName, "network_domain_name", nDm),
					resource.TestCheckResourceAttr(resourceName, "vpc_id", os.Getenv("AWS_VPC_ID")),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAwsTgwVpcAttachmentConfigBasic(rName string, awsSideAsNumber string, nDm string) string {
	return fmt.Sprintf(`
resource "aviatrix_account" "test_account" {
	account_name       = "tfa-%s"
	cloud_type         = 1
	aws_account_number = "%s"
	aws_iam            = false
	aws_access_key     = "%s"
	aws_secret_key     = "%s"
}

resource "aviatrix_aws_tgw" "test_aws_tgw" {
	account_name           = aviatrix_account.test_account.account_name
	aws_side_as_number     = "%s"
	manage_vpc_attachment  = false
	region                 = "%s"
	tgw_name               = "tft-%s"
	manage_security_domain = false
	manage_vpc_attachment  = false
}

resource "aviatrix_aws_tgw_network_domain" "Default_Domain" {
	name     = "Default_Domain"
	tgw_name = aviatrix_aws_tgw.test_aws_tgw.tgw_name
}

resource "aviatrix_aws_tgw_network_domain" "Shared_Service_Domain" {
	name     = "Shared_Service_Domain"
	tgw_name = aviatrix_aws_tgw.test_aws_tgw.tgw_name
}

resource "aviatrix_aws_tgw_network_domain" "Aviatrix_Edge_Domain" {
	name     = "Aviatrix_Edge_Domain"
	tgw_name = aviatrix_aws_tgw.test_aws_tgw.tgw_name
}

resource "aviatrix_aws_tgw_network_domain" "test" {
	name       = "%s"
	tgw_name   = aviatrix_aws_tgw.test_aws_tgw.tgw_name
	depends_on = [
    	aviatrix_aws_tgw_network_domain.Default_Domain,
    	aviatrix_aws_tgw_network_domain.Shared_Service_Domain,
    	aviatrix_aws_tgw_network_domain.Aviatrix_Edge_Domain
  ]
}

resource "aviatrix_aws_tgw_vpc_attachment" "test" {
	tgw_name            = aviatrix_aws_tgw.test_aws_tgw.tgw_name
	region              = "%s"
	network_domain_name = "%s"
	vpc_account_name    = aviatrix_account.test_account.account_name
	vpc_id              = "%s"
}
	`, rName, os.Getenv("AWS_ACCOUNT_NUMBER"), os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"),
		awsSideAsNumber, os.Getenv("AWS_REGION"), rName, nDm, os.Getenv("AWS_REGION"), nDm,
		os.Getenv("AWS_VPC_ID"))
}

func tesAccCheckAwsTgwVpcAttachmentExists(n string, awsTgwVpcAttachment *goaviatrix.AwsTgwVpcAttachment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("AWS TGW VPC ATTACH Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no AWS TGW VPC ATTACH ID is set")
		}

		client := testAccProvider.Meta().(*goaviatrix.Client)

		foundAwsTgwVpcAttachment := &goaviatrix.AwsTgwVpcAttachment{
			TgwName:            rs.Primary.Attributes["tgw_name"],
			SecurityDomainName: rs.Primary.Attributes["network_domain_name"],
			VpcID:              rs.Primary.Attributes["vpc_id"],
		}

		foundAwsTgwVpcAttachment2, err := client.GetAwsTgwVpcAttachment(foundAwsTgwVpcAttachment)
		if err != nil {
			return err
		}
		if foundAwsTgwVpcAttachment2.TgwName != rs.Primary.Attributes["tgw_name"] {
			return fmt.Errorf("tgw_name Not found in created attributes")
		}
		if foundAwsTgwVpcAttachment2.SecurityDomainName != rs.Primary.Attributes["network_domain_name"] {
			return fmt.Errorf("network_domain_name Not found in created attributes")
		}
		if foundAwsTgwVpcAttachment2.VpcID != rs.Primary.Attributes["vpc_id"] {
			return fmt.Errorf("vpc_id Not found in created attributes")
		}

		*awsTgwVpcAttachment = *foundAwsTgwVpcAttachment2
		return nil
	}
}

func testAccCheckAwsTgwVpcAttachmentDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*goaviatrix.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aviatrix_aws_tgw_vpc_attachment" {
			continue
		}

		foundAwsTgwVpcAttachment := &goaviatrix.AwsTgwVpcAttachment{
			TgwName:            rs.Primary.Attributes["tgw_name"],
			SecurityDomainName: rs.Primary.Attributes["network_domain_name"],
			VpcID:              rs.Primary.Attributes["vpc_id"],
		}

		_, err := client.GetAwsTgwVpcAttachment(foundAwsTgwVpcAttachment)
		if err == nil {
			return fmt.Errorf("aviatrix AWS TGW VPC ATTACH still exists")
		}

		return nil
	}

	return nil
}
