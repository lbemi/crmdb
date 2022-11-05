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
      @size-change="handlePageChange"
      @current-change="handlePageChange"
    />
  </div>
</template>

<script setup lang="ts">
import { PageInfo } from "@/type/sys";
import { reactive } from "vue";

defineProps({ total: Number });

const page = reactive<PageInfo>({
  page: 1,
  limit: 10,
});

const emit = defineEmits(["handlePageChange"]);
const handlePageChange = () => {
  emit("handlePageChange", page);
};

</script>

<style scoped></style>
