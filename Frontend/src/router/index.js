import {createRouter, createWebHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Layout from '../components/Layout.vue'
import Home from '../views/Home.vue'
import Main from '../views/Main.vue'
import User from '../views/User.vue'
import Profile from '../views/Profile.vue'
import Role from "../views/Role.vue";
import FileView from "../views/FileView.vue";
import Forum from '../views/Forum.vue'
import ChatRoom from '../views/ChatRoom.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'login',
            component: Login,
        },
        {
            path: '/file-view',
            name: 'file-view',
            component: FileView
        },
        {
            path: '/',
            component: Layout,
            children: [
                {
                    path: 'main',
                    name: 'main',
                    component: Home,
                },
                {
                    path: 'rooms',
                    name: 'rooms',
                    component: Main,
                },
                {
                    path: 'users',
                    name: 'users',
                    component: User,
                },
                {
                    path: 'profile',
                    name: 'profile',
                    component: Profile,
                },  
                {
                    path: 'roles',
                    name: 'roles',
                    component: Role
                },
                // 评论
                {
                    path: 'forum',
                    name: 'forum',
                    component: Forum
                },
                // 聊天室
                {
                    path: 'chat/:roomId',
                    name: 'chat-room',
                    component: ChatRoom
                }
            ]
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/',
        },
    ],
})

// 简单路由守卫
router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('ta_token')
    // 文件查看页面不需要token验证
    if (to.name === 'file-view') {
        next()
    } else if (to.name !== 'login' && !token) {
        next({name: 'login'})
    } else if (to.name === 'users') {
        // 人员档案仅 admin 可见
        const userType = localStorage.getItem('ta_user_type')
        if (userType !== 'admin') {
            next({name: 'main'})
        } else {
            next()
        }
    } else {
        next()
    }
})

export default router
