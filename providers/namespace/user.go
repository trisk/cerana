package namespace

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/cerana/cerana/acomm"
	"github.com/cerana/cerana/pkg/errors"
	"github.com/cerana/cerana/pkg/logrusx"
)

// UserArgs are arguments for SetUser.
type UserArgs struct {
	PID  uint64  `json:"pid"`
	UIDs []IDMap `json:"uids"`
	GIDs []IDMap `json:"gids"`
}

// IDMap is a map of id in container to id on host and length of a range.
type IDMap struct {
	ID     uint64 `json:"id"`
	HostID uint64 `json:"hostID"`
	Length uint64 `json:"length"`
}

// SetUser sets the user and group id mapping for a process.
func (n *Namespace) SetUser(req *acomm.Request) (interface{}, *url.URL, error) {
	var args UserArgs
	if err := req.UnmarshalArgs(&args); err != nil {
		return nil, nil, err
	}

	uidMapPath := fmt.Sprintf("/proc/%d/uid_map", args.PID)
	if err := writeIDMapFile(uidMapPath, args.UIDs); err != nil {
		return nil, nil, err
	}

	gidMapPath := fmt.Sprintf("/proc/%d/gid_map", args.PID)
	if err := writeIDMapFile(gidMapPath, args.GIDs); err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func writeIDMapFile(path string, idMaps []IDMap) error {
	mode := os.O_TRUNC | os.O_RDWR
	perms := os.FileMode(0644)
	mapFile, err := os.OpenFile(path, mode, perms)
	if err != nil {
		return errors.Wrapv(err, map[string]interface{}{"path": path, "mode": mode, "perms": perms})
	}
	defer logrusx.LogReturnedErr(mapFile.Close, map[string]interface{}{"path": path}, "failed to close map file")

	content := make([]string, len(idMaps))
	for i, idMap := range idMaps {
		content[i] = fmt.Sprintf("%d %d %d", idMap.ID, idMap.HostID, idMap.Length)
	}

	data := strings.Join(content, "\n")
	_, err = fmt.Fprintln(mapFile, data)
	return errors.Wrapv(err, map[string]interface{}{"path": path, "data": data})
}
