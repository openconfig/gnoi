# Bootconfig API

The bootconfig service is responsible for the setting of bootstrap configuration for the device after initial Bootz bootstrap.
This message may include both the base device configuration but also the new gNSI artifacts.

While this is service which can be used out of band it also allows for the development of an industry standard CLI for recovery
of devices.

## RPC use case

* From remote rpc client call `gnoi.Bootconfig.Set` using bootconfig message containing elements to override.

## CLI use cases

The CLI should be bundled with the router and exist in the default path.
The CLI must be named `oc_cli` it must have the following heirarchy for calling specific services.
This CLI must not use RPC's to services to interact with services as the use of the service would be
for emergency recovery of the device.

```
oc_cli
  gnoi
    bootconfig
    factoryreset
  gnmi
    get
    set
    subscribe
```

### Recovery using bootz.BootConfig.Set

* Access the device CLI using SSH or console.
* Use a CLI command to invoke a gnoi bootconfig.

```
cli> scp remotehost:/filename ./local_bootzdata
cli> oc_cli gnoi bootconfig --reboot ./local_bootzdata
Successfully set boot config to ./local_bootzdata
Are you sure? (y/n): y
Rebooting…
```

### Recovery using gnoi.FactoryReset

* Access the device CLI using SSH or console.
* Use a CLI command to invoke a gnoi factory reset.

```
cli> oc_cli gnoi factoryreset
Are you sure? (y/n): y
Factory reset in progress…
```

* This will cause the device to perform the gnoi.FactoryReset process.  The device will return to its factory default configuration and reboot.  When booting, the bootz process should be performed, resulting in a gRPC accessible device.  
