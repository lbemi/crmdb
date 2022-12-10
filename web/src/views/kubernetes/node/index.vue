/** * Created by lei on 2022/12/10 */
<template>
  <div>
    <el-table :data="table.tableData" style="width: 100%; font-size: 12px">
      <el-table-column
        :prop="item.prop"
        :label="item.label"
        v-for="item in table.tableHeader"
        :key="item.prop"
      >
        <template #default="scope">
          <div v-show="item.editable || scope.row.editable">
            <template v-if="item.type === 'input'">
              <el-input
                style="width: 62px"
                size="small"
                v-model="scope.row[item.prop]"
                :placeholder="`${item.label}`"
                @change="handleEdit(scope.$index, scope.row)"
              />
            </template>
          </div>
          <div
            v-show="!item.editable && !scope.row.editable"
            class="editable-row"
          >
            <span class="editable-row-span">{{ scope.row[item.prop] }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        width="120px"
        style="display: flex; padding-left: 2px"
      >
        <template #default="scope">
          <el-button
            v-show="!scope.row.editable || !scope.row.addStatus"
            size="small"
            link
            type="primary"
            @click="scope.row.editable = true"
            >编辑</el-button
          >
          <el-button
            v-show="scope.row.addStatus"
            :disabled="scope.row.key === ''"
            size="small"
            link
            type="primary"
            @click="append(scope.$index)"
            >添加</el-button
          >
          <el-button
            v-show="scope.row.editable && !scope.row.addStatus"
            size="small"
            link
            type="primary"
            @click="add(scope.row)"
            >确定</el-button
          >
          <el-button
            v-show="!scope.row.addStatus && !scope.row.editable"
            size="small"
            type="danger"
            link
            @click="handleDelete(scope.$index)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue'

const item = reactive({
  key: '',
  editable: false,
  value: '',
  addStatus: false
})

interface label {
  key: string
  value: string
}
interface header {
  prop: string
  label: string
  editable: boolean
  type: string
}
const table = reactive({
  tableHeader: [
    {
      prop: 'key',
      label: '变量名称',
      editable: false,
      type: 'input'
    },
    {
      prop: 'value',
      label: '变量值',
      editable: false,
      type: 'input'
    }
  ],
  tableData: [] as label[]
})
onMounted(() => {
  item.editable = true
  item.addStatus = true
  table.tableData.splice(table.tableData.length, 0, item)
})
const handleEdit = (index: number, row: header) => {
  row.editable = true
}
const handleDelete = (index: number) => {
  table.tableData.splice(index, 1)
}
const append = (index: number) => {
  item.editable = false
  table.tableData.push({
    key: item.key,
    value: item.value
  })
  table.tableData.splice(index, 1)
  item.key = ''
  item.value = ''
  item.editable = true
  table.tableData.splice(table.tableData.length + 1, 0, item)
  console.log('+++++++', table.tableData)
}
const add = (header: header) => {
  header.editable = false
}
</script>

<style scoped lang="less">
.editable-row {
  margin-left: 5px;
}
.el-table__header .el-table__header-wrapper {
  font-size: 10px;
  background-color: red;
}
</style>
