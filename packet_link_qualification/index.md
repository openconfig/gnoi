## Link Qualification Service v2 Definition

[TOC]

### Background

The link qualification service is provides a way to certify link
quality on two devices. This service defines a protocol for setting up the peer
devices to allow for the generation of packets from one device to another and
validate those packets are sent and received, then restoring the devices to
their previous state. This service's intent is to allow for different
generation and reflection modes of operation based on the hardware capabilities
of the device. There is a standard report generated regardless of modes
selected and that common report can be used by upstream services to aggregate
network wide link quality.

### TLDR

#### General flow

#### Specific requirements around the service

##### Connectivity to devices during link qualification maybe interrupted

Upon calling Create, the interfaces on the device will be put into a forwarding
mode which must not contain other traffic, control or data. This may cause the
generator or reflector endpoint to become unreachable.  The service
implementation must gracefully handle this state.


##### Devices must return to pre-link qualification state after the link qualification has completed or errored

The service is not expected to persist state across reboots of the device.
Since no state is changed in the configuration of the device, on boot the
system should boot normally and there will be no side effects left from
previous link qualifications.

During the link qualification the device must put the interfaces into a
`TESTING` `oper-status` for the duration of the qualification. Once complete,
the `oper-status` should be restored to the previous state.

##### Devices must garbage collect results after some period to keep from filling up storage on the device.

The service should store at least 2 qualification results for each interface.
The results are not expected to persist across reboots of the device.


##### The service implementation must support multiple generations of forwarding hardware which have different capabilities with regards to both generation and reflection of packets.

*   Packet Generators
The preferred method for the generation of frames for the link qualification is
via a packet generator. This mode provides the most flexibility around packet
rates, packet sizes and packet content. 

*   Packet Injectors
To support hardware which doesn't have a built in packet generator, link
qualification provides a packet injector mode to support a loop between both
endpoints in the qualification for packets to be injected and looped.

*   ASIC Loopback
The ASIC loopback mode allows for the loopback of frames to happen in the
forwarding path of the device.  This enables the most comprehensive testing
results as all counters should be available to the qualification.  This is the
preferred mode of reflection.

*   PMD Loopback
For devices which cannot implement an ASIC based loopback. The PMD loopback
mode can be supported.  This mode allows for the interface to be looped back at
either the PMD or in the interface.  This will likely limit what counters are
available for results on the reflector but can still validate the link.

### Definition

#### Service Definition


```
service LinkQualification {
  // Create will dispatch a create operation for each interface and return.
  // The rpc will only return an error in the case that gNOI service cannot
  // handle the RPC request. Create will return an error on failure to
  // create the 
  rpc Create(CreateRequest) returns (CreateResponse);

  // Get will return the status for the provided qualification ids.
  rpc Get(GetRequest) returns (GetResponse);

  // Capabilities will return the  capabilities of the gnoi.LinkQualification
  // service implementation.  This RPC is used to allow the caller to 
  // orchestrate the peer requirements of the service to complete a link
  // qualification between two devices.
  rpc Capabilities(CapabilitiesRequest) returns (CapabilitiesResponse);
  
  // Delete will remove the qualification results for the provided based on the
  // ids provided. If the qualification is not in QUALIFICATION_STATE_COMPLETED
  // or QUALIFICATION_STATE_ERROR, the qualification will be canceled then
  // deleted as requested.
  // If the qualification cannot be stopped or deleted a status will be returned
  // with the error. 
  // If the id is not found NOT_FOUND will be returned. 
  rpc Delete(DeleteRequest) returns (DeleteResponse);

  // List qualifications currently on the target.
  rpc List(ListRequest) returns (ListResponse);
}

```

#### Create Messages 

```
// CreateRequest contains the list of interfaces to be Qualified.
// Each interface must only appear once and all IDs to be used
// by qualification must be unique on the device.
message CreateRequest {
  repeated QualificationConfiguration interfaces = 1;
}

// CreateResponse will return a map of the status of each CreateRequest.
// The map key is the qualification id requested in the CreateRequest.
// Valid Status responses are:
// OK: create has been accepted.
// NOT_FOUND: interface_name could not be found on the service.
// INVALID_ARGUMENT: if any of the configuration is not supported.
// ALREADY_EXISTS: if the qualification id is already in use.
message CreateResponse {
  map<string, google.rpc.Status> status = 2;
}

message NTPSyncedTiming {
  // The timestamp for the start of the qualification. Based on the
  // provided configuration the qualification setup must
  // at least start at (start_time - min_setup_time). 
  google.protobuf.Timestamp start_time = 1;
  // The timestamp for the end of qualification.
  google.protobuf.Timestamp end_time = 2;
  // Timestamp to begin teardown. If unset teardown will start
  // immediately after end_time.
  // This value allows for a peer to wait for before tearing down
  // the port under test.
  // teardown_time must occur after end_time.
  google.protobuf.Timestamp teardown_time = 3;
}

// RPCSyncedTiming will be all synchronization by assuming the start rpc's are
// sent very close temporally with enough overlap in duration to get valid
// results.
// The pre_qual_wait will allow the caller to adjust any setup timing
// differences between peers. The post_qual_wait will allow for the caller to
//  adjust any teardown differences in timing between peers.
// pre_qual_wait cannot be less than selected endpoint type's min_setup_time.
// post_qual_wait cannot be less than the selected endpoint type's
// min_teardown_time.
message RPCSyncedTiming {
  // pre_sync_duration is the time the service should wait before starting
  // setup.
  // For generators this would be the time the remote side needs to put itself
  // into loopback before generating packets.
  // For loopbacks this value would be expected to be zero or unset.
  google.protobuf.Duration pre_sync_duration = 1;

  // The requested setup time for the endpoint. setup_duration must be >=
  // min_setup_duration in the service capabilities. If the service
  // completes setup before setup_duration it must wait setup_duration
  // before moving to qualification. 
  google.protobuf.Duration setup_duration = 2;

  // duration is the length of the qualification.
  google.protobuf.Duration duration = 3;

  // post_sync_duration is the amount time a side should wait before starting
  // its teardown.
  // For generators this value would be expected to be zero or unset.
  // For loopbacks this would be the duration it takes for the generator
  // to stop sending packets before starting a teardown.
  google.protobuf.Duration post_sync_duration = 4;
 
  // This requested teardown duration for the endpoint. teardown_duration
  // must be >= min_teardown_duration in the service capabilities. If the
  // service completes before the teardown_duration it must wait teardown
  //_duration before completing.
  google.protobuf.Duration teardown_duration = 5;
}

message QualificationConfiguration {
  // Id to be used for this interface qualification run.
  string id = 1;

  // interface name on the device must be unique to the device.
  string interface_name = 2;

  // timing allows for specifying either NTP or by using the Create RPC as 
  // the synchronization method for the qualification.
  oneof timing {
    NTPSyncedTiming ntp = 101;
    RPCSyncedTiming rpc = 102;
  }

  // endpoint_type is how this side of the link will be configured for the
  // qualification.
  oneof endpoint_type {
    PacketGeneratorConfiguration packet_generator = 111;
    PacketInjectorConfiguration packet_injector = 112;
    PmdLoopbackConfiguration pmd_loopback = 113;
    AsicLoopbackConfiguration asic_loopback = 114;
  }
}

// A packet generator implementation defines that the generator of the side of
// the link will be responsible for generating a stream of packets at the
// provided rate and size.
message PacketGeneratorConfiguration {
  // Packet per second rate to use for this test.
  uint64 packet_rate = 1;

  // Size of packets to inject.
  // if unspecified, the default value is 1500 bytes.
  uint32 packet_size = 2;
}

// A packet injector implementation defines that the generator side of the link
// will be responsible for both setting the interface into a loopback as well
// as injecting individual packets up to packet_count into the closed loop at
// the packet_size. These packets will form a closed loop as both sides of the
// loop will forward the same set of packets for the duration of the
// qualification.
message PacketInjectorConfiguration {
  // Number of packets to inject into the closed loop.
  uint32 packet_count = 1;

  // Size of packets to inject.
  // if unspecified, the default value is 1500 bytes.
  uint32 packet_size = 2;

  // Loopback mode for this qualification.
  oneof loopback_mode  {
    // PMD based loopback encompass either PHY based port
    // loopbacks or port based ASIC loopbacks which do
    // use forwarding engine processing.
    // Their use may limit the stats available for the
    // qualification.
    PmdLoopbackConfiguration pmd_loopback = 101;
    // ASIC based loops are done inside the forwarding
    // engine and must have stats available to the
    // qualification.
    AsicLoopbackConfiguration asic_loopback = 102;    
  }
}


message PmdLoopbackConfiguration {
}

message AsicLoopbackConfiguration {
  // This is where any l2/l3/l4 match criteria would be described.
}
```

#### Get Messages

```
// GetRequest returns the status for the provided ids.
message GetRequest{
  repeated string ids =1;
}

// GetResponse returns a map containing the values for all requested
// Qualification results. If the QualificationResult state is
// QUALIFICATION_STATE_ERROR the caller should inspect the status field for
// the exact error code and message.
// Expected errors codes:
// NOT_FOUND when the requested id was not found by the service.
// INVALID_ARGUMENT for any configuration parameter which is unsupported.
message GetResponse{
  map<string, QualificationResult> results = 1;
}

// States of qualification.
enum QualificationState {
  QUALIFICATION_STATE_UNSPECIFIED = 0;
  QUALIFICATION_STATE_ERROR = 1;  // The qualification has errored.
  QUALIFICATION_STATE_IDLE = 2;       // Initial state for the qualification.
  QUALIFICATION_STATE_SETUP = 3;       // Interface is being configured.
  QUALIFICATION_STATE_RUNNING = 4;    // Qualification underway.
  QUALIFICATION_STATE_TEARDOWN = 5;    // Interface is being reset.
  QUALIFICATION_STATE_COMPLETED = 6;  // Qualification is complete.
}

message QualificationResult {
  // ID of the qualification.
  string id = 1;

  // Interface name of the qualification.
  string interface_name = 2;

  // The state the test was in when the results were snapshotted.
  QualificationState state = 3;

  // The number of qualification packets sent.
  uint64 packets_sent = 4;

  // The number of qualification packets received.
  uint64 packets_received = 5;

  // The number of packets transmitted that experienced corruption.
  uint64 packets_error = 6;

  // The number of packets dropped by the device due to internal drop,
  // lookup or forwarding errors.
  uint64 packets_dropped = 7;

  // The qualification start time. Also when the first snapshot of
  // results are taken.
  google.protobuf.Timestamp start_time = 8;

  // The qualification end time or the current snapshot time since epoch.
  google.protobuf.Timestamp end_time = 9;

  // Expected rate for the qualification. This is the computed or
  // observed rate that the service expected to be maintained
  // throughout the qualification duration.
  uint64 expected_rate_bytes_per_second = 10;

  // The qualification rate observed during the qualification. 
  // It is updated every update_interval in bytes per second.
  uint64 qualification_rate_bytes_per_second = 11;

  // Status response for the Qualification result. Will only be set if the
  // state is QUALIFICATION_STATE_ERROR.
  google.rpc.Status status = 12;
}
```

#### Delete Messages

```
// DeleteRequest will delete the qualification results for the provided id.
message DeleteRequest{
  repeated string ids = 1;
}

// Delete response will contain a list of all id's in the request to be deleted.
// If the id was not found NOT_FOUND will be returned.
message DeleteResponse{
  map<string, google.rpc.Status> results = 1;
}

```

#### Capabilities Messages

```
message CapabilitiesRequest {
}

message CapabilitiesResponse {
  // Current timestamp on the service.
  google.protobuf.Timestamp time = 1;
  // Does the service support NTP based synchronization.
  bool ntp_synced = 2;
  
  // Capabilities that the generator and reflect support on the
  // service. If the top level field is unset the service cannot act
  // as any defined generator or reflector.
  GeneratorCapabilities generator = 3;
  ReflectorCapabilities reflector = 4;
  
  // Maximum number of results allowed to be stored per interface.
  // The minimum supported just be 2 or greater.
  uint64 max_historical_results_per_interface = 5;
}

message PacketGeneratorCapabilities {
  uint64 max_bps = 1;  // bits per second
  uint64 max_pps = 2;  // packets per second
  uint32 min_mtu = 3;  // minimum MTU supported.
  uint32 max_mtu = 4;  // max MTU supported.
  google.protobuf.Duration min_setup_duration = 5;
  google.protobuf.Duration min_teardown_duration = 6;
  google.protobuf.Duration min_sample_interval = 7;
}

enum PacketInjectorLoopbackMode {
  PACKET_INJECTOR_LOOPBACK_MODE_UNSPECIFIED = 0;
  PACKET_INJECTOR_LOOPBACK_MODE_PMD = 1;
  PACKET_INJECTOR_LOOPBACK_MODE_ASIC = 2;
}

message PacketInjectorCapabilities {
  uint32 min_mtu = 1;
  uint32 max_mtu =2;
  uint32 min_injected_packets = 3;
  uint32 max_injected_packets = 4;
  google.protobuf.Duration min_setup_duration = 5;
  google.protobuf.Duration min_teardown_duration = 6;
  google.protobuf.Duration min_sample_interval = 7;
  repeated PacketInjectorLoopbackMode loopback_modes = 8;  
}

// If the service does not support any of the defined
// modes then the message should be unset.
message GeneratorCapabilities {
  PacketGeneratorCapabilities packet_generator = 1;
  PacketInjectorCapabilities packet_injector = 2;
}

// PMD or port based loopbacks are expect to have limited ability
// to report packet counters. Packet rate/errors/transmitted/received
// may not be available on the remote side for these types of loopbacks.
message PmdLoopbackCapabilities {
  google.protobuf.Duration min_setup_duration = 1;
  google.protobuf.Duration min_teardown_duration = 2;
}

enum HeaderMatchField {
  HEADER_MATCH_FIELD_UNSPECIFIED = 0;
  HEADER_MATCH_FIELD_L2 = 1;
  HEADER_MATCH_FIELD_L3 = 2;
  HEADER_MATCH_FIELD_L4 = 3;
}

message AsicLoopbackCapabilities {
  google.protobuf.Duration min_setup_duration = 1;
  google.protobuf.Duration min_teardown_duration = 2;
  repeated HeaderMatchField fields = 3;
}

// If the service does not support any of the defined
// modes then the message should be unset.
message ReflectorCapabilities {
  PmdLoopbackCapabilities pmd_loopback = 1;
  AsicLoopbackCapabilities asic_loopback = 2;
}


```

#### List Messages


```
message ListRequest {
}

message ListResponse {
  repeated ListResult results = 1;
}

message ListResult {
  string id = 1;
  QualificationState state = 2;
  string interface_name = 3;
}

```

### Use Cases

#### Duplicate test ID called

*   Create with config.id=”Test1” for interface Ethernet1
*   While the above is still ongoing, Create with config.id=”Test1” for interface Ethernet2
*   Second call to Create will return AlreadyExists error code

#### Existing Qualification running on interface

*   StartPacketQualification with config.id=”Test1” for interface Ethernet1
*   After the above is completed, StartPacketQualification with config.id=”Test1” for interface Ethernet 2
*   Second call to Create will return AlreadyExists error code

#### Duplicate id but called after previous id is deleted

*   Create with config.id=”Test1” for interface Ethernet1
*   Delete with id=”Test1”
*   Create with config.id=”Test1” for interface Ethernet 2
*   Second call to StartPacketQualification is expected to go through without errors.

#### How a based generator walks through setting itself up

*   Create call is made
*   Service validates the interface\_name is valid
    *   if not return `INVALID_PARAMETER`
*   Service validates the `interface_name` is found
    *   if not return `NOT_FOUND`
*   Service validates the test id doesn't conflict with existing tests
    *   if not return `ALREADY_EXISTS`
*   Service validates the endpoint type is supported
    *   if not return `INVALID_PARAMETER`
*   Service validates endpoint configuration
    *   if not return `INVALID_PARAMETER`
    *   status should return what parameters are invalid
*   If the synchronization mode is NTP
    *   Validate that the service can start at that time it must be at least endpoint type setup time + current time away
    *   if not return `INVALID_PARAMETER`
    *   Mark status as `QUALIFICATION_STATE_IDLE`
    *   Wait for `start_time` - `min_setup_time`
    *   Begin Setup
        *   Mark status as `QUALIFICATION_STATE_SETUP`
        *   if the setup cannot be completed before `start_time`
            *   Mark status as `QUALIFICATION_STATE_ERROR`
    *   Begin Qualification at `start_time`
    *   Start the qualification run for duration providing updates to the result every `update_interval`
        *   Mark status as `QUALIFICATION_STATE_RUNNING`
    *   Wait for `end_time`
        *   Mark status as `QUALIFICATION_STATE_TEARDOWN`
    *   If `teardown_time` is set
        *   Wait until `teardown_time`
    *   Begin teardown
        *   if the teardown cannot be completed
            *   Mark status as `QUALIFICATION_STATE_ERROR`
    *   Mark status as `QUALIFICATION_STATE_COMPLETED`
*   if the synchronization mode is RPC
    *   Validate the synchronization parameters for setup / teardown
    *   If any errors have occurred
        *   Mark status as `QUALIFICATION_STATE_ERROR`
    *   if any failures occur during the following steps also mark status as `QUALIFICATION_STATE_ERROR`
    *   Mark status as `QUALIFICATION_STATE_IDLE`
    *   Mark status as `QUALIFICATION_STATE_SETUP`
    *   Wait for `pre_sync_duration`
    *   Start the setup for the setup
        *   if the setup cannot be completed in `setup_duration`
            *   Mark status as `QUALIFICATION_STATE_ERROR`
    *   Once setup, wait for any remaining time based on the setup
    *   Start the qualification run for duration providing updates to the result every `update_interval`
        *   Mark status as `QUALIFICATION_STATE_RUNNING`
    *   Mark status as `QUALIFICATION_STATE_TEARDOWN`
    *   Wait for `post_sync_duration`
    *   Begin teardown
        *   Mark status as `QUALIFICATION_STATE_TEARDOWN`
        *   if the teardown cannot be completed in `teardown_duration`
            *   Mark status as `QUALIFICATION_STATE_ERROR`
    *   Wait for the remaining `teardown_duration`
    *   Mark status as `QUALIFICATION_STATE_COMPLETED`


#### Call graph for NTP synchronization

#### Call example for RPC synchronization with packet injector and PMD loopback


*   Caller gets Capabilities
    *   generator supports:

        ```
message PacketInjectorCapabilities {
  min_mtu: 64
  max_mtu: 1500
  min_injected_packets: 1
  max_injected_packets: 1000
  min_setup_duration: "30s"
  min_teardown_duration: "30s"
  min_sample_interval: "60s"
  loopback_modes: PACKET_INJECTOR_LOOPBACK_MODE_PMD
}
```


*   loopback supports:

        ```
message PmdLoopbackCapabilities {
  min_setup_duration: "10s"
  min_teardown_duration: "10s"
}
```


*   Caller calculates the `pre_sync_duration`, `setup_duration`, `teardown_duration` and `post_sync_duration` for each side
    *   **This assumes that both rpc's are sent in parallel**
    *   Generator
        *   `pre_sync_duration` = 10s
            *   remote setup time so we know the loopback is ready for us to generate packets
        *   `setup_duration` = 30s
            *   Setup time desired for this endpoint type.  cannot be less than `min_setup_duration` provided as a capability.
            *   Start sending packets
        *   `duration` = 180s
            *   qualification duration
            *   stats must be updated every `update_interval`
        *   `post_sync_duration` = 0s
            *   Generators should not need to set this value but may for synchronization issues with remote sides.
        *   `teardown_duration` = 30s
            *   The generator side teardown duration
            *   The `min_teardown_duration` will be used if left unset
    *   Reflector
        *   `pre_sync_duration` = 0s
            *   Loopbacks should not need to set this value but may for synchronization issues with remote sides.
        *   `setup_duration` = 40s
            *   This value should be the sum of the remote `setup_duration` and the local `min_setup_duration`. This value allows for the sides to be synchronized at the qualification start.
        *   `duration` = 180s
            *   qualification duration
            *   stats must be updated every `update_interval`
        *   `post_sync_duration` = 30s
            *   the reflector side should hold up the loopback until the remote side has finished teardown
        *   `teardown_duration` = 10s
            *   The reflector side teardown duration
            *   The `min_teardown_duration` will be used if left unset
*   Call generator and reflector Create()
*   Generator (PacketInjector)
    *   Validate that `setup_duration` is >= `min_setup_duration`
    *   Validate that `teardown_duration` is >= `min_teardown_duration`
    *   Schedule the work with backend which should being with waiting for `pre_sync_duration` if set
    *   Once the `pre_sync_duration` duration is reached
        *   Begin putting interface into generator mode
        *   If `setup_duration` has not been reached wait for the remaining duration
    *   Take first snapshot of packets and rates for the initial result
    *   Repeat taking snapshots of results every `update_interval`
    *   Once duration is reached wait for `post_sync_duration` if set (for generators this value is not expected to be set)
    *   After `post_sync_duration` is reached begin teardown
*   Reflector (PmdLoopback)
    *   Validate that setup\_duration is >= min\_setup\_duration
    *   Validate that teardown\_duration is >= min\_teardown\_duration
    *   Schedule the work with backend which should being with waiting for pre\_sync\_duration if set (for reflector this value is not expected to be set)
    *   Once the pre\_sync\_duration duration is reached
        *   Begin putting interface into reflector mode
        *   If `setup_duration` has not been reached wait for the remaining duration
    *   Take first snapshot of packets and rates for the initial result
    *   Repeat taking snapshots of results every update\_interval
    *   Once duration is reached wait for `post_sync_duration` if set
    *   After `post_sync_duration` is reached begin teardown


#### Call graph for Mixed synchronization generator RPC / loopback NTP



*   Caller gets Capabilities
    *   generator supports:

        ```
message PacketGeneratorCapabilities {
  max_bps: 250000000000
  max_pps: 200000000
  min_mtu: 64
  max_mtu: 9216
  min_setup_duration: "2s"
  min_teardown_duration: "2s"
  min_sample_interval: "1s"
}
```


*   loopback supports:

        ```
message PmdLoopbackCapabilities {
  min_setup_duration: "10s"
  min_teardown_duration: "10s"
}
```


*   Caller calculates the `pre_sync_duration`, `setup_duration`, `teardown_duration` and `post_sync_duration` for each side
    *   **This assumes that both rpc's are sent in parallel**
    *   Generator

        ```
message NTPSyncedTiming {
  start_time: time.Now()+"60s"
  end_time: time.Now()+"180s"
}
```


        *   `start_time`
            *   adding 60 secs just to show how to balance out the loopback side via RPC
        *   `end_time` = `start_time` + "180s"
            *   qualification duration
            *   stats must be updated every `update_interval`
        *   `teardown_time`
            *   left unset as we want to the generator to teardown immediately after `end_time`
    *   Reflector
        *   `pre_sync_duration` = 0s
            *   Reflectors should not need to set this value unless the RPC's are dispatched synchronously.
        *   `setup_duration` = 60s
            *   This value should be the sum of the remote `setup_duration` and the local `setup_duration`. This value allows for the sides to be synchronized at the qualification start.
        *   `duration` = 180s
            *   qualification duration
            *   stats must be updated every `update_interval`
        *   `post_sync_duration` = 2s
            *   the reflector side should hold up the loopback until the remote side has finished teardown
        *   `teardown_duration` = 10s
            *   The reflector side teardown duration
            *   The `min_teardown_duration` will be used if left unset
*   Call generator and reflector Create()
*   Generator
    *   Validate that `start_time >= time.Now() + min_setup_duration`
    *   if `teardown_time` is set
        *   `teardown_time` must be >= `end_time`
        *   If not, `teardown_time` will be `end_time`
    *   Begin setup at `start_time` - `min_setup_duration`
    *   Take first snapshot of packets and rates for the initial result
    *   Repeat taking snapshots of results every `update_interval` until `end_time`
    *   if `teardown_time` is not set, begin teardown.
        *   if set wait until `teardown_time` then begin teardown
*   Reflector
    *   Validate that `setup_duration` is >= `min_setup_duration`
    *   Validate that `teardown_duration` is >= `min_teardown_duration`
    *   Schedule the work with backend which should being with waiting for `pre_sync_duration` if set (for reflector this value is not expected to be set)
    *   Once the `pre_sync_duration` duration is reached
        *   Begin putting interface into reflector mode
        *   If `setup_duration` has not been reached wait for the remaining duration
    *   Take first snapshot of packets and rates for the initial result
    *   Repeat taking snapshots of results every `update_interva`l
    *   Once duration is reached wait for `post_sync_duration` if set
    *   After `post_sync_duration` is reached begin teardown


#### Workflow for Delete in the case of canceling early



*   A caller may want cancel a running qualification early due to several administrative reasons
    *   Noticing the errors have already exceeded a threshold 
    *   Wanting to change the rate or mtu for a running test
    *   Canceling it as might have been scheduled at the wrong time
*   The expectation would be that calling delete before completion would cause the service to immediately cancel the packet ongoing test by
    *   Setting the status into `QUALIFICATION_STATE_ERROR`
    *   The status would be `CANCELLED` in the embedded status message
    *   The service would immediately start a teardown on the qualification
    *   Once the qualification is torn down the results would be deleted.


### Future work


#### Inband signaled link qualification

Using this API as a general purpose framework, it could be extended to support
an inband signaled mode of operation.  In the currently defined implementations
the calling service must orchestrate both ends of the link via create / get /
delete calls.  An inband version could allow the caller to only need to
orchestrate the one side of the link. Also there could even be internal use
cases in which the devices themselves could initiate a link qualification
without external prompting.
