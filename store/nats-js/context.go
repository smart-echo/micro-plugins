package natsjs

import (
	"context"

	"github.com/smart-echo/micro/store"
)

// setStoreOption returns a function to setup a context with given value.
func setStoreOption(k, v interface{}) store.Option {
	return func(o *store.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
