import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import { App, createVNode, render } from "vue";
import { useStore } from "@/store/usestore";
import { storeToRefs } from "pinia";
import { MenuObj } from "@/store/usestore";
import loadingBar from "@/component/loadingBar/index.vue";

const Vnode = createVNode(loadingBar);
render(Vnode, document.body);

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    name: "login",
    component: () => import("../views/login/index.vue"),
    meta: {
      title: "登陆",
    },
  },

  {
    path: "/404",
    name: "notFound",
    component: () => import("@/views/error/404.vue"),
    meta: {
      title: "找不到此页面",
    },
  },
  {
    path: "/401",
    name: "noPower",
    component: () => import("@/views/error/401.vue"),
    meta: {
      title: "没有权限",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  // 记录路由页面位置
  scrollBehavior: (to, from, savePostion) => {
    if (savePostion) {
      return savePostion;
    } else {
      return {
        top: 0,
      };
    }
  },
  routes,
});

const genRouters = (menus: MenuObj[]) => {
  for (let key in menus) {
    const newRoute: RouteRecordRaw = {
      path: menus[key].url,
      name: menus[key].name,
      component: () => import("../views/layout/index.vue"),
      // redirect: menus[key].url + menus[key].children[0]?.url,
      children: [],
    };
    if (menus[key].children != null) {
      for (let i = 0; i < menus[key].children.length; i++) {
        let vueUrl = `../views${menus[key].url}${menus[key].children[i].url}${menus[key].children[i].url}`;
        newRoute.children?.push({
          path: menus[key].url + menus[key].children[i].url,
          name: menus[key].children[i].name,
          component: () => import(`${vueUrl}.vue`),
        });
      }
    }
    // 动态添加路由规则
    router.addRoute(newRoute);
  }
  router.addRoute({
    path: "/",
    name: "",
    component: () => import("../views/layout/index.vue"),
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: () => import("../views/dashboard/dashboard.vue"),
      },
    ],
  });
};

//前置导航守卫
router.beforeEach((to, from, next) => {
  Vnode.component?.exposed?.startLoading();
  const store = useStore();
  const { menus } = storeToRefs(store);

  const token = localStorage.getItem("token");

  if (token && menus.value.length === 0) {
    // 异步请求,then是异步完成后操作
    store.getLeftMenus().then(() => {
      genRouters(menus.value);
      next(to);
    });
  } else if (
    token &&
    menus.value.length !== 0 &&
    from.path === "/login" &&
    to.path === "/home"
  ) {
    // genRouters(menus.value)
    next("/dashboard");
  } else if (!token && to.path !== "/login") {
    //token不存在,跳转到登录页面
    next("/login");
  } else if (token && to.path === "/login") {
    //登录后禁止访问login
    next(from);
  } else if (router.getRoutes().length <= routes.length) {
    // 刷新后重新生成路由
    genRouters(menus.value);
    next(to);
  } else {
    console.log(router.getRoutes().length,"-",routes.length);
    
    next();
  }
});

router.afterEach((to, from) => {
  Vnode.component?.exposed?.endLoading();
});
export const initRouter = (app: App<Element>) => {
  app.use(router);
};

export default router;
