
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="包名:" prop="package">
          <el-input v-model="formData.package" :clearable="true"  placeholder="请输入包名" />
       </el-form-item>
        <el-form-item label="服务名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入服务名" />
       </el-form-item>
        <el-form-item label="服务器描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入服务器描述" />
       </el-form-item>
        <el-form-item label="分类:" prop="category">
           <el-select v-model="formData.category" placeholder="请选择分类" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mcp_typeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="标签(JSON数组):" prop="tags">
          <ArrayCtrl v-model="formData.tags" editable/>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入状态" />
       </el-form-item>
        <el-form-item label="调用方法(JSON数组):" prop="callMethod">
          <ArrayCtrl v-model="formData.callMethod" editable/>
       </el-form-item>
        <el-form-item label="所有者:" prop="owner">
          <el-input v-model="formData.owner" :clearable="true"  placeholder="请输入所有者" />
       </el-form-item>
        <el-form-item label="在线使用次数:" prop="onlineUsageCount">
          <el-input v-model.number="formData.onlineUsageCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="离线使用次数:" prop="offlineUsageCount">
          <el-input v-model.number="formData.offlineUsageCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="点赞数:" prop="likes">
          <el-input v-model.number="formData.likes" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="GitHub链接:" prop="github">
          <el-input v-model="formData.github" :clearable="true"  placeholder="请输入GitHub链接" />
       </el-form-item>
        <el-form-item label="网站链接:" prop="website">
          <el-input v-model="formData.website" :clearable="true"  placeholder="请输入网站链接" />
       </el-form-item>
        <el-form-item label="详细说明:" prop="detail">
          <el-input v-model="formData.detail" :clearable="true"  placeholder="请输入详细说明" />
       </el-form-item>
        <el-form-item label="logo图片路径:" prop="logo">
          <el-input v-model="formData.logo" :clearable="true"  placeholder="请输入logo图片路径" />
       </el-form-item>
        <el-form-item label="工具:" prop="tools">
          <ArrayCtrl v-model="formData.tools" editable/>
       </el-form-item>
        <el-form-item label="uuid:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入uuid" />
       </el-form-item>
        <el-form-item label="服务名称:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入服务名称" />
       </el-form-item>
        <el-form-item label="头像地址:" prop="avatarUrl">
          <el-input v-model="formData.avatarUrl" :clearable="true"  placeholder="请输入头像地址" />
       </el-form-item>
        <el-form-item label="作者:" prop="authorName">
          <el-input v-model="formData.authorName" :clearable="true"  placeholder="请输入作者" />
       </el-form-item>
        <el-form-item label="作者头像地址:" prop="authorAvatarUrl">
          <el-input v-model="formData.authorAvatarUrl" :clearable="true"  placeholder="请输入作者头像地址" />
       </el-form-item>
        <el-form-item label="特色:" prop="isFeatured">
          <el-switch v-model="formData.isFeatured" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true"  placeholder="请输入类型" />
       </el-form-item>
        <el-form-item label="用户uuid:" prop="userUuid">
          <el-input v-model="formData.userUuid" :clearable="true"  placeholder="请输入用户uuid" />
       </el-form-item>
        <el-form-item label="sseUrl:" prop="sseUrl">
          <el-input v-model="formData.sseUrl" :clearable="true"  placeholder="请输入sseUrl" />
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
  createMcpServer,
  updateMcpServer,
  findMcpServer
} from '@/api/mcp/mcpServer'

defineOptions({
    name: 'McpServerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const mcp_typeOptions = ref([])
const formData = ref({
            package: '',
            name: '',
            description: '',
            category: '',
            tags: [],
            status: '',
            callMethod: [],
            owner: '',
            onlineUsageCount: undefined,
            offlineUsageCount: undefined,
            likes: undefined,
            github: '',
            website: '',
            detail: '',
            logo: '',
            tools: [],
            uuid: '',
            title: '',
            avatarUrl: '',
            authorName: '',
            authorAvatarUrl: '',
            isFeatured: false,
            sort: undefined,
            url: '',
            type: '',
            userUuid: '',
            sseUrl: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMcpServer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mcp_typeOptions.value = await getDictFunc('mcp_type')
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
               res = await createMcpServer(formData.value)
               break
             case 'update':
               res = await updateMcpServer(formData.value)
               break
             default:
               res = await createMcpServer(formData.value)
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
