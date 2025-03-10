output "AVIATRIX_CONTROLLER_IP" {
  value = module.aviatrix-controller-build.public_ip
}

output "AVIATRIX_USERNAME" {
  value = "admin"
}

output "AVIATRIX_PASSWORD" {
  value = var.admin_password
}

output "AWS_ACCOUNT_NUMBER" {
  value = data.aws_caller_identity.current.account_id
}

output "AWS_ACCESS_KEY" {
  value = var.aws_access_key
}

output "AWS_SECRET_KEY" {
  value = var.aws_secret_key
}

output "AWSGOV_ACCOUNT_NUMBER" {
  value = var.awsgov_account_number
}

output "AWSGOV_ACCESS_KEY" {
  value = var.awsgov_access_key
}

output "AWSGOV_SECRET_KEY" {
  value = var.awsgov_secret_key
}

output "ARM_SUBSCRIPTION_ID" {
  value = var.azure_subscription_id
}

output "ARM_DIRECTORY_ID" {
  value = var.azure_tenant_id
}

output "ARM_APPLICATION_ID" {
  value = var.azure_client_id
}

output "ARM_APPLICATION_KEY"{
  value = var.azure_client_secret
}

output "GCP_ID" {
  value = var.gcp_project_id1
}

output "GCP_CREDENTIALS_FILEPATH"{
  value = var.gcp_credentials_file_path
}

output "AWS_BGP_VGW_ID" {
  value = aws_vpn_gateway.vgw.id
}

output "GCP_VPC_ID" {
  value = module.aviatrix_gcp_vpc1.vpc_id
}

output "GCP_SUBNET" {
  value = module.aviatrix_gcp_vpc1.subnet
}

output "GCP_ZONE" {
  value = var.gcp_zone1
}

output "AZURE_REGION" {
  value = var.azure_region1
}

output "AZURE_VNET_ID" {
  value = "${module.aviatrix_azure_vpc1.vnet}:${module.aviatrix_azure_vpc1.group}:${module.aviatrix_azure_vpc1.guid}"
}

output "AZURE_SUBNET" {
  value = module.aviatrix_azure_vpc1.subnet
}

output "AZURE_REGION2" {
  value = var.azure_region2
}

output "AZURE_VNET_ID2" {
  value = "${module.aviatrix_azure_vpc2.vnet}:${module.aviatrix_azure_vpc2.group}:${module.aviatrix_azure_vpc2.guid}"
}

output "AZURE_SUBNET2" {
  value = module.aviatrix_azure_vpc2.subnet
}

output "AZURE_GW_SIZE" {
  value = var.azure_gw_size
}

output "AWS_VPC_ID" {
  value = module.aviatrix_aws_vpc1.vpc
}

output "AWS_SUBNET" {
  value = module.aviatrix_aws_vpc1.subnet
}

output "AWS_VPC_ID2" {
  value = module.aviatrix_aws_vpc2.vpc
}

output "AWS_SUBNET2" {
  value = module.aviatrix_aws_vpc2.subnet
}

output "AWS_VPC_ID3" {
  value = module.aviatrix_aws_vpc3.vpc
}

output "AWS_SUBNET3" {
  value = module.aviatrix_aws_vpc3.subnet
}

output "AWS_VPC_ID4" {
  value = module.aviatrix_aws_vpc4.vpc
}

output "AWS_SUBNET4" {
  value = module.aviatrix_aws_vpc4.subnet
}

output "AWS_REGION" {
  value = data.aws_region.current.name
}

output "AWS_REGION2" {
  value = data.aws_region.current.name
}

output "AWS_DX_GATEWAY_ID" {
  value = aws_dx_gateway.dx-gateway.id
}

output "DOMAIN_NAME" {
  value = var.domain_name
}

output "AWSGOV_VPC_ID" {
  value = module.aviatrix_aws_vpc1.vpc
}

output "AWSGOV_SUBNET" {
  value = module.aviatrix_aws_vpc1.subnet
}

output "AWSGOV_REGION" {
  value = data.aws_region.current_awsgov.name
}

output "OCI_VPC_ID" {
  value = module.aviatrix_oci_vpc1.vpc_id
}

output "OCI_REGION" {
  value = var.oci_region1
}

output "OCI_SUBNET" {
  value = module.aviatrix_oci_vpc1.subnet
}

output "OCI_TENANCY_ID" {
  value = var.oci_tenancy_id
}

output "OCI_USER_ID" {
  value = var.oci_user_id
}

output "OCI_COMPARTMENT_ID" {
  value = var.oci_compartment_id
}

output "OCI_API_KEY_FILEPATH" {
  value = var.oci_api_key_filepath
}

output "IDP_METADATA" {
  value = var.IDP_METADATA
}

output "IDP_METADATA_TYPE" {
  value = var.IDP_METADATA_TYPE
}

output "DEVICE_PUBLIC_IP" {
  value = module.cisco-csr.DEVICE_PUBLIC_IP
}

output "DEVICE_KEY_FILE_PATH" {
  value = module.cisco-csr.DEVICE_KEY_FILE_PATH
}

output "TRANSIT_GATEWAY_NAME" {
  value = aviatrix_transit_gateway.cwan-transitgw.gw_name
}

output "ARM_RESOURCE_GROUP" {
  value = module.azure-vwan.azure_resource_group
}

output "ARM_HUB_NAME" {
  value = module.azure-vwan.azure_hub_name
}

output "AWS_TGW_NAME" {
  value = aviatrix_aws_tgw.cwan-awstgw.tgw_name
}

output "DATADOG_API_KEY" {
  value = var.datadog_api_key
}

output "AZURE_VNG_VNET_ID" {
  value = "${module.azure-vng.vnet}:${module.azure-vng.resource_group}"
}

output "AZURE_VNG_SUBNET" {
  value = module.azure-vng.subnet
}

output "AZURE_VNG" {
  value = module.azure-vng.vng
}

output "EDGE_SPOKE_NAME" {
  value = var.edge_spoke_name
}

output "EDGE_SPOKE_SITE_ID" {
  value = var.edge_spoke_site_id
}

output "EDGE_CSP_USERNAME" {
  value = var.edge_spoke_site_id
}

output "EDGE_CSP_PASSWORD" {
  value = var.edge_spoke_site_id
}

output "EDGE_CSP_PROJECT_UUID" {
  value = var.edge_spoke_site_id
}

output "EDGE_CSP_COMPUTE_UUID" {
  value = var.edge_spoke_site_id
}

output "EDGE_CSP_TEMPLATE_UUID" {
  value = var.edge_spoke_site_id
}
