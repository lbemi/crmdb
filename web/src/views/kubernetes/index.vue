/** * Created by lei on 2022/11/29 */
<template>
  <el-card class="body">
    <el-container>
      <el-aside class="in-asibe">
        <el-menu
          active-text-color="#409EFF"
          :unique-opened="true"
          :router="true"
          class="in-menu"
        >
          <div class="cluster-info">
            <el-select
              v-model="kube.activeCluster"
              class="m-2"
              placeholder="Select"
              @change="flush"
            >
              <el-option
                v-for="item in clusters"
                :key="item.name"
                :label="item.name + ' - 集群'"
                :value="item.name"
                :disabled="!item.status"
                style="align-items: center"
              />
            </el-select>
            <el-divider />
          </div>

          <template v-for="menu in kubernetesRoutes">
            <el-menu-item
              v-if="!menu.children"
              :index="menu.path"
              :key="menu.id"
            >
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
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-card>
</template>

<script setup lang="ts">
import { inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { nsStore } from '@/store/kubernetes/namespace'

const namespace = nsStore()
const kube = kubeStore()

const clusters = kube.clusters

const flush = inject('reload')
namespace.activeNamespace = 'default'

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

onMounted(() => {
  getNamespace()
})

const getNamespace = () => {
  namespace.listNamespace()
}
</script>

<style scoped lang="less">
.in-asibe {
  display: flex;
  align-items: center;
  justify-items: center;
  width: auto !important;
  text-align: center;
  flex-direction: column;
  height: 100%;
  padding: 0;
  .in-menu:not(.el-menu--collapse) {
    width: 200px;
    height: calc(100vh - 170px);
  }
}
.cluster-info {
  display: flex;
  align-items: center;
  flex-direction: column;
  margin-top: 0px;
  .el-divider--horizontal {
    margin-top: 1px;
    margin-bottom: 2px;
    width: 140px;
  }
  .m-2 {
    margin-top: 0px;
  }
}
.el-card {
  --el-card-padding: 0px;
  .el-container {
    padding-top: 20px;
    padding-right: 20px;
    padding-bottom: 20px;
  }
}
.el-main {
  padding-top: 0px;
}
</style>
