package sys

import (
	"github.com/lbemi/lbemi/pkg/model"
)

type Tenant struct {
	model.Model
	TenantName string `json:"tenant_Name"`
}
