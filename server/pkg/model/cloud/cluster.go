package cloud

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	basemodel.Model
	Name        string  `json:"name" gorm:"column:name;not null;unique_index:kube_name;comment:集群名称"`
	KubeConfig  string  `json:"-" gorm:"column:kube_config;not null;comment:cloud config"`
	Version     string  `json:"version" gorm:"version:name;comment:kubernetes版本"`
	RunTime     string  `json:"runtime" gorm:"colum:runtime;comment:运行时"`
	ServiceCidr string  `json:"service_cidr" gorm:"colum:service_cidr;comment:service cloud ip"`
	PodCidr     string  `json:"pod_cidr" gorm:"column:pod_cidr;comment:pod id"`
	CNI         string  `json:"cni" gorm:"column:cni;comment:cni网络插件"`
	ProxyMode   string  `json:"proxy_mode" gorm:"column:proxy_mode;comment:网络模式（iptables，ipvs)"`
	Status      bool    `json:"status" gorm:"column:status;comment:集群状态"`
	Nodes       int     `json:"nodes" gorm:"column:nodes;comment:节点数量"`
	InternalIP  string  `json:"internal_ip" gorm:"column:internal_ip"`
	CPU         float64 `json:"cpu" gorm:"column:cpu"`
	Memory      float64 `json:"memory" gorm:"column:memory"`
}

// TableName 表名
func (m *Config) TableName() string {
	return "clusters"
}

// BeforeCreate 添加前
func (m *Config) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *Config) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
