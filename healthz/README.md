# Healthz Overview

The purpose of Healthz is to allow a component inside of a system to report its
health. The concept of asking "Are you healthy?" is a general design principle
in distributed systems.

The ability to also include "implementation specific" details about the health
of the system or component is very standard for network devices in the form for
"show tech", "diag" or "debug" commands. These are very useful for getting
useful diagnostics about the system state to the implementers. Healthz exposes
these interfaces as queriable endpoints to allow operators to validate "health"
on components and if "unhealthy" gather implementor specific data to help triage
or reproduce issues.

## Design

*   The service definition provides a Get method which is correlated with a
    /component in the OC tree. The Get method is responsible for getting the
    current status of the requested component path. This will cause the system
    to fetch the current state of component this may entail running specific
    validations and gathering any relevant status of the component to reflect
    the "health check"

*   The internal design details of the "health check" and the frequency of the
    health check are up to the implementor to determine. This presentation of
    the status in OC is proposed to be modeled as

```yang
component/healthz/status = Enum
component/healthz/last-unhealthy = Timestamp
component/healthz/unhealthy_count = Int64
```

## Use Cases

### Evaluate the long term health of the component

The health of a component is critical to overall system health. The ability to
monitor the status of the component in the system has traditionally been tracked
via 'show commands' and some form of alarms via syslog or snmptraps.

Healthz provides the implementtor the ability to react to issues and "store" the
metadata then provide a signal via streaming telemetry that this component has
experienced an unhealthy segment. The operator can then perform a 'Get' on the
component to fetch associated metadata for the event. If the event is transient
in nature, the counter will continue to increment and allow for operator to
observe this component is consistently unhealthy.
