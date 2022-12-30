import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import Home from "../views/home.vue";

const routes:RouteRecordRaw[] = [
    {
        path:"/",
        component: Home
    }
]
// 创建路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes //路由表
})
export default router