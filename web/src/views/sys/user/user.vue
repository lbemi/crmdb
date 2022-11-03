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
      <el-table-column prop="id" label="ID" width="180" />
      <el-table-column prop="user_name" label="名称" width="180" />
      <el-table-column prop="description" label="描述" width="180" />
      <el-table-column prop="created_at" label="创建时间" width="240" />
      <el-table-column prop="email" label="Email" width="160" />
      <el-table-column prop="status" label="状态" width="160">
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            class="ml-2"
            style="
              --el-switch-on-color: #409eff;
              --el-switch-off-color: #ff4949;
            "
            :active-value="1"
            :inactive-value="0"
            size="small"
            inline-prompt
            active-text="启用"
            inactive-text="禁用"
            width="45px"
          />
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template #default="scope">
          <el-button
            v-auth="'sys:user:edit'"
            type="primary"
            size="small"
            :icon="Edit"
            >编辑</el-button
          >
          <el-button
            v-auth="'sys:user:del'"
            type="danger"
            size="small"
            :icon="Delete"
            >删除</el-button
          >
          <el-button
            v-auth="'sys:user:set'"
            type="primary"
            size="small"
            :icon="Edit"
            >分配角色</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页区域 -->
    <pagination :total="data.total"></pagination>
  </el-card>
  <UserDialog v-model:visble="userFormVisible" @value-change="getUserList" />

</template>

<script setup lang="ts">
import { reactive, toRefs, ref, onMounted } from "vue";
import { Delete, Edit } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { userApi } from "@/views/sys/api";
import UserDialog from './componet/userDialog.vue'

const loading = ref(false);
const userFormVisible = ref(false)
const data = reactive({
  userList: [],
  total: 0,
});
const { userList, total } = toRefs(data);

onMounted(() => {
  loading.value = true;
  getUserList();
  loading.value = false;
});

const getUserList = async () => {
  const res = await userApi.listUser.request();
  userList.value = res.data.users;
  total.value = res.data.total;
};

const addUser =()=>{
  userFormVisible.value = true
}
</script>

<style scoped lang="less"></style>
