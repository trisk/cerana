package zfs

import (
	"syscall"

	"github.com/cerana/cerana/pkg/errors"
)

// dmuType corresponds to the  dmu_objset_type enum in libzfs sys/fs/zfs.h
type dmuType int32

const (
	dmuNone dmuType = iota
	dmuMeta
	dmuZFS
	dmuZVOL
	dmuOther /* For testing only! */
	dmuAny   /* Be careful! */
	dmuNumtypes
)

var dmuTypes = map[string]dmuType{
	"none":  dmuNone,
	"meta":  dmuMeta,
	"zfs":   dmuZFS,
	"zvol":  dmuZVOL,
	"other": dmuOther,
	"any":   dmuAny,
}

func getDMUType(name string) (dmuType, error) {
	d, ok := dmuTypes[name]
	if !ok {
		return dmuNone, errors.Wrapv(syscall.EINVAL, map[string]interface{}{"type": name})
	}
	return d, nil
}
