package core

import (
	"github.com/lbemi/lbemi/app/option"
	"github.com/lbemi/lbemi/pkg/handler"
)

var V1 handler.IController

func Setup(options *option.Options) {
	V1 = handler.NewController(options.Factory, options.Redis)
}
