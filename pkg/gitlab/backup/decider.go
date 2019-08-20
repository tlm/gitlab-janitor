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

func WithAggregateAgree(deciders ...Decider) Decider {
	return DeciderFn(func(b *Backup) bool {
		lastDecision := true
		for i, d := range deciders {
			decision := d.Keep(b)
			if i == 0 {
				lastDecision = decision
			}
			if decision != lastDecision {
				return true
			}
		}
		return lastDecision
	})
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
