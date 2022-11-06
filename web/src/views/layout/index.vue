/** * Created by lei on 2022/09/24 */
<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
      <el-card class="box-card" style="height: 100vh" :body-style="{ padding: '0px' }">
        <el-image style="width: 100px; height: 70px; margin-left: 30px;" src="https://element-plus.gitee.io/images/element-plus-logo.svg"  />
          <el-menu
            active-text-color="#409EFF"
            class="el-menu-vertical-demo"
            :default-active="routeActive"
            :unique-opened="true"
            :router="true"
            style="height: 100vh"
          >
            <el-sub-menu
              :index="menu.id + ''"
              v-for="menu in store.menus"
              :key="menu.id"
            >
              <template #title>
                <el-icon>
                  <setting />
                </el-icon>
                <span>{{ menu.name }}</span>
              </template>
              <el-menu-item v-if="menu.children != null"
                :index="menu.url + child.url"
                v-for="child in menu.children"
                :key="child.id"
              >
                <el-icon>
                  <setting />
                </el-icon>
                {{ child.name }}
              </el-menu-item>
            </el-sub-menu>
          </el-menu>
      </el-card>
      </el-aside>
      <el-container>
        <el-header style="padding: 0px 5px">
          <Header />
        </el-header>

        <el-main style="padding: 10px 5px;margin-top: 10px;">
          <Bredcrumb/>
          <router-view></router-view>
        </el-main>
        <el-footer
          style="
            display: flex;
            align-items: stretch;

            justify-content: center;
            height: 50px;
            padding: 0px 10px;
          "
          >Wirte by Lbemi</el-footer
        >
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useStore } from "@/store/usestore";
import Header from "./header/index.vue";
import Bredcrumb from "@/component/breadcrumb/index.vue";
import { useRoute } from "vue-router";
import { Menu as IconMenu, Setting } from "@element-plus/icons-vue";
const data = reactive({
  isCollapse: false,
});
const route = useRoute()
const  routeActive=ref()
routeActive.value = route.path
const store = useStore();
</script>

<style scoped lang="less">
.expand-icon {
  position: absolute;
  width: 100%;
  bottom: 0;
}
</style>
