import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Posts from '../views/Posts.vue'
import PostEdit from '../views/PostEdit';
import PostDetail from '../views/PostDetail';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Index',
    redirect: {name: "Posts"}
  },
  {
    path: '/posts',
    name: 'Posts',
    component: Posts
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/post/add',
    name: 'PostAdd',
    component: PostEdit,
    meta: {
      requireAuth: true
    }
  },
  {
    path: '/post/:postId',
    name: 'PostDetail',
    component: PostDetail
  },
  {
    path: '/post/:postId/edit',
    name: 'PostEdit',
    component: PostEdit,
    meta: {
      requireAuth: true
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
