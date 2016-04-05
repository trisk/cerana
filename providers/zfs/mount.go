package zfs

import (
	"errors"
	"net/url"
	"strings"
	"syscall"

	"github.com/cerana/cerana/acomm"
	gzfs "github.com/mistifyio/go-zfs"
)

// TODO: Update this method once `zfs` supports Mount

// MountArgs are arguments for the Mount handler.
type MountArgs struct {
	Name    string `json:"name"`
	Overlay bool   `json:"overlay"`
}

// Mount mounts a zfs filesystem.
func (z *ZFS) Mount(req *acomm.Request) (interface{}, *url.URL, error) {
	var args MountArgs
	if err := req.UnmarshalArgs(&args); err != nil {
		return nil, nil, err
	}

	if args.Name == "" {
		return nil, nil, errors.New("missing arg: name")
	}

	ds, err := gzfs.GetDataset(args.Name)
	if err != nil {
		// Fix errors to be more like what zfs will probably return
		if strings.Contains(err.Error(), "dataset does not exist") {
			err = syscall.ENOENT
		}
		return nil, nil, err
	}

	ds, err = ds.Mount(args.Overlay, nil)
	if err != nil {
		// Fix errors to be more like what zfs will probably return
		if strings.Contains(err.Error(), "already mounted") {
			err = syscall.EBUSY
		}

		return nil, nil, err
	}

	return nil, nil, nil
}
