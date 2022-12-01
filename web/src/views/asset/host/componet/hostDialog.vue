/** * Created by lei on 2022/11/03 */
<template>
  <el-dialog
    v-model="visible"
    @close="handleClose(hostFormRef)"
    :title="title"
    style="width: 400px;"
  >
    <el-form
      label-width="100px"
      ref="hostFormRef"
      :rules="hsotFormRules"
      :model="form"
      style="max-width: 300px;"
    >
      <el-form-item label="主机IP:" prop="ip">
        <el-input v-model="form.ip" />
      </el-form-item>
      <el-form-item label="端口:" prop="port">
        <el-input v-model="form.port" />
      </el-form-item>
      <el-form-item label="用户名:" prop="username">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="认证方式:">
        <el-select
          ref="selectRef"
          v-model="form.auth_method"
          clearable
          placeholder="请选择"
          value-key="id"
        >
          <el-option
            v-for="item in authOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="密码:" prop="password" v-if="form.auth_method == 1">
        <el-input v-model="form.password" type="password" />
      </el-form-item>
      <el-form-item label="密钥:" prop="secret" v-if="form.auth_method == 2">
        <el-input
          v-model="form.secret"
          autosize
          type="textarea"
          placeholder="输入或上传密钥"
        />
      </el-form-item>
      <el-form-item label="描述:" prop="remark">
        <el-input v-model="form.remark" />
      </el-form-item>
      <el-form-item label="标签:" prop="label">
        <el-tag
          v-for="tag in dynamicTags"
          :key="tag"
          class="mx-1"
          closable
          :disable-transitions="false"
          @close="handleTagClose(tag)"
        >
          {{ tag }}
        </el-tag>
        <el-input
          v-if="inputVisible"
          ref="InputRef"
          v-model="inputValue"
          class="ml-1 w-20"
          size="small"
          @keyup.enter="handleInputConfirm"
          @blur="handleInputConfirm"
        />
        <el-button
          v-else
          class="button-new-tag ml-1"
          size="small"
          @click="showInput"
        >
          + 添加标签
        </el-button>
      </el-form-item>
      <div style="display: flex;">
        <el-form-item label="状态:">
          <el-switch
            v-model="form.status"
            class="ml-2"
            style="

              --el-switch-on-color: #409eff;
              --el-switch-off-color: #ff4949;
"
            :active-value="1"
            :inactive-value="2"
            size="large"
            inline-prompt
            active-text="启用"
            inactive-text="禁用"
          />
        </el-form-item>
        <el-form-item label="SSH:">
          <el-switch
            v-model="form.enable_ssh"
            class="ml-2"
            style="

              --el-switch-on-color: #409eff;
              --el-switch-off-color: #ff4949;
"
            :active-value="1"
            :inactive-value="2"
            size="large"
            inline-prompt
            active-text="启用"
            inactive-text="禁用"
          />
        </el-form-item>
      </div>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose(hostFormRef)">取消</el-button>
        <el-button type="primary" @click="btnOk(hostFormRef)"> 确定 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { HostForm, HostInfo } from '@/type/host'
import { FormInstance, FormRules, ElMessage } from 'element-plus'
import { ref, reactive, watch, nextTick } from 'vue'
import { hostApi } from '../../api'
import { useStore } from '@/store/usestore'
import { ElInput } from 'element-plus'

const use = useStore()
const hostFormRef = ref<FormInstance>()

const authOptions = [
  {
    value: 1,
    label: '密码认证'
  },
  {
    value: 2,
    label: '密钥认证'
  }
]
const props = defineProps<{
  visible: boolean
  title: string
  data?: HostInfo | undefined
}>()

const emits = defineEmits(['update:visible', 'valueChange'])

const handleClose = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  emits('update:visible', false)
}

let form = reactive<HostForm>({
  label: '',
  remark: '',
  ip: '',
  port: 22,
  username: '',
  auth_method: 1,
  status: 1,
  enable_ssh: 1
})

const inputValue = ref('')
const dynamicTags = ref<Array<string>>([])
const inputVisible = ref<boolean>(false)
const InputRef = ref<InstanceType<typeof ElInput>>()

const handleTagClose = (tag: string) => {
  dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const showInput = () => {
  inputVisible.value = true
  nextTick(() => {
    InputRef.value!.input!.focus()
  })
}

const handleInputConfirm = () => {
  if (inputValue.value) {
    dynamicTags.value.push(inputValue.value)
  }
  inputVisible.value = false
  inputValue.value = ''
}

const hsotFormRules = reactive<FormRules>({
  ip: [{ required: true, message: '请输入IP', trigger: 'blur' }]
})

const btnOk = async (formEl: FormInstance | undefined) => {
  form.label = dynamicTags.value.toString()
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (!props.data) {
        hostApi.add.request(form).then((res) => {
          handleClose(hostFormRef.value)
          emits('valueChange')
          ElMessage.success(res.message)
        })
      } else {
        hostApi.update.request({ id: props.data.id }, form).then((res) => {
          handleClose(hostFormRef.value)
          emits('valueChange')
          ElMessage.success(res.message)
          use.getLeftMenus()
          use.getUserPermissions()
        })
      }
    } else {
      ElMessage.error('请正确填写')
    }
  })
}

watch(
  () => props.data,
  () => {
    form.remark = props.data!.remark
    form.status = props.data!.status
    form.ip = props.data!.ip
    form.username = props.data!.username
    form.port = props.data!.port
    form.auth_method = props.data!.auth_method
    form.ip = props.data!.ip
  }
)
</script>

<style scoped lang="less"></style>
