<template>
  <div class="user-page">
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="page-title">人员档案 // USERS</h1>
          <p class="page-subtitle">机构成员信息管理</p>
        </div>
        <div class="actions">
          <a-input-search
            v-model:value="searchText"
            placeholder="搜索特工ID/昵称..."
            style="width: 300px"
            @search="onSearch"
            class="custom-search"
          />
          <a-button type="primary" class="action-btn create-btn" @click="showModal('create')">
            <template #icon><UserAddOutlined /></template>
            注册新成员
          </a-button>
        </div>
      </div>
    </div>

    <div class="table-container">
      <a-table 
        :columns="userType === 'vip' ? admincolumns : columns"
        :data-source="users" 
        :loading="loading"
        :pagination="pagination"
        @change="handleTableChange"
        rowKey="id"
        class="custom-table"
      >
        <template #bodyCell="{ column, record }">
          <template v-if="column.key === 'gender'">
            <span>{{ formatGender(record.gender) }}</span>
          </template>
          
          <template v-else-if="column.key === 'userType'">
            <span class="user-type-text">
              {{ record.userType || 'normal' }}
            </span>
          </template>

          <template v-else-if="column.key === 'createdAt'">
            <span>{{ formatTime(record.createdAt) }}</span>
          </template>

          <template v-else-if="column.key === 'action'">
            <span class="action-links">
              <a @click="showModal('edit', record)">编辑</a>
              <a-divider type="vertical" />
              <a-popconfirm
                title="确认移除该成员档案？"
                @confirm="handleDelete(record)"
                ok-text="确认"
                cancel-text="取消"
              >
                <a class="danger-link">移除</a>
              </a-popconfirm>
            </span>
          </template>
        </template>
      </a-table>

    </div>

    <!-- 用户编辑弹窗 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalType === 'create' ? '注册新成员' : '更新档案'"
      @ok="handleModalOk"
      :confirmLoading="modalLoading"
      class="custom-modal"
    >
      <a-form :model="form" layout="vertical" ref="formRef">
        <a-form-item label="特工ID (Account)" name="account" required>
          <a-input v-model:value="form.account" :disabled="modalType === 'edit'" />
        </a-form-item>
        <a-form-item 
          label="密钥 (Password)" 
          name="password" 
          :required="modalType === 'create'"
        >
          <a-input-password v-model:value="form.password" :placeholder="modalType === 'edit' ? '留空则不修改' : ''" />
        </a-form-item>
        <a-form-item label="代号 (Nickname)" name="nickname">
          <a-input v-model:value="form.nickname" />
        </a-form-item>
        <a-form-item label="性别" name="gender">
          <a-select v-model:value="form.gender">
            <a-select-option :value="0">未知</a-select-option>
            <a-select-option :value="1">男</a-select-option>
            <a-select-option :value="2">女</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="用户类型" name="userType">
          <a-select v-model:value="form.userType">
            <a-select-option value="normal">normal</a-select-option>
            <a-select-option value="vip">vip</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { message } from 'ant-design-vue'
import { UserAddOutlined } from '@ant-design/icons-vue'
import { getApiUserList } from '../api/controller/YongHu/getApiUserList'
import { putApiUserUpdate } from '../api/controller/YongHu/putApiUserUpdate'
import { deleteApiUserDelete } from '../api/controller/YongHu/deleteApiUserDelete'
import request from '../api/http'

// 手动定义创建用户接口（OpenAPI 中未生成）
const createUser = (data) => request.post('/user/create', data)

const users = ref([])
const loading = ref(false)
const searchText = ref('')
const userType = ref('normal')
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true
})

const modalVisible = ref(false)
const modalType = ref('create')
const modalLoading = ref(false)
const formRef = ref(null)
const form = reactive({
  id: null,
  account: '',
  password: '',
  nickname: '',
  gender: 0,
  userType: 'normal'
})

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '特工ID', dataIndex: 'account', key: 'account' },
  { title: '代号', dataIndex: 'nickname', key: 'nickname' },
  { title: '性别', dataIndex: 'gender', key: 'gender', width: 80 },
  { title: '用户类型', dataIndex: 'userType', key: 'userType', width: 100 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
]
const admincolumns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '特工ID', dataIndex: 'account', key: 'account' },
  { title: '代号', dataIndex: 'nickname', key: 'nickname' },
  { title: '性别', dataIndex: 'gender', key: 'gender', width: 80 },
  { title: '用户类型', dataIndex: 'userType', key: 'userType', width: 100 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', key: 'action', width: 150 }
]

const fetchUsers = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.current,
      pageSize: pagination.pageSize,
    }
    if (searchText.value) {
      params.account = searchText.value
      params.nickname = searchText.value
    }
    
    const res = await getApiUserList(params)
    
    if (res.code && res.code !== 0) {
      message.error(res.message || '获取档案列表失败')
      return
    }
    
    const data = res.data || res || {}
    users.value = data.users || data.Users || data.list || []
    pagination.total = data.total || 0
  } catch (err) {
    message.error(err.message || '获取档案列表失败')
  } finally {
    loading.value = false
  }
}

const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  fetchUsers()
}

const onSearch = () => {
  pagination.current = 1
  fetchUsers()
}

const showModal = (type, record) => {
  modalType.value = type
  modalVisible.value = true
  if (type === 'edit') {
    Object.assign(form, {
      id: record.id,
      account: record.account,
      password: '',
      nickname: record.nickname || '',
      gender: record.gender || 0,
      userType: record.userType || 'normal'
    })
  } else {
    Object.assign(form, {
      id: null,
      account: '',
      password: '',
      nickname: '',
      gender: 0,
      userType: 'normal'
    })
  }
}

const handleModalOk = async () => {
  modalLoading.value = true
  try {
    if (modalType.value === 'create') {
      await createUser(form)
      message.success('新成员注册成功')
    } else {
      await putApiUserUpdate(form)
      message.success('档案更新成功')
    }
    modalVisible.value = false
    fetchUsers()
  } catch (err) {
    message.error(err.message || '操作失败')
  } finally {
    modalLoading.value = false
  }
}

const formatGender = (gender) => {
  if (gender === 1) return '男'
  if (gender === 2) return '女'
  return '未知'
}

const formatTime = (time) => {
  if (!time) return '—'
  if (typeof time === 'string') {
    return time.replace('T', ' ').split('+')[0].split('.')[0]
  }
  return '—'
}

const handleDelete = async (record) => {
  try {
    await deleteApiUserDelete({ account: record.account })
    message.success('档案已移除')
    fetchUsers()
  } catch (err) {
    message.error(err.message || '删除失败')
  }
}

onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
.user-page {
  height: 100%;
}

.page-header {
  margin-bottom: 24px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  padding-bottom: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.page-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  color: #e2e8f0;
  margin: 0;
  letter-spacing: 2px;
}

.page-subtitle {
  color: #64748b;
  margin: 4px 0 0;
  font-size: 14px;
}

.actions {
  display: flex;
  gap: 16px;
}

.table-container {
  background: rgba(15, 23, 42, 0.6);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  padding: 1px; /* for border */
}

.avatar-cell {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  overflow: hidden;
  background: #0f172a;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.avatar-cell img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.danger-link {
  color: #ef4444;
}

.danger-link:hover {
  color: #dc2626;
}

/* Custom Table Styles */
:deep(.custom-table) .ant-table {
  background: transparent;
  color: #e2e8f0;
}

:deep(.custom-table) .ant-table-thead > tr > th {
  background: rgba(15, 23, 42, 0.8);
  color: #94a3b8;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  font-family: 'Roboto Mono', monospace;
  font-size: 12px;
}

:deep(.custom-table) .ant-table-tbody > tr > td {
  border-bottom: 1px solid rgba(255, 255, 255, 0.04);
}

:deep(.custom-table) .ant-table-tbody > tr:hover > td {
  background: rgba(255, 255, 255, 0.04) !important;
}

:deep(.ant-pagination) {
  margin: 16px;
}

:deep(.ant-pagination-item-active) {
  background: transparent;
  border-color: #d50000;
}

:deep(.ant-pagination-item-active a) {
  color: #d50000;
}

:deep(.create-btn) {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
}

.user-type-text {
  color: #ffffff;
  font-weight: 500;
}
</style>
