package pgo2

import (
	"github.com/pinguo/pgo2/config"
	"os"
)

var defaultOpt *Options

func init() {
	defaultOpt = NewOpt()
}

//默认参数
func Opt(opts ...Option) *Options {
	for _, opt := range opts {
		opt(defaultOpt)
	}
	return defaultOpt
}

//新建参数配置
func NewOpt(opts ...Option) *Options {
	if defaultOpt == nil {
		defaultOpt = &Options{
			Args:       os.Args,
			ConfigData: make(map[string]interface{}),
		}
	}
	for _, opt := range opts {
		opt(defaultOpt)
	}
	return defaultOpt
}

type Option func(opts *Options)

type Options struct {
	Args           []string                      //运行参数
	NewApp         bool                          //是否创建新的App
	PostConfigInit []func(config config.IConfig) //配置信息初始化完成后回调函数
	ConfigData     map[string]interface{}        //配置信息
	AppInit        []func(app *Application)      //app初始化回调函数
}

func (s *Options) SetupConfig(cfg config.IConfig) error {
	for k, v := range s.ConfigData {
		cfg.Set(k, v)
	}
	return nil
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

//本地配置加载成功后回调函数
func PostConfigInitOption(v func(config config.IConfig)) Option {
	return func(opts *Options) {
		opts.PostConfigInit = append(opts.PostConfigInit, v)
	}
}

func AppInitOption(v func(app *Application)) Option {
	return func(opts *Options) {
		opts.AppInit = append(opts.AppInit, v)
	}
}

//启动时动态调整配置参数
func ConfigValueOption(key string, value interface{}) Option {
	return func(opts *Options) {
		opts.ConfigData[key] = value
	}
}
