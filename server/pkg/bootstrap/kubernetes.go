package bootstrap

//
//import (
//	"github.com/lbemi/lbemi/pkg/bootstrap/log"
//	"github.com/lbemi/lbemi/pkg/core"
//	"github.com/lbemi/lbemi/pkg/util"
//)
//
//func LoadKubernetes() {
//	//clusterList := f.Cluster().List()
//	//for _, cluster := range *clusterList {
//	//	config := util.Decrypt(cluster.KubeConfig)
//	//	_, _, err := f.Cluster().GenerateClient(cluster.Name, config)
//	//
//	//	if err != nil {
//	//		//如果初始化失败且数据库状态是正常，则修改为异常
//	//		if cluster.Status {
//	//			f.Cluster().ChangeStatus(cluster.ID, false)
//	//		}
//	//		log.Logger.Errorf("%s 集群异常，请检查集群. %v", cluster.Name, err)
//	//		// TODO 是否手设置手动启动监听  启动informer监听
//	//		//go f.Cluster().StartInformer(cluster.Name)
//	//	} else {
//	//		if !cluster.Status {
//	//			// 初始化成功后，如果当前集群状态异常的则改为正常
//	//			f.Cluster().ChangeStatus(cluster.ID, true)
//	//		}
//	//
//	//	}
//	//}
//	list := core.V1.Cluster("").List()
//	for _, cluster := range *list {
//		config := util.Decrypt(cluster.KubeConfig)
//		_, _, err := core.V1.Cluster("").GenerateClient(cluster.Name, config)
//
//		if err != nil {
//			//如果初始化失败且数据库状态是正常，则修改为异常
//			if cluster.Status {
//				core.V1.Cluster("").ChangeStatus(cluster.ID, false)
//			}
//			log.Logger.Errorf("%s 集群异常，请检查集群. %v", cluster.Name, err)
//			// TODO 是否手设置手动启动监听  启动informer监听
//			//go f.Cluster().StartInformer(cluster.Name)
//		} else {
//			if !cluster.Status {
//				// 初始化成功后，如果当前集群状态异常的则改为正常
//				core.V1.Cluster("").ChangeStatus(cluster.ID, true)
//			}
//
//		}
//	}
//}
