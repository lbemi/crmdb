/** * Created by lei on 2022/09/25 */
<template>
  <el-card class="box-card">
    <el-button
      v-auth="'sys:role:add'"
      type="primary"
      :icon="Edit"
      style="margin-bottom: 10px"
      @click="addRole"
      >添加角色</el-button
    >
    <el-table
      stripe
      :data="roleList"
      border
      style="width: 100%"
      v-loading="loading"
      row-key="id"
    >
      <el-table-column prop="id" label="ID" min-width="180px" />
      <el-table-column prop="name" label="名称" min-width="180px" />
      <el-table-column prop="memo" label="描述" min-width="180px" />
      <el-table-column prop="sequence" label="排序" min-width="180px" />
      <el-table-column prop="created_at" label="创建时间" min-width="240px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" min-width="160px">
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
      <el-table-column fixed="right" label="操作" min-width="300px">
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
            @click="handleSetMenu(scope.row)"
            >授权</el-button
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
            @click="deleteRole(scope.row)"
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
  <RoleDialog
    v-model:visible="roleAdd.visible"
    v-model:roleList="roleList"
    :title="roleAdd.title"
    @value-change="getRoleList"
    v-if="roleAdd.visible"
  />
  <RoleDialog
    v-model:visible="roleEdit.visible"
    v-model:roleList="roleList"
    :title="roleEdit.title"
    v-model:data="roleEdit.data"
    @value-change="getRoleList"
  />

  <RoleSetMenu
    v-model:visible="setMenu.visible"
    :defaultCheckedMenus="setMenu.defaultCheckedMenus"
    :title="setMenu.title"
    v-model:roleID="setMenu.roleID"
    v-model:menuList="setMenu.menuList"
    v-if="setMenu.visible"
  />
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, onMounted } from "vue";
import { Delete, Edit,View } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { menuApi, roleApi } from "@/views/sys/api";
import RoleDialog from "./componet/roleDialog.vue";
import { MenuInfo, PageInfo, RoleInfo } from "@/type/sys";
import { ElMessage, ElMessageBox } from "element-plus";
import RoleSetMenu from "./roleSetMenu.vue";
import { useStore } from "@/store/usestore";
const use = useStore();

const setMenu = reactive({
  visible: false,
  title: "分配权限",
  roleID: 0,
  menuList: [],
  defaultCheckedMenus: [] as Array<number>,
});

const loading = ref<boolean>(false);
const data = reactive({
  roleList: [] as Array<RoleInfo>,
  total: 0,
});
const { roleList, total } = toRefs(data);

const roleAdd = reactive({
  visible: false,
  title: "添加角色",
});

const roleEdit = reactive({
  visible: false,
  title: "编辑角色",
  data: {} as RoleInfo,
});

onMounted(() => {
  getRoleList();
  getMenuList();
});

const query = reactive<PageInfo>({
  page: 1,
  limit: 10,
});

const getRoleList = async () => {
  loading.value = true;
  const res = await roleApi.list.request(query);
  roleList.value = res.data.roles;
  total.value = res.data.total;
  loading.value = false;
};

const addRole = () => {
  roleAdd.visible = true;
};

const handlePageChange = (pageInfo: PageInfo) => {
  query.page = pageInfo.page;
  query.limit = pageInfo.limit;
  getRoleList();
};

const changeStatus = async (role: RoleInfo) => {
  await roleApi.changeStatus
    .request({
      id: role.id,
      status: role.status,
    })
    .then((res) => {
      getRoleList();

      ElMessage.success(res.message);
    })
    .catch(() => {
      role.status = 1;
    });
};

const deleteRole = (role: RoleInfo) => {
  ElMessageBox.confirm(`此操作将删除[ ${role.name} ]角色 . 是否继续?`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    draggable: true,
  })
    .then(() => {
      roleApi.delete
        .request({ id: role.id })
        .then((res) => {
          getRoleList();
          use.getUserPermissions();
          ElMessage.success(res.message);
        })
        .catch();
    })
    .catch(() => {}); // 取消
};

const handleEdit = (role: RoleInfo) => {
  roleEdit.data = JSON.parse(JSON.stringify(role));
  roleEdit.visible = true;
};

const getMenuList = async () => {
  await menuApi.list.request().then((res) => {
    setMenu.menuList = res.data.menus;
  });
};

const handleSetMenu = async (role: RoleInfo) => {
  setMenu.title = `为【${role.name}】分配角色：`;
  setMenu.roleID = role.id;
  setMenu.defaultCheckedMenus = [];

  await roleApi.getRoleMenus.request({ id: role.id }).then((res) => {
    const menuList: Array<MenuInfo> = res.data;
    if (menuList !== null) {
      for (let i = 0; i < menuList.length; i++) {
        if (menuList[i].children !== null) {
          for (let j = 0; j < menuList[i].children.length; j++) {
            setMenu.defaultCheckedMenus.push(menuList[i].children[j].id);
          }
        }
        setMenu.defaultCheckedMenus.push(menuList[i].id);
      }
    }
  });
  setMenu.visible = true;
};
</script>

<style scoped lang="less"></style>
