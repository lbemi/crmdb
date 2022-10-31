import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import { App } from "vue";
import pinia from "@/store/index";
import { useStore } from "@/store/usestore";
import { storeToRefs } from "pinia";
import { MenuObj } from "@/store/usestore";

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    name: "login",
    component: () => import("../views/login/index.vue"),
    meta: {
      title: "登陆",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

const genRouters = (menus: MenuObj[]) => {
  for (let key in menus) {
    const newRoute: RouteRecordRaw = {
      path: menus[key].url,
      name: menus[key].name,
      component: () => import("../views/home/home.vue"),
      // redirect: menus[key].url + menus[key].children[0]?.url,
      children: [],
    };
    if (menus[key].children != null) {
      for (let i = 0; i < menus[key].children.length; i++) {
        let vueUrl = `../views${menus[key].url}${menus[key].children[i].url}`;
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
    component: () => import("../views/home/home.vue"),
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
  const store = useStore();
  const { menus } = storeToRefs(store);

  const token = localStorage.getItem("token");

  if (token && menus.value.length === 0) {
    console.log("-------", menus.value);
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
    console.log("????????????");

    // genRouters(menus.value)
    next("/dashboard");
  } else if (!token && to.path !== "/login") {
    //token不存在,跳转到登录页面
    next("/login");
  } else if (token && to.path === "/login") {
    //登录后禁止访问login
    next(from);
  }  else {
    console.log("11111111111111", menus.value, "*****", router.getRoutes);
    next();
  }
});

export const initRouter = (app: App<Element>) => {
  app.use(router);
};

export default router;
