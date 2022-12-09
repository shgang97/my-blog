import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router/index.js'
import '../public/style/iconfont/iconfont.css'


import axios from 'axios'

const app = createApp(App)

axios.defaults.baseURL = 'http://localhost:8080/api/blog'
app.config.globalProperties.$http = axios

app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(router)

app.mount('#app')
