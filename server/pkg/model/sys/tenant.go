package sys

import "github.com/lbemi/lbemi/pkg/model/basemodel"

type Tenant struct {
	basemodel.Model
	TenantName string `json:"tenant_Name"`
}
