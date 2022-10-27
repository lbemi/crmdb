package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
)

func Register(c *gin.Context) {
	var registerForm form.RegisterUserForm
	if err := c.ShouldBind(&registerForm); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(registerForm, err))
		return
	}

	if core.Core.User().CheckUserExist(c, registerForm.UserName) {
		response.Fail(c, response.ErrCodeUserExist)
		return
	}

	if err := core.Core.User().Register(c, &registerForm); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeRegisterFail, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)

}
