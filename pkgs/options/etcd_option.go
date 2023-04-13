package options

import (
	"github.com/HyetPang/go-frame/internal/components/etcd"
	"go.uber.org/fx"
)

func WithEtcd() Option {
	return func(o *Options) {
		o.FxOptions = append(o.FxOptions, fx.Provide(etcd.New))
	}
}
