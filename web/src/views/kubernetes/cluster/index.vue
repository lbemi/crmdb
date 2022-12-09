/** * Created by lei on 2022/11/29 */
<template>
  <div class="body">
    <!-- <div>Cluster-集群</div> -->
    <el-tabs type="border-card">
      <el-tab-pane lazy label="概览"><Dashboard /></el-tab-pane>
      <el-tab-pane lazy label="基本信息"><Info /></el-tab-pane>
      <el-tab-pane label="连接信息">连接信息</el-tab-pane>
      <el-tab-pane label="集群资源">集群资源</el-tab-pane>
      <el-tab-pane lazy label="集群日志"><Log /></el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import Dashboard from './dashboard.vue'
import Info from './info.vue'
import Log from './log.vue'
</script>

<style scoped lang="less">
.body {
  margin-left: 10px;
  width: 100%;
  // height: 100;
}
.el-tabs--border-card {
  // border-bottom: 0px;
  // border-left: 0px;
  // border-right: 0px;
  border: 0;
}
</style>
