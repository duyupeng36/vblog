import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'

// 创还能一个 router
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        //   前台页面
        {
            path: "/frontend",
            name: "FrontendLayout",
            component: () => import('@/views/frontend/FrontendLayout.vue'),
            children: [
                {
                    path: 'blogs/list',
                    name: 'BlogsList',
                    component: () => import("@/views/frontend/BlosList.vue")

                }
            ]
        },
        //   后台页面
        {
            path: "/backend",
            name: "BackendLayout",
            component: () => import('@/views/backend/BackendLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'Statistics',
                    component: () => import("@/views/backend/Statistics.vue")

                },
                {
                    path: 'blogs',
                    name: 'BlogsList',
                    component: () => import("@/views/backend/blogs/BlogList.vue")

                },
                {
                    path: 'tags',
                    name: 'TagList',
                    component: () => import("@/views/backend/tags/TagList.vue")

                },
                {
                    path: 'comments',
                    name: 'CommentList',
                    component: () => import("@/views/backend/comments/CommentList.vue")
                }
            ]
        },
        {
            path: '/error/403',
            name: 'PermissionDeny',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('@/views/errors/PermissionDeny.vue')
        },
        {
            path: "/:pathMatch(.*)*",  // 匹配所有前面没有命中的路由
            name: "NotFound",
            component: () => import('@/views/errors/NotFound.vue'),
        }
    ]
})

// 将 router 暴露出去
export default router

