import Api from '@/request/api'

/*
   namespace.GET("", cloud.ListNamespace)
   namespace.GET("/:name", cloud.GetNamespace)
   namespace.POST("", cloud.CreateNamespace)
   namespace.DELETE("/:name", cloud.DeleteNamespace)
*/

export const namespacerApi = {
  list: Api.create('/namespace', 'get'),
  get: Api.create('/namespace/{name}', 'get'),
  create: Api.create('/namespace', 'post'),
  update: Api.create('/namespace', 'put'),
  delete: Api.create('/namespace/{name}', 'delete')
}

/*
    deployment.GET("/:namespace", cloud.ListDeployments)
    deployment.GET("/:namespace/:deploymentName", cloud.GetDeployment)
    deployment.POST("", cloud.CreateDeployment)
    deployment.PUT("", cloud.UpdateDeployment)
    deployment.DELETE("/:namespace/:deploymentName", cloud.DeleteDeployment)
*/
export const deploymentApi = {
  list: Api.create('/deployment/{namespace}', 'get'),
  get: Api.create('/deployment/{namespace}/{deploymentName}', 'get'),
  create: Api.create('/deployment', 'post'),
  delete: Api.create('/deployment/{deploymentName}', 'delete')
}

/*
	node := group.Group("/node")
	{
		node.GET("", cloud.ListNodes)
		node.GET("/:nodeName", cloud.GetNode)
	}
*/
export const nodeApi = {
  list: Api.create('/node', 'get'),
  get: Api.create('/node/{nodeName}', 'get'),
  update: Api.create('/node', 'put'),
}

/*
	// event 事件路由
	event := group.Group("/event")
	{
		event.GET("/:namespace", cloud.ListEvents)
		event.GET("/:namespace/:name", cloud.GetEvent)
	}
*/

export const eventApi = {
  list: Api.create('/event/{namespace}', 'get'),
  get: Api.create('/event/{namespace}/{name}', 'get')
}
