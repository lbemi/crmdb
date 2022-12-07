<template>
  <div style="margin-left: 5px">
    <div>
      namespace:
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
      <el-button type="primary">添加</el-button>
      <el-button ref="delete_button" type="danger" disabled>删除</el-button>
    </div>

    <el-table
      :data="delploys.Deployments"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      v-loading="loading"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column prop="metadata.name" label="名称" width="220px" />
      
      <el-table-column label="镜像" width="400px">
        <template #default="scope">
          <el-tag
            type="success"
            v-for="(item, index) in scope.row.spec.template.spec.containers"
            :key="index"
            >{{ item.image }}</el-tag
          >
        </template>
      </el-table-column>

      <el-table-column label="标签" width="180px">
        <template #default="scope">
          <el-tag
            type="info"
            v-for="(item, index) in scope.row.metadata.labels"
            :key="index"
            >{{ item }}</el-tag
          >
        </template>
      </el-table-column>

      <el-table-column prop="spec.replicas" label="Pods" width="80px">
        <template #default="scope">
          {{ scope.row.status.readyReplicas }}/{{ scope.row.status.replicas }}
        </template>
      </el-table-column>
      <el-table-column
        prop="metadata.creationTimestamp"
        label="创建时间"
        width="180px"
      />
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted, ref } from 'vue'
import { deploymentApi } from '../api'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { nsStore } from '@/store/kubernetes/namespace'
import { deploymentData } from '@/type/deployment'
const loading = ref(false)
const ns = nsStore()
const kube = kubeStore()
onMounted(() => {
  listDeployment()
})

const query = reactive({
  namespace: ns.activeNamespace,
  cloud: kube.activeCluster
})

const delploys = reactive(new deploymentData())

const delete_button = ref()
const handleSelectionChange = () => {
  console.log(delete_button.value.disabled)
}
const listDeployment = async () => {
  try {
    loading.value = true
    await deploymentApi.list.request(query).then((res) => {
      delploys.Deployments = res.data.items
      loading.value = false
    })
  } catch (error) {
    console.log(error)
  }
}
const handleChange = () => {
  query.namespace = ns.activeNamespace
  listDeployment()
}
</script>

<style scoped lang="less"></style>
