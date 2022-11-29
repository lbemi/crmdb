import Api from "@/request/api"

/*
   namespace.GET("", cloud.ListNamespace)
   namespace.GET("/:name", cloud.GetNamespace)
   namespace.POST("", cloud.CreateNamespace)
   namespace.DELETE("/:name", cloud.DeleteNamespace)
*/

export const namespacerApi = {
    list: Api.create('/namespace', 'get'),
    get: Api.create('/namespace/{name}', 'get'),
    create: Api.create("/namespace", 'post'),
    delete: Api.create('/namespace/{name}', 'delete'),
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
    create: Api.create("/deployment", 'post'),
    delete: Api.create('/deployment/{deploymentName}', 'delete'),
}

