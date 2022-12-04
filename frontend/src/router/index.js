import {createRouter, createWebHashHistory} from 'vue-router';
import Home from '../views/Home.vue';
import Header from '../components/Header.vue';

const router = createRouter({
    // 指定路由的工作模式 hash
    history: createWebHashHistory(),
    routes: [
        { path: '/', redirect: '/home' },
        { path: '/home', component: Home },
        { path: '/category', component: Header },
        { path: '/tag', component: Home },
        { path: '/archive', component: Home },
        { path: '/about', component: Home },
        { path: '/link', component: Home },
    ],
})

export default router