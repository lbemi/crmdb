package bootstrap

import (
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
	"k8s.io/klog/v2"
)

func LoadKubernetes(f services.IDbFactory) {
	clusterList, err := f.Cluster().List()
	if err != nil {
		klog.Error(err)
	}
	for _, cluster := range *clusterList {
		config := util.Decrypt(cluster.KubeConfig)
		if _, err := f.Cluster().GenerateClient(cluster.Name, config); err != nil {
			klog.Error(err)
		}
	}
}
