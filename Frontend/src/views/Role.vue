<template>
  <div class="role-container">
    <div class="header">
      <h2>角色管理</h2>
      <a-button type="primary" @click="handleCreate">新增角色卡</a-button>
    </div>

    <a-spin :spinning="loading">
      <div v-if="roles.length === 0 && !loading" class="empty-state">暂无角色数据</div>
      <div class="role-grid">
        <div v-for="role in roles" :key="role.id" class="role-card">
          <div class="card-triangle"></div>
          <div class="card-content">
            <div class="card-actions-top">
              <eye-outlined class="action-icon" @click="handleDetail(role)" />
              <edit-outlined class="action-icon" @click="handleEdit(role)" />
              <delete-outlined class="action-icon" @click="handleDelete(role)" />
            </div>
            <h3 class="role-name">{{ role.agentName || '未命名特工' }}</h3>
            <div class="role-info">
              <div class="info-item">
                <span class="label">代号:</span>
                <span class="value">{{ role.codeName || '-' }}</span>
              </div>
              <div class="info-item">
                <span class="label">部门:</span>
                <span class="value">{{ getDepartmentName(role.departmentId) }}</span>
              </div>
            </div>
            <div class="card-actions">
              <a-button type="primary" block class="detail-btn" @click="handleDetail(role)">
                查看详情
              </a-button>
            </div>
          </div>
        </div>
      </div>

      <div class="pagination-container" v-if="total > 0">
        <a-pagination
          v-model:current="pagination.current"
          v-model:pageSize="pagination.pageSize"
          :total="total"
          @change="onPaginationChange"
          show-size-changer
        />
      </div>
    </a-spin>

    <!-- Modal for Create/Edit/View -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      :footer="modalType === 'view' ? null : undefined"
      :closable="true"
      :maskClosable="modalType === 'view'"
      width="900px"
      class="report-modal"
    >
      <a-form
        v-if="modalType !== 'view'"
        :model="formState"
        layout="vertical"
        ref="formRef"
        class="report-form"
      >
        <div class="report-header">
          <div class="report-title">特工档案报告</div>
          <div class="report-subtitle">AGENT PROFILE REPORT</div>
        </div>
        
        <div class="report-section">
          <div class="section-title">基础信息</div>
          <a-row :gutter="24">
            <a-col :span="12">
              <a-form-item label="特工名字" name="agentName" :rules="[{ required: true, message: '请输入特工名字' }]">
                <a-input v-model:value="formState.agentName" placeholder="Agent Name" class="report-input" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="代号" name="codeName" :rules="[{ required: true, message: '请输入代号' }]">
                <a-input v-model:value="formState.codeName" placeholder="Code Name" class="report-input" />
              </a-form-item>
            </a-col>
          </a-row>
          
          <a-row :gutter="24">
            <a-col :span="12">
              <a-form-item label="所属部门ID" name="departmentId">
                <a-input-number v-model:value="formState.departmentId" style="width: 100%" placeholder="Department ID" class="report-input" />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="性别" name="gender">
                <a-select v-model:value="formState.gender" style="width: 100%" class="report-input" @change="handleGenderChange">
                  <a-select-option value="male">男</a-select-option>
                  <a-select-option value="female">女</a-select-option>
                  <a-select-option value="other">其他</a-select-option>
                </a-select>
                <a-input 
                  v-if="formState.gender === 'other'" 
                  v-model:value="formState.genderCustom" 
                  placeholder="请输入性别" 
                  class="report-input" 
                  style="margin-top: 8px;"
                />
              </a-form-item>
            </a-col>
          </a-row>
        </div>

        <div class="report-section">
          <div class="section-title">QA 素质保障</div>
          <div class="qa-fixed-grid">
            <div v-for="(label, key) in qaFields" :key="key" class="qa-fixed-item">
              <div class="qa-fixed-label">{{ label }}</div>
              <div class="qa-fixed-value">{{ formState[key] || 0 }}</div>
            </div>
          </div>
        </div>
      </a-form>

      <div v-else class="detail-view report-detail">
        <div class="report-layout">
          <div class="report-content">
            <div class="report-header">
              <div class="report-title">特工档案报告</div>
              <div class="report-subtitle">AGENT PROFILE REPORT</div>
            </div>
            
            <div id="section-basic" class="report-card">
              <div class="card-header">
                <div class="card-title">基础信息</div>
              </div>
              <div class="card-body">
                <div class="report-info-grid">
                  <div class="report-info-item">
                    <div class="info-label">特工名字</div>
                    <div class="info-value">{{ currentRole?.agentName || '-' }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">代号</div>
                    <div class="info-value">{{ currentRole?.codeName || '-' }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">部门</div>
                    <div class="info-value">{{ getDepartmentName(currentRole?.departmentId) }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">用户ID</div>
                    <div class="info-value">{{ currentRole?.userId || '-' }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">性别</div>
                    <div class="info-value">{{ formatGender(currentRole?.gender) || '-' }}</div>
                  </div>
                </div>
              </div>
            </div>
            
            <div id="section-arc" class="report-card">
              <div class="card-header">
                <div class="card-title">ARC状态</div>
              </div>
              <div class="card-body">
                <div class="report-info-grid">
                  <div class="report-info-item">
                    <div class="info-label">异常</div>
                    <div class="info-value">{{ currentRole?.arcAbnormal || '-' }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">现实</div>
                    <div class="info-value">{{ currentRole?.arcReality || '-' }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">职位</div>
                    <div class="info-value">{{ currentRole?.arcPosition || '-' }}</div>
                  </div>
                </div>
              </div>
            </div>

            <div id="section-qa" class="report-card">
              <div class="card-header">
                <div class="card-title">QA 素质保障</div>
              </div>
              <div class="card-body">
                <div class="qa-fixed-grid">
                  <div v-for="(val, key) in getQaProperties(currentRole)" :key="key" class="qa-fixed-item">
                    <div class="qa-fixed-label">{{ formatQaLabel(key) }}</div>
                    <div class="qa-fixed-value">{{ val }}</div>
                  </div>
                </div>
              </div>
            </div>
            
            <div id="section-reward" class="report-card">
              <div class="card-header">
                <div class="card-title">奖惩与轨迹</div>
              </div>
              <div class="card-body">
                <div class="report-info-grid">
                  <div class="report-info-item">
                    <div class="info-label">嘉奖</div>
                    <div class="info-value">{{ currentRole?.commendation || 0 }}</div>
                  </div>
                  <div class="report-info-item">
                    <div class="info-label">申戒</div>
                    <div class="info-value">{{ currentRole?.reprimand || 0 }}</div>
                  </div>
                </div>
                
                <!-- 轨迹显示 -->
                <div class="track-display">
                  <div class="track-section">
                    <div class="track-title blue-title">异常轨道（蓝轨）</div>
                    <div class="track-grid-30">
                      <div 
                        v-for="i in 30" 
                        :key="i"
                        class="track-cell-wrapper"
                      >
                        <div 
                          class="track-cell"
                          :class="{ 
                            'active': (currentRole?.blueTrack || 0) >= i,
                            'blue-track': true,
                            'has-file': getTrackFileName('blue', i)
                          }"
                          @click="handleImageClick('blue', i)"
                        >
                          <span v-if="getTrackFileName('blue', i)" class="file-label">
                            {{ getTrackFileName('blue', i) }}
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  <div class="track-section">
                    <div class="track-title red-title">职能轨道（红轨）</div>
                    <div class="track-grid-30">
                      <div 
                        v-for="i in 30" 
                        :key="i"
                        class="track-cell-wrapper"
                      >
                        <div 
                          class="track-cell"
                          :class="{ 
                            'active': (currentRole?.redTrack || 0) >= i,
                            'red-track': true,
                            'has-file': getTrackFileName('red', i)
                          }"
                          @click="handleImageClick('red', i)"
                        >
                          <span v-if="getTrackFileName('red', i)" class="file-label">
                            {{ getTrackFileName('red', i) }}
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  <div class="track-section">
                    <div class="track-title yellow-title">现实轨道（黄轨）</div>
                    <div class="track-grid-30">
                      <div 
                        v-for="i in 30" 
                        :key="i"
                        class="track-cell-wrapper"
                      >
                        <div 
                          class="track-cell"
                          :class="{ 
                            'active': (currentRole?.yellowTrack || 0) >= i,
                            'yellow-track': true,
                            'has-file': getTrackFileName('yellow', i)
                          }"
                          @click="handleImageClick('yellow', i)"
                        >
                          <span v-if="getTrackFileName('yellow', i)" class="file-label">
                            {{ getTrackFileName('yellow', i) }}
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <div class="report-bookmarks">
            <div 
              class="bookmark-item" 
              :class="{ active: activeBookmark === 'basic' }"
              @click="scrollToSection('basic')"
            >
              <div class="bookmark-label">基础信息</div>
            </div>
            <div 
              class="bookmark-item" 
              :class="{ active: activeBookmark === 'arc' }"
              @click="scrollToSection('arc')"
            >
              <div class="bookmark-label">ARC状态</div>
            </div>
            <div 
              class="bookmark-item" 
              :class="{ active: activeBookmark === 'qa' }"
              @click="scrollToSection('qa')"
            >
              <div class="bookmark-label">QA 素质保障</div>
            </div>
            <div 
              class="bookmark-item" 
              :class="{ active: activeBookmark === 'reward' }"
              @click="scrollToSection('reward')"
            >
              <div class="bookmark-label">奖惩与轨迹</div>
            </div>
          </div>
        </div>
      </div>
    </a-modal>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed } from 'vue';
import { message, Modal } from 'ant-design-vue';
import { EyeOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons-vue';
import { getApiRoleList } from '../api/controller/JiaoSe/getApiRoleList';
import { postApiRoleCreate } from '../api/controller/JiaoSe/postApiRoleCreate';
import { putApiRoleUpdate } from '../api/controller/JiaoSe/putApiRoleUpdate';
import { deleteApiRoleDelete } from '../api/controller/JiaoSe/deleteApiRoleDelete';
import { getApiDepartmentView } from '../api/controller/BuMen/getApiDepartmentView';
import { getApiUserView } from '../api/controller/YongHu/getApiUserView';
// 定义接口类型 (简单复用，实际应从 generated types 导入)
interface Role {
  id: number;
  userId?: number;
  departmentId?: number;
  agentName?: string;
  codeName?: string;
  [key: string]: any;
}

// 状态
const roles = ref<Role[]>([]);
const loading = ref(false);
const total = ref(0);
const currentUserId = ref<number | undefined>(undefined);
const pagination = reactive({
  current: 1,
  pageSize: 10,
});

const departmentCache = reactive(new Map<number, string>());

// 弹窗状态
const modalVisible = ref(false);
const modalType = ref<'view' | 'create' | 'edit'>('view');
const currentRole = ref<Role | null>(null);
const formRef = ref();
const activeBookmark = ref('basic');
const showTrackCard = ref(false);
// 添加一个响应式变量来存储需要在新窗口显示的图像信息
const showFileImage = ref<{trackType: string, position: number, filename: string} | null>(null);

const formState = reactive<Record<string, any>>({
    agentName: '',
    codeName: '',
    departmentId: undefined,
    userId: undefined,
    gender: 'other' as 'male' | 'female' | 'other',
    genderCustom: '',
    qaMeticulous: 0,
    qaDeception: 0,
    qaVigor: 0,
    qaEmpathy: 0,
    qaInitiative: 0,
    qaResilience: 0,
    qaPresence: 0,
    qaProfessional: 0,
    qaDiscretion: 0,
    commendation: 0,
    reprimand: 0,
    blueTrack: 0,
    yellowTrack: 0,
    redTrack: 0
});

// QA 字段映射
const qaFields: Record<string, string> = {
    qaMeticulous: '细致',
    qaDeception: '欺诈',
    qaVigor: '活力',
    qaEmpathy: '共情',
    qaInitiative: '主动',
    qaResilience: '韧性',
    qaPresence: '气场',
    qaProfessional: '专业',
    qaDiscretion: '谨慎'
};

// 轨迹文件映射：位置 => 文件名
// 蓝轨：位置 1,2,5,7,11,13,17,19,23 对应文件 H4,H3,U2,X2,N1,Q2,L10,G8,A7
const blueTrackFileMap: Record<number, string> = {
  1: 'H4',
  2: 'H3',
  5: 'U2',
  7: 'X2',
  11: 'A7',
  13: 'Q2',
  17: 'L10',
  19: 'G8',
  23: 'N1'
};

// 黄轨：位置 1,4,8,10,14,16,20,22,26 对应文件 C4,L11,E2,O4,T6,V2,E3,H5,X3
const yellowTrackFileMap: Record<number, string> = {
  1: 'C4',
  4: 'L11',
  8: 'E2',
  10: 'O4',
  14: 'T6',
  16: 'V2',
  20: 'E3',
  22: 'H5',
  26: 'X3'
};

// 红轨：位置 3,6,9,12,15,18,21,24,27 对应文件 A3,D4,G3,J3,N3,Y2,W8,T3,Q3
const redTrackFileMap: Record<number, string> = {
  3: 'A3',
  6: 'D4',
  9: 'G3',
  12: 'J3',
  15: 'N3',
  18: 'Y2',
  21: 'W8',
  24: 'T3',
  27: 'Q3'
};

// 获取轨迹文件名
const getTrackFileName = (trackType: string, position: number): string => {
  let map: Record<number, string>;
  if (trackType === 'blue') {
    map = blueTrackFileMap;
  } else if (trackType === 'yellow') {
    map = yellowTrackFileMap;
  } else if (trackType === 'red') {
    map = redTrackFileMap;
  } else {
    return '';
  }
  return map[position] || '';
};

// 图片加载错误处理
const handleImageError = (event: Event) => {
  const img = event.target as HTMLImageElement;
  img.style.display = 'none';
};

const modalTitle = computed(() => {
    switch (modalType.value) {
        case 'create': return '新增角色卡';
        case 'edit': return '编辑角色卡';
        case 'view': return '';
        default: return '';
    }
});

// 获取当前用户ID
const fetchCurrentUserId = async () => {
    try {
        const storedId = localStorage.getItem('ta_user_id');
        const params: any = {};
        if (storedId) {
            params.id = Number(storedId);
        } else {
            const account = localStorage.getItem('ta_account');
            if (!account) {
                message.error('未找到用户信息');
                return;
            }
            params.account = account;
        }
        const res = await getApiUserView(params);
        const data = (res as any)?.data || res || {};
        if (data.id) {
            currentUserId.value = data.id;
            localStorage.setItem('ta_user_id', String(data.id));
        }
    } catch (error) {
        console.error('获取用户ID失败:', error);
    }
};

// 获取列表
const fetchRoles = async () => {
    if (!currentUserId.value) {
        await fetchCurrentUserId();
    }
    
    loading.value = true;
    try {
        const res = await getApiRoleList({
            page: pagination.current,
            pageSize: pagination.pageSize,
            userId: currentUserId.value,
        });
        const data = (res as any)?.data || res || {};
        roles.value = data.list || [];
        total.value = data.total || 0;
        
        // 获取部门名称
        fetchDepartmentNames(roles.value);
    } catch (error) {
        message.error('获取角色列表失败');
        console.error(error);
    } finally {
        loading.value = false;
    }
};

const onPaginationChange = (page: number, pageSize: number) => {
    pagination.current = page;
    pagination.pageSize = pageSize;
    fetchRoles();
};

// 获取部门名称逻辑
const fetchDepartmentNames = async (roleList: Role[]) => {
    const ids = new Set<number>();
    roleList.forEach(r => {
        if (r.departmentId && !departmentCache.has(r.departmentId)) {
            ids.add(r.departmentId);
        }
    });

    if (ids.size === 0) return;

    // 并发请求部门信息 (简单处理)
    const promises = Array.from(ids).map(id => 
        getApiDepartmentView({ id })
            .then(res => {
                // http.js 已经处理了 response.data，所以 res 就是数据本身
                const data = (res as any)?.data || res || {};
                if (data.branchName) {
                    departmentCache.set(id, data.branchName);
                }
            })
            .catch(() => {
                departmentCache.set(id, '未知部门');
            })
    );

    await Promise.all(promises);
};

const getDepartmentName = (id?: number) => {
    if (!id) return '-';
    return departmentCache.get(id) || `ID:${id}`;
};

// 操作处理
const handleCreate = async () => {
    if (!currentUserId.value) {
        await fetchCurrentUserId();
    }
    modalType.value = 'create';
    resetForm();
    // 自动填充当前用户ID
    formState.userId = currentUserId.value;
    modalVisible.value = true;
};

const handleGenderChange = () => {
    if (formState.gender !== 'other') {
        formState.genderCustom = '';
    }
};

const handleEdit = (role: Role) => {
    modalType.value = 'edit';
    currentRole.value = role;
    // 编辑时只能修改：特工名字、代号、所属部门ID、性别
    formState.id = role.id;
    formState.agentName = role.agentName || '';
    formState.codeName = role.codeName || '';
    formState.departmentId = role.departmentId;
    // 处理性别：如果是自定义的，设置为 other 并填充 custom
    const gender = role.gender || 'other';
    if (gender !== 'male' && gender !== 'female' && gender !== 'other') {
        formState.gender = 'other';
        formState.genderCustom = gender;
    } else {
        formState.gender = gender;
        formState.genderCustom = '';
    }
    // 其他字段保持原值（只读）
    Object.keys(qaFields).forEach(k => {
        formState[k] = role[k] || 0;
    });
    modalVisible.value = true;
};

const handleDetail = (role: Role) => {
    console.log('handleDetail called', role);
    try {
        modalType.value = 'view';
        currentRole.value = role;
        activeBookmark.value = 'basic';
        showTrackCard.value = false;
        modalVisible.value = true;
        // 等待DOM更新后滚动到顶部
        setTimeout(() => {
            const content = document.querySelector('.report-content');
            if (content) {
                content.scrollTop = 0;
            }
        }, 100);
    } catch (error) {
        console.error('Error in handleDetail:', error);
        message.error('打开详情失败');
    }
};

const handleModalCancel = () => {
    modalVisible.value = false;
    if (modalType.value === 'view') {
        currentRole.value = null;
        showTrackCard.value = false;
    }
};

const scrollToSection = (section: string) => {
    activeBookmark.value = section;
    const element = document.getElementById(`section-${section}`);
    if (element) {
        const content = document.querySelector('.report-content');
        if (content) {
            const offsetTop = element.offsetTop - 20;
            content.scrollTo({
                top: offsetTop,
                behavior: 'smooth'
            });
        }
    }
};

const handleDelete = (role: Role) => {
    Modal.confirm({
        title: '确认删除',
        content: `确定要删除特工 "${role.agentName}" 的角色卡吗？`,
        onOk: async () => {
            try {
                await deleteApiRoleDelete({ id: role.id });
                message.success('删除成功');
                fetchRoles();
            } catch (error) {
                message.error('删除失败');
            }
        }
    });
};

const handleModalOk = async () => {
    if (modalType.value === 'view') {
        modalVisible.value = false;
        return;
    }

    try {
        await formRef.value.validate();
        
        if (modalType.value === 'create') {
            // 确保必填字段存在
            const createData = {
                userId: formState.userId || currentUserId.value,
                agentName: formState.agentName,
                codeName: formState.codeName,
                gender: formState.gender === 'other' && formState.genderCustom ? formState.genderCustom : (formState.gender || 'other'),
                departmentId: formState.departmentId,
                ...formState
            };
            await postApiRoleCreate(createData as any);
            message.success('创建成功');
        } else {
            // 编辑时只能修改：特工名字、代号、所属部门ID、性别
            const updateData = {
                id: formState.id,
                agentName: formState.agentName,
                codeName: formState.codeName,
                gender: formState.gender === 'other' && formState.genderCustom ? formState.genderCustom : (formState.gender || 'other'),
                departmentId: formState.departmentId,
            };
            await putApiRoleUpdate(updateData as any);
            message.success('更新成功');
        }
        
        modalVisible.value = false;
        fetchRoles();
    } catch (error) {
        console.error(error);
        // Form validation error or API error
        if (!(error as any).errorFields) { // not validation error
             message.error(modalType.value === 'create' ? '创建失败' : '更新失败');
        }
    }
};

const resetForm = () => {
    formState.id = undefined;
    formState.agentName = '';
    formState.codeName = '';
    formState.departmentId = undefined;
    formState.userId = currentUserId.value;
    formState.gender = 'other';
    Object.keys(qaFields).forEach(k => formState[k] = 0);
    formState.commendation = 0;
    formState.reprimand = 0;
    formState.blueTrack = 0;
    formState.yellowTrack = 0;
    formState.redTrack = 0;
};

// 辅助函数：格式化性别
const formatGender = (gender?: string) => {
    if (!gender) return '-';
    if (gender === 'male') return '男';
    if (gender === 'female') return '女';
    return gender; // 其他或自定义
};

// 辅助函数：格式化 QA 标签
const formatQaLabel = (key: string) => {
    // 尝试直接匹配
    if (qaFields[key]) return qaFields[key];

    // 模糊匹配（兼容 snake_case 和 camelCase）
    const lowerKey = key.toLowerCase();
    if (lowerKey.includes('meticulous')) return '细致';
    if (lowerKey.includes('deception')) return '欺诈';
    if (lowerKey.includes('vigor')) return '活力';
    if (lowerKey.includes('empathy')) return '共情';
    if (lowerKey.includes('initiative')) return '主动';
    if (lowerKey.includes('resilience')) return '韧性';
    if (lowerKey.includes('presence')) return '气场';
    if (lowerKey.includes('professional')) return '专业';
    if (lowerKey.includes('discretion')) return '谨慎';

    return key;
};

// 辅助函数：获取 QA 属性列表
const getQaProperties = (role: Role | null) => {
    if (!role) return {};
    const result: Record<string, number> = {};
    for (const key in role) {
        // 匹配 qaXxx 或 qa_xxx
        if (key.startsWith('qa') && typeof role[key] === 'number') {
             result[key] = role[key];
        }
    }
    // 如果 role 中没有 QA 字段（可能是新对象），填充默认 0
    Object.keys(qaFields).forEach(k => {
        if (result[k] === undefined) {
             result[k] = 0;
        }
    });
    return result;
};

// 添加 handleImageClick 方法用于处理点击事件
const handleImageClick = (trackType: string, position: number) => {
  const filename = getTrackFileName(trackType, position);
  if (filename && currentRole.value) {
    // 检查是否已达到该位置
    let trackValue = 0;
    switch(trackType) {
      case 'blue':
        trackValue = currentRole.value.blueTrack || 0;
        break;
      case 'red':
        trackValue = currentRole.value.redTrack || 0;
        break;
      case 'yellow':
        trackValue = currentRole.value.yellowTrack || 0;
        break;
    }
    
    // 只有当轨道值大于等于当前位置时才能查看文件
    if (trackValue >= position) {
      // 在新窗口打开文件查看页面
      const url = `/file-view?trackType=${trackType}&filename=${filename}`;
      window.open(url, '_blank', 'width=1200,height=800');
    } else {
      message.warning('该轨道位置尚未解锁');
    }
  }
};

onMounted(async () => {
    await fetchCurrentUserId();
    fetchRoles();
});
</script>

<style scoped>
.role-container {
  padding: 20px;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.role-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 24px;
}

.role-card {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  position: relative;
  overflow: hidden;
  transition: all 0.3s;
  backdrop-filter: blur(4px);
}

.role-card:hover {
  transform: translateY(-4px);
  border-color: rgba(213, 0, 0, 0.5);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.card-triangle {
  position: absolute;
  top: 0;
  left: 0;
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 24px 24px 0 0;
  border-color: #d50000 transparent transparent transparent;
  z-index: 2;
}

.card-content {
  padding: 20px;
}

.card-actions-top {
  position: absolute;
  top: 12px;
  right: 12px;
  display: flex;
  gap: 8px;
  z-index: 3;
}

.action-icon {
  color: #cbd5e1;
  font-size: 16px;
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  transition: all 0.2s;
}

.action-icon:hover {
  color: #d50000;
  background: rgba(213, 0, 0, 0.1);
}

.action-icon:first-child {
  color: #40a9ff;
}

.action-icon:first-child:hover {
  color: #1890ff;
  background: rgba(24, 144, 255, 0.1);
}

.role-name {
  color: #e2e8f0;
  font-size: 18px;
  margin: 0 0 16px;
  font-family: 'Roboto Mono', monospace;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding-right: 60px;
}

.role-info {
  margin-bottom: 16px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 14px;
}

.info-item .label {
  color: #64748b;
}

.info-item .value {
  color: #cbd5e1;
  font-weight: 500;
}

.detail-btn {
  background: linear-gradient(120deg, #d50000, #7f1d1d);
  border: none;
  font-weight: 600;
  letter-spacing: 1px;
}

.detail-btn:hover {
  background: linear-gradient(120deg, #f43f5e, #991b1b);
}

.pagination-container {
    margin-top: 20px;
    text-align: right;
}

.empty-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 40px;
  color: #64748b;
}

/* 报告样式 */
:deep(.report-modal .ant-modal-content) {
  background: linear-gradient(135deg, #0f172a 0%, #1e293b 100%);
  border: 1px solid rgba(213, 0, 0, 0.3);
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.8);
}

:deep(.report-modal .ant-modal-header) {
  background: transparent;
  border-bottom: 1px solid rgba(213, 0, 0, 0.3);
  padding: 20px 24px;
}

:deep(.report-modal .ant-modal-title) {
  color: #e2e8f0;
  font-family: 'Roboto Mono', monospace;
  font-size: 18px;
  letter-spacing: 2px;
}

:deep(.report-modal .ant-modal-body) {
  padding: 24px;
  background: rgba(15, 23, 42, 0.8);
  color: #e2e8f0;
}

:deep(.report-modal .ant-modal-close) {
  color: #cbd5e1;
}

:deep(.report-modal .ant-modal-close:hover) {
  color: #d50000;
}

.report-header {
  text-align: center;
  margin-bottom: 32px;
  padding-bottom: 20px;
  border-bottom: 2px solid rgba(213, 0, 0, 0.5);
  position: relative;
}

.report-header::before {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 50%;
  transform: translateX(-50%);
  width: 60px;
  height: 2px;
  background: #d50000;
}

.report-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 24px;
  font-weight: 800;
  color: #d50000;
  letter-spacing: 3px;
  margin-bottom: 8px;
}

.report-subtitle {
  font-family: 'Roboto Mono', monospace;
  font-size: 12px;
  color: #d50000;
  letter-spacing: 4px;
  opacity: 0.8;
}

.report-section {
  margin-bottom: 32px;
  padding: 20px;
  background: rgba(30, 41, 59, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  border-left: 3px solid #d50000;
}

.section-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  font-weight: 700;
  color: #d50000;
  margin-bottom: 20px;
  letter-spacing: 2px;
  padding-bottom: 10px;
  border-bottom: 1px solid rgba(213, 0, 0, 0.3);
}

.report-form :deep(.ant-form-item-label > label) {
  color: #94a3b8;
  font-weight: 600;
  font-size: 13px;
  letter-spacing: 1px;
}

.report-input :deep(.ant-input),
.report-input :deep(.ant-input-number),
.report-input :deep(.ant-select-selector) {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #e2e8f0;
}

.report-input :deep(.ant-input:focus),
.report-input :deep(.ant-input-number:focus),
.report-input :deep(.ant-select-focused .ant-select-selector) {
  border-color: #d50000;
  box-shadow: 0 0 0 2px rgba(213, 0, 0, 0.2);
}

.report-info-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.report-info-item {
  padding: 12px;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 4px;
}

.info-label {
  font-size: 12px;
  color: #64748b;
  margin-bottom: 6px;
  letter-spacing: 1px;
  font-weight: 600;
}

.info-value {
  font-size: 16px;
  color: #e2e8f0;
  font-weight: 500;
  font-family: 'Roboto Mono', monospace;
}

.qa-fixed-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.qa-fixed-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 16px;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 4px;
  border-left: 3px solid #40a9ff;
  transition: all 0.2s;
}

.qa-fixed-item:hover {
  background: rgba(15, 23, 42, 0.8);
  border-left-color: #1890ff;
}

.qa-fixed-label {
  font-size: 14px;
  color: #cbd5e1;
  font-weight: 500;
  letter-spacing: 0.5px;
}

.qa-fixed-value {
  font-size: 18px;
  color: #40a9ff;
  font-weight: 700;
  font-family: 'Roboto Mono', monospace;
  min-width: 50px;
  text-align: right;
}

.report-detail {
  color: #e2e8f0;
}

.report-layout {
  display: flex;
  gap: 24px;
  height: 70vh;
  max-height: 800px;
}

.report-content {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
}

.report-content::-webkit-scrollbar {
  width: 6px;
}

.report-content::-webkit-scrollbar-thumb {
  background: rgba(213, 0, 0, 0.3);
  border-radius: 3px;
}

.report-content::-webkit-scrollbar-thumb:hover {
  background: rgba(213, 0, 0, 0.5);
}

.report-card {
  background: rgba(30, 41, 59, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  border-left: 3px solid #d50000;
  margin-bottom: 20px;
  transition: all 0.3s;
}

.report-card:hover {
  border-left-color: #f43f5e;
  box-shadow: 0 4px 12px rgba(213, 0, 0, 0.2);
}

.card-header {
  padding: 16px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  background: rgba(15, 23, 42, 0.3);
}

.card-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  font-weight: 700;
  color: #d50000;
  letter-spacing: 2px;
}

.card-body {
  padding: 20px;
}

.report-bookmarks {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 120px;
  flex-shrink: 0;
}

.bookmark-item {
  position: relative;
  padding: 12px 16px;
  background: rgba(30, 41, 59, 0.4);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-left: 3px solid transparent;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s;
  writing-mode: vertical-rl;
  text-orientation: mixed;
  height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.bookmark-item::before {
  content: '';
  position: absolute;
  right: -1px;
  top: 50%;
  transform: translateY(-50%);
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 8px 0 8px 8px;
  border-color: transparent transparent transparent rgba(213, 0, 0, 0.3);
  opacity: 0;
  transition: opacity 0.3s;
}

.bookmark-item:hover {
  background: rgba(30, 41, 59, 0.6);
  border-left-color: rgba(213, 0, 0, 0.5);
  transform: translateX(-4px);
}

.bookmark-item:hover::before {
  opacity: 1;
}

.bookmark-item.active {
  background: rgba(213, 0, 0, 0.15);
  border-left-color: #d50000;
  box-shadow: 0 2px 8px rgba(213, 0, 0, 0.3);
}

.bookmark-item.active::before {
  opacity: 1;
  border-color: transparent transparent transparent #d50000;
}

.bookmark-label {
  font-family: 'Roboto Mono', monospace;
  font-size: 13px;
  font-weight: 600;
  color: #cbd5e1;
  letter-spacing: 1px;
  transform: rotate(180deg);
}

.bookmark-item.active .bookmark-label {
  color: #d50000;
  font-weight: 700;
}

/* 轨迹显示样式 */
.track-display {
  margin-top: 24px;
}

.track-section {
  margin-bottom: 32px;
}

.track-title {
  font-family: 'Roboto Mono', monospace;
  font-size: 14px;
  font-weight: 800;
  letter-spacing: 1px;
  margin-bottom: 16px;
  text-transform: uppercase;
}

.track-title.blue-title {
  color: #3498db;
  border-bottom: 2px solid #3498db;
  padding-bottom: 8px;
}

.track-title.yellow-title {
  color: #f1c40f;
  border-bottom: 2px solid #f1c40f;
  padding-bottom: 8px;
}

.track-title.red-title {
  color: #e74c3c;
  border-bottom: 2px solid #e74c3c;
  padding-bottom: 8px;
}

.track-grid-30 {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 8px;
  margin-bottom: 20px;
}

.track-cell-wrapper {
  position: relative;
  aspect-ratio: 1;
}

.track-cell {
  width: 100%;
  height: 100%;
  aspect-ratio: 1;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: bold;
  color: #333;
  transition: all 0.2s;
  position: relative;
}

.track-cell.active {
  border-width: 2px;
}

.track-cell.blue-track.active {
  background: #2980b9;
  border-color: #2980b9;
  color: white;
}

.track-cell.yellow-track.active {
  background: #f39c12;
  border-color: #f39c12;
  color: white;
}

.track-cell.red-track.active {
  background: #c0392b;
  border-color: #c0392b;
  color: white;
}

.track-cell.has-file {
  border-style: dashed;
}

.file-label {
  font-size: 9px;
  font-weight: 700;
  text-align: center;
}

.track-file-card {
  position: absolute;
  top: calc(100% + 8px);
  left: 50%;
  transform: translateX(-50%);
  width: 120px;
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  padding: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 10;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.track-file-card::before {
  content: '';
  position: absolute;
  top: -6px;
  left: 50%;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 0 6px 6px 6px;
  border-color: transparent transparent white transparent;
}

.track-file-image {
  width: 100%;
  height: auto;
  border-radius: 4px;
  display: block;
}

/* 文件查看模态框样式 */
.file-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.9);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.file-modal-content {
  text-align: center;
}

.loading-container {
  position: relative;
  width: 200px;
  height: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

/* 三角形加载动画 */
.triangle-loader {
  width: 0;
  height: 0;
  border-left: 30px solid transparent;
  border-right: 30px solid transparent;
  border-bottom: 50px solid #d50000;
  animation: pulse 1.5s infinite;
  margin-bottom: 40px;
}

/* 脉冲动画 */
@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.2);
    opacity: 0.7;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

/* 扫描线动画 */
.scan-line {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 4px;
  background: linear-gradient(90deg, transparent, #ff0000, transparent);
  box-shadow: 0 0 10px #ff0000;
  animation: scan 2s linear 1s forwards;
  transform: translateY(-100%);
}

/* 扫描动画 */
@keyframes scan {
  0% {
    transform: translateY(-100%);
  }
  100% {
    transform: translateY(200px);
  }
}

/* 验证文本 */
.verification-text {
  color: #4ade80;
  font-family: 'Roboto Mono', monospace;
  font-size: 16px;
  margin-top: 60px;
  opacity: 0;
  animation: fadeIn 0.5s forwards 2.5s;
}

/* 淡入动画 */
@keyframes fadeIn {
  to {
    opacity: 1;
  }
}


</style>
