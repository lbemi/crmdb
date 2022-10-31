<template>
  <div class="root">
    <!-- 分页区域 -->
    <el-pagination
      style="
        float: left;
        margin-right: 40px;
        margin-top: 20px;
        margin-bottom: 20px;
      "
      v-model:current-page="page.page"
      v-model:page-size="page.limit"
      :page-sizes="[10, 20, 50]"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="onChange"
      @current-change="onChange"
    />
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
defineProps({ total: Number });
const page = reactive({
  page: 1,
  limit: 10,
});
const emit = defineEmits(["onChange"]);
const onChange = () => {
  emit("onChange", page);
};
// 初始化回调,页面进入的时候需要回调一次！
onChange();
</script>

<style scoped></style>
