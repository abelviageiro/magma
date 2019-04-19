---
id: configurations
sidebar_label: Configurations
title: Configurations in Magma
hide_title: true
---
# Configurations in Magma
### Cloud 
On the cloud side, service related configuration files are stored in 
`magma/{MODULE_NAME}/cloud/configs/`.

#### Service Registry 
`magma/{MODULE_NAME}/cloud/configs/service_registry.yml` lists all services in 
the module and stores configurations that all services must have (i.e. port, host).
The information is used for services routing. 

#### Service Specific Configs
All service specific configurations are stored in 
`magma/{MODULE_NAME}/cloud/configs/{SERVICE_NAME}.yml`. 

#### How To Modify Configs
When the cloud VM is provisioned, the service specific configuration files are 
sym-linked into `/etc/magma/configs/{MODULE_NAME}/`. 
The configs can be modified in both `/etc/magma/configs/` and `/var/opt/magma/configs/`, but
the latter takes priority over the other.

Every time a service starts, both the per-service configs and service registries 
are loaded. Restart the corresponding service after modifying configs to see the changes.

### Gateway 
On the Gateway side, service related configuration files are stored in 
`magma/lte/gateway/configs/`

#### Service Registry 
`magma/lte/gateway/configs/service_registry.yml` lists all services on the gateway
 and stores configurations that all services must have. 
 The information is used for services routing. 
 
#### Service Specific Configs
All service specific configurations are stored in 
`magma/lte/gateway/configs/{SERVICE_NAME}.yml`. 

#### How To Modify Configs
When the magma VM is provisioned, the service specific configuration files are 
sym-linked into `/etc/magma/configs/`. 
The configs can be modified in both `/etc/magma/configs/` and `/var/opt/magma/configs/`, but
the latter takes priority over the other.

Restart the service to see the the changes reflected.