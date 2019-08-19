package backup

import (
	"time"
)

type Decider interface {
	Keep(*Backup) bool
}

type DeciderFn func(*Backup) bool

func (d DeciderFn) Keep(b *Backup) bool {
	return d(b)
}

func WithKeepAfterDuration(duration time.Duration) Decider {
	return WithKeepAfterTime(time.Now().Add(-duration))
}

func WithKeepAfterTime(after time.Time) Decider {
	return DeciderFn(func(b *Backup) bool {
		return b.Time.After(after)
	})
}

func WithKeepNumberOfVersions(count int) Decider {
	counter := map[string]struct{}{}
	return DeciderFn(func(b *Backup) bool {
		if len(counter) == count {
			_, exists := counter[b.Version.String()]
			return exists
		}
		counter[b.Version.String()] = struct{}{}
		return true
	})
}
