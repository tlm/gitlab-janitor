package config

import (
	"github.com/tlmiller/gitlab-janitor/pkg/gitlab/backup"
)

type DeciderMapper func(conf interface{}) (backup.Decider, error)

type DeciderMapperConfig func() interface{}

type DeciderMapping struct {
	Config DeciderMapperConfig
	Mapper DeciderMapper
}

func DeciderMapperFactory() map[string]DeciderMapping {
	return map[string]DeciderMapping{
		"keepAggregateAgree": DeciderMapping{
			Config: func() interface{} { return &DeciderAggregateAgreeConfig{} },
			Mapper: deciderAggregateAgreeMapper,
		},
		"keepFirstMatch": DeciderMapping{
			Config: func() interface{} { return &DeciderFirstKeepMatchConfig{} },
			Mapper: deciderFirstKeepMatchMapper,
		},
		"keepAfterDuration": DeciderMapping{
			Config: func() interface{} { return &DeciderKeepAfterDurationConfig{} },
			Mapper: deciderKeepAfterDurationMapper,
		},
		"keepAfterTime": DeciderMapping{
			Config: func() interface{} { return &DeciderKeepAfterTimeConfig{} },
			Mapper: deciderKeepAfterTimeMapper,
		},
		"keepPerVersion": DeciderMapping{
			Config: func() interface{} { return &DeciderKeepPerVersionConfig{} },
			Mapper: deciderKeepPerVersionMapper,
		},
		"keepNumberVersions": DeciderMapping{
			Config: func() interface{} { return &DeciderKeepNumberOfVersionsConfig{} },
			Mapper: deciderKeepNumberOfVersionsMapper,
		},
	}
}
