import {createRouter, createWebHashHistory} from 'vue-router';
import Index from '../views/Index.vue'
import ArticleDetail from "../views/ArticleDetail.vue";
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
        { path: '/writing', name: 'writing', component: Writing },
        { path: '/writing/:id', component: Writing },
        { path: '/login', name: 'login', component: Login },
        { path: '/article/:id', component: ArticleDetail, name: 'article' },
        // { path: '/tag/:id', component: ArticleDetail, name: 'article' },
        // { path: '/category/:id', component: ArticleDetail, name: 'article' },
    ],
})

router.beforeEach((to, from, next) => {
    console.log(to)
    console.log(sessionStorage.getItem("userInfo"))
    if (to.name === 'writing' && !sessionStorage.getItem("userInfo")) {
        next({name: 'login'})
    } else {
        next()
    }
})

export default router