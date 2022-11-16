/** * Created by lei on 2022/11/16 */
<template>
  <el-dialog v-model="dialogVisable" :close="handleClose">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">{{ title }}</h4>
        <el-divider />
      </div>
    </template>
    <div class="dialog-body">
      <el-radio-group v-model="radio1" class="ml-4">
        <!-- <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Card name</span>
            <el-button class="button" text>导入现有集群</el-button>
          </div>
        </template>
        <div v-for="o in 4" :key="o" class="text item">
          {{ "List item " + o }}
        </div>
        
      </el-card> -->
        <el-radio label="1" border class="r-1">
          <h1>asd</h1>
          <h1>asd</h1>
          <h1>asd</h1>
        </el-radio>
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span>Card name</span>
              <el-button class="button" text>新建集群</el-button>
            </div>
          </template>
          <div v-for="o in 4" :key="o" class="text item">
            {{ "List item " + o }}
          </div>
          <el-radio label="2" size="large">Option 2</el-radio>
        </el-card>
      </el-radio-group>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue";

const radio1 = ref("1");
const props = defineProps<{
  dialogVisable: Boolean;
  title: string;
}>();
const emits = defineEmits(["update:dialogVisable"]);

const handleClose = () => {
  emits("update:dialogVisable", false);
};
</script>

<style scoped lang="less">
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 360px;
  border-color: #409eff;
}
.ml-4 {
  display: flex;
  justify-content: space-evenly;
}
.r-1 {
  display: flex;
  height: 300px;
  width: 360px;
  text-align: start;
}
.el-radio__inner {
  margin-top: 0px;
}
</style>
