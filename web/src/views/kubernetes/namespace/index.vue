/** * Created by lei on 2022/12/09 */
<template>
  <div class="main">
    <el-button type="primary" @click="createNamespace()"
      >创建Namespace</el-button
    >
    <el-button type="danger" :disabled="data.selectData.length == 0"
      >批量删除</el-button
    >
    <el-table
      :data="ns.namespace"
      style="width: 100%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="35" />
      <el-table-column fixed prop="metadata.name" label="名称" width="150" />
      <el-table-column prop="metadata.labels" label="标签" width="500">
        <template #default="scope">
          <el-tag
            type="info"
            v-for="(item, key, index) in scope.row.metadata.labels"
            :key="index"
          >
            {{ key }}:{{ item }}
          </el-tag>
        </template></el-table-column
      >
      <el-table-column label="状态" width="120">
        <template #default="scope">
          <div v-if="scope.row.status.phase === 'Active'">
            <div
              style="
                display: inline-block;
                width: 12px;
                height: 12px;
                background: #67c23a;
                border-radius: 50%;
              "
            ></div>
            <span style="margin-left: 5px; font-size: 12px; color: #67c23a"
              >就绪
            </span>
          </div>
          <div v-else>
            <div
              style="
                display: inline-block;
                width: 12px;
                height: 12px;
                background: #f56c6c;
                border-radius: 50%;
              "
            ></div>
            <span style="margin-left: 5px; font-size: 12px; color: #f56c6c"
              >未就绪
            </span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="160">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.metadata.creationTimestamp) }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="220">
        <template #default>
          <el-button link type="primary" size="small">资源配额与限制</el-button
          ><el-divider direction="vertical" />
          <el-button link type="primary" size="small">编辑</el-button
          ><el-divider direction="vertical" />
          <el-button link type="primary" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

  <NamespaceDialog
    :title="data.titile"
    v-model:visible="data.visible"
    v-model:namespace="data.namespace"
    @value-change="ns.listNamespace()"
  />
</template>

<script setup lang="ts">
import { nsStore } from '@/store/kubernetes/namespace'
import { reactive } from 'vue'
import { NamespaceData, Namespace } from '@/type/namespace'
import NamespaceDialog from './component/dialog.vue'

const ns = nsStore()
const data = reactive(new NamespaceData())
const handleSelectionChange = (value: Namespace[]) => {
  data.selectData = value
}

const createNamespace = () => {
  data.titile = '创建命名空间'
  console.log(data.visible)

  data.visible = true
}
</script>

<style scoped lang="less">
.main {
  margin-left: 15px;

  .el-table {
    margin-top: 10px;
  }
}
</style>
