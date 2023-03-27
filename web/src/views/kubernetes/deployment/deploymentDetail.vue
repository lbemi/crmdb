<template>
  <el-descriptions
    :title="depStore.activeDeployment?.metadata.name"
    :column="3"
    border
  >
    <el-descriptions-item
      label="名称"
      label-align="right"
      align="center"
      label-class-name="my-label"
      class-name="my-content"
      width="150px"
      >{{ depStore.activeDeployment?.metadata.name }}</el-descriptions-item
    >
    <el-descriptions-item label="命名空间" label-align="right" align="center">{{
      depStore.activeDeployment?.metadata.namespace
    }}</el-descriptions-item>
    <el-descriptions-item label="副本数" label-align="right" align="center">{{
      depStore.activeDeployment?.spec.replicas
    }}</el-descriptions-item>
    <el-descriptions-item label="创建时间" label-align="right" align="center">{{
      $filters.dateFormat(depStore.activeDeployment?.metadata.creationTimestamp)
    }}</el-descriptions-item>
    <el-descriptions-item label="选择器" label-align="right" align="center">
      <div class="tag-center">
        <el-tag
          effect="plain"
          round
          v-for="(item, key, index) in depStore.activeDeployment?.spec.selector
            .matchLabels"
          :key="index"
        >
          {{ key }}:{{ item }}
        </el-tag>
      </div>
    </el-descriptions-item>
    <el-descriptions-item label="镜像" label-align="right" align="center">
      <div class="tag-center">
        <el-tag
          round
          effect="plain"
          v-for="(item, index) in depStore.activeDeployment?.spec.template.spec
            .containers"
          :key="index"
          >{{ item.image.split('@')[0] }}</el-tag
        >
      </div>
    </el-descriptions-item>
    <el-descriptions-item label="注解" label-align="right" align="center">
      <div class="tag-center">
        <el-tag
          effect="plain"
          type="info"
          v-for="(item, key, index) in depStore.activeDeployment?.metadata
            .annotations"
          :key="index"
        >
          {{ key }}:{{ item }}</el-tag
        >
      </div>
    </el-descriptions-item>
    <el-descriptions-item
      label="滚动升级策略"
      label-align="right"
      align="center"
    >
      <div>
        超过期望的Pod数量：
        {{ depStore.activeDeployment?.spec.strategy.rollingUpdate.maxSurge }}
      </div>
      <div>
        不可用Pod最大数量：
        {{
          depStore.activeDeployment?.spec.strategy.rollingUpdate.maxUnavailable
        }}
      </div>
    </el-descriptions-item>
    <el-descriptions-item label="策略" label-align="right" align="center">{{
      depStore.activeDeployment?.spec.strategy.type
    }}</el-descriptions-item>
    <el-descriptions-item label="状态" label-align="right" align="center">
      就绪：<a v-if="depStore.activeDeployment?.status.readyReplicas">{{
        depStore.activeDeployment?.status.readyReplicas
      }}</a>
      <a style="color: red" v-else>0</a> /{{
        depStore.activeDeployment?.status.replicas
      }}
      个，已更新：{{ depStore.activeDeployment?.status.updatedReplicas }}
      个，可用：
      <a v-if="depStore.activeDeployment?.status.readyReplicas">{{
        depStore.activeDeployment?.status.readyReplicas
      }}</a>
      <a style="color: red" v-else>0</a>

      个
      <el-link
        type="primary"
        :underline="false"
        @click="iShow = !iShow"
        style="font-size: 10px; margin-left: 5px"
        >展开现状现状详情<el-icon> <CaretBottom /> </el-icon
      ></el-link>
    </el-descriptions-item>
  </el-descriptions>

  <div iShow="iShow">
    <el-divider />
    <el-table
      :data="depStore.activeDeployment?.status.conditions"
      stripe
      style="width: 100%"
    >
      <el-table-column prop="type" label="类型" />
      <el-table-column prop="status" label="状态" />
      <el-table-column prop="lastUpdateTime" label="更新时间">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.lastUpdateTime) }}
        </template>
      </el-table-column>
      <el-table-column prop="reason" label="内容" />
      <el-table-column prop="message" label="消息" />
    </el-table>
  </div>

  <!-- <el-divider /> -->
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
    <el-tab-pane label="容器组" name="first">
      <el-table :data="pods" stripe style="width: 100%">
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
            <span
              v-if="scope.row.status.phase == 'Running'"
              style="color: green"
            >
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
        <el-table-column label="标签" width="280px">
          <template #default="scope">
            <el-tooltip placement="top" effect="light">
              <template #content>
                <div style="display: flex; flex-direction: column">
                  <el-tag
                    class="label"
                    type="info"
                    v-for="(item, key, index) in scope.row.metadata.labels"
                    :key="index"
                  >
                    {{ key }}:{{ item }}
                  </el-tag>
                </div>
              </template>
              <el-tag
                type="info"
                v-for="(item, key, index) in scope.row.metadata.labels"
                :key="index"
              >
                <div>{{ key }}:{{ item }}</div>
              </el-tag>
            </el-tooltip>
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
        <el-table-column fixed="right" label="操作" width="160">
          <template #default="scope">
            <el-button link type="primary" size="small" @click="handleClick"
              >详情</el-button
            ><el-divider direction="vertical" />
            <el-button link type="primary" size="small">编辑</el-button
            ><el-divider direction="vertical" />
            <el-button
              link
              type="primary"
              size="small"
              @click="deletePod(scope.row)"
              >删除</el-button
            >
            <el-button link type="primary" size="small">终端</el-button
            ><el-divider direction="vertical" />
            <el-button link type="primary" size="small">日志</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="Config" name="second">Config</el-tab-pane>
    <el-tab-pane label="Role" name="third">Role</el-tab-pane>
    <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
  </el-tabs>
</template>
<script lang="ts" setup>
import { deployStore } from '@/store/kubernetes/deployment'
import { Data } from '@/type/deployment'
import { reactive, onMounted, ref, onBeforeUnmount } from 'vue'
import { deploymentApi } from '../api'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { Pod } from '@/type/pod'
import type { TabsPaneContext } from 'element-plus'
import { CaretBottom } from '@element-plus/icons-vue'
import { podStore } from '@/store/kubernetes/pods'
import { ElMessage, ElMessageBox } from 'element-plus'
import { webSocketURL } from '@/request/request'

const podFunc = podStore()
const iShow = ref(false)
const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
const cluster = kubeStore()
const depStore = deployStore()

const data = reactive(new Data())
const timer = ref()
const pods = ref<Pod[]>()
onMounted(() => {
  getPods()
  buildwebsocket()

  timer.value = window.setInterval(() => {
    getPods()
  }, 5000)
  onBeforeUnmount(() => {
    window.clearInterval(timer.value)
  })
})

const getPods = async () => {
  data.query.deploymentName = depStore.activeDeployment?.metadata.name
  data.query.namespace = depStore.activeDeployment!.metadata.namespace
  data.query.cloud = cluster.activeCluster
  const res = await deploymentApi.detail.request(data.query)
  pods.value = res.data
}

const deletePod = async (pod: Pod) => {
  ElMessageBox.confirm(
    `此操作将删除[ ${pod.metadata.name} ] 容器 . 是否继续?`,
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )
    .then(() => {
      podFunc.deletPod(pod.metadata.namespace, pod.metadata.name)
      getPods()
      ElMessage({
        type: 'success',
        message: `${pod.metadata.name}` + ' 已删除'
      })
    })
    .catch() // 取消
}

const buildwebsocket = async () => {
  var dns = webSocketURL + cluster.activeCluster + '/deployment'
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
        object.result.namespace ===
          depStore.activeDeployment?.metadata.namespace &&
        object.cluster == cluster.activeCluster
      ) {
        data.Deployments = object.result.data
        data.Deployments.forEach((item) => {
          if (item.metadata.name == depStore.activeDeployment?.metadata.name) {
            depStore.activeDeployment = item
            return
          }
        })
      }
    }
  }

  ws.onclose = () => {
    console.log('close')
  }
}
</script>
<style lang="less">
.tag-center {
  display: flex;
  flex-direction: column;
  align-items: center;

  .el-tag {
    margin-bottom: 3px;
    margin-bottom: 3px;
    white-space: normal;
    height: auto;
  }
}
</style>
