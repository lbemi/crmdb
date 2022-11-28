import Api from "@/request/api"

 /*
    cluster.POST("", cloud.CreateCluster)
    cluster.GET("", cloud.ListCluster)
    cluster.DELETE("/:id", cloud.DeleteCluster)
*/

 export const clusterApi = {
    list: Api.create('/cluster', 'get'),
    create: Api.create('/cluster', 'post'),
    delete: Api.create('/cluster/{id}', 'delete'),
  }
  
