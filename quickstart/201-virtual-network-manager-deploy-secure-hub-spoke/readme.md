# Azure resource group

This template deploys an Azure Virtual Network Manager instance with a connectivity configuration for a Mesh topology. It includes resources including virtual networks, subnets, and more.

## Terraform resource types

- [random_pet](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet)
- [azurerm_resource_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_group)
- [azurerm_virtual_network](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network)
- [azurerm_subnet](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet)
- [azurerm_virtual_network_manager](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager)
- [azurerm_network_manager_network_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_network_group)
- [azurerm_network_manager_static_member](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_static_member)
- [azurerm_network_manager_connectivity_configuration](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_connectivity_configuration)
- [azurerm_network_manager_deployment](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_deployment)
- [azurerm_public_ip](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/public_ip)
- [azurerm_public_ip_prefix](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/public_ip_prefix)
- [azurerm_nat_gateway](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/nat_gateway)
- [azurerm_nat_gateway_public_ip_association](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/nat_gateway_public_ip_association)
- [azurerm_subnet_nat_gateway_association](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet_nat_gateway_association)
- [azurerm_nat_gateway_public_ip_prefix_association](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/nat_gateway_public_ip_prefix_association)
- [azurerm_policy_definition](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/policy_definition)
- [azurerm_subscription_policy_assignment](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subscription_policy_assignment)
- [azurerm_network_manager_admin_rule](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_admin_rule)
- [azurerm_network_manager_admin_rule_collection](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_manager_admin_rule_collection)
- [azurerm_network_security_admin_configuration](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_admin_configuration)

## Variables

| **Name** | **Description** | **Default** |
|---|---|---|
| `resource_group_name_prefix` | Prefix of the resource group name that's combined with a random ID so name is unique in your Azure subscription. | rg |
| `resource_group_location` | Location of the resource group. | eastus |

## Example

To see how to run this example, see [Quickstart: Deploy a Virtual Network Manager in Azure using Terraform](https://learn.microsoft.com/azure/virtual-network-manager/create-virtual-network-manager-terraform).