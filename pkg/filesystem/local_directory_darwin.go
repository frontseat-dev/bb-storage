//go:build darwin
// +build darwin

package filesystem

import (
	"os"

	"github.com/buildbarn/bb-storage/pkg/filesystem/path"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// rawDeviceNumber is the equivalent of POSIX dev_t.
type rawDeviceNumber = int32

// O_SEARCH is not exported by x/sys for darwin. macOS defines it in
// <sys/fcntl.h> as a COMPOSITE macro rather than a literal:
//
//	#define O_EXEC      0x40000000
//	#define O_DIRECTORY 0x00100000
//	#define O_SEARCH    (O_EXEC | O_DIRECTORY)
//
// and x/sys's constant generator extracts literal values, not macro
// compositions — which is why it has O_SEARCH for freebsd/solaris/aix but
// not darwin. Buildbarn builds with Bazel, so `go build` on darwin never
// exercised this. Spelled out here so the package compiles from the public
// module path; drop it if x/sys ever exports the constant.
const oflagExec = 0x40000000

const oflagSearch = oflagExec | unix.O_DIRECTORY

func (localDirectory) Mknod(name path.Component, perm os.FileMode, deviceNumber DeviceNumber) error {
	return status.Error(codes.Unimplemented, "Creation of device nodes is not supported on Darwin")
}

func clonefileImpl(oldFD int, oldName string, newFD int, newName string) error {
	return unix.Clonefileat(oldFD, oldName, newFD, newName, unix.CLONE_NOFOLLOW)
}

func (localDirectory) Mount(mountpoint path.Component, source, fstype string) error {
	return status.Error(codes.Unimplemented, "Mount is not supported")
}
