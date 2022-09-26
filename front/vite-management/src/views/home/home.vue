/**
 * Created by lei on 2022/09/24
 */
<template>
  <div class="home_container">
    <div class="home_header">头部</div>
    <div class="home_menus">
      <el-menu active-text-color="#ffd04b"  class="el-menu-vertical-demo" background-color="#545c64" default-active="2"
         :unique-opened="true"  text-color="#fff" :router="true">
        <el-sub-menu :index="menu.id+''" v-for="menu in menus" :key="menu.id">
          <template #title>
            <el-icon><setting /></el-icon>
            <span>{{menu.name}}</span>
          </template>
            <el-menu-item :index="menu.url+ child.url" v-for="child in menu.children" :key="child.id">
              <el-icon><setting /></el-icon>
              {{child.name}}
            </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </div>
    <div class="home_content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from '@vue/reactivity';
import { useStore } from 'vuex';
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

const store = useStore()
const menus = computed(() => store.state.menus) 
</script>

<style scoped lang="less">
.home_container {
  position: relative;
  height: 100%;

  .home_header {
    height: 70px;
    background-color: antiquewhite;
  }

  .home_menus {
    position: absolute;
    top: 70px;
    left: 0;
    bottom: 0;
    width: 250px;
  }

  .home_content {
    position: absolute;
    top: 70px;
    right: 0;
    left: 250px;
    bottom: 0;
  }
}
</style>
