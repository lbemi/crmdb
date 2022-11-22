/** * Created by lei on 2022/09/25 */
<template>
  <el-card class="box-card">
    <el-button
      v-auth="'sys:menu:add'"
      type="primary"
      :icon="Edit"
      style="margin-bottom: 10px"
      @click="addMenu"
      >添加菜单权限</el-button
    >
    <el-table
      stripe
      :data="menuList"
      border
      style="width: 100%"
      v-loading="loading"
      row-key="id"
    >
      <el-table-column prop="id" label="菜单ID" width="180" />
      <el-table-column prop="name" label="菜单名称" width="120" />
      <el-table-column prop="menu_type" label="类型" width="80">
        <template #default="scope">
          <span v-if="scope.row.menu_type === 1" style="font-size: 13px">
            <el-tag class="ml-2" type="success" effect="dark">菜单</el-tag>
          </span>
          <span v-if="scope.row.menu_type === 2" style="font-size: 13px">
            <el-tag class="ml-2" type="warning" effect="dark">按钮</el-tag>
          </span>
          <span v-if="scope.row.menu_type === 3" style="font-size: 13px">
            <el-tag class="ml-2" effect="dark">API</el-tag>
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="memo" label="描述" width="120" />
      <el-table-column prop="url" label="URL" width="250" />
      <el-table-column prop="code" label="Code" width="130" />

      <el-table-column prop="method" label="请求方式" width="100">
        <template #default="scope">
          <span
            v-if="scope.row.method === 'GET'"
            style="font-size: 13px"
            text="bold"
          >
            <el-tag class="ml-2">{{ scope.row.method }}</el-tag>
          </span>
          <span v-if="scope.row.method === 'POST'" style="font-size: 13px">
            <el-tag class="ml-2" type="success">{{ scope.row.method }}</el-tag>
          </span>
          <span v-if="scope.row.method === 'DELETE'" style="font-size: 13px">
            <el-tag class="ml-2" type="danger">{{ scope.row.method }}</el-tag>
          </span>
          <span v-if="scope.row.method === 'PUT'" style="font-size: 13px">
            <el-tag class="ml-2" type="warning">{{ scope.row.method }}</el-tag>
          </span>
        </template>
      </el-table-column>

      <el-table-column prop="icon" label="图标" width="60"> 
        <template #default="scope">
          <SvgIcon :iconName="scope.row.icon" :color="'#79bbff'" />
        </template>
        
      </el-table-column>
      <el-table-column prop="status" label="状态" width="80">
        <template #default="scope">
          <el-switch
            v-auth="'sys:menu:status'"
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
            @change="changeStatus(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column prop="sequence" label="排序" min-width="80px" />
      <el-table-column prop="created_at" label="创建时间" min-width="240px">
        <template #default="scope">
          {{ $filters.dateFormat(scope.row.created_at) }}
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" min-width="300px">
        <template #default="scope">
          <el-button
            v-auth="'sys:menu:edit'"
            type="primary"
            size="small"
            :icon="Edit"
            @click="handleEdit(scope.row)"
            >编辑</el-button
          >
          <el-button
            v-auth="'sys:menu:del'"
            type="danger"
            size="small"
            :icon="Delete"
            @click="deleteMenu(scope.row)"
            >删除</el-button
          >
          <el-button
            v-auth="'sys:menu:set'"
            type="success"
            size="small"
            :icon="View"
            >查看</el-button
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
  <MenuDialog
    v-model:visible="menuAdd.visible"
    v-model:menuList="menuList"
    :title="menuAdd.title"
    @value-change="getMenuList"
    v-if="menuAdd.visible"
  />
  <MenuDialog
    v-model:visible="menuEdit.visible"
    v-model:menuList="menuList"
    :title="menuEdit.title"
    v-model:data="menuEdit.data"
    @value-change="getMenuList"
  />
</template>

<script setup lang="ts">
import { reactive, toRefs, ref, onMounted, onUpdated } from "vue";
import { Delete, Edit, View } from "@element-plus/icons-vue";
import pagination from "@/component/pagination/pagination.vue";
import { menuApi } from "@/views/sys/api";
import MenuDialog from "./componet/menuDialog.vue";
import { PageInfo, MenuInfo } from "@/type/sys";
import { ElMessage, ElMessageBox } from "element-plus";
import { useStore } from "@/store/usestore";
const user = useStore();
const loading = ref<boolean>(false);
const data = reactive({
  menuList: [] as Array<MenuInfo>,
  total: 0,
});
const { menuList, total } = toRefs(data);

const menuAdd = reactive({
  visible: false,
  title: "添加菜单权限",
});

const menuEdit = reactive({
  visible: false,
  title: "",
  data: {} as MenuInfo,
});

onMounted(() => {
  getMenuList();
});

const query = reactive<PageInfo>({
  page: 1,
  limit: 10,
});

const getMenuList = async () => {
  loading.value = true;
  const res = await menuApi.list.request(query);
  menuList.value = res.data.menus;
  total.value = res.data.total;
  loading.value = false;
};

const addMenu = () => {
  menuAdd.visible = true;
};

const handlePageChange = (pageInfo: PageInfo) => {
  query.page = pageInfo.page;
  query.limit = pageInfo.limit;
  getMenuList();
};

const changeStatus = async (menu: MenuInfo) => {
  await menuApi.changeStatus
    .request({
      id: menu.id,
      status: menu.status,
    })
    .then((res) => {
      ElMessage.success(res.message);
      getMenuList();
      user.getUserPermissions();
      user.getLeftMenus()
    })
    .catch(() => {
      menu.status = 1;
    });
};

const deleteMenu = (menu: MenuInfo) => {
  ElMessageBox.confirm(`此操作将删除[ ${menu.name} ]角色 . 是否继续?`, "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    draggable: true,
  })
    .then(() => {
      menuApi.delete
        .request({ id: menu.id })
        .then((res) => {
          getMenuList();
          user.getUserPermissions();
          user.getLeftMenus();
          ElMessage.success(res.message);
        })
        .catch();
    })
    .catch(() => {}); // 取消
};

const handleEdit = (menu: MenuInfo) => {
  menuEdit.title = `编辑 [${menu.name}] ：`;
  menuEdit.data = JSON.parse(JSON.stringify(menu));
  menuEdit.visible = true;
};

onUpdated(() => {});
</script>

<style scoped lang="less"></style>
