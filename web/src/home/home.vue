/** * Created by lei on 2022/09/24 */
<template>

  <div class="common-layout">
    <el-container>
      <el-header
        style="
          display: flex;
          background-color: rgb(40, 43, 51);
          align-items: stretch;
          justify-content: space-between;
          height: 50px;
          padding: 0px 10px;
          color: rgb(173, 176, 184);
          fill: rgb(173, 176, 184);
        "
      >
        <Header />
      </el-header>
      <el-container>
        <!-- <el-card class="box-card" style="height: 100vh"> -->
        <el-aside width="200px">
          <el-scrollbar>
            <el-menu
              active-text-color="#ffd04b"
              class="el-menu-vertical-demo"
              background-color="rgb(51, 55, 68)"
              default-active="2"
              :unique-opened="true"
              text-color="#fff"
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
                <el-menu-item
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
          </el-scrollbar>
        </el-aside>
        <!-- </el-card> -->

        <el-container>
          <el-main>
            <Bredcrumb />
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
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useStore } from "@/store/usestore";
import Header from "@/component/header/index.vue";
import Bredcrumb from "@/component/breadcrumb/index.vue";


import {
  Menu as IconMenu,
  Setting,
} from "@element-plus/icons-vue";
const data = reactive({
  isCollapse: false,
});
const store = useStore();

</script>

<style scoped lang="less">
.expand-icon {
  position: absolute;
  width: 100%;
  bottom: 0;
}
</style>
