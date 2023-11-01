package entity

import (
	"github.com/lbemi/lbemi/pkg/common/entity"
)

type Tenant struct {
	entity.Model
	TenantName string `json:"tenant_Name"`
}
