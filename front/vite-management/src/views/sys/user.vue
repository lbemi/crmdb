/**
 * Created by lei on 2022/09/25
 */
<template>
  <el-table :data="tableData" border >
    <el-table-column prop="id" label="ID" style="width: 100px;"  />
    <el-table-column prop="name" label="用户名" />
    <el-table-column prop="gmt_create" label="创建时间" />
    <el-table-column prop="gmt_modified" label="修改时间" />
    <el-table-column prop="status" label="状态" />
    <el-table-column label="操作" style="width: 800px;"> 
        <el-button type="primary">编辑</el-button>
      <el-button type="primary">删除</el-button>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import  {getUserListApi} from '@/request/api'
const Props = reactive({
  tableData: []
})
getUserListApi().then(res => {
  if (res.code === 200) {
    tableData.value = res.data
  }
  
})
const { tableData } = toRefs(Props)

</script>

<style scoped lang="less">

</style>
