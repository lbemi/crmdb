/** * Created by lei on 2022/11/08 */

<template>
  <el-card class="box-card">
    <el-button
      v-auth="'asset:host:add'"
      type="primary"
      :icon="Edit"
      style="margin-bottom: 10px"
      @click="addHost"
      >添加主机</el-button
    >
    <el-table
      stripe
      :data="hostList"
      border
      style="width: 100%;"
      v-loading="loading"
      row-key="id"

    >
      <el-table-column header-align="center" align="center" prop="id" label="ID" min-width="40px" />
      <el-table-column header-align="center" align="center" prop="ip" label="主机IP" min-width="130px" />
      <el-table-column  header-align="center" align="center" prop="port" label="端口" min-width="80px" />
      <el-table-column header-align="center" align="center" prop="remark" label="描述" min-width="80px" />
      <el-table-column header-align="center" align="center" prop="label" label="标签" min-width="40px">
        <template #default="scope">
          <el-tooltip placement="top" effect="light">
            <template #content>
              <el-tag
                class="ml-2"
                type="info"
                effect="plain"
                v-for="item in labels(scope.row.label)"
                >{{ item }}</el-tag
              >
            </template>
            <SvgIcon iconName="icon-biaoqian" className="icon-1-4em" />
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column header-align="center" align="center" prop="username" label="用户名" min-width="80px" />
      <el-table-column header-align="center" align="center" prop="auth_method" label="认证方式" min-width="80px" />
      <el-table-column header-align="center" align="center" prop="created_at" label="创建时间" min-width="140px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column header-align="center" align="center" prop="status" label="状态" min-width="60px">
        <template #default="scope">
          <el-switch
            v-auth="'sys:role:status'"
            v-model="scope.row.status"
            class="ml-2"
            style="
              --el-switch-on-color: #409eff;
              --el-switch-off-color: #ff4949;
            "
            :active-value="1"
            :inactive-value="2"
            size="small"
            inline-prompt
            active-text="启用"
            inactive-text="禁用"
            width="45px"
            @click="changeStatus(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column header-align="center" align="center" prop="enable_ssh" label="SSH" min-width="60px">
        <template #default="scope">
          <el-switch
            v-auth="'sys:role:status'"
            v-model="scope.row.status"
            class="ml-2"
            style="
              --el-switch-on-color: #409eff;
              --el-switch-off-color: #ff4949;
            "
            :active-value="1"
            :inactive-value="2"
            size="small"
            inline-prompt
            active-text="启用"
            inactive-text="禁用"
            width="45px"
            @click="changeStatus(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column header-align="center" align="center" fixed="right" label="操作" min-width="300px">
        <template #default="scope">
          <el-button
            v-auth="'sys:role:edit'"
            type="primary"
            size="small"
            :icon="Edit"
            @click="handleEdit(scope.row)"
            >编辑</el-button
          >

          <el-button
            v-auth="'sys:role:set'"
            type="warning"
            size="small"
            :icon="Edit"
            @click="handleTerminal(scope.row)"
            >Terminal</el-button
          >
          <el-button
            v-auth="'sys:role:view'"
            type="success"
            size="small"
            :icon="View"
            @click="handleSetMenu(scope.row)"
            >查看</el-button
          >
          <el-button
            v-auth="'sys:role:del'"
            type="danger"
            size="small"
            :icon="Delete"
            @click="deleteHost(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页区域 -->
    <pagination
      :total="data.total"
      @handlePageChange="handlePageChange"
    ></pagination>
  </el-card>
  <HostDialogVue
    v-model:visible="hostAdd.visible"
    :title="hostAdd.title"
    @value-change="getHostList"
  />
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, onMounted } from "vue";
import { Delete, Edit, View } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { menuApi, roleApi } from "@/views/sys/api";
import { PageInfo, RoleInfo } from "@/type/sys";
import { HostInfo } from "@/type/host";
import { ElMessage, ElMessageBox } from "element-plus";
import { hostApi } from "../api";
import HostDialogVue from "./componet/hostDialog.vue";
import router from "@/router/index";

const labels = (str: string) => {
  return str.split(",");
};


const setMenu = reactive({
  visible: false,
  title: "分配权限",
  roleID: 0,
  menuList: [],
  defaultCheckedMenus: [] as Array<number>,
});

const loading = ref<boolean>(false);
const data = reactive({
  hostList: [] as Array<HostInfo>,
  total: 0,
});
const { hostList, total } = toRefs(data);

const hostAdd = reactive({
  visible: false,
  title: "添加主机",
});

const roleEdit = reactive({
  visible: false,
  title: "编辑主机",
  data: {} as HostInfo,
});

onMounted(() => {
  getHostList();
});

const query = reactive<PageInfo>({
  page: 1,
  limit: 10,
});

const getHostList = async () => {
  loading.value = true;
  const res = await hostApi.list.request(query);
  hostList.value = res.data.hosts;
  total.value = res.data.total;
  loading.value = false;
};

const addHost = () => {
  hostAdd.visible = true;
};

const handlePageChange = (pageInfo: PageInfo) => {
  query.page = pageInfo.page;
  query.limit = pageInfo.limit;
  getHostList();
};

const changeStatus = async (role: RoleInfo) => {
  await roleApi.changeStatus
    .request({
      id: role.id,
      status: role.status,
    })
    .then((res) => {
      ElMessage.success(res.message);
    })
    .catch(() => {
      role.status = 1;
    });
};

const deleteHost = (host: HostInfo) => {
  ElMessageBox.confirm(`此操作将删除[ ${host.ip} ]主机 . 是否继续?`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    draggable: true,
  })
    .then(() => {
      hostApi.delete
        .request({ id: host.id })
        .then((res) => {
          getHostList();
          ElMessage.success(res.message);
        })
        .catch();
    })
    .catch(() => {}); // 取消
};

const handleEdit = (host: HostInfo) => {
  roleEdit.data = JSON.parse(JSON.stringify(host));
  roleEdit.visible = true;
};

const handleSetMenu = async (host: HostInfo) => {
  setMenu.title = `为【${host.ip}】分配角色：`;
  setMenu.roleID = host.id;
  setMenu.defaultCheckedMenus = [];

  setMenu.visible = true;
};

const handleTerminal =(host: HostInfo) => {
  const terminalPage = router.resolve({
    path: '/termial',
    query: {
      id: host.id,
      ip: host.ip
    }
  })
  
  window.open(terminalPage.href,"_blank")
}
</script>

<style scoped lang="less"></style>
