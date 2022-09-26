package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
)

func Register(c *gin.Context) {
	var registerForm form.RegisterUserForm
	if err := c.ShouldBind(&registerForm); err != nil {
		global.App.Log.Error(err.Error())
		util.GetErrorMsg(registerForm, err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if err, _ := services.Register(c, registerForm); err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.ErrCodeRegisterFail)
		return
	} else {
		response.Success(c, response.StatusOK, nil)
	}
}
