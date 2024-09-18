import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'

// 创还能一个 router
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            redirect: '/frontend',
        },
        //   前台页面
        {
            path: "/frontend",
            name: "frontend",
            component: () => import('@/views/frontend/FrontendLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'frontend-blogs',
                    component: () => import("@/views/frontend/BlogList.vue")
                },
                {
                    path: ':id',
                    name: 'blog-content',
                    component: () => import("@/views/frontend/BlogView.vue")
                },

            ]
        },
        //   后台页面
        {
            path: "/backend",
            name: "backend",
            component: () => import('@/views/backend/BackendLayout.vue'),
            children: [
                {
                    path: '',
                    name: 'statistics',
                    component: () => import("@/views/backend/Statistics.vue")

                },
                {
                    path: 'blogs',
                    name: 'backend-blogs',
                    component: () => import("@/views/backend/blogs/BlogList.vue")

                },
                {
                    path: 'tags',
                    name: 'backend-tags',
                    component: () => import("@/views/backend/tags/TagList.vue")

                },
                {
                    path: 'comments',
                    name: 'backend-comments',
                    component: () => import("@/views/backend/comments/CommentList.vue")
                }
            ]
        },
        {
            path: '/error/403',
            name: 'permission-deny',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('@/views/errors/PermissionDeny.vue')
        },
        {
            path: "/:pathMatch(.*)*",  // 匹配所有前面没有命中的路由
            name: "not-found",
            component: () => import('@/views/errors/NotFound.vue'),
        }
    ]
})

// 将 router 暴露出去
export default router

