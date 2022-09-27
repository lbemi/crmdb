import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import { App } from 'vue'
import pinia from '@/store/index'
import {useStore} from '@/store/usestore';
const store = useStore(pinia)
const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login/index.vue'),
    meta: {
      title: '登陆',
    },
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

const genRouters = () => {
  const menus = store.$state.menus
  for (let key in menus) {
    const newRoute: RouteRecordRaw = {
      path: menus[key].url,
      name: menus[key].name,
      component: () => import('../views/home/home.vue'),
      // redirect: menus[key].url + menus[key].children[0]?.url,
      children: []
    }
    if (menus[key].children != null) {
      for (let i = 0; i < menus[key].children.length; i++) {
        let vueUrl = `../views${menus[key].url}${menus[key].children[i].url}.vue`
        newRoute.children?.push(
          {
            path: menus[key].url + menus[key].children[i].url,
            name: menus[key].children[i].name,
            component: () => import(`../views${menus[key].url}${menus[key].children[i].url}.vue`),
          }
        )
      }
    }
    // 动态添加路由规则
    router.addRoute(newRoute)
  }
  router.addRoute(
    {
      path: '/',
      name: 'homepage',
      component: () => import('../views/home/home.vue'),
      redirect: '/dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('../views/dashboard/dashboard.vue'),
        },
      ]

    },
  )
}

//前置导航守卫
router.beforeEach((to, from, next) => {
  // const store = useStore()
  const token = sessionStorage.getItem("token")
  if (token && store.menus.length === 0) {
    // 异步请求,then是异步完成后操作
    store.getLeftMenusApi().then(() => {
      // const newRoutes:RouteRecordRaw[] = []
      genRouters()
      next(to)
    })
  } else if (token && store.menus.length !== 0 && from.path === '/login' && to.path === '/home') {
    genRouters()
    next("/dashboard")
  } else if (!token && to.path !== '/login') { //token不存在,跳转到登录页面
    next('/login')
  } else if (token && to.path === '/login') { //登录后禁止访问login
    next(from)
  } else {
    next()
  }
})

export const initRouter = (app: App<Element>) => {
  app.use(router)
}


export default router