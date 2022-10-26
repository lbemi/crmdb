package core

import (
	"github.com/lbemi/lbemi/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/controller"
)

var Core controller.IController

func Setup(options option.Options) {
	Core = controller.NewController(options.Factory)
}