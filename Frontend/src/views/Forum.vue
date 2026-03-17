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

    <!-- 帖子列表 / 帖子详情 切换 -->
    <div class="forum-content">
      <!-- 列表视图 -->
      <div v-show="viewMode === 'list'" class="post-list-container" v-loading="loading">
        <a-empty v-if="!loading && posts.length === 0" description="暂无情报记录" />
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
            {{ stripMarkdownPreview(post.content) }}
          </div>
          <div class="post-footer">
            <div class="stats">
              <span class="stat-item"><EyeOutlined /> {{ post.viewCount || 0 }}</span>
              <span class="stat-item stat-like" @click.stop="handleLike(post)" :class="{ liked: post.liked }">
                <LikeOutlined /> {{ post.likeCount || 0 }}
              </span>
              <span class="stat-item"><MessageOutlined /> {{ post.commentCount || 0 }}</span>
            </div>
            <div class="status-tag" :class="post.status">{{ post.status || 'NORMAL' }}</div>
          </div>
        </div>
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

      <!-- 帖子详情全页视图 -->
      <div v-show="viewMode === 'detail'" class="post-detail-page">
        <a-button class="back-btn" @click="viewMode = 'list'">
          <template #icon><ArrowLeftOutlined /></template>
          返回列表
        </a-button>
        <div v-if="currentPost" class="detail-main">
          <h1 class="detail-title">{{ currentPost.title }}</h1>
          <div class="detail-meta-row">
            <span class="detail-author"><UserOutlined /> {{ currentPost.name || '匿名用户' }}</span>
            <span class="detail-time"><ClockCircleOutlined /> {{ formatTime(currentPost.createdAt) }}</span>
            <div class="detail-stats">
              <span><EyeOutlined /> {{ currentPost.viewCount || 0 }}</span>
              <span class="stat-like" @click="handleLike(currentPost)" :class="{ liked: currentPost.liked }">
                <LikeOutlined /> {{ currentPost.likeCount || 0 }}
              </span>
              <span><MessageOutlined /> {{ currentPost.commentCount || 0 }}</span>
            </div>
          </div>
          <div class="detail-content markdown-body">
            <MdPreview :model-value="currentPost.content || ''" :theme="isDark ? 'dark' : 'light'" />
          </div>
          <a-divider class="detail-divider">评论通讯</a-divider>
          <div class="comment-list">
            <a-empty v-if="comments.length === 0" description="暂无评论" />
            <div v-for="comment in comments" :key="comment.id" class="comment-item">
              <div class="comment-header">
                <span class="comment-time">{{ formatTime(comment.createdAt) }}</span>
              </div>
              <div class="comment-body">{{ comment.content }}</div>
            </div>
          </div>
          <div class="comment-input-area">
            <a-textarea
              v-model:value="newCommentContent"
              placeholder="输入评论内容..."
              :rows="3"
              class="custom-textarea"
            />
            <a-button type="primary" class="send-btn" @click="handleSendComment" :loading="commentSending">
              发送
            </a-button>
          </div>
        </div>
      </div>
    </div>

    <!-- 发布帖子弹窗（可拖拽、可缩放） -->
    <a-modal
      v-model:open="createModalVisible"
      :title="null"
      :footer="null"
      class="forum-create-modal"
      :width="modalSize.w"
      :styles="{ body: { maxHeight: modalSize.h + 'px', overflow: 'auto' } }"
      :closable="true"
      @cancel="createModalVisible = false"
      :afterOpenChange="onModalAfterOpenChange"
    >
      <div class="create-modal-inner">
        <div class="create-modal-header" @mousedown.prevent="onDragStart">
          <div class="create-modal-icon">
            <FormOutlined />
          </div>
          <div class="create-modal-title-wrap">
            <h2 class="create-modal-title">发布情报</h2>
            <p class="create-modal-subtitle">支持 Markdown，图片直接上传至存储</p>
          </div>
        </div>

        <a-form :model="createForm" layout="vertical" class="create-form">
          <a-form-item class="form-item-title">
            <template #label>
              <span class="form-label">情报标题</span>
              <span class="form-required">*</span>
            </template>
            <a-input
              v-model:value="createForm.title"
              placeholder="简明扼要地概括情报主题..."
              size="large"
              class="create-input"
              :maxlength="120"
              show-count
            />
          </a-form-item>

          <a-form-item class="form-item-content">
            <template #label>
              <span class="form-label">情报内容</span>
              <span class="form-required">*</span>
              <span class="form-optional">（支持 Markdown，粘贴/拖拽图片自动上传）</span>
            </template>
            <MdEditor
              v-model="createForm.content"
              :theme="isDark ? 'dark' : 'light'"
              :language="'zh-CN'"
              :preview="true"
              :height="320"
              :on-upload-img="onUploadImg"
              placeholder="支持 **粗体**、*斜体*、图片、代码块等..."
            />
          </a-form-item>

          <div class="create-modal-footer">
            <a-button class="cancel-btn" @click="createModalVisible = false">
              取消
            </a-button>
            <a-button
              type="primary"
              class="submit-btn"
              :loading="createLoading"
              @click="handleCreatePost"
            >
              <FormOutlined v-if="!createLoading" />
              发布情报
            </a-button>
          </div>
        </a-form>
        <div class="modal-resize-handle" @mousedown.prevent="onResizeStart" />
      </div>
    </a-modal>

  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick } from 'vue'
import { message } from 'ant-design-vue'
import { MdEditor, MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import 'md-editor-v3/lib/preview.css'
import { 
  FormOutlined, 
  UserOutlined, 
  ClockCircleOutlined, 
  EyeOutlined, 
  LikeOutlined, 
  MessageOutlined,
  ArrowLeftOutlined
} from '@ant-design/icons-vue'
import { useTheme } from '../composables/useTheme'

// API Imports
import { getApiForumPostsList } from '../api/controller/LunTanTieZi/getApiForumPostsList'
import { getApiForumPostsView } from '../api/controller/LunTanTieZi/getApiForumPostsView'
import { postApiForumPostsCreate } from '../api/controller/LunTanTieZi/postApiForumPostsCreate'
import { postApiForumPostsLike } from '../api/controller/LunTanTieZi/postApiForumPostsLike'
import { getApiForumCommentsList } from '../api/controller/LunTanPingLun/getApiForumCommentsList'
import { postApiForumCommentsCreate } from '../api/controller/LunTanPingLun/postApiForumCommentsCreate'
import { getApiUserView } from '../api/controller/YongHu/getApiUserView'
import { postApiUploadImage } from '../api/controller/ShangChuan/postApiUploadImage'

const { isDark } = useTheme()

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

const viewMode = ref<'list' | 'detail'>('list')
const createModalVisible = ref(false)
const createLoading = ref(false)
const createForm = reactive({
  title: '',
  content: '',
  coverImage: '' // 可选，可从内容首图提取
})

// 弹窗拖拽与缩放
const modalSize = reactive({ w: 720, h: 520 })
let modalEl: HTMLElement | null = null
let dragStart = { x: 0, y: 0, modalX: 0, modalY: 0 }
let resizeStart = { x: 0, y: 0, w: 720, h: 520 }

const getModalEl = () =>
  document.querySelector('.create-modal-header')?.closest('.ant-modal') as HTMLElement | null

const onModalAfterOpenChange = (open: boolean) => {
  if (open) {
    modalSize.w = 720
    modalSize.h = 520
    nextTick(() => {
      modalEl = getModalEl()
      if (modalEl) {
        modalEl.style.left = ''
        modalEl.style.top = ''
        modalEl.style.transform = ''
        modalEl.style.margin = ''
      }
    })
  } else {
    modalEl = null
  }
}

const onDragStart = (e: MouseEvent) => {
  const el = getModalEl()
  if (!el) return
  const rect = el.getBoundingClientRect()
  dragStart = { x: e.clientX, y: e.clientY, modalX: rect.left, modalY: rect.top }
  const onMove = (ev: MouseEvent) => {
    const dx = ev.clientX - dragStart.x
    const dy = ev.clientY - dragStart.y
    el.style.left = (dragStart.modalX + dx) + 'px'
    el.style.top = (dragStart.modalY + dy) + 'px'
    el.style.transform = 'none'
    el.style.margin = '0'
  }
  const onUp = () => {
    document.removeEventListener('mousemove', onMove)
    document.removeEventListener('mouseup', onUp)
    document.body.style.cursor = ''
    document.body.style.userSelect = ''
  }
  document.body.style.cursor = 'move'
  document.body.style.userSelect = 'none'
  document.addEventListener('mousemove', onMove)
  document.addEventListener('mouseup', onUp)
}

const onResizeStart = (e: MouseEvent) => {
  e.stopPropagation()
  const el = getModalEl()
  if (!el) return
  resizeStart = { x: e.clientX, y: e.clientY, w: modalSize.w, h: modalSize.h }
  const onMove = (ev: MouseEvent) => {
    const dw = ev.clientX - resizeStart.x
    const dh = ev.clientY - resizeStart.y
    modalSize.w = Math.max(480, Math.min(1200, resizeStart.w + dw))
    modalSize.h = Math.max(400, Math.min(800, resizeStart.h + dh))
    el.style.width = modalSize.w + 'px'
    el.style.height = modalSize.h + 'px'
  }
  const onUp = () => {
    document.removeEventListener('mousemove', onMove)
    document.removeEventListener('mouseup', onUp)
    document.body.style.cursor = ''
    document.body.style.userSelect = ''
  }
  document.body.style.cursor = 'se-resize'
  document.body.style.userSelect = 'none'
  document.addEventListener('mousemove', onMove)
  document.addEventListener('mouseup', onUp)
}

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

const handleLike = async (post: TrangleAgentApiForumV1ForumPostsViewRes) => {
  if (!post.id) return
  try {
    await postApiForumPostsLike(post.id)
    post.likeCount = (post.likeCount || 0) + 1
    post.liked = true
    message.success('点赞成功')
  } catch {
    message.error('点赞失败')
  }
}

const openPostDetail = async (post: TrangleAgentApiForumV1ForumPostsViewRes) => {
  try {
    const res = await getApiForumPostsView({ id: post.id! })
    const data = (res as any).data || res || {}
    currentPost.value = { ...post, ...data }
  } catch {
    currentPost.value = post
  }
  viewMode.value = 'detail'
  comments.value = []
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
    if (currentPost.value) {
      currentPost.value.commentCount = (currentPost.value.commentCount || 0) + 1
    }
    fetchComments(currentPost.value!.id)
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

// 列表预览：去除 Markdown 语法，取纯文本
const stripMarkdownPreview = (md?: string) => {
  if (!md) return ''
  let text = md
    .replace(/!\[.*?\]\(.*?\)/g, '[图片]')
    .replace(/\[([^\]]+)\]\([^)]+\)/g, '$1')
    .replace(/#{1,6}\s*/g, '')
    .replace(/\*\*([^*]+)\*\*/g, '$1')
    .replace(/\*([^*]+)\*/g, '$1')
    .replace(/`([^`]+)`/g, '$1')
    .replace(/```[\s\S]*?```/g, '')
    .replace(/\n+/g, ' ')
    .trim()
  return text.length > 150 ? text.slice(0, 150) + '...' : text
}

// 图片上传：粘贴/拖拽时调用，上传到 MinIO
const onUploadImg = async (files: File[], callback: (urls: string[]) => void) => {
  try {
    const urls = await Promise.all(
      files.map(async (file) => {
        const res = await postApiUploadImage(file) as any
        const data = res?.data || res || {}
        return data.url || ''
      })
    )
    callback(urls.filter(Boolean))
  } catch (err) {
    message.error('图片上传失败，请检查 MinIO 服务')
    callback([])
  }
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
  border-bottom: 1px solid var(--ta-border);
}

.page-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  color: var(--ta-text);
  margin: 0;
  letter-spacing: 2px;
}

.page-subtitle {
  color: var(--ta-text-dim);
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
  background: var(--ta-surface);
  border: 1px solid var(--ta-border);
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 16px;
  cursor: pointer;
  transition: all 0.3s;
}

.post-card:hover {
  background: var(--ta-surface-hover);
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
  color: var(--ta-text);
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
  color: var(--ta-text-muted);
  font-size: 13px;
  margin-right: 16px;
  background: var(--ta-overlay);
  padding: 4px 10px;
  border-radius: 12px;
  border: 1px solid var(--ta-border);
  transition: all 0.3s ease;
}

.post-author:hover {
  background: var(--ta-overlay-hover);
  border-color: var(--ta-overlay-hover);
  color: var(--ta-text);
}

.author-name {
  font-weight: 500;
}

.post-meta {
  font-size: 12px;
  color: var(--ta-text-dim);
  display: flex;
  gap: 12px;
}

.post-preview {
  color: var(--ta-text-muted);
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
  color: var(--ta-text-dim);
  font-size: 13px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.stat-like {
  cursor: pointer;
  padding: 2px 8px;
  border-radius: 4px;
  transition: all 0.2s;
}
.stat-like:hover {
  color: var(--ta-accent);
  background: rgba(213, 0, 0, 0.1);
}
.stat-like.liked {
  color: var(--ta-accent);
}

.status-tag {
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 4px;
  background: var(--ta-overlay-hover);
  color: var(--ta-text-muted);
}

/* 帖子详情全页 */
.post-detail-page {
  max-width: 900px;
  margin: 0 auto;
  padding: 0 0 40px;
}

.back-btn {
  margin-bottom: 24px;
  color: var(--ta-text-muted);
}

.back-btn:hover {
  color: var(--ta-accent);
}

.detail-main {
  background: var(--ta-surface);
  border: 1px solid var(--ta-border);
  border-radius: 12px;
  padding: 32px 40px;
}

.detail-title {
  color: var(--ta-text);
  font-size: 26px;
  font-weight: 700;
  margin: 0 0 20px;
  line-height: 1.4;
}

.detail-meta-row {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 16px 24px;
  color: var(--ta-text-dim);
  font-size: 14px;
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--ta-border);
}

.detail-author,
.detail-time {
  display: flex;
  align-items: center;
  gap: 6px;
}

.detail-stats {
  display: flex;
  gap: 20px;
  margin-left: auto;
}

.detail-stats span {
  display: flex;
  align-items: center;
  gap: 6px;
}

.detail-meta {
  color: var(--ta-text-dim);
  font-size: 13px;
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.detail-content {
  color: var(--ta-text-muted);
  font-size: 15px;
  line-height: 1.8;
  margin-bottom: 30px;
}
.detail-content:not(.markdown-body) {
  white-space: pre-wrap;
}

.comment-list {
  margin-bottom: 20px;
}

.comment-item {
  background: var(--ta-card-alt);
  padding: 12px;
  border-radius: 6px;
  margin-bottom: 12px;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: var(--ta-text-dim);
  margin-bottom: 6px;
}

.comment-body {
  color: var(--ta-text-muted);
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
  background: var(--ta-scrollbar-track);
}

.forum-content::-webkit-scrollbar-thumb {
  background: var(--ta-scrollbar-thumb);
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

/* 发布情报弹窗 - 论坛风格 */
:deep(.forum-create-modal .ant-modal-content) {
  background: linear-gradient(160deg, #0f172a 0%, #1e293b 50%, #0f172a 100%);
  border: 1px solid rgba(213, 0, 0, 0.25);
  box-shadow: 0 25px 80px rgba(0, 0, 0, 0.6), 0 0 40px rgba(213, 0, 0, 0.08);
  border-radius: 12px;
  overflow: hidden;
  position: relative;
}

.modal-resize-handle {
  position: absolute;
  right: 0;
  bottom: 0;
  width: 20px;
  height: 20px;
  cursor: se-resize;
  z-index: 10;
}
.modal-resize-handle::after {
  content: '';
  position: absolute;
  right: 4px;
  bottom: 4px;
  width: 10px;
  height: 10px;
  border-right: 2px solid var(--ta-text-muted);
  border-bottom: 2px solid var(--ta-text-muted);
  opacity: 0.6;
}

:deep(.forum-create-modal .ant-modal-body) {
  padding: 0;
  background: transparent;
}

/* Markdown 预览样式 */
.detail-content.markdown-body {
  color: var(--ta-text);
}

:deep(.detail-divider) {
  border-color: var(--ta-border) !important;
}
.detail-content.markdown-body :deep(img) {
  max-width: 100%;
  height: auto;
  border-radius: 8px;
  margin: 12px 0;
}

:deep(.forum-create-modal .ant-modal-close) {
  top: 16px;
  right: 16px;
  width: 36px;
  height: 36px;
  color: #94a3b8;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

:deep(.forum-create-modal .ant-modal-close:hover) {
  color: #e2e8f0;
  background: rgba(213, 0, 0, 0.2);
}

.create-modal-inner {
  padding: 0;
  position: relative;
  min-height: 200px;
}

.create-modal-header {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 28px 32px 24px;
  background: linear-gradient(90deg, rgba(213, 0, 0, 0.12) 0%, transparent 100%);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  cursor: move;
  user-select: none;
}

.create-modal-icon {
  width: 56px;
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #d50000, #7f1d1d);
  border-radius: 12px;
  font-size: 28px;
  color: #fff;
  box-shadow: 0 8px 24px rgba(213, 0, 0, 0.35);
}

.create-modal-title-wrap {
  flex: 1;
}

.create-modal-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 22px;
  font-weight: 700;
  color: #e2e8f0;
  margin: 0;
  letter-spacing: 2px;
}

.create-modal-subtitle {
  font-size: 13px;
  color: #64748b;
  margin: 6px 0 0;
  letter-spacing: 1px;
}

.create-form {
  padding: 24px 32px 28px;
}

.create-form .form-item-title,
.create-form .form-item-content,
.create-form .form-item-cover {
  margin-bottom: 20px;
}

.create-form .form-label {
  font-family: 'Roboto Mono', monospace;
  font-size: 13px;
  font-weight: 600;
  color: #94a3b8;
  letter-spacing: 1px;
}

.create-form .form-required {
  color: #f87171;
  margin-left: 2px;
}

.create-form .form-optional {
  color: #64748b;
  font-weight: 400;
  margin-left: 4px;
}

.create-input,
.create-textarea,
.custom-textarea {
  background: var(--ta-input-bg) !important;
  border: 1px solid var(--ta-input-border) !important;
  color: var(--ta-text) !important;
  border-radius: 8px !important;
}

.create-input:focus,
.create-textarea:focus,
.custom-textarea:focus,
.create-input:hover,
.create-textarea:hover,
.custom-textarea:hover {
  border-color: var(--ta-input-focus-border) !important;
  box-shadow: 0 0 0 2px var(--ta-input-focus-shadow) !important;
}

.create-input::placeholder,
.create-textarea::placeholder,
.custom-textarea::placeholder {
  color: var(--ta-text-dim);
}

.create-textarea {
  resize: vertical;
  min-height: 180px;
}

:deep(.create-input .ant-input),
:deep(.create-textarea .ant-input),
:deep(.custom-textarea .ant-input) {
  background: transparent !important;
  color: var(--ta-text) !important;
}

:deep(.create-input .ant-input-show-count-suffix),
:deep(.create-textarea .ant-input-data-count) {
  color: var(--ta-text-dim);
  font-size: 12px;
}

.create-modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 28px;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
}

.cancel-btn {
  background: rgba(255, 255, 255, 0.05) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  color: #94a3b8 !important;
}

.cancel-btn:hover {
  background: rgba(255, 255, 255, 0.08) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  color: #e2e8f0 !important;
}

.submit-btn {
  background: linear-gradient(120deg, #d50000, #7f1d1d) !important;
  border: none !important;
  font-weight: 600 !important;
  letter-spacing: 1px !important;
  padding: 0 24px !important;
}

.submit-btn:hover {
  background: linear-gradient(120deg, #f43f5e, #991b1b) !important;
  box-shadow: 0 4px 20px rgba(213, 0, 0, 0.4) !important;
}
</style>
