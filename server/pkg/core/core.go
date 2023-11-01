package core

import (
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
)

var V1 Getter

func Register(options *option.Options) {
	V1 = NewHandler(
		options.Redis,
		options.DB,
		options.Enforcer,
		options.ClientStore)
}
