import {createRouter, createWebHashHistory} from 'vue-router';
import Index from '../views/Index.vue'
import ArticleDetail from "../components/ArticleDetail.vue";
import Writing from '../views/Writing.vue';
import Login from '../views/Login.vue';

const router = createRouter({
    // 指定路由的工作模式 hash
    history: createWebHashHistory(),
    routes: [
        { path: '/', redirect: '/home' },
        { path: '/home', component: Index },
        { path: '/category', component: Index },
        { path: '/tag', component: Index },
        { path: '/archive', component: Index },
        { path: '/about', component: Index },
        { path: '/link', component: Index },
        { path: '/writing', component: Writing },
        { path: '/writing/:id', component: Writing },
        { path: '/login', component: Login },
        { path: '/article/:id', component: ArticleDetail, name: 'article' },
        // { path: '/tag/:id', component: ArticleDetail, name: 'article' },
        // { path: '/category/:id', component: ArticleDetail, name: 'article' },
    ],
})

export default router