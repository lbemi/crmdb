<template>
  <el-breadcrumb :separator-icon="ArrowRight" class="bread">
    <el-breadcrumb-item
      v-for="item in routers"
      :key="item.path"
      :to="{ path: item.path }"
      >{{ item.name }}</el-breadcrumb-item
    >
  </el-breadcrumb>
</template>

<script lang="ts" setup>
import { ArrowRight } from '@element-plus/icons-vue'
import { computed } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const routers = computed(() => {
  return router.currentRoute.value.matched.filter((item) => item.name)
})
</script>

<style>
/* 这里为了跟收起图标对其 */
.bread {
  margin-bottom: 15px;
  margin-left: 15px;
}
</style>
