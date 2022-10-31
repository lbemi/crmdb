package lbemi

import (
	"github.com/lbemi/lbemi/app/option"
	"github.com/lbemi/lbemi/pkg/controller"
)

var CoreV1 controller.IController

func Setup(options *option.Options) {
	CoreV1 = controller.NewController(options.Factory, options.Redis)
}
