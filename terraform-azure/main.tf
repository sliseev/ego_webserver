terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "ego-rg" {
  name     = "ego-resources"
  location = "West Europe"
  tags = {
    environment = "dev"
  }
}

resource "azurerm_container_group" "ego-cg" {
  name                = "ego-containers"
  location            = azurerm_resource_group.ego-rg.location
  resource_group_name = azurerm_resource_group.ego-rg.name
  ip_address_type     = "Public"
  os_type             = "Linux"

  container {
    name     = "egoserv"
    image    = "sergcpp/egoserv:latest"
    cpu      = "0.25"
    memory   = "0.5"
    commands = ["./ego_server", "-standalone"]

    ports {
      port     = 8080
      protocol = "TCP"
    }
  }

  tags = {
    environment = "dev"
  }
}

output "container_ip_address" {
  value = azurerm_container_group.ego-cg.ip_address
}
