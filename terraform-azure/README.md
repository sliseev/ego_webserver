Deploy to Azure Cloud
=====================

This is the simplest possible deploy of the only one
container egoserv from Docker Hub in standalone mode.

Requirements:
- Terraform (https://developer.hashicorp.com/terraform/downloads)
- Azure-cli (https://learn.microsoft.com/en-us/cli/azure/install-azure-cli-linux)

```
$ cd terraform-azure
$ az login
$ terraform init
$ terraform plan
$ terraform apply

# get 'container_ip_address' from output
# check the service is up and running:
$ curl http://<container_ip_address>:8080/driver/count
```
