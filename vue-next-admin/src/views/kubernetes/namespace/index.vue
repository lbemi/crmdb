/** * Created by lei on 2022/12/09 */
<template>
  <el-button type="primary" @click="createNamespace()">创建Namespace</el-button>
  <el-button type="danger" :disabled="data.selectData.length == 0"
    >批量删除</el-button
  >
  <el-table
    :data="ns.namespace"
    @selection-change="handleSelectionChange"
    style="width: 100%"
  >
    <el-table-column type="selection" width="35" />
    <el-table-column prop="metadata.name" label="名称" width="220" />
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
            >{{ scope.row.status.phase }}
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
            >{{ scope.row.status.phase }}
          </span>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="创建时间" width="170">
      <template #default="scope">
        {{ $filters.dateFormat(scope.row.metadata.creationTimestamp) }}
      </template>
    </el-table-column>

    <el-table-column fixed="right" label="操作" width="220">
      <template #default="scope">
        <el-button link type="primary" size="small">资源配额与限制</el-button
        ><el-divider direction="vertical" />
        <el-button
          link
          type="primary"
          size="small"
          @click="updateNamespace(scope.row)"
          >编辑</el-button
        ><el-divider direction="vertical" />
        <el-button
          :disabled="scope.row.metadata.name === 'default'"
          link
          type="danger"
          size="small"
          @click="deleteNamespace(scope.row)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
  <!-- 分页区域 -->
  <pagination
    :total="ns.total"
    @handlePageChange="handlePageChange"
  ></pagination>
  <NamespaceDialog
    :title="data.titile"
    v-model:visible="data.visible"
    :namespace="data.namespace"
    @value-change="ns.listNamespace()"
    v-if="data.visible"
  />
</template>

<script setup lang="ts">
import { nsStore } from '@/store/kubernetes/namespace'
import { reactive } from 'vue'
import pagination from '@/component/pagination/pagination.vue'
import { NamespaceData, Namespace } from '@/type/namespace'
import NamespaceDialog from './component/dialog.vue'
import { namespacerApi } from '../api'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { ElMessage, ElMessageBox } from 'element-plus'
import { PageInfo } from '@/type/sys'

const ns = nsStore()
const kube = kubeStore()
const data = reactive(new NamespaceData())
const handleSelectionChange = (value: Namespace[]) => {
  data.selectData = value
}

const createNamespace = () => {
  data.titile = '创建命名空间'
  data.visible = true
}

const updateNamespace = (namespace: Namespace) => {
  data.titile = '更新命名空间'
  data.namespace = namespace
  data.visible = true
}
const deleteNamespace = (namespace: Namespace) => {
  ElMessageBox.confirm(
    `此操作将删除[ ${namespace.metadata.name} ]命名空间 . 是否继续?`,
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
      draggable: true
    }
  )
    .then(() => {
      namespacerApi.delete
        .request({ cloud: kube.activeCluster, name: namespace.metadata.name })
        .then((res) => {
          ns.listNamespace()
          ElMessage.success(res.message)
        })
        .catch()
    })
    .catch() // 取消
}
const handlePageChange = (pageInfo: PageInfo) => {
  ns.query.page = pageInfo.page
  ns.query.limit = pageInfo.limit
  ns.listNamespace()
}
</script>

<style scoped lang="less">
.main {
  display: flex;
  margin-left: 15px;

  // .el-table {
  //   margin-top: 10px;
  // }
}
</style>
