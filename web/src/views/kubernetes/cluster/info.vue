/** * Created by lei on 2022/12/08 */
<template>
  <div>
    <table class="info">
      <tr>
        <td colspan="3" class="title">基本信息</td>
      </tr>
      <tr>
        <td>集群ID: {{ data.cluster.id }}</td>
        <td colspan="2">
          <div v-if="data.cluster.status == true">
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
              >运行中
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
              >故障
            </span>
          </div>
        </td>
      </tr>
    </table>
  </div>
  <div style="margin-top: 15px">
    <table class="info">
      <tr>
        <td colspan="3" class="title">基本信息</td>
      </tr>
      <tr>
        <td>API Server 内网连接端点</td>
        <td colspan="2">http://{{ data.cluster.internal_ip }}</td>
      </tr>
      <tr>
        <td>Pod 网络 CIDR</td>
        <td colspan="2">{{ data.cluster.pod_cidr }}</td>
      </tr>
      <tr>
        <td>Service CIDR</td>
        <td colspan="2">10.244.0.0/24</td>
      </tr>
      <tr>
        <td>运行时</td>
        <td colspan="2">{{ data.cluster.runtime }}</td>
      </tr>
      <tr>
        <td>节点 IP 数量</td>
        <td colspan="2">{{ data.cluster.nodes }}</td>
      </tr>
      <tr>
        <td>网络插件</td>
        <td colspan="2">Calico</td>
      </tr>
      <tr>
        <td>创建时间</td>
        <td colspan="2">{{ $filters.dateFormat(data.cluster.created_at) }}</td>
      </tr>
    </table>
  </div>
</template>

<script setup lang="ts">
import { clusterInfo } from '@/type/cluster'
import { onMounted, reactive } from 'vue'
import { clusterApi } from '@/views/container/api'
import { kubeStore } from '@/store/kubernetes/kubernetes'

const clu = kubeStore()
const data = reactive({
  cluster: {} as clusterInfo,
  query: {
    name: ''
  }
})

onMounted(() => {
  getCluster()
})
const getCluster = async () => {
  data.query.name = clu.activeCluster
  await clusterApi.get.request(data.query).then((res) => {
    data.cluster = res.data
  })
}
</script>

<style scoped lang="less">
table,
table tr th,
table tr td {
  table-layout: fixed;
  border: 1px solid #dbdbdbe5;
  height: 45px;
  width: 100%;
  font-size: 12px;
  color: rgb(145, 143, 143);
  padding-left: 16px;
}
table {
  table-layout: fixed;
  width: 100%;
  min-height: 25px;
  line-height: 25px;
  text-align: left;
  border-collapse: collapse;
}
.title {
  background-color: #f5f7fa;
  color: #333333;
  font-size: 14px;
}
.info {
  max-height: calc(100vh - 200px);
}
</style>
