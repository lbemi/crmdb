/**
 * Created by lei on 2022/09/25
 */
 <template>
  <el-table :data="tableData" border style="width: 100%">
    <el-table-column prop="date" label="Date" width="180" />
    <el-table-column prop="name" label="Name" width="180" />
    <el-table-column prop="address" label="Address" />
  </el-table>
</template>

<script setup lang="ts">
import { reactive, toRefs } from 'vue';

  const Props = reactive({
    tableData: []
  })
  const {tableData} = toRefs(Props)
</script>

<style scoped lang="less">

</style>
