provider "azurerm" {
  version = "=2.20.0"
  features {}
}

data "azurerm_subscription" "current" {}

variable "workload_name" {
  type = string
}

resource "azurerm_resource_group" "rg" {
  name     = "${var.workload_name}-rg"
  location = "West Europe"
}

resource "azurerm_virtual_network" "vnet" {
  name                = "${var.workload_name}-vnet"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  address_space       = ["10.20.0.0/16"]
}

output "resource_group_name" {
  value = azurerm_resource_group.rg.name
}

output "virtual_network_name" {
  value = azurerm_virtual_network.vnet.name
}

output "subscription_id" {
  value = data.azurerm_subscription.current.subscription_id
}