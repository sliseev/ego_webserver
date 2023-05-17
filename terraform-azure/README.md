Deploy to Azure Cloud
=====================

This is the second version of deployment of the only one
container egoserv from Docker Hub in standalone mode. It
works with paid subscription due to logs.

Requirements:
- Terraform (https://developer.hashicorp.com/terraform/downloads)
- Azure-cli (https://learn.microsoft.com/en-us/cli/azure/install-azure-cli-linux)

```
$ cd terraform-azure
$ az login
$ terraform init
$ terraform plan
$ terraform apply

# get 'container_address' from output
# check the service is up and running:
$ curl http://<container_address>/driver/count
```
