
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="服务名" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="分类" prop="category">
            <el-select v-model="searchInfo.category" clearable placeholder="请选择" @clear="()=>{searchInfo.category=undefined}">
              <el-option v-for="(item,key) in mcp_typeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="包名" prop="package" width="120" />
          <el-table-column align="left" label="服务名" prop="name" width="120" />
          <el-table-column align="left" label="服务器描述" prop="description" width="120" />
        <el-table-column align="left" label="分类" prop="category" width="120">
            <template #default="scope">
                
                {{ filterDict(scope.row.category,mcp_typeOptions) }}
                
            </template>
        </el-table-column>
           <el-table-column label="标签(JSON数组)" prop="tags" width="200">
               <template #default="scope">
                  <ArrayCtrl v-model="scope.row.tags"/>
               </template>
           </el-table-column>
          <el-table-column align="left" label="状态" prop="status" width="120" />
           <el-table-column label="调用方法(JSON数组)" prop="callMethod" width="200">
               <template #default="scope">
                  <ArrayCtrl v-model="scope.row.callMethod"/>
               </template>
           </el-table-column>
          <el-table-column align="left" label="所有者" prop="owner" width="120" />
          <el-table-column align="left" label="在线使用次数" prop="onlineUsageCount" width="120" />
          <el-table-column align="left" label="离线使用次数" prop="offlineUsageCount" width="120" />
          <el-table-column align="left" label="点赞数" prop="likes" width="120" />
          <el-table-column align="left" label="GitHub链接" prop="github" width="120" />
          <el-table-column align="left" label="网站链接" prop="website" width="120" />
          <el-table-column align="left" label="详细说明" prop="detail" width="120" />
          <el-table-column align="left" label="logo图片路径" prop="logo" width="120" />
           <el-table-column label="工具" prop="tools" width="200">
               <template #default="scope">
                  <ArrayCtrl v-model="scope.row.tools"/>
               </template>
           </el-table-column>
          <el-table-column align="left" label="uuid" prop="uuid" width="120" />
          <el-table-column align="left" label="服务名称" prop="title" width="120" />
          <el-table-column align="left" label="头像地址" prop="avatarUrl" width="120" />
          <el-table-column align="left" label="作者" prop="authorName" width="120" />
          <el-table-column align="left" label="作者头像地址" prop="authorAvatarUrl" width="120" />
        <el-table-column align="left" label="特色" prop="isFeatured" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isFeatured) }}</template>
        </el-table-column>
          <el-table-column align="left" label="排序" prop="sort" width="120" />
          <el-table-column align="left" label="地址" prop="url" width="120" />
          <el-table-column align="left" label="类型" prop="type" width="120" />
          <el-table-column align="left" label="用户uuid" prop="userUuid" width="120" />
          <el-table-column align="left" label="sseUrl" prop="sseUrl" width="120" />
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateProjectsFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="包名:"  prop="package" >
              <el-input v-model="formData.package" :clearable="true"  placeholder="请输入包名" />
            </el-form-item>
            <el-form-item label="服务名:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入服务名" />
            </el-form-item>
            <el-form-item label="服务器描述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入服务器描述" />
            </el-form-item>
            <el-form-item label="分类:"  prop="category" >
              <el-select v-model="formData.category" placeholder="请选择分类" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in mcp_typeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="标签(JSON数组):"  prop="tags" >
              <ArrayCtrl v-model="formData.tags" editable/>
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-input v-model="formData.status" :clearable="true"  placeholder="请输入状态" />
            </el-form-item>
            <el-form-item label="调用方法(JSON数组):"  prop="callMethod" >
              <ArrayCtrl v-model="formData.callMethod" editable/>
            </el-form-item>
            <el-form-item label="所有者:"  prop="owner" >
              <el-input v-model="formData.owner" :clearable="true"  placeholder="请输入所有者" />
            </el-form-item>
            <el-form-item label="在线使用次数:"  prop="onlineUsageCount" >
              <el-input v-model.number="formData.onlineUsageCount" :clearable="true" placeholder="请输入在线使用次数" />
            </el-form-item>
            <el-form-item label="离线使用次数:"  prop="offlineUsageCount" >
              <el-input v-model.number="formData.offlineUsageCount" :clearable="true" placeholder="请输入离线使用次数" />
            </el-form-item>
            <el-form-item label="点赞数:"  prop="likes" >
              <el-input v-model.number="formData.likes" :clearable="true" placeholder="请输入点赞数" />
            </el-form-item>
            <el-form-item label="GitHub链接:"  prop="github" >
              <el-input v-model="formData.github" :clearable="true"  placeholder="请输入GitHub链接" />
            </el-form-item>
            <el-form-item label="网站链接:"  prop="website" >
              <el-input v-model="formData.website" :clearable="true"  placeholder="请输入网站链接" />
            </el-form-item>
            <el-form-item label="详细说明:"  prop="detail" >
              <el-input v-model="formData.detail" :clearable="true"  placeholder="请输入详细说明" />
            </el-form-item>
            <el-form-item label="logo图片路径:"  prop="logo" >
              <el-input v-model="formData.logo" :clearable="true"  placeholder="请输入logo图片路径" />
            </el-form-item>
            <el-form-item label="工具:"  prop="tools" >
              <ArrayCtrl v-model="formData.tools" editable/>
            </el-form-item>
            <el-form-item label="uuid:"  prop="uuid" >
              <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入uuid" />
            </el-form-item>
            <el-form-item label="服务名称:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入服务名称" />
            </el-form-item>
            <el-form-item label="头像地址:"  prop="avatarUrl" >
              <el-input v-model="formData.avatarUrl" :clearable="true"  placeholder="请输入头像地址" />
            </el-form-item>
            <el-form-item label="作者:"  prop="authorName" >
              <el-input v-model="formData.authorName" :clearable="true"  placeholder="请输入作者" />
            </el-form-item>
            <el-form-item label="作者头像地址:"  prop="authorAvatarUrl" >
              <el-input v-model="formData.authorAvatarUrl" :clearable="true"  placeholder="请输入作者头像地址" />
            </el-form-item>
            <el-form-item label="特色:"  prop="isFeatured" >
              <el-switch v-model="formData.isFeatured" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="排序:"  prop="sort" >
              <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序" />
            </el-form-item>
            <el-form-item label="地址:"  prop="url" >
              <el-input v-model="formData.url" :clearable="true"  placeholder="请输入地址" />
            </el-form-item>
            <el-form-item label="类型:"  prop="type" >
              <el-input v-model="formData.type" :clearable="true"  placeholder="请输入类型" />
            </el-form-item>
            <el-form-item label="用户uuid:"  prop="userUuid" >
              <el-input v-model="formData.userUuid" :clearable="true"  placeholder="请输入用户uuid" />
            </el-form-item>
            <el-form-item label="sseUrl:"  prop="sseUrl" >
              <el-input v-model="formData.sseUrl" :clearable="true"  placeholder="请输入sseUrl" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="包名">
                        {{ detailFrom.package }}
                    </el-descriptions-item>
                    <el-descriptions-item label="服务名">
                        {{ detailFrom.name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="服务器描述">
                        {{ detailFrom.description }}
                    </el-descriptions-item>
                    <el-descriptions-item label="分类">
                        
                        {{ filterDict(detailFrom.category,mcp_typeOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="标签(JSON数组)">
                            <ArrayCtrl v-model="detailFrom.tags"/>
                    </el-descriptions-item>
                    <el-descriptions-item label="状态">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item label="调用方法(JSON数组)">
                            <ArrayCtrl v-model="detailFrom.callMethod"/>
                    </el-descriptions-item>
                    <el-descriptions-item label="所有者">
                        {{ detailFrom.owner }}
                    </el-descriptions-item>
                    <el-descriptions-item label="在线使用次数">
                        {{ detailFrom.onlineUsageCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="离线使用次数">
                        {{ detailFrom.offlineUsageCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="点赞数">
                        {{ detailFrom.likes }}
                    </el-descriptions-item>
                    <el-descriptions-item label="GitHub链接">
                        {{ detailFrom.github }}
                    </el-descriptions-item>
                    <el-descriptions-item label="网站链接">
                        {{ detailFrom.website }}
                    </el-descriptions-item>
                    <el-descriptions-item label="详细说明">
                        {{ detailFrom.detail }}
                    </el-descriptions-item>
                    <el-descriptions-item label="logo图片路径">
                        {{ detailFrom.logo }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工具">
                            <ArrayCtrl v-model="detailFrom.tools"/>
                    </el-descriptions-item>
                    <el-descriptions-item label="uuid">
                        {{ detailFrom.uuid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="服务名称">
                        {{ detailFrom.title }}
                    </el-descriptions-item>
                    <el-descriptions-item label="头像地址">
                        {{ detailFrom.avatarUrl }}
                    </el-descriptions-item>
                    <el-descriptions-item label="作者">
                        {{ detailFrom.authorName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="作者头像地址">
                        {{ detailFrom.authorAvatarUrl }}
                    </el-descriptions-item>
                    <el-descriptions-item label="特色">
                        {{ detailFrom.isFeatured }}
                    </el-descriptions-item>
                    <el-descriptions-item label="排序">
                        {{ detailFrom.sort }}
                    </el-descriptions-item>
                    <el-descriptions-item label="地址">
                        {{ detailFrom.url }}
                    </el-descriptions-item>
                    <el-descriptions-item label="类型">
                        {{ detailFrom.type }}
                    </el-descriptions-item>
                    <el-descriptions-item label="用户uuid">
                        {{ detailFrom.userUuid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="sseUrl">
                        {{ detailFrom.sseUrl }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createProjects,
  deleteProjects,
  deleteProjectsByIds,
  updateProjects,
  findProjects,
  getProjectsList
} from '@/api/mcp/projects'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Projects'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.isFeatured === ""){
        searchInfo.value.isFeatured=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getProjectsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    mcp_typeOptions.value = await getDictFunc('mcp_type')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteProjectsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteProjectsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProjectsFunc = async(row) => {
    const res = await findProjects({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProjectsFunc = async (row) => {
    const res = await deleteProjects({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createProjects(formData.value)
                  break
                case 'update':
                  res = await updateProjects(formData.value)
                  break
                default:
                  res = await createProjects(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findProjects({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
