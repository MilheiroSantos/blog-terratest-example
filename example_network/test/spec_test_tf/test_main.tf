provider "azurerm" {
  version = "=2.20.0"
  features {}
}

variable "workload_name" {
  type = string
}

data "azurerm_subscription" "current" {}

module "network" {
  source        = "../../module"
  workload_name = var.workload_name
}

output "resource_group_name" {
  value = module.network.resource_group_name
}

output "virtual_network_name" {
  value = module.network.virtual_network_name
}

output "subscription_id" {
  value = data.azurerm_subscription.current.subscription_id
}
