package log

import "github.com/wlbwlbwlb/log/feishu"

type Options struct {
	Name string
}

type OptionFunc func(*Options)

func Feishu(token string) OptionFunc {
	return func(opts *Options) {
		feishu.Writer.With(token)
	}
}
