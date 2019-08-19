package backup

import (
	"time"
)

type Backup struct {
	Key     string
	Time    time.Time
	Version *version.Version
}

type BackupList []*Backup

func (l BackupList) Len() int {
	return len(l)
}

func (l BackupList) Less(i, j int) bool {
	if l[i].Time.After(l[j].Time) {
		return true
	} else if l[i].Time.Equal(l[j].Time) {
		return l[i].Version.GreaterThan(l[j].Version)
	}
	return false
}

func (l BackupList) Swap(i, j int) {
	tmp := l[i]
	l[i] = l[j]
	l[j] = tmp
}
