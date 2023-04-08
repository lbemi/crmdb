<template>
	<div>
		<el-table :data="tableData" style="width: 100%">
			<el-table-column :prop="item.prop" :label="item.label" v-for="item in tableHeader" :key="item.prop">
				<template #default="scope">
					<div v-show="item.editable || scope.row.editable || tableData.length == 0" class="editable-row">
						<template v-if="item.type === 'input'">
							<el-input
								size="small"
								v-model="scope.row[item.prop]"
								:placeholder="`请输入${item.label}`"
								@change="handleEdit(scope.$index, scope.row)"
							/>
						</template>
					</div>
					<div v-show="(!item.editable && !scope.row.editable) || tableData.length == 0" class="editable-row">
						<span class="editable-row-span">{{ scope.row[item.prop] }}</span>
					</div>
				</template>
			</el-table-column>
			<el-table-column label="操作">
				<template #default="scope">
					<el-button v-show="!scope.row.editable" size="small" @click="scope.row.editable = true">编辑</el-button>
					<el-button v-show="tableData.length == 0" size="small" @click="append(scope.$index)">添加</el-button>
					<el-button v-show="scope.row.editable" size="small" type="success" @click="scope.row.editable = false">确定</el-button>
					<el-button size="small" type="danger" @click="handleDelete(scope.$index)">删除</el-button>
				</template>
			</el-table-column>
		</el-table>
	</div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
const item = {
	key: '',
	value: '',
};
const header = {
	prop: 'key',
	label: '自定义',
	editable: false,
	type: 'input',
};

const tableHeader = [
	{
		prop: 'key',
		label: 'KEY',
		editable: false,
		type: 'input',
	},
	{
		prop: 'value',
		label: 'VALUE',
		editable: false,
		type: 'input',
	},
];

const tableData = [
	{
		key: '张三',
		value: '上海市',
	},
	{
		key: '李四',
		value: '上海市',
	},
];

const handleEdit = (row, scope) => {
	row.editable = true;
};
const handleDelete = (index) => {
	tableData.splice(index, 1);
};
// prepend(index) {
//   item.editable = true
//   this.tableData.splice(index, 0, item)
// },
const append = (index) => {
	item.editable = true;
	tableData.splice(index + 1, 0, item);
};
const deleteCurrentColumn = (index) => {
	tableHeader.splice(index, 1);
};

// insertBefore(index) {
//   header.editable = true
//   this.tableHeader.splice(index, 0, header)
// },
// insertAfter(index) {
//   header.editable = true
//   this.tableHeader.splice(index + 1, 0, header)
// }
</script>

<style scoped>
.edit-icon {
	cursor: pointer;
}

.editable-row {
	display: flex;
	align-items: center;
}

.editable-row-span {
	display: inline-block;
	margin-right: 6px;
}

.menu-item {
	height: 32px;
	line-height: 32px;
	padding-left: 12px;
}

.menu-item:hover {
	color: #fff;
	background: #409eff;
	cursor: pointer;
}
</style>
