package core

import (
	"github.com/lbemi/lbemi/pkg/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/handler"
)

var V1 handler.Getter

func Register(options *option.Options) {
	V1 = handler.NewHandler(options.Factory, options.Redis, options.DB, options.Enforcer)
}
