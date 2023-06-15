package sys

import (
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type CaptchaInfo struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}

func GetCaptcha(rc *rctx.ReqCtx) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.4, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	restfulx.ErrNotNilDebug(err, restfulx.ServerErr)
	var cap CaptchaInfo
	cap.CaptchaID = id
	cap.PicPath = b64s
	rc.Set("captcha_id", id)
	rc.ResData = cap
}
