package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
	"net/http"
)

func Login(c *gin.Context) {
	userForm := form.UserLoginForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": util.GetErrorMsg(userForm, err),
		})
		return
	}
	c.JSON(http.StatusOK, "login success")
}
func Logout(c *gin.Context) {
	c.String(http.StatusOK, "Logout")
}
