/** * Created by lei on 2022/11/29 */
<template>
  <el-card>
    <el-container>
      <el-aside width="200px" style="height: 100%;">
        <el-menu
          active-text-color="#409EFF"
          :unique-opened="true"
          :router="true"
          @open="handleOpen"
          @close="handleClose"
        >
          <el-select
            v-model.number="activeCluster!.name"
            class="m-2"
            placeholder="Select"
            @change="flush"
          >
            <el-option
              v-for="item in clusters"
              :key="item.name"
              :label="item.name + ' - 集群'"
              :value="item.name"
              style="align-items: center;"
            />
          </el-select>
          <hr />
          <template v-for="menu in kubernetesRoutes">
            <el-menu-item v-if="!menu.children" :index="menu.path">
              <template #title>{{ menu.name }}</template>
            </el-menu-item>
            <el-sub-menu v-else :index="menu.id + ''" :key="menu.path">
              <template #title>
                <span>{{ menu.name }}</span>
              </template>
              <el-menu-item
                :index="child.path"
                v-for="child in menu.children"
                :key="child.id"
              >
                {{ child.name }}
              </el-menu-item>
            </el-sub-menu>
          </template>
        </el-menu>
      </el-aside>
      <router-view></router-view>
    </el-container>
  </el-card>
</template>

<script setup lang="ts">
import { reactive, inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { namespacerApi } from './api'
import { Namespace } from '@/type/namespace'
const kube = kubeStore()

const activeCluster = kube.activeCluster
const clusters = kube.clusters

const handleOpen = (key: string, keyPath: string[]) => {}
const handleClose = (key: string, keyPath: string[]) => {}

const flush = inject('reload')
const kubernetesRoutes = [
  {
    id: 1,
    name: '集群信息',
    path: '/cluster'
  },
  {
    id: 2,
    name: '命名空间',
    path: '/namespace'
  },
  {
    id: 3,
    name: '节点管理',
    path: '/nodes',
    children: [
      {
        id: 3.1,
        name: '节点信息',
        path: '/node'
      }
    ]
  },
  {
    id: 5,
    name: '工作负载',
    path: '/deployment',
    children: [
      {
        id: 5.1,
        name: 'Deployments',
        path: '/deployment'
      },
      {
        id: 5.2,
        name: 'Stateful Sets',
        path: '/statefulset'
      },
      {
        id: 5.3,
        name: 'Daemon Sets',
        path: '/daemonset'
      }
    ]
  }
]

const query = reactive({
  cloud: ''
})

const state = reactive({
  namespaces: <Namespace[]>[]
})
onMounted(() => {
  getNamespace()
})

const getNamespace = async () => {
  query.cloud = kube.activeCluster!.name
  const res = await namespacerApi.list.request(query)
  state.namespaces = res.data.items
  console.log(state.namespaces)
}
</script>

<style scoped lang="less"></style>
