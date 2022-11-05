/** * Created by lei on 2022/09/25 */
<template>
  <el-card class="box-card">
    <el-button
      v-auth="'sys:user:add'"
      type="primary"
      :icon="Edit"
      style="margin-bottom: 10px"
      @click="addUser"
      >添加用户</el-button
    >
    <el-table :data="userList" border style="width: 100%" v-loading="loading">
      <el-table-column prop="id" label="ID" min-width="180px" />
      <el-table-column prop="user_name" label="名称" min-width="180px" />
      <el-table-column prop="description" label="描述" min-width="180px" />
      <el-table-column prop="created_at" label="创建时间" min-width="240px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.created_at) }}
        </template>
      </el-table-column>

      <el-table-column prop="email" label="Email" min-width="160px" />
      <el-table-column prop="status" label="状态" min-width="160px">
        <template #default="scope">
          <el-switch
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
            v-auth="'sys:user:edit'"
            type="primary"
            size="small"
            :icon="Edit"
            @click="handleEdit(scope.row)"
            >编辑</el-button
          >
          <el-button
            v-auth="'sys:user:del'"
            type="danger"
            size="small"
            :icon="Delete"
            @click="deleteUser(scope.row)"
            >删除</el-button
          >
          <el-button
            v-auth="'sys:user:set'"
            type="primary"
            size="small"
            :icon="Edit"
            @click="handleSetRole(scope.row)"
            >分配角色</el-button
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
  <UserDialog
    v-model:visible="userAdd.visible"
    :title="userAdd.title"
    @value-change="getUserList"
    v-if="userAdd.visible"
  />
  <UserDialog
    v-model:visible="userEdit.visible"
    :title="userEdit.title"
    v-model:data="userEdit.data"
    @value-change="getUserList"
    v-if="userEdit.visible"
  />
  <UserSetRole
    v-model:visible="setRole.visible"
    :defaultCheckedRoles="setRole.defaultCheckedRoles"
    :title="setRole.title"
    v-model:userID="setRole.userID"
    v-model:roleList="setRole.roleList"
    v-if="setRole.visible"
  />
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, onMounted } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { userApi, roleApi } from "@/views/sys/api";
import UserDialog from "./componet/userDialog.vue";
import { PageInfo, UserInfo, UserForm, RoleInfo } from "@/type/sys";
import { ElMessage, ElMessageBox } from "element-plus";
import UserSetRole from "./userSetRole.vue";

const loading = ref<boolean>(false);
const data = reactive({
  userList: [] as Array<UserInfo>,
  total: 0,
});
const { userList, total } = toRefs(data);

const userAdd = reactive({
  visible: false,
  title: "添加用户",
});

const userEdit = reactive({
  visible: false,
  title: "编辑用户",
  data: {} as UserInfo,
});
const setRole = reactive({
  visible: false,
  title: "分配角色",
  userID: 0,
  roleList: [],
  defaultCheckedRoles: [] as Array<number>,
});
onMounted(() => {
  loading.value = true;
  getUserList();
  loading.value = false;
  getRoleList();
});

const query = reactive<PageInfo>({
  page: 1,
  limit: 10,
});

const getUserList = async () => {
  const res = await userApi.listUser.request(query);
  userList.value = res.data.users;
  total.value = res.data.total;
};
const getRoleList = async () => {
  const res = await roleApi.list.request();
  setRole.roleList = res.data.roles;
};
const addUser = () => {
  userAdd.visible = true;
};

const handlePageChange = (pageInfo: PageInfo) => {
  query.page = pageInfo.page;
  query.limit = pageInfo.limit;
  getUserList();
};

const handleSetRole = async (user: UserInfo) => {
  setRole.title = `为【${user.user_name}】分配角色：`;
  setRole.userID = user.id;
  setRole.defaultCheckedRoles = [];

  await userApi.listUserRole.request({ id: user.id }).then((res) => {
    const roleList: Array<RoleInfo> = res.data;
    if (roleList !== null) {
      for (let i = 0; i < roleList.length; i++) {
        if (roleList[i].children !== null) {
          for (let j = 0; j < roleList[i].children.length; j++) {
            setRole.defaultCheckedRoles.push(roleList[i].children[j].id);
          }
        }
        setRole.defaultCheckedRoles.push(roleList[i].id);
      }
    }
  });
  setRole.visible = true;
};
const changeStatus = async (user: UserInfo) => {
  await userApi.chageStaus
    .request({
      id: user.id,
      status: user.status,
    })
    .then((res) => {
      getUserList();
      ElMessage.success(res.message);
    })
    .catch(() => {
      user.status = 1;
    });
};

const deleteUser = (user: UserInfo) => {
  ElMessageBox.confirm(
    `此操作将删除[ ${user.user_name} ]用户 . 是否继续?`,
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      draggable: true,
    }
  )
    .then(() => {
      userApi.deleteUser
        .request({ id: user.id })
        .then((res) => {
          getUserList();
          ElMessage.success(res.message);
        })
        .catch();
    })
    .catch(() => {}); // 取消
};

const handleEdit = (user: UserInfo) => {
  userEdit.title = `编辑【${user.user_name}】角色`;
  userEdit.data = JSON.parse(JSON.stringify(user));
  userEdit.visible = true;
};
</script>

<style scoped lang="less"></style>
