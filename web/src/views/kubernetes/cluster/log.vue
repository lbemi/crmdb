/** * Created by lei on 2022/12/09 */
<template>
  <div>
    <el-table
      :data="data.events"
      :default-sort="{
        prop: 'metadata.creationTimestamp',
        order: 'descending'
      }"
      v-loading="data.loading"
    >
      <el-table-column
        prop="metadata.creationTimestamp"
        label="时间"
        width="180"
        sortable
      >
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.metadata.creationTimestamp) }}
        </template>
      </el-table-column>
      <el-table-column
        label="类型"
        width="100"
        :filters="[
          { text: 'Normal', value: 'Normal' },
          { text: 'Warning', value: 'Warning' }
        ]"
        :filter-method="filterTag"
      >
        <template #default="scope">
          <el-tag v-if="scope.row.type === 'Normal'" type="success">
            {{ scope.row.type }}
          </el-tag>
          <el-tag v-else-if="scope.row.type === 'Warning'" type="warning">
            {{ scope.row.type }}
          </el-tag>
          <el-tag v-else type="danger">
            {{ scope.row.type }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="信息">
        <template #default="scope">
          {{ scope.row.metadata.uid.split('-')[0]
          }}<el-divider direction="vertical" /> {{ scope.row.reason }}|
          {{ scope.row.message }}<el-divider direction="vertical" />
          <span style="color: #409eff">{{
            scope.row.metadata.name.split('.')[0]
          }}</span>
        </template></el-table-column
      >
    </el-table>
    <!-- 分页区域 -->
    <pagination
      :total="data.total"
      @handlePageChange="handlePageChange"
    ></pagination>
  </div>
</template>

<script setup lang="ts">
import { EventData, Event } from '@/type/event'
import { onMounted, reactive } from 'vue'
import { eventApi } from '../api'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import pagination from '@/component/pagination/pagination.vue'
import type { ElTableColumn } from 'element-plus'
import { PageInfo } from '@/type/sys'

const kube = kubeStore()
const data = reactive(new EventData())

const filterTag = (value: string, row: Event) => {
  return row.type === value
}
const listEvent = () => {
  data.loading = true
  data.query.cloud = kube.activeCluster
  data.query.namespace = 'all'
  eventApi.list.request(data.query).then((res) => {
    data.events = res.data.data
    data.total = res.data.total
  })
  data.loading = false
}
const handlePageChange = (pageInfo: PageInfo) => {
  data.query.page = pageInfo.page
  data.query.limit = pageInfo.limit
  listEvent()
}
onMounted(() => {
  listEvent()
})
</script>

<style scoped lang="less">
.el-table {
  height: calc(100vh - 260px);
}
</style>
