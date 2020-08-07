package options

import (
	"github.com/pinguo/pgo2/config"
	"os"
)

func Opt() *Options {
	opts := &Options{
		Args: os.Args,
	}
	return opts
}

type Option func(opts *Options)

type Options struct {
	Args            []string
	NewApp          bool
	AfterConfigInit func(config config.IConfig)
}

func ArgsOption(v []string) Option {
	return func(opts *Options) {
		opts.Args = v
	}
}

func IsNewAppOption(v bool) Option {
	return func(opts *Options) {
		opts.NewApp = v
	}
}

func AfterConfigInitOption(v func(config config.IConfig)) Option {
	return func(opts *Options) {
		opts.AfterConfigInit = v
	}
}
