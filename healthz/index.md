# gNOI `Healthz` Streaming RPC Design
**Contributors**: hines@google.com, robjs@google.com, bneville@arista.com  
**Last Updated**: 2022-05-17

## Background

 * [gNOI Repository](https://github.com/openconfig/gnoi)
 * [gNOI `Healthz` service](https://github.com/openconfig/gnoi/tree/master/healthz)

The current Healthz.Get() provides a transaction RPC for a single gnmi single component.


## Motivation

[Original comment](https://github.com/openconfig/gnoi/pull/65#issuecomment-1090547841)

Currently the GNOI Healthz endpoint defines a single transactional API for
getting the "healthz" status of a component.  As discussed in the above issue,
due to large data volume or to an interactive debug output users have requested
the ability to stream a status from the device.  

Ideally we would like to provide a secure facility to stream data which might
contain sensitive data over the channel with proper ACL and encryption. Since
this data may be larger than a single process memory footprint, providing a
stream will allow us to address this.

In discussing this design, I believe another benefit would be to add the
concept of a healthz event history to the infrastructure to allow for the
server to store previous events and allow them to be exported at a later time.
This could be internally generated or externally "checked" via the check api
described below.


## Architecture

#### Healthz.Get() 

Healthz.Get() continues to work as currently described.

We would add an "artifacts" repeated into the proto to describe what if any
artifacts are provided at this endpoint for the healthz event.


```
message ComponentStatus {
  types.Path path = 1; // path of subcomponent.

  // Subcomponents that are aggregated by this status.
  repeated ComponentStatus subcomponents = 2;

  // Status of this component.
  Status status = 3;

  // Opaque data for how the healthcheck is implemented.  This can be any proto
  // defined by the vendor.  This could be the equivalent to outputs like
  // "show tech" or core files or any other diagnostic data.
  google.protobuf.Any healthz = 4 [deprecated=true];

  // Artifacts provides links to all artifacts contained in this event.
  // The individual artifacts can be retrieved via the Artifact() RPC.
  repeated ArtifactHeader artifacts = 5;

  // ID is the unique key for this event in the system.
  string id = 6; 

  // Acknowledged is set when a caller has processed the event.
  bool acknowledged = 7;

  // Create is the timestamp when this event was created.
  google.protobuf.Timestamp created = 8;
  
  // Expires is the timestamp when the system will clean up the
  // artifact. If unset, the artifact is not scheduled for garbage
  // collection.
  google.protobuf.Timestamp expires = 9;
}


message ArtifactHeader {
  // Id is the uuid of the artifact for this event.
  string id = 1;
  // Artifact type describes data contents in the artifact.
  // File artifacts should use the defined FileArtifactType.
  // Proto artifacts should either use the generic ProtoArtifactType
  // or the implementer can provide a specific artifact type
  // which can add any additional metadata the implementor wants.
  ArtifactType artifact_type = 4;
}  

```

#### Healthz.List()

Healthz.List() will list all events for the provided component.


```
service Healthz {
  // List returns all events for the provided component.
  rpc List(ListRequest) returns (ListResponse) {}

}

message ListRequest {
  types.Path path = 1;
  // By default only the unacknowledged events for the component will be returned.
  bool include_acknowledged = 2;
}

message ListResponse {
  repeated ComponentStatus status = 1;
}

```

#### Healthz.Acknowledge()

Healthz.Acknowledge() will acknowledge a specific component and event id.


```
service Healthz {
  // Acknowledge will set the acknowledged field for the event. This is an idempotent operation.
  rpc Acknowledge(AcknowlegeRequest) returns (AcknowledgeResponse) {}

}

message AcknowlegeRequest {
  types.Path path = 1;
  // Healthz event id.
  string id = 2;
}

message AcknowledgeResponse {
  ComponentStatus status = 1;
}

```

#### Healthz.Artifact() 

Healthz.Artifact() introduces a new artifact stream implementation for
streaming artifacts related to Healthz. 


```
service Healthz {
  // Get will return health status for a gNMI path.  If no status is available for
  // the requested path an error will be returned.
  rpc Get(GetRequest) returns (GetResponse) {}

  rpc Artifact(ArtifactRequest) returns (stream ArtifactResponse) {}
}

message ArtifactRequest {
  string id = 1;
}

message ArtifactResponse {
  oneof contents {
    // Header is the first message in the stream. It contains
    // the id of the artifact and metadata for the artifact
    // based on the type of the artifact.
    // OC defines FileArtifactType and ProtoArtifactType.
    ArtifactHeader header = 1;
    ArtifactTrailer trailer = 2;
    bytes bytes = 3;    
    google.protobuf.Any proto = 4;
  }
}

message FileArtifactType {
  // Local file name of the artifact.
  string name = 1;
  // Path to file on the local file system. (optional)
  string path = 2;
  // Mimetype of the file.
  string mimetype = 3;
  // Size of the file.
  int64 size = 4;
  // Hash of the file.
  types.HashType hash = 5;
}

// Generic proto message artifact stream.
// This proto tells the caller that the artifact stream 
// will be a stream of proto encoded messages that make up
// the artifact. Each message must be deserialized by the caller
// and there are no other assumptions about the number of
// messages or length of the stream or how those messages are to
// be reassembled.
message ProtoArtifactType {
}

// ArtifactTrailer is the last message in the artifact stream.
message ArtifactTrailer {
}
```

The general flow would be to check healthz for an endpoint and it would then
return a list of artifacts for that endpoint. Then the user can call the
Artifact endpoint to retrieve those artifacts.  The main reason wanted to
separate would be due to not being clear if the user really wants to directly
handle the artifacts right then and there or to provide a signal for some other
system to fetch and handle the artifacts separately.

This also raises issues around the history of artifacts and how to clean up
those artifacts over time.  Likely we should look at how to extend the healthz
to include a history of events and a way to "acknowledge" and "clean up"
previous artifacts.

#### Healthz.Check()

Healthz.Check() introduces a new endpoint to execute a set of "validations" against the gnmi path.

The result of the Check() will produce a gnmi healthz status message which will
contain a list of the generated artifacts used in the validation process.


```
service Healthz {
  rpc Check(CheckRequest) returns (CheckResponse) {}
}

message CheckRequest{
  types.Path path = 1;
}

message CheckResponse{
  ComponentStatus status =1;
}

```

The Check() endpoint will allow for "very expensive" debugging commands such as
causing the device to pause and snapshot its system database or to examine the
state of a protocol.  These commands will likely be considered "service
impacting" and should have a clear security ACL restricting their use during
normal operations of the device.

## User Experience

#### A BGP routing process goes unhealthy due to a crash.

##### Expected Action

Write the core dump to a file system.

A call to Healthz.Get should respond with the current state of the component
and provide feedback that there is an artifact available to be requested.

A Healthz call should stream the core dump over the service back to the caller.

#### A chassis linecard goes unhealthy due to hardware failure.

##### Expected Action

Logs / dumps collected and the gNMI components are marked unhealthy.

A call to Healthz.Get should respond with the current state of the component
and provide feedback that there is an artifact available to be requested.

A Healthz call should stream the core dump over the service back to the caller.

A Healthz call should stream the logs over the service back to the caller.


#### A chassis file storage is over a configured capacity and reports itself unhealthy.

Since there are no artifacts the Healthz.Get should just return the status of
the component and no other action is necessary

#### A user would like to stream console debugging or other shell output for collection

For a component this could be a just io dump of a specific debugging command
redirected as a byte stream.  It could also be output like a top or show
interfaces where the data is periodically updated (each "frame" would be sent
over the stream as a separate message).  This would allow for tech like
information to be collected and sent to vendors in a programmatic secure way
but not completely reinventing the wheel.


#### Collect a snapshot of a particular component.

For components which can provide a "snapshot" of state the healthz artifact
endpoint can be provided to take a snapshot of state (hopefully serialized as a
proto or other typed data structure) and provided to the caller. 

Examples:

 * Get optical data from i2c bus on optical systems
 * Get system database state for the core pub/sub of a device


#### Exporting internal or specific debug logs

For a component in the system the developers could provide a debug log for that
specific component or other trace / dump information.  The Check() and
Artifact() apis could be used in these cases to create the setup and then
download the specific data in a secure, reliable way vs just a general file
scp. Additionally if the device can "automatically" generate these specific
data, the Get api could also be used to check for existence and to download the
artifacts. To support this we would also likely need additional gnmi paths to
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


## Open Questions/Considerations

*   Should the stream support suspend / resume?

This feels expensive to implement on the server but would be open to it if
vendors think it would be a useful mode

*   Should the stream be completely opaque to the server?
*   What are the stream types that should be "typed"?
    *   Byte data?
    *   Proto data?
    *   Interactive stream data?

I think the protocol should be standard and the data frames can be opaque as
either any proto's or bytes

*   Should this be exposed as a separate RPC endpoint from Get to allow for
    "chunking" up the data resources provided to the consumer.

I think we should have a separate stream api from the top level get status

*   Should interactive debug data be able to be streamed back over this endpoint?

Currently only planning on the artifact creation and then sending over the artifact interface.

The problem with supporting an arbitrary protocol over the single endpoint
would be confusing to the user and likely to have significant 

*   how do we force or keep the system clean - if multiple service are all
    "saving artifacts" how do we not just create our own outages by filling up
    storage

Added Acknowledge for the support of a way to clean up old events and
artifacts.

There is still a question on adding a "ttl" signal but we need to probably look
at how serverside we would expect this to be implemented.
