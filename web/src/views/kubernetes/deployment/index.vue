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
      <el-button type="danger" :disabled="data.selectData.length == 0"
        >批量删除</el-button
      >
    </div>

    <el-table
      :data="data.Deployments"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      v-loading="data.loading"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="metadata.name" label="名称" width="220px">
        <template #default="scope">
          <el-button link type="primary" @click="deployDetail(scope.row)">
            {{ scope.row.metadata.name }}</el-button
          >
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
          {{ scope.row.status.readyReplicas }}/{{ scope.row.status.replicas }}
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="180px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.metadata.creationTimestamp) }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted } from 'vue'
import { deploymentApi } from '../api'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { nsStore } from '@/store/kubernetes/namespace'
import { Deployment, Data } from '@/type/deployment'
const ns = nsStore()
const kube = kubeStore()
onMounted(() => {
  listDeployment()
})

const data = reactive(new Data())

const handleSelectionChange = (value: Deployment[]) => {
  data.selectData = value
}
const listDeployment = async () => {
  data.query.namespace = ns.activeNamespace
  data.query.cloud = kube.activeCluster
  try {
    data.loading = true
    await deploymentApi.list.request(data.query).then((res) => {
      data.Deployments = res.data.items
    })
    data.loading = false
  } catch (error) {
    console.log(error)
  }
}
const handleChange = () => {
  data.query.namespace = ns.activeNamespace
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
