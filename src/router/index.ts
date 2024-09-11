import {createRouter, createWebHistory} from 'vue-router'
import NProgress from "nprogress";
import "nprogress/nprogress.css";


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
