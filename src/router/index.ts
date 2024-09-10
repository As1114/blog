import {createRouter, createWebHistory} from 'vue-router'
import NProgress from "nprogress";
import "nprogress/nprogress.css";


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
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
