package config

import (
	"github.com/spf13/viper"
)

type Builder struct {
	Viper *viper.Viper
}

type Config struct {
	Bucket  *Bucket  `json:"bucket" yaml:"bucket"`
	Decider *Decider `json:"decider" yaml:"decider"`
	DryRun  bool     `json:"dry_run" yaml:"dryRun"`
}

const (
	EnvPrefix = "JANITOR"
	KeyDryRun = "DryRun"
)

func (b *Builder) BuildWithConfFile(confFile string) (*Config, error) {
	b.Viper.SetConfigFile(confFile)
	if err := b.Viper.MergeInConfig(); err != nil {
		return nil, err
	}
	return b.Build()
}

func (b *Builder) Build() (*Config, error) {
	c := New()
	b.Viper.AutomaticEnv()
	return c, b.Viper.Unmarshal(&c)
}

func New() *Config {
	return &Config{
		Bucket:  NewBucket(),
		Decider: NewDecider(),
		DryRun:  false,
	}
}

func NewBuilder() (builder *Builder) {
	builder = &Builder{
		Viper: viper.New(),
	}
	builder.Viper.SetEnvPrefix(EnvPrefix)
	return
}
