# Modernising Authentication and Authorization for Network Devices
**Contributors:** danielwong<sup><small>†</small></sup>, johnwestbrook<sup><small>†</small></sup>, robjs<sup><small>†</small></sup>, morrowc<sup><small>†</small></sup>, aashaikh<sup><small>†</small></sup>  
<sup><small>†</small></sup> @google.com  
**Published:** November 2020

## Problem Statement
Historically, network devices have had a single store of configuration data,
users with sufficient privilege levels are allowed to edit this data.
Implementation-specific mechanisms have been used to restrict the subsets of
configuration and telemetry that a particular user can interact with, such that
it is possible to enforce particular roles interacting with a network device in
specific manners. Often, the authorization policies that network operators have
implemented have been coarse - for example, high privilege users being able to
edit any configuration, or a single SNMP community having full access to all
telemetry on the device. Typically, only lawful intercept credentials and
configuration has been subject to a 'protected' configuration store, which was
treated differently by the device.

The OpenConfig project initially provided means to inter-work with existing
means of authorization and authentication on network devices - providing means
to supply a username and password via gRPC metadata for gNMI and gNOI - along
with mutual TLS authentication. However, this approach has a number of
limitations:

* It relies on an external run-time dependency to ensure that
  authentication/authorization is performed by the switch - particularly, the
  external AAA server being accessible. Thus introducing new failure modes that
  need to handled - such as fall-back credentials.
* It introduces a potential vulnerability if the remote AAA is subject to MITM
  attacks.
* In high-volume deployments (hundreds of thousands of switches), it introduces
  a new service that must be scaled.
* It can introduce external factors which impact the performance of
  programmatic RPCs to the device - e.g., the latency between the switch and
  the AAA server.
* It requires extension and definition of new patterns to be able to handle
  more granular authorization of access to the configuration and telemetry
  stored on a network device.
* It requires a networking-specific approach to be taken for secure access to
  device endpoints rather than building on the wider industry open source
  developments, and their surrounding ecosystem.

## Proposed Direction

We propose to develop an approach which addresses a number of these concerns,
particularly to provide:

* A secure means by which authentication and authorization policy can be
  provided to a network device and stored locally.
 
* Means by which this policy can interact with user identifiers that are
  specified as part of the gRPC connection - ideally building upon wider
  industry developments.

* Means by which this policy can be used as a mechanism for authenticating
  users against other access mechanisms (e.g., local console, SSH) such that an
  operator can unify their authentication and authorization policy for network
  devices into a single definition.

* Means within the policy to provide for granular access to a data tree which
  is described by a specified data model. This will particularly focus around
  the YANG tree that is defined by the OpenConfig data models.

We envisage that this approach can coexist with authentication and
authorization mechanisms that are already implemented by typical NOS today,
whilst also evolving towards a unified, modernised approach.

### Next Steps

We envisage that the work to define the security mechanism defined above
involves a number of disparate steps:

1. Defining an approach for authentication, and the interface via which it is
   to be instantiated on a local network element.
2. Defining an approach for authorization within specific services - both based
   on RPC called, as well as payload (particularly, path-based within gNMI).
3. Defining means by which the services defined can coexist with existing
   mechanisms implemented in common NOSes, whilst ensuring the device remains
   secure.

We propose to use the [gNOI](https://github.com/openconfig/gnoi) repository as
an initial starting point for this development work, and invite contributions
from operator and vendor engineers within this branch and the corresponding
pull requests.
