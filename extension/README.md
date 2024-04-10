### Extension management

#### What is an extension?
An extension is a single file (e.g. RPM, squashfs, zip) which can be installed
on the system to extends the functionality of the system in some way.   
An extension may be installed to:
- fix some bug/CVE on the system, by updating the programs affected by that
  bug/CVE
- add a new feature to the system
- add/update some software utility (e.g. perl, ncurses)
- update some dependency (e.g. updating some python site-package)

### Extension installation workflow

#### 1. Extension Transfer
Optionally transfer the extension file to the system, or select an
extension which had been previously transferred but which was not installed.

#### 2. Extension Installation
Install the extension.
Optionally:
- set the signature verification option.
  This signature verification guarantees that the extension is unaltered
  after being signed, and that it originates from some trusted source.   
  Either:
  - ignore signature verification for this extension
  - accept extension signatures which are cryptographically signed by the switch vendor
  - accept extension signatures which are signed by some trusted certificates (specified on-device).
  
- mark the extension as an extension which should be applied on boot.   
  When the system reboots, it should apply the extension from boot.
- ignore any dependency issues.    
  (e.g. a later extension may be installed which would provide the correct
  dependencies)

Programs which are running will not be affected by the newly installed
extension, and will need to be reloaded before any changes take effect.    
E.g. if an extension contains a bug fix for OpenConfig, then this bug fix
will not be applied until OpenConfig is reloaded

(optionally) multiple extensions can be installed in sequence, to minimise
the amount of times we have to reload affected programs

#### 3. Service reload

Once any required extensions are installed, the system needs to reload any
currently-running programs which would be affected by those extensions.
When those programs reload, they will reload with the extensions applied.

### Comparison to existing SetPackage RPC

Right now SetPackage is
```protobuf
 // SetPackage places a software package (possibly including bootable images)
  // on the target. The file is sent in sequential messages, each message
  // up to 64KB of data. A final message must be sent that includes the hash
  // of the data sent. An error is returned if the location does not exist or
  // there is an error writing the data. If no checksum is received, the target
  // must assume the operation is incomplete and remove the partially
  // transmitted file. The target should initially write the file to a temporary
  // location so a failure does not destroy the original file.
  rpc SetPackage(stream SetPackageRequest) returns (SetPackageResponse) {}
```
with
```protobuf

// Package defines a single package file to be placed on the target.
message Package  {
  // Destination path and filename of the package.
  string filename = 1;
  // Version of the package. (vendor internal name)
  string version = 4;
  // Indicates that the package should be made active after receipt on
  // the device. For system image packages, the new image is expected to
  // be active after a reboot.
  bool activate = 5;
  // Details for the device to download the package from a remote location.
  common.RemoteDownload remote_download = 6;
}

// SetPackageRequest will place the package onto the target and optionally mark
// it as the next bootable image. The initial message must be a package
// message containing the filename and information about the file. Following the
// initial message the contents are then streamed in maximum 64k chunks. The
// final message must be a hash message contains the hash of the file contents.
message SetPackageRequest {
  oneof request {
    Package package = 1;
    bytes contents = 2;
    types.HashType hash = 3;    // Verification hash of data.
  }
}

message SetPackageResponse {
}
```

#### Comparison with the current version of SetPackage:
Under `SetPackage`, a "package" is very loosely defined.    
From the comments like `SetPackageRequest will place the package onto the
target and optionally mark it as the next bootable image` or
`SetPackage places a software package (possibly including bootable
images)` or that the `Package.activate` field being set to true means that
`For system image packages, the new image is expected to be active
after a reboot.`   
this makes it seem that this RPC was intended to be used primarily for
transferring new bootable images to the system.

In 2019 the gNOI OS service was introduced, and when `SetPackage` was
pointed to as an alternative to adding a new service, the idea was
rejected based on SetPackage being too generic to capture the workflow:    
https://github.com/openconfig/gnoi/pull/19#issuecomment-479083625

The generic nature of SetPackage therefore means it could be seen as
potentially being used to handle extensions.
However, the same case that was made for gNOI OS can be made here for Extensions -
extension management is complex enough to deserve its own service and RPCs.

A workflow for installing an extension is more complex than
adding a file onto a system and either "activating" it or not.
- The system should be configured to do signature verification.      
  SetPackage has no concept of this.
- The installation may be performed with an extension which is currently
  present on the system.        
  SetPackage only offers the option to transfer and
  then either immediately activate, or leave the package on the system.
  If the package is left on the system then SetPackage does not offer
  any capability to ever interact with that package again.
- The network operator may need/want to install multiple extensions
  before reloading all affected programs at once.   
  SetPackage doesn't have any mechanism for handling the reload/activation
  step separately.
- Activation may affect the gNOI service, if the program running
  the gNOI service needs to reload.      
  SetPackage doesn't have any expectation that the RPC could be interrupted
  due to activation.
- SetPackage allows for packages to be transferred to arbitrary locations
  on the filesystem, while this should be abstracted away from the client
  when installing extensions to allow for the system to group all extension
  files under a single directory (the system shouldn't search its entire
  filesystem for extension files)

