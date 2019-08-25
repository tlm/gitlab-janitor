package config

import (
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/tlmiller/gitlab-janitor/pkg/gitlab/backup"
)

type Decider struct {
	Type    string                 `json:"type" yaml:"type"`
	Options map[string]interface{} `json:"options" yaml:"options"`
}

func NewDecider() *Decider {
	return &Decider{}
}

func ToDecider(conf *Decider) (backup.Decider, error) {
	factory := DeciderMapperFactory()
	if conf.Type == "" {
		return nil, errors.New("decider type cannot be null")
	}

	mapping, found := factory[conf.Type]
	if !found {
		return nil, fmt.Errorf("no decider mapping for type: %s", conf.Type)
	}

	rawConf := mapping.Config()
	if err := mapstructure.Decode(conf.Options, rawConf); err != nil {
		return nil, fmt.Errorf("decoding  decider configuration: %v", err)
	}

	return mapping.Mapper(rawConf)
}
