locals {
  tags = {
    "environment" = "dev"
  }
}

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.54.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "ego-rg" {
  name     = "ego-resources"
  location = "West Europe"
  tags     = local.tags
}

resource "azurerm_log_analytics_workspace" "ego-aw" {
  name                = "ego-analytics"
  location            = azurerm_resource_group.ego-rg.location
  resource_group_name = azurerm_resource_group.ego-rg.name
  retention_in_days   = 30
  tags                = local.tags
}

resource "azurerm_container_app_environment" "ego-cenv" {
  name                       = "ego-container-environment"
  location                   = azurerm_resource_group.ego-rg.location
  resource_group_name        = azurerm_resource_group.ego-rg.name
  log_analytics_workspace_id = azurerm_log_analytics_workspace.ego-aw.id
  tags                       = local.tags
}

resource "azurerm_container_app" "ego-app" {
  name                         = "ego-app"
  container_app_environment_id = azurerm_container_app_environment.ego-cenv.id
  resource_group_name          = azurerm_resource_group.ego-rg.name
  revision_mode                = "Single"
  tags                         = local.tags

  template {
    container {
      name   = "egoserv"
      image  = "sergcpp/egoserv:latest"
      cpu    = 0.25
      memory = "0.5Gi"
    }
  }

  ingress {
    allow_insecure_connections = true
    external_enabled           = true
    target_port                = 8080
    traffic_weight {
      percentage = 100
    }
  }
}

output "container_address" {
  value = azurerm_container_app.ego-app.ingress[0].fqdn
}
