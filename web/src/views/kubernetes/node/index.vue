/** * Created by lei on 2022/12/11 */
<template>
  <el-table :data="data.nodes" style="width: 100%" align="center">
    <el-table-column label="名称/IP地址/UID" width="190">
      <template #default="scope">
        <div>
          <div>{{ scope.row.metadata.name }}</div>
          <el-button link type="primary">{{
            scope.row.status.addresses[0].address
          }}</el-button>
          <div>{{ scope.row.metadata.uid.split('-')[0] }}</div>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="角色" width="120" align="center">
      <template #default="scope">
        <div v-if="scope.row.metadata.labels['kubernetes.io/role']">Master</div>
        <div v-else>Worker</div>
        <div
          v-if="scope.row.status.conditions.slice(-1)[0]['status'] == 'True'"
          style="display: flex; align-items: center; justify-content: center"
        >
          <div
            style="
              width: 12px;
              height: 12px;
              background: #67c23a;
              border-radius: 50%;
            "
          ></div>
          <span style="margin-left: 5px; font-size: 12px; color: #67c23a"
            >运行中
          </span>
          <el-tooltip content="Right center" placement="right" effect="light">
            <el-icon size="17px" color="#909399" style="margin-left: 3px"
              ><InfoFilled
            /></el-icon>
          </el-tooltip>
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
          <el-tooltip content="Right center" placement="right" effect="light">
            <el-icon size="17px" color="#909399" style="margin-left: 3px"
              ><InfoFilled
            /></el-icon>
          </el-tooltip>
        </div>
        <div v-if="!scope.row.spec.unschedulable">可调度</div>
        <div v-else>不可调度</div>
      </template>
    </el-table-column>
    <el-table-column prop="state" label="State" width="120" />
    <el-table-column prop="city" label="City" width="120" />
    <el-table-column prop="address" label="Address" width="600" />
    <el-table-column prop="zip" label="Zip" width="120" />
    <el-table-column fixed="right" label="Operations" width="120">
      <template #default>
        <el-button link type="primary" size="small">Detail</el-button>
        <el-button link type="primary" size="small">Edit</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import { kubeStore } from '@/store/kubernetes/kubernetes'
import { reactive } from 'vue'
import { nodeApi } from '../api'
import { NodeData } from '@/type/node'
import { InfoFilled } from '@element-plus/icons-vue'
const kube = kubeStore()
const data = reactive(new NodeData())
nodeApi.list
  .request({ cloud: kube.activeCluster })
  .then((result) => {
    data.nodes = result.data.items
    console.log(data.nodes)
  })
  .catch((err) => {
    console.log(err)
  })
</script>

<style scoped lang="less">
.tooltip-base-box {
  width: 600px;
}
.tooltip-base-box .row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.tooltip-base-box .center {
  justify-content: center;
}
.tooltip-base-box .box-item {
  width: 110px;
  margin-top: 10px;
}
</style>
