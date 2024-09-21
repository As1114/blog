import {createRouter, createWebHistory} from 'vue-router'
import "nprogress/nprogress.css";
import NProgress from "nprogress";
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";


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
                {
                    path: "article/:id",
                    name: "article",
                    component: () => import('../views/web/article.vue'),
                },
                {
                    path: "article/404",
                    name: "article_notfound",
                    component: () => import('../views/web/article_notfound.vue'),
                }
            ]
        },
        {
            path: "/admin",
            name: "admin_index",
            meta: {
                title: "首页",
                isAdmin: true,
            },
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
    const store = useStore()
    const meta = to.meta
    if (store.userStoreInfo.role != 1 && (meta.isAdmin)) {
        Message.warning("权限不足")
        router.push({name: "web_home"})
        return
    }
    NProgress.start()
    next()
})

router.afterEach(() => {
    NProgress.done()
})
