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
		return
	}

	for _, cluster := range *clusterList {
		config := util.Decrypt(cluster.KubeConfig)
		_, _, err := f.Cluster().GenerateClient(cluster.Name, config)
		if err != nil {
			if cluster.Status {
				err := f.Cluster().ChangeStatus(cluster.ID, false)
				if err != nil {
					klog.Error(err)
				}
			}
			klog.Errorf("%s集群异常", cluster.Name)
		} else {
			go f.Cluster().StartInformer(cluster.Name)
			if !cluster.Status {
				err := f.Cluster().ChangeStatus(cluster.ID, true)
				if err != nil {
					klog.Error(err)
				}
			}
		}
	}
}
