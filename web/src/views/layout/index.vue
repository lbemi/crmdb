/** * Created by lei on 2022/09/24 */
<template>
  <div class="common-layout">
    <el-container>
      <el-aside style="">
        <el-menu
          active-text-color="#409EFF"
          :default-active="routeActive"
          :unique-opened="true"
          :router="true"
          :collapse="isCollapse"
        >
          <el-icon style="width: 100%">
            <img
              style="
                margin-top: 50px;
                margin-right: 30px;
                width: 100px;
                height: 70px;
              "
              src="@/assets/image/element-plus-logo.svg"
            />
          </el-icon>

          <template v-for="menu in store.menus">
            <el-menu-item
              v-if="menu.children === null"
              :index="menu.url"
              :key="menu.id"
            >
              <el-icon style="margin-right: 10px; width: 0.5em; height: 0.5em">
                <SvgIcon :iconName="menu.icon" />
              </el-icon>
              <template #title>{{ menu.name }}</template>
            </el-menu-item>
            <el-sub-menu v-else :index="menu.id + ''" :key="menu.url">
              <template #title>
                <el-icon
                  style="margin-right: 10px; width: 0.5em; height: 0.5em"
                >
                  <SvgIcon :iconName="menu.icon" />
                </el-icon>
                <span>{{ menu.name }}</span>
              </template>
              <el-menu-item
                :index="menu.url + child.url"
                v-for="child in menu.children"
                :key="child.id"
              >
                <el-icon
                  style="margin-right: 10px; width: 0.5em; height: 0.5em"
                >
                  <SvgIcon :iconName="child.icon" />
                </el-icon>
                {{ child.name }}
              </el-menu-item>
            </el-sub-menu>
          </template>
          <el-affix target=".el-menu" position="bottom">
            <el-icon @click="isCollapse = !isCollapse">
              <SvgIcon iconName="icon-zhankaicaidan" className="icon-1-4em" />
            </el-icon>
          </el-affix>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header style="padding: 0 5px">
          <Header />
        </el-header>

        <el-main style="padding: 10px 5px; margin-top: 5px">
          <Bredcrumb />
          <router-view v-if="isRouterActive"></router-view>
        </el-main>
        <el-footer style="height: 30px; text-align: center">底线....</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick, provide } from 'vue'
import { useStore } from '@/store/usestore'
import Header from './header/index.vue'
import Bredcrumb from '@/component/breadcrumb/index.vue'
import { useRoute } from 'vue-router'
const isRouterActive = ref(true)
const reload = () => {
  isRouterActive.value = false
  nextTick(() => {
    isRouterActive.value = true
  })
}
provide('reload', reload)

const route = useRoute()
const routeActive = ref()
routeActive.value = route.path
const store = useStore()
const isCollapse = ref<boolean>(false)
</script>

<style scoped lang="less">
.el-aside {
  width: auto !important;
  /** 宽度自适应 */
  height: 100vh;
  text-align: center;
  flex-direction: column;
  .el-menu:not(.el-menu--collapse) {
    width: 200px;
    height: 100%;
  }
}
.el-header {
  height: 40px;
}
</style>
