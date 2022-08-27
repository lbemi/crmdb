package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
)

func Register(c *gin.Context) {
	var registerForm form.RegisterUserForm
	if err := c.ShouldBind(&registerForm); err != nil {
		fmt.Println(err)
		response.Fail(c, 202, util.GetErrorMsg(registerForm, err))
		return
	}
	if err, _ := services.Register(registerForm); err != nil {
		response.Fail(c, 201, err.Error())
		return
	} else {
		response.Success(c, 200, "注册成功", nil)
	}
}
