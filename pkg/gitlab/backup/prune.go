package backup

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"

	"gocloud.dev/blob"
)

func CreatePruneList(bucket *blob.Bucket, d Decider) ([]string, error) {
	it := bucket.List(&blob.ListOptions{})
	ctx := context.Background()
	var (
		backups BackupList = BackupList{}
		err     error
		obj     *blob.ListObject
	)
	for obj, err = it.Next(ctx); obj != nil && err == nil; obj, err = it.Next(ctx) {
		parts := strings.Split(obj.Key, "_")
		if len(parts) != 7 {
			continue
		}
		unixTime, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert backup time to unix int: %v", err)
		}
		ver, err := version.NewVersion(parts[4])
		if err != nil {
			return nil, fmt.Errorf("failed to parse backup version %s: %v", parts[4], err)
		}
		backups = append(backups, &Backup{
			Key:     obj.Key,
			Time:    time.Unix(unixTime, 0),
			Version: ver,
		})
	}
	rval := []string{}
	sort.Sort(backups)
	for _, b := range backups {
		if !d.Keep(b) {
			rval = append(rval, b.Key)
		}
	}

	return rval, nil
}
