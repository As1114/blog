import {createRouter, createWebHistory} from 'vue-router'
import "nprogress/nprogress.css";
import NProgress from "nprogress";


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "web_index",
            component: () => import('../views/web/index.vue'),
            children: [
                {
                    path: "",
                    name: "web_home",
                    component: () => import('../views/web/home.vue'),
                },
            ]
        },
        {
            path: "/admin",
            name: "admin_index",
            component: () => import('../views/admin/index.vue'),
            children: [
                {
                    path: "",
                    name: "admin_home",
                    component: () => import('../views/admin/home/home.vue'),
                },
            ]
        }
    ]
})

export default router

router.beforeEach((to, from, next) => {
    NProgress.start()
    next()
})

router.afterEach(() => {
    NProgress.done()
})
