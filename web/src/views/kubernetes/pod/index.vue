/** * Created by lei on 2023/03/21 */
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
      <el-button type="danger" :disabled="pod.data.selectData.length == 0"
        >批量删除</el-button
      >
    </div>

    <el-table
      :data="pod.data.pods"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      v-loading="pod.data.loading"
      max-height="100vh - 235px"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column prop="metadata.name" label="名称" width="320px">
        <template #default="scope">
          <el-button link type="primary">{{
            scope.row.metadata.name
          }}</el-button>
          <div v-if="scope.row.status.phase != 'Running'" style="color: red">
            {{ scope.row.status.containerStatuses[0].state }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态" width="80px">
        <template #default="scope">
          <span v-if="scope.row.status.phase == 'Running'" style="color: green">
            {{ scope.row.status.phase }}</span
          >
          <span v-else style="color: red"> {{ scope.row.status.phase }}</span>
        </template>
      </el-table-column>
      <el-table-column label="重启次数" width="80px">
        <template #default="scope">
          {{ scope.row.status.containerStatuses[0].restartCount }}
        </template>
      </el-table-column>
      <el-table-column label="标签" width="280px" show-overflow-tooltip>
        <template #default="scope">
          <div class="ellipsis">
            <el-tag
              type="info"
              v-for="(item, key, index) in scope.row.metadata.labels"
              :key="index"
            >
              <span>{{ key }}:{{ item }}</span>
            </el-tag>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="status.podIP" label="IP" width="220px">
        <template #default="scope">
          {{ scope.row.status.podIP }}
        </template>
      </el-table-column>
      <el-table-column prop="spec.nodeName" label="所在节点" width="220px">
        <template #default="scope">
          <div>{{ scope.row.spec.nodeName }}</div>
          <div>{{ scope.row.status.hostIP }}</div>
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
      :total="pod.data.total"
      @handlePageChange="handlePageChange"
    ></pagination>
  </div>
</template>

<script setup lang="ts">
import pagination from '@/component/pagination/pagination.vue'
import { nsStore } from '@/store/kubernetes/namespace'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { podStore } from '@/store/kubernetes/pods'
import { onMounted } from 'vue'
import { PageInfo } from '@/type/sys'
import { webSocketURL } from '@/request/request'
const pod = podStore()
const ns = nsStore()
const kube = kubeStore()
onMounted(() => {
  pod.data.loading = true
  pod.listPods()
  pod.data.loading = false
})

const handleSelectionChange = () => {}

const handleChange = () => {
  pod.data.loading = true
  pod.listPods()
  pod.data.loading = false
}

const handlePageChange = (pageInfo: PageInfo) => {
  pod.data.query.page = pageInfo.page
  pod.data.query.limit = pageInfo.limit
  pod.data.loading = true
  pod.listPods()
  pod.data.loading = false
}

var dns = webSocketURL + kube.activeCluster + '/pod'
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
      object.type === 'pod' &&
      object.result.namespace === ns.activeNamespace &&
      object.cluster == kube.activeCluster
    ) {
      pod.data.pods = object.result.data
    }
  }
}
ws.onclose = () => {
  console.log('close')
}
</script>

<style scoped lang="less">
.ellipsis {
  height: 60px;
  white-space: nowrap;
  overflow: hidden;
  overflow-y: auto;
  // text-overflow: ellipsis;
}
</style>
