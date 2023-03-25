package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type CaptchaInfo struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.4, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		log.Logger.Error("生成验证码错误", err)
		response.Fail(c, response.StatusInternalServerError)
		return
	}
	var cap CaptchaInfo
	cap.CaptchaID = id
	cap.PicPath = b64s
	c.Set("captcha_id", id)
	response.Success(c, response.StatusOK, cap)
}
