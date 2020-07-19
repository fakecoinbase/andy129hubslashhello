package registry

import "time"

// 定义初始化需要的参数
type Options struct {
	// 地址
	Addrs []string
	// 超时时间
	Timeout time.Duration
	// 定义心跳时间
	HeartBeat int64
	//   /a/b/c/xxx/10.xxxx
	RegistryPath string
}

type Option func(opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout2(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartBeat(heartBeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}
func WithRegistryPath(path string) Option {
	return func(opts *Options) {
		opts.RegistryPath = path
	}
}

