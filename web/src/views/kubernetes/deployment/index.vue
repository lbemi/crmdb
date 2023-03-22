<template>
  <div style="margin-left: 5px">
    <div>
      Namespace:
      <el-select
        v-model="ns.activeNamespace"
        class="m-2"
        placeholder="Select"
        @change="handleChange"
        ><el-option key="all" label="所有命名空间" value="all"></el-option>
        <el-option
          v-for="item in ns.namespace"
          :key="item.metadata.name"
          :label="item.metadata.name"
          :value="item.metadata.name"
        />
      </el-select>
      <el-button type="primary">创建Deployment</el-button>
      <el-button
        type="danger"
        :disabled="data.selectData.length == 0"
        @click="depoly.deleteDeoloyments(data.selectData)"
        >批量删除</el-button
      >
    </div>

    <el-table
      :data="data.Deployments"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      v-loading="data.loading"
      max-height="100vh - 235px"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column prop="metadata.name" label="名称" width="220px">
        <template #default="scope">
          <el-button link type="primary" @click="deployDetail(scope.row)">
            {{ scope.row.metadata.name }}</el-button
          >
        </template>
      </el-table-column>
      <el-table-column label="状态" width="220px">
        <template #default="scope">
          <el-button
            v-if="scope.row.status.conditions[0].status === 'True'"
            type="success"
            :icon="Check"
            size="small"
            circle
          />
          <el-button v-else type="danger" :icon="Close" size="small" circle />
        </template>
      </el-table-column>

      <el-table-column label="镜像" width="540px">
        <template #default="scope">
          <el-tag
            type="success"
            v-for="(item, index) in scope.row.spec.template.spec.containers"
            :key="index"
            >{{ item.image.split('@')[0] }}</el-tag
          >
        </template>
      </el-table-column>

      <el-table-column label="标签" width="280px">
        <template #default="scope">
          <el-tag
            type="info"
            v-for="(item, key, index) in scope.row.metadata.labels"
            :key="index"
          >
            {{ key }}:{{ item }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column prop="spec.replicas" label="Pods" width="80px">
        <template #default="scope">
          <a style="color: green">{{ scope.row.status.readyReplicas || '0' }}</a
          >/ <a style="color: green">{{ scope.row.status.replicas }}</a
          >/
          <a style="color: red">{{
            scope.row.status.unavailableReplicas || '0'
          }}</a>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="180px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.metadata.creationTimestamp) }}
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页区域 -->
    <pagination
      :total="data.total"
      @handlePageChange="handlePageChange"
    ></pagination>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted, onBeforeUnmount } from 'vue'
import { deploymentApi } from '../api'
import pagination from '@/component/pagination/pagination.vue'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { nsStore } from '@/store/kubernetes/namespace'
import { Deployment, Data } from '@/type/deployment'
import { PageInfo } from '@/type/sys'
import { Check, Close } from '@element-plus/icons-vue'
import { deployStore } from '@/store/kubernetes/deployment'
import { webSocketURL } from '@/request/request'
const depoly = deployStore()
const ns = nsStore()
const kube = kubeStore()
onMounted(() => {
  listDeployment()
})

var dns = webSocketURL + kube.activeCluster + '/deployment'
var ws = new WebSocket(dns)
ws.onopen = () => {
  console.log('ws connected.')
}
ws.onmessage = (e) => {
  if (e.data === 'ping') {
    return
  } else {
    const object = JSON.parse(e.data)
    if (
      object.type === 'deployment' &&
      object.result.namespace === ns.activeNamespace &&
      object.cluster == kube.activeCluster
    ) {
      data.Deployments = object.result.data
    }
  }
}
ws.onclose = () => {
  console.log('close')
}

const data = reactive(new Data())

const handleSelectionChange = (value: Deployment[]) => {
  data.selectData = value
  console.log(data.selectData)
}

onBeforeUnmount(() => {
  console.log('关闭....')
  ws.close()
})
const listDeployment = async () => {
  data.query.namespace = ns.activeNamespace
  data.query.cloud = kube.activeCluster
  try {
    data.loading = true
    await deploymentApi.list.request(data.query).then((res) => {
      data.Deployments = res.data.data
      data.total = res.data.total
    })
  } catch (error) {
    console.log(error)
  }
  data.loading = false
}
const handleChange = () => {
  data.query.namespace = ns.activeNamespace
  listDeployment()
}
const handlePageChange = (pageInfo: PageInfo) => {
  data.query.page = pageInfo.page
  data.query.limit = pageInfo.limit
  listDeployment()
}
const deployDetail = (deploy: Deployment) => {
  data.query.deploymentName = deploy.metadata.name
  data.query.namespace = deploy.metadata.namespace
  console.log(data.query)

  const res = deploymentApi.get.request(data.query)

  console.log(res)
}
</script>

<style scoped lang="less"></style>
