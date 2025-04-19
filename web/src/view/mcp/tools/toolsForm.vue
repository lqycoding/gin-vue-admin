
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="服务id:" prop="serverId">
        <el-select  v-model="formData.serverId" placeholder="请选择服务id" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.serverId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="工具名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入工具名称" />
       </el-form-item>
        <el-form-item label="工具描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入工具描述" />
       </el-form-item>
        <el-form-item label="积分消耗:" prop="points">
          <el-input v-model.number="formData.points" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="输入参数模式(JSON):" prop="input_schema">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.input_schema 后端会按照json的类型进行存取
          {{ formData.input_schema }}
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getToolsDataSource,
  createTools,
  updateTools,
  findTools
} from '@/api/mcp/tools'

defineOptions({
    name: 'ToolsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            serverId: undefined,
            name: '',
            description: '',
            points: undefined,
            input_schema: {},
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getToolsDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTools({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createTools(formData.value)
               break
             case 'update':
               res = await updateTools(formData.value)
               break
             default:
               res = await createTools(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
