/**
 * Created by lei on 2022/11/01
 */
<template>
  <div>
    <el-card>
      <el-dialog
    v-model="dialogVisible"
    title="Tips"
    width="30%"
    :before-close="handleClose"
  >
    <span>This is a message</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
    </el-card>
  </div>
</template>

<script setup lang="ts">
  
</script>

<style scoped lang="less">

</style>
