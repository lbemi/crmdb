/** * Created by lei on 2022/09/24 */
<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="160px" >
        <!-- <el-card
          class="box-card"
          style="height: 100vh"
          :body-style="{ padding: '0px' }"
        >  -->
        <el-menu
          active-text-color="#409EFF"
          class="el-menu-vertical-demo"
          :default-active="routeActive"
          :unique-opened="true"
          :router="true"
          style="height: 100vh"
          :collapse="isCollapse"
          @open="handleOpen"
          @close="handleClose"
        >
          <el-icon style="width: 100%">
            <img
              style="
                width: 100px;
                height: 70px;
                margin-top: 50px;
                margin-right: 30px;
              "
              src="@/assets/image/element-plus-logo.svg"
            />
          </el-icon>
          <el-icon @click="isCollapse = !isCollapse">
            <SvgIcon iconName="icon-zhankaicaidan" className="expand-icon" />
          </el-icon>
          <template v-for="menu in store.menus">
            <el-menu-item v-if="menu.children === null" :index="menu.url">
              <el-icon style="width: 0.5em; height: 0.5em; margin-right: 10px">
                <SvgIcon :iconName="menu.icon" />
              </el-icon>
              <template #title>{{ menu.name }}</template>
            </el-menu-item>
            <el-sub-menu v-else :index="menu.id + ''" :key="menu.id">
              <template #title>
                <el-icon
                  style="width: 0.5em; height: 0.5em; margin-right: 10px"
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
                  style="width: 0.5em; height: 0.5em; margin-right: 10px"
                >
                  <SvgIcon :iconName="child.icon" />
                </el-icon>
                {{ child.name }}
              </el-menu-item>
            </el-sub-menu>
          </template>
        </el-menu>
        <!-- </el-card> -->
      </el-aside>
      <el-container>
        <el-header style="padding: 0px 5px">
          <Header />
        </el-header>

        <el-main style="padding: 10px 5px; margin-top: 12px">
          <Bredcrumb />
          <router-view></router-view>
        </el-main>
        <el-footer
        style="text-align:center; height: 30px;"
        >底线....</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useStore } from "@/store/usestore";
import Header from "./header/index.vue";
import Bredcrumb from "@/component/breadcrumb/index.vue";
import { useRoute } from "vue-router";

const route = useRoute();
const routeActive = ref();
routeActive.value = route.path;
const store = useStore();
const isCollapse = ref<boolean>(false);
const handleOpen = (key: string, keyPath: string[]) => {};
const handleClose = (key: string, keyPath: string[]) => {};
</script>

<style scoped lang="less">
.expand-icon {
  position: relative;
  width: 20px;
  height: 20px;
  bottom: 0;
  margin-top: 15px;
}
</style>
