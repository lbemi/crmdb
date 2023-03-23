export const kubernetesRoutes = [
  {
    id: 1,
    name: '集群信息',
    path: '/cluster'
  },
  {
    id: 2,
    name: '命名空间',
    path: '/namespace'
  },
  {
    id: 3,
    name: '节点管理',
    path: '/cluster/nodes',
    children: [
      {
        id: 3.1,
        name: '节点信息',
        path: '/nodes'
      }
    ]
  },
  {
    id: 5,
    name: '工作负载',
    path: '/',
    children: [
      {
        id: 5.1,
        name: 'deployments',
        path: '/deployment'
      },
      {
        id: 5.2,
        name: 'Stateful Sets',
        path: '/statefulsets'
      },
      {
        id: 5.3,
        name: 'Daemon Sets',
        path: '/daemonsets'
      }
    ]
  }
]
