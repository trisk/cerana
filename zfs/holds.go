package zfs

import (
	"bytes"

	"github.com/cerana/cerana/pkg/errors"
	"github.com/cerana/cerana/zfs/nv"
)

const emptyList = "\x00\x01\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

func holds(name string) ([]string, error) {
	m := map[string]interface{}{
		"cmd":     "zfs_get_holds",
		"version": uint64(0),
	}

	encoded := &bytes.Buffer{}
	err := nv.NewNativeEncoder(encoded).Encode(m)
	if err != nil {
		return nil, errors.Wrapv(err, map[string]interface{}{"name": name, "args": m})
	}

	out := make([]byte, 1024)
	copy(out, emptyList)

	err = ioctl(zfs(), name, encoded.Bytes(), out)
	if err != nil {
		return nil, err
	}

	m = map[string]interface{}{}

	err = nv.NewNativeDecoder(bytes.NewReader(out)).Decode(&m)
	if err != nil {
		return nil, errors.Wrapv(err, map[string]interface{}{"name": name, "args": m})
	}

	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}

	return names, nil
}
