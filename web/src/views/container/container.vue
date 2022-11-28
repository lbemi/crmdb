/** * Created by lei on 2022/11/16 */
<template>
  <el-card style="height: 100%">
    <el-button type="primary" style="margin-bottom: 10px" @click="handleCreate"
      >创建集群</el-button
    >

    <el-table :data="data.clusters" style="width: 100%">
      <el-table-column fixed prop="id" label="ID" width="100" />
      <el-table-column prop="name" label="Name" width="120" />
      <el-table-column prop="status" label="状态" width="120">
        <template #default="scope">
          <div v-if="scope.row.status == true">
            <div
              style="
                display: inline-block;
                background: #67c23a;
                border-radius: 50%;
                width: 12px;
                height: 12px;
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
                background: #f56c6c;
                border-radius: 50%;
                width: 12px;
                height: 12px;
              "
            ></div>
            <span style="margin-left: 5px; font-size: 12px; color: #f56c6c"
              >故障
            </span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="nodes" label="节点数量" width="120" />
      <el-table-column prop="version" label="Version" width="120" />
      <el-table-column prop="runtime" label="运行时" width="160" />
      <el-table-column prop="pod_cidr" label="Pod_CIDR" width="120" />
      <el-table-column prop="service_cidr" label="Service_CIDR" width="120" />
      <el-table-column prop="cpu" label="CPU" width="120" />
      <el-table-column prop="memory" label="内存" width="160" >
        <template #default="scope">
            {{scope.row.memory/1024}}M
        </template>
      </el-table-column>
      <el-table-column  label="操作" width="160">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleClick"
            >Detail</el-button
          >
          <el-button link type="primary" size="small">Edit</el-button>
          <el-button link type="danger" size="small" @click="handlerDelete(scope.row)">Detele</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
  <CreateCluster
    :dialogVisable="createData.dialogVisable"
    @value-change="getCluster()"
    :title="createData.title"
  />
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import CreateCluster from "./compent/create.vue";
import { clusterApi } from "./api";
import { clusterInfo } from "@/type/container";
import { ElMessage, ElMessageBox } from "element-plus";

const createData = reactive({
  dialogVisable: false,
  title: "创建集群",
});

const handleClick = () => {
  console.log("click");
};

onMounted(() => {
  getCluster();
});

const handleCreate = () => {
  createData.dialogVisable = true;
};

const data = reactive({
  clusters: [] as Array<clusterInfo>,
});

const getCluster = async () => {
  const res = await clusterApi.list.request();
  data.clusters = res.data;
};


const handlerDelete = async (cluster:clusterInfo) =>{

    ElMessageBox.confirm(
    `此操作将删除[ ${cluster.name} ]集群 . 是否继续?`,
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      draggable: true,
    }
  )
    .then(() => {
      clusterApi.delete
        .request({ id: cluster.id })
        .then((res) => {
          getCluster();
          ElMessage.success(res.message);
        })
        .catch();
    })
    .catch(() => {}); // 取消
}
</script>

<style scoped lang="less"></style>
