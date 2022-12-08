import {createRouter, createWebHashHistory} from 'vue-router';
import Index from '../views/Index.vue'

const router = createRouter({
    // 指定路由的工作模式 hash
    history: createWebHashHistory(),
    routes: [
        { path: '/', component: Index },
        { path: '/home', component: Index },
        { path: '/category', component: Index },
        { path: '/tag', component: Index },
        { path: '/archive', component: Index },
        { path: '/about', component: Index },
        { path: '/link', component: Index },
    ],
})

export default router