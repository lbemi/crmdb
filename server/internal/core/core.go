package core

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common"
	"github.com/lbemi/lbemi/pkg/model/configs"
	"gorm.io/gorm"
)

type CoreV1 interface {
	InitConfig(string) error
	InintDb() error
	InitServer() error
}
type register struct {
	config    *configs.Config
	ginEngine *gin.Engine
}

func NewRegister(db *gorm.DB) *register {
	return &register{}
}

func (r *register) InitConfig(path string) error {
	config, err := common.InitConfig(path)
	if err != nil {
		return err
	}
	r.config = config
	return nil
}
