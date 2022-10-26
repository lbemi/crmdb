package sys

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
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(registerForm, err))
		return
	}

	if services.CheckUserExist(registerForm.UserName) {
		response.Fail(c, response.ErrCodeUserExist)
		return
	}

	if err, _ := services.Register(c, registerForm); err != nil {
		global.App.Log.Error(err.Error())
		response.FailWithMessage(c, response.ErrCodeRegisterFail, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)

}
