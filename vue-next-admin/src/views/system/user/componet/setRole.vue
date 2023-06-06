<template>
	<el-dialog v-model="data.visible" @close="handleClose()" :title="title" style="width: 400px">
		角色：
		<el-select v-model="data.defaultCheckedRoles" multiple placeholder="Select" style="width: 240px" size="small">
			<el-option v-for="item in data.roleList" :key="item.id" :label="item.name" :value="item.id" />
		</el-select>
		<!-- <el-tree ref="menusRef" node-key="id" :data="data.roleList" :default-checked-keys="data.defaultCheckedRoles" check-strictly show-checkbox>
			<template #default="{ data: { name } }"> {{ name }}</template>
		</el-tree> -->

		<template #footer>
			<span class="dialog-footer">
				<el-button size="small" @click="handleClose()">取消</el-button>
				<el-button size="small" type="primary" @click="btnOk()"> 确定 </el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus';
import { ref, reactive, watch, onMounted } from 'vue';
import { RoleType } from '/@/types/views';
import { useUserApi } from '/@/api/system/user';
import { useRoleApi } from '/@/api/system/role';

const menusRef = ref();
const userApi = useUserApi();
const roleApi = useRoleApi();
const props = defineProps<{
	visible: boolean;
	title: string;
	userID: number;
	roleList: Array<RoleType>;
	defaultCheckedRoles: Array<number>;
}>();

const data = reactive({
	visible: false,
	roleForm: {
		role_ids: [] as Array<number>,
	},
	userID: 0,
	roleList: [] as Array<RoleType>,
	role: {} as RoleType,
	defaultCheckedRoles: [] as Array<number>,
});

const emits = defineEmits(['update:visible', 'valueChange']);
const handleClose = () => {
	emits('update:visible', false);
};

const btnOk = async () => {
	console.log('----', data.defaultCheckedRoles);

	// const roleIds = menusRef.value.getCheckedKeys();
	data.roleForm.role_ids = data.defaultCheckedRoles;
	userApi
		.setUserRole(data.userID, data.roleForm)
		.then((res) => {
			handleClose();
			emits('valueChange');
			ElMessage.success(res.message);
		})
		.catch((e) => {
			ElMessage.error(e.message);
		});
};
const getRoleList = async () => {
	await roleApi.listRole({ page: 0, limit: 0 }).then((res: any) => {
		data.roleList = res.data.data;
	});
};
onMounted(() => {
	getRoleList();
});

watch(
	() => props,
	() => {
		data.visible = props.visible;
		data.defaultCheckedRoles = props.defaultCheckedRoles;
		data.userID = props.userID;
	},
	{
		immediate: true,
	}
);
</script>

<style scoped lang="less"></style>
