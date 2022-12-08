/** * Created by lei on 2022/11/29 */
<template>
  <div class="body">
    <div>Cluster-集群</div>
    <el-tabs type="border-card">
      <el-tab-pane label="概览"><Dashboard /></el-tab-pane>
      <el-tab-pane label="基本信息">基本信息</el-tab-pane>
      <el-tab-pane label="连接信息">连接信息</el-tab-pane>
      <el-tab-pane label="集群资源">集群资源</el-tab-pane>
      <el-tab-pane label="集群日志">集群日志</el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import Dashboard from './dashboard.vue'
// import { clusterInfo } from '@/type/cluster'
// import { onMounted, reactive } from 'vue'
// import { clusterApi } from '@/views/container/api'
// import { kubeStore } from '@/store/kubernetes/kubernetes'

// const clu = kubeStore()
// const data = reactive({
//   cluster: {} as clusterInfo,
//   query: {
//     name: ''
//   }
// })

// onMounted(() => {
//   getCluster()
// })
// const getCluster = () => {
//   data.query.name = clu.activeCluster
//   console.log(clu.activeCluster)

//   const res = clusterApi.get.request(data.query)
//   console.log(res)
// }
</script>

<style scoped lang="less">
.body {
  margin-left: 10px;
  width: 100%;
  height: 100%;
}
</style>
