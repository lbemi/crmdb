import Api from '@/request/api'
/*
  host.POST("", asset.AddHost)        // 添加主机
  host.GET("", asset.ListHosts)       // 获取主机列表
  host.GET("/:id", asset.GetHostById) // 根据id获取主机
  host.PUT("/:id", asset.UpdateHost)  // 根据id修改主机
*/

export const hostApi = {
  add: Api.create('/host', 'post'),
  list: Api.create('/host', 'get'),
  delete: Api.create('/host/{id}', 'delete'),
  update: Api.create('/host/{id}', 'put'),
  listHostByID: Api.create('/host/{id}', 'get'),

  ws: Api.create('/host/{id}/ws', 'get')
}
