/** * Created by lei on 2022/09/25 */
<template>
  <el-card class="box-card">
    <el-button
      v-auth="'sys:user:add'"
      type="primary"
      :icon="Edit"
      style="margin-bottom: 10px"
      >添加用户</el-button
    >
    <el-table :data="userList" border style="width: 100%">
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
            v-auth="'sys:user:edit2'"
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
</template>

<script setup lang="ts">
import { reactive, toRefs, onMounted } from "vue";
import { Delete, Edit, Search, Share, Upload } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { useStore} from "@/store/usestore"
import { userApi } from "@/request/sys/user";
const store = useStore()
console.log("------",store.permissions);

const data = reactive({
  userList: [],
  total: 0,
});
const { userList, total } = toRefs(data);

onMounted(() => {
  getUserList();
});

const getUserList = async () => {
  const res = await userApi.listUser.request();
    userList.value = res.data.users;
    total.value = res.data.total;
};
</script>

<style scoped lang="less"></style>
