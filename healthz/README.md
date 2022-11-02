# gNOI `Healthz` Streaming RPC Design
**Contributors**: hines@google.com, robjs@google.com, bneville@arista.com  
**Last Updated**: 2022-05-17

## Background

 * [gNOI Repository](https://github.com/openconfig/gnoi)
 * [gNOI `Healthz` service](https://github.com/openconfig/gnoi/tree/master/healthz)

The purpose of Healthz is to allow a component inside of a system to report its
health. The concept of asking "Are you healthy?" is a general design principle
in distributed systems.

The ability to also include "implementation specific" details about the health
of the system or component is very standard for network devices in the form for
"show tech", "diag" or "debug" commands. These are very useful for getting
diagnostics about the system state to the implementers. Healthz exposes
these interfaces as queriable endpoints to allow operators to validate "health"
on components and if "unhealthy" gather implementor specific data to help
triage or reproduce issues.

## Design Background

`Healthz` provides a means by which a user may initiate checks of health
(through the `Check` RPC) or a system may report the results of a check that it
has initiated of its own accord. 

Following a health check occurring, a caller uses the `List` or `Get` RPC to discover the
artifacts that are assoicated with a component and its corresponding
subcomponents. Each artifact reflects an entity that is required to debug or
further root cause the fault that occurs with it.

The `Artifact` RPC is used to retrieve specific artifacts that are listed by
the target system in the `List` RPC. Once retrieved each artifact can be
'acknowledged' by the client of the RPC. Acknowledged artifacts are no longer
returned in a list of the corresponding artifacts, and a device may use this
status as a hint to allow garbage collection of artifacts that are no longer
relevant. The device itself is responsible for garbage collection any may,
if necessary, garbage collect artifacts that are not yet acknowledged.

Whilst the system may initiate health checks itself, these should be safe to
perform operations that do not impact the device's functionality. Expensive
checks (e.g., pausing protocols, or dumping internal database state) that are
potentially service impacting should require use of the `Check` RPC.

`Healthz` works in conjunction with telemetry streamed via gNMI, OpenConfig
paths for a specific component are streamed to indicate when components become
unhealthy, allowing the receiving system to determine that further inspection
of the component's health is required. The following paths (defined in the
`openconfig-platform` model) are used:

```
component/healthz/state/status          (enumerated value)
component/healthz/state/last-unhealthy  (timestamp)
component/healthz/state/unhealthy-count (int64)
```

### Choice of Streaming for Artifacts


[Background comment](https://github.com/openconfig/gnoi/pull/65#issuecomment-1090547841)

Originally, the gNOI Healthz endpoint defined single transactional API for
getting the "healthz" status of a component.  As discussed in the above issue,
due to large data volume or to an interactive debug output there is a
requirement that it support the ability to stream a status from the device.  

To ensure that this can be done using a secure channel with the relevant
ACL and encryption, this streaming is implemented within the `Healthz` service.
This approach further allows the concept of a historical set of collected
articacts to be reported by the device.


## Architecture

### Use of gNMI paths

Where a gNMI path (`gnoi.types.Path`) is used in `Healthz`, the path specified
should be the complete path to a specific component - i.e.,
`/components/component[name=FOO]`.

### Healthz.Get() 

```protobuf
rpc Get(GetRequest) returns (GetResponse) {}
```
[(Definition)](https://github.com/openconfig/gnoi/blob/671c4286820310a6fa7b016124ea248bf5cfbe76/healthz/healthz.proto#L44)

`Get` allows a user to retrieve the latest set of health statuses that are associated
with a specific component and any of its corresponding subcomponents. In contrast to
`List` the `Get` RPC should return only the latest event.

The `GetResponse` returned includes a `ComponentStatus` message which
corresponds to the latest health event for the  component itself and each subcomponent
being reported on. Thus, multiple `ComponentStatus` messages may be reported for
a single component. Each message includes a set of `ArtifactHeader` messages
that correspond to the health event -- and provide identifiers and types for the
artifacts that can be returned by the system.

All artifacts that are listed within the same `ComponentStatus` message are
expected to have a shared acknowledgement state, and expiry time.

### Healthz.List()

```protobuf
rpc List(ListRequest) returns (ListResponse) {}
```

`List` returns all events that are associated with a particular component. In
contrast to `Get`, it returns all events that the target has for the device. By
default, events that are acknowledged are not returned.

Similarly to `Get`, a series of `ComponentStatus` messages are returned with
the same semantics as are described above.

### Healthz.Acknowledge()

```protobuf
rpc Acknowledge(AcknowledgeRequest) returns (AcknowledgeResponse) {}
```

`Acknowledge` is used by a client to indicate to the target device that a
particular (component, event) tuple has been retrieved by the client. This
allows a device to intelligently determine whether to retain artifacts. Devices
MUST ensure that artifact storage for healthz does not cause resource exhaustion
and SHOULD remove acknowledged artifacts before those that have not yet received
an acknowledgement.

### Healthz.Artifact() 

```protobuf
rpc Artifact(ArtifactRequest) returns (stream ArtifactResponse) {}
```

`Artifact` allows a user to retrieve a specific artifact that is related with
an event that has occurred. Since these artifacts may be large, the `Artifact`
RPC is implemented as a server-side streaming RPC. Use of the `Artifact` RPC
ensures that the resources to send these potentially large artifacts only
when explicitly requested by the client.

#### Healthz.Check()

```protobuf
rpc Check(CheckRequest) returns (CheckResponse) {}
```

`Check` allows a client to execute a set of "validations" against the specified
component. The component, as with other operations, is specified in terms of
the gNMI path.  

The result of the `Check` produces a healthz `ComponentStatus` message which
will contain a list of the generated artifacts used in the validation process.


A call to the `Check` endpoint will allow for "very expensive" debugging
commands such as causing the device to pause and snapshot its system database
or to examine the state of a protocol.  These commands will likely be
considered "service impacting" and should have a clear security ACL restricting
their use during normal operations of the device.

## User Experience

#### A BGP routing process goes unhealthy due to a crash.

##### Expected Action

Write the core dump to a file system.

A call to `Healthz.Get` should respond with the current state of the component
and provide feedback that there is an artifact available to be requested.

A Healthz call should stream the core dump over the service back to the caller.

The component that is specified should correspond to the software process that
implements the BGP protocol in the system.

#### A chassis linecard goes unhealthy due to hardware failure.

##### Expected Action

Logs / dumps collected and the gNMI components are marked unhealthy.

A call to Healthz.Get should respond with the current state of the component
and provide feedback that there is an artifact available to be requested.

A Healthz call should stream the core dump over the service back to the caller.

A Healthz call should stream the logs over the service back to the caller.

The component that is reported on should be of type
[`CHASSIS`](https://openconfig.net/projects/models/schemadocs/yangdoc/openconfig-platform.html#ident-chassis).


#### A chassis file storage is over a configured capacity and reports itself unhealthy.

Since there are no artifacts the Healthz.Get should just return the status of
the component and no other action is necessary

The component reported on via telemetry should correspond to the device's chassis
and be of type
[`CHASSIS`](https://openconfig.net/projects/models/schemadocs/yangdoc/openconfig-platform.html#ident-chassis).

#### A user would like to stream console debugging or other shell output for collection

For a component healthz could be used to provide an I/O dump of a specific
debugging command redirected as a byte stream.  It could also be output like a
top or show interfaces where the data is periodically updated (each "frame"
would be sent over the stream as a separate message).  This would allow for
show tech like information to be collected and sent to vendors in a
programmatic secure way but not completely reinventing the wheel.

To achieve this use case, extensions to the existing specification are required.

#### Collect a snapshot of a particular component.

For components which can provide a "snapshot" of state the healthz artifact
endpoint can be provided to take a snapshot of state (ideally serialized as a
proto or other typed data structure) and provided to the caller. 

Examples:

 * Get optical data from i2c bus on optical systems
 * Get system database state for the core pub/sub of a device

#### Exporting internal or specific debug logs

For a component in the system the developers could provide a debug log for that
specific component or other trace / dump information.  The `Check()` and
`Artifact()` APIs could be used in these cases to create the setup and then
download the specific data in a secure, reliable way (rather than a general file
scp). If the device can "automatically" generate these specific
data the `Get` API can also be used to check for existence and to download the
artifacts. Note that to support this use case, asdditional gNMI paths are likely required to
signal existence of this type of data.

#### Event lifecycle

An event is created internally by the system for a linecard rebooting
unexpectedly. The chassis process will take a core of the component.  Snapshot
it's log for the component. Snapshot any relevant state in sysdb for the
component and create a healthz event which includes pointers to these 3
artifacts.

The external monitoring system will a gnmi update with the healthz of the
component going unhealthy.

An event monitoring agent can then call healthz get on the component and see
the current event and the associated artifacts.  Since this externally
reportable event (i.e. we might want to rma this component) the system decides
to grab the artifacts and store them with a buganizer case.  Once the artifacts
are retrieved the system will mark the event as acknowledged and the server can
then use whatever retention policy it has configured to clean up those events.
