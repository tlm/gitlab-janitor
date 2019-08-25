package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/tlmiller/janitor/pkg/gitlab/backup"
)

type DeciderAggregateAgreeConfig struct {
	Deciders []Decider `mapstructure:"deciders"`
}

type DeciderFirstKeepMatchConfig struct {
	Match    bool      `mapstructure:"match"`
	Deciders []Decider `mapstructure:"deciders"`
}

type DeciderKeepAfterDurationConfig struct {
	Duration string `mapstructure:"duration"`
}

type DeciderKeepAfterTimeConfig struct {
	Time string `mapstructure:"time"`
}

type DeciderKeepPerVersionConfig struct {
	Count int `mapstructure:"count"`
}

type DeciderKeepNumberOfVersionsConfig struct {
	Keep int `mapstructure:"keep"`
}

func deciderAggregateAgreeMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderAggregateAgreeConfig)
	if !ok {
		return nil, errors.New("decider aggregate agree config is not of type DeciderAggregateAgreeConfig")
	}

	deciders, err := deciderAggregateMapper(conf.Deciders)
	if err != nil {
		return nil, err
	}
	return backup.WithAggregateAgree(deciders...), nil
}

func deciderAggregateMapper(conf []Decider) ([]backup.Decider, error) {
	deciders := make([]backup.Decider, len(conf))
	for i, deciderConf := range conf {
		decider, err := ToDecider(&deciderConf)
		if err != nil {
			return nil, fmt.Errorf("building decider aggregate aggree for decider %d: %v", i, err)
		}
		deciders[i] = decider
	}
	return deciders, nil
}

func deciderFirstKeepMatchMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderFirstKeepMatchConfig)
	if !ok {
		return nil, errors.New("decider first keep match config is not of type DeciderFirstKeepMatchConfig")
	}

	deciders, err := deciderAggregateMapper(conf.Deciders)
	if err != nil {
		return nil, err
	}
	return backup.WithFirstKeepMatch(conf.Match, deciders...), nil
}

func deciderKeepAfterDurationMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderKeepAfterDurationConfig)
	if !ok {
		return nil, errors.New("decider keep after duration config is not of type DeciderKeepAfterDurationConfig")
	}

	duration, err := time.ParseDuration(conf.Duration)
	if err != nil {
		return nil, fmt.Errorf("parsing decider keep after duration: %v", err)
	}
	if duration < time.Duration(0) {
		return nil, errors.New("decider keep after duration value cannot be less than zero")
	}
	return backup.WithKeepAfterDuration(duration), nil
}

func deciderKeepAfterTimeMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderKeepAfterTimeConfig)
	if !ok {
		return nil, errors.New("decider keep after time config is not of type DeciderKeepAfterTimeConfig")
	}

	time, err := time.Parse(time.RFC1123Z, conf.Time)
	if err != nil {
		return nil, fmt.Errorf("parsing decider keep after time: %v", err)
	}
	return backup.WithKeepAfterTime(time), nil
}

func deciderKeepPerVersionMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderKeepPerVersionConfig)
	if !ok {
		return nil, errors.New("decider keep per version config is not of type DeciderKeepPerVersionConfig")
	}

	if conf.Count < 0 {
		return nil, errors.New("decider keep per versions cannot be less than zero ")
	}
	return backup.WithKeepPerVersion(conf.Count), nil
}

func deciderKeepNumberOfVersionsMapper(raw interface{}) (backup.Decider, error) {
	conf, ok := raw.(*DeciderKeepNumberOfVersionsConfig)
	if !ok {
		return nil, errors.New("decider keep number of versions config is not of type DeciderKeepNumberOfVersionsConfig")
	}

	if conf.Keep < 1 {
		return nil, errors.New("decider keep number of versions cannot be less than one")
	}
	return backup.WithKeepNumberOfVersions(conf.Keep), nil
}
