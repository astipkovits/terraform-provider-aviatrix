---
subcategory: "TGW Orchestrator"
layout: "aviatrix"
page_title: "Aviatrix: aviatrix_aws_tgw_connect"
description: |- 
  Creates and manages Aviatrix AWS TGW Connect connections
---

# aviatrix_aws_tgw_connect

The **aviatrix_aws_tgw_connect** resource allows the creation and management of AWS TGW Connect connections. To create
and manage TGW Connect peers, please use `aviatrix_aws_tgw_connect_peer` resources. This resource is available as of
provider version R2.18.1+.

~> **NOTE:** Before creating an AWS TGW Connect, the AWS TGW must have an attached VPC via
the `aviatrix_aws_tgw_vpc_attachment` resource. Also, the AWS TGW must have configured CIDRs via
the `aviatrix_aws_tgw` `cidrs` attribute.

## Example Usage

```hcl
# Create an Aviatrix AWS TGW Connect
resource "aviatrix_aws_tgw_connect" "test_aws_tgw_connect" {
  tgw_name            = aviatrix_aws_tgw.test_aws_tgw.tgw_name
  connection_name     = "aws-tgw-connect"
  transport_vpc_id    = aviatrix_aws_tgw_vpc_attachment.test_aws_tgw_vpc_attachment.vpc_id
  network_domain_name = aviatrix_aws_tgw_vpc_attachment.test_aws_tgw_vpc_attachment.network_domain_name
}
```

## Argument Reference

The following arguments are supported:

### Required

* `tgw_name` - (Required) AWS TGW name.
* `connection_name` - (Required) Connection name.
* `transport_vpc_id` - (Required) Transport Attachment VPC ID.

!> **WARNING:** Attribute `security_domain_name` will be deprecated in future releases. Please use the attribute `network_domain_name` instead. Either `security_domain_name` or `network_domain_name` must be configured.

* `security_domain_name` - (Optional) Security Domain name.
* `network_domain_name` - (Optional) Network Domain name.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `connect_attachment_id` - Connect Attachment ID.
* `transport_attachment_id` - Transport Attachment ID.

## Import

**aws_tgw_connect** can be imported using the `tgw_name` and `connection_name`, e.g.

```
$ terraform import aviatrix_aws_tgw_connect.test tgw_name~~connection_name
```
