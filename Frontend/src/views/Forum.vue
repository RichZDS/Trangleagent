<template>
  <div class="forum-page">
    <div class="page-header">
      <div class="header-content">
        <div class="title-section">
          <h1 class="page-title">冷冻酸奶茶水间</h1>
          <p class="page-subtitle">内部情报交换与讨论终端</p>
        </div>
        <div class="actions">
          <a-input-search
            v-model:value="searchText"
            placeholder="搜索情报关键词..."
            style="width: 300px"
            @search="onSearch"
            class="custom-search"
          />
          <a-button type="primary" class="action-btn create-btn" @click="showCreateModal">
            <template #icon><FormOutlined /></template>
            发布情报
          </a-button>
        </div>
      </div>
    </div>

    <!-- 帖子列表 -->
    <div class="forum-content">
      <div class="post-list-container" v-loading="loading">
        <a-empty v-if="!loading && posts.length === 0" description="暂无情报记录" />
        <!-- 帖子卡片 -->
        <div 
          v-for="post in posts" 
          :key="post.id" 
          class="post-card"
          @click="openPostDetail(post)"
        >
          <div class="post-header">
            <div class="post-title">{{ post.title }}</div>
            <div class="post-author">
              <UserOutlined />
              <span class="author-name">{{ post.name || '匿名用户' }}</span>
            </div>
            <div class="post-meta">
              <span class="meta-item"><ClockCircleOutlined /> {{ formatTime(post.createdAt) }}</span>
            </div>
          </div>
          <div class="post-preview">
            {{ post.content }}
          </div>
          <div class="post-footer">
            <div class="stats">
              <span class="stat-item"><EyeOutlined /> {{ post.viewCount || 0 }}</span>
              <span class="stat-item"><LikeOutlined /> {{ post.likeCount || 0 }}</span>
              <span class="stat-item"><MessageOutlined /> {{ post.commentCount || 0 }}</span>
            </div>
            <div class="status-tag" :class="post.status">{{ post.status || 'NORMAL' }}</div>
          </div>
        </div>
        <!-- 分页器 -->
        <div class="pagination-wrapper" v-if="total > 0">
          <a-pagination
            v-model:current="pagination.page"
            v-model:pageSize="pagination.pageSize"
            :total="total"
            @change="handlePageChange"
            show-less-items
          />
        </div>
      </div>
    </div>

    <!-- 发布帖子弹窗 -->
    <a-modal
      v-model:open="createModalVisible"
      title="发布新情报"
      @ok="handleCreatePost"
      :confirmLoading="createLoading"
      class="custom-modal"
      width="600px"
    >
      <a-form :model="createForm" layout="vertical">
        <a-form-item label="情报标题" required>
          <a-input v-model:value="createForm.title" placeholder="请输入标题..." />
        </a-form-item>
        <a-form-item label="情报内容" required>
          <a-textarea 
            v-model:value="createForm.content" 
            placeholder="请输入详细内容..." 
            :rows="6" 
          />
        </a-form-item>
        <a-form-item label="封面图 (可选)">
          <a-input v-model:value="createForm.coverImage" placeholder="图片URL..." />
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 帖子详情抽屉 -->
    <a-drawer
      v-model:open="drawerVisible"
      title="情报详情"
      placement="right"
      width="600px"
      class="custom-drawer"
    >
      <div v-if="currentPost" class="post-detail">
        <h2 class="detail-title">{{ currentPost.title }}</h2>
        <div class="detail-meta">
          <span>{{ formatTime(currentPost.createdAt) }}</span>
        </div>
        
        <div class="detail-content">
          {{ currentPost.content }}
        </div>

        <a-divider style="border-color: rgba(255,255,255,0.1)">评论通讯</a-divider>

        <div class="comment-list">
          <a-empty v-if="comments.length === 0" description="暂无评论" />
          <div v-for="comment in comments" :key="comment.id" class="comment-item">
            <div class="comment-header">
              <span class="comment-time">{{ formatTime(comment.createdAt) }}</span>
            </div>
            <div class="comment-body">
              {{ comment.content }}
            </div>
          </div>
        </div>

        <div class="comment-input-area">
          <a-textarea
            v-model:value="newCommentContent"
            placeholder="输入评论内容..."
            :rows="3"
            class="custom-textarea"
          />
          <a-button 
            type="primary" 
            class="send-btn" 
            @click="handleSendComment"
            :loading="commentSending"
          >
            发送
          </a-button>
        </div>
      </div>
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { 
  FormOutlined, 
  UserOutlined, 
  ClockCircleOutlined, 
  EyeOutlined, 
  LikeOutlined, 
  MessageOutlined 
} from '@ant-design/icons-vue'

// API Imports
import { getApiForumPostsList } from '../api/controller/LunTanTieZi/getApiForumPostsList'
import { postApiForumPostsCreate } from '../api/controller/LunTanTieZi/postApiForumPostsCreate'
import { getApiForumCommentsList } from '../api/controller/LunTanPingLun/getApiForumCommentsList'
import { postApiForumCommentsCreate } from '../api/controller/LunTanPingLun/postApiForumCommentsCreate'
import { getApiUserView } from '../api/controller/YongHu/getApiUserView'

// Types
import type { TrangleAgentApiForumV1ForumPostsViewRes } from '../api/interface/apiTypes/TrangleAgentApiForumV1ForumPostsViewRes'
import type { TrangleAgentApiForumV1ForumCommentsViewRes } from '../api/interface/apiTypes/TrangleAgentApiForumV1ForumCommentsViewRes'

// State
const loading = ref(false)
const posts = ref<TrangleAgentApiForumV1ForumPostsViewRes[]>([])
const total = ref(0)
const searchText = ref('')
const pagination = reactive({
  page: 1,
  pageSize: 10
})

const createModalVisible = ref(false)
const createLoading = ref(false)
const createForm = reactive({
  title: '',
  content: '',
  coverImage: ''
})

const drawerVisible = ref(false)
const currentPost = ref<TrangleAgentApiForumV1ForumPostsViewRes | null>(null)
const comments = ref<TrangleAgentApiForumV1ForumCommentsViewRes[]>([])
const newCommentContent = ref('')
const commentSending = ref(false)
const currentUserId = ref<number>(0)

// 获取当前用户ID
const fetchCurrentUser = async () => {
  try {
    const storedId = localStorage.getItem('ta_user_id')
    const params: any = {}
    if (storedId) {
      params.id = Number(storedId)
    } else {
      const account = localStorage.getItem('ta_account')
      if (!account) return
      params.account = account
    }

    const res = await getApiUserView(params)
    const data = (res as any).data || res || {}
    if (data.id) {
      currentUserId.value = data.id
      localStorage.setItem('ta_user_id', String(data.id))
    }
  } catch (error) {
    console.error('获取用户信息失败', error)
  }
}

// Methods
const fetchPosts = async () => {
  loading.value = true
  try {
    const res = await getApiForumPostsList({
      page: pagination.page,
      pageSize: pagination.pageSize,
      title: searchText.value || undefined
    })
    const data = (res as any).data || res || {}
    const errMsg = (res as any).err
    if (errMsg) {
      message.error(errMsg)
      return
    }
    posts.value = data.list || []
    total.value = data.total || 0
  } catch (error) {
    message.error('获取情报列表失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const onSearch = () => {
  pagination.page = 1
  fetchPosts()
}

const handlePageChange = () => {
  fetchPosts()
}

const showCreateModal = () => {
  createForm.title = ''
  createForm.content = ''
  createForm.coverImage = ''
  createModalVisible.value = true
}

const handleCreatePost = async () => {
  if (!createForm.title || !createForm.content) {
    message.warning('请填写完整的标题和内容')
    return
  }

  if (!currentUserId.value) {
    message.error('无法获取用户信息，请重新登录')
    return
  }

  createLoading.value = true
  try {
    const res = await postApiForumPostsCreate({
      boardId: Number("1"),//默认板块IDtodo
      userId: currentUserId.value,
      title: createForm.title,
      content: createForm.content,
      coverImage: createForm.coverImage,
      status: 'NORMAL'
    })
    const data = (res as any).data || res || {}
    const errMsg = (res as any).err
    if (errMsg) {
      message.error(errMsg)
      return
    }
    message.success('情报发布成功')
    createModalVisible.value = false
    fetchPosts()
  } catch (error) {
    message.error('发布失败')
    console.error(error)
  } finally {
    createLoading.value = false
  }
}

const openPostDetail = async (post: TrangleAgentApiForumV1ForumPostsViewRes) => {
  currentPost.value = post
  drawerVisible.value = true
  comments.value = [] // 清空旧评论
  fetchComments(post.id!)
}

const fetchComments = async (postId: number) => {
  try {
    const res = await getApiForumCommentsList({
      postId: postId,
      page: 1,
      pageSize: 100
    })
    const data = (res as any).data || res || {}
    comments.value = data.list || []
  } catch (error) {
    console.error('获取评论失败', error)
  }
}

const handleSendComment = async () => {
  if (!newCommentContent.value.trim()) return
  if (!currentPost.value?.id) return

  if (!currentUserId.value) {
    message.error('无法获取用户信息，请重新登录')
    return
  }

  commentSending.value = true
  try {
    await postApiForumCommentsCreate({
      userId: currentUserId.value,
      postId: currentPost.value.id,
      content: newCommentContent.value,
      status: 'NORMAL'
    })
    message.success('评论发送成功')
    newCommentContent.value = ''
    fetchComments(currentPost.value.id)
    // 可以在这里更新 post.commentCount
  } catch (error) {
    message.error('发送失败')
    console.error(error)
  } finally {
    commentSending.value = false
  }
}

const formatTime = (timeStr?: string) => {
  if (!timeStr) return '-'
  return new Date(timeStr).toLocaleString()
}

onMounted(async () => {
  await fetchCurrentUser()
  fetchPosts()
})
</script>

<style scoped>
.forum-page {
  height: 100%;
  display: flex;
  flex-direction: column;
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

.forum-content {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
}

.post-card {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 16px;
  cursor: pointer;
  transition: all 0.3s;
}

.post-card:hover {
  background: rgba(30, 41, 59, 0.8);
  border-color: rgba(213, 0, 0, 0.5);
  transform: translateY(-2px);
}

.post-header {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.post-title {
  font-size: 18px;
  font-weight: 600;
  color: #e2e8f0;
  flex: 1;
  margin-right: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.post-author {
  display: flex;
  align-items: center;
  gap: 6px;
  color: #94a3b8;
  font-size: 13px;
  margin-right: 16px;
  background: rgba(255, 255, 255, 0.05);
  padding: 4px 10px;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  transition: all 0.3s ease;
}

.post-author:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.1);
  color: #e2e8f0;
}

.author-name {
  font-weight: 500;
}

.post-meta {
  font-size: 12px;
  color: #64748b;
  display: flex;
  gap: 12px;
}

.post-preview {
  color: #94a3b8;
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 16px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stats {
  display: flex;
  gap: 20px;
  color: #64748b;
  font-size: 13px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.status-tag {
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
}

/* Drawer Styles */
.detail-title {
  color: #e2e8f0;
  margin-bottom: 12px;
}

.detail-meta {
  color: #64748b;
  font-size: 13px;
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.detail-content {
  color: #cbd5e1;
  font-size: 15px;
  line-height: 1.8;
  white-space: pre-wrap;
  margin-bottom: 30px;
}

.comment-list {
  margin-bottom: 20px;
}

.comment-item {
  background: rgba(0, 0, 0, 0.2);
  padding: 12px;
  border-radius: 6px;
  margin-bottom: 12px;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #64748b;
  margin-bottom: 6px;
}

.comment-body {
  color: #cbd5e1;
  font-size: 14px;
}

.comment-input-area {
  margin-top: auto;
  padding-top: 16px;
}

.send-btn {
  margin-top: 12px;
  float: right;
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
}

/* Custom Scrollbar */
.forum-content::-webkit-scrollbar {
  width: 6px;
}

.forum-content::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
}

.forum-content::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
}

:deep(.create-btn) {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
}

:deep(.ant-pagination-item-active) {
  background: transparent;
  border-color: #d50000;
}

:deep(.ant-pagination-item-active a) {
  color: #d50000;
}
</style>
