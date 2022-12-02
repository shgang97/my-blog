import { createApp } from 'vue'
import App from './App.vue'
import './index.css'

import axios from 'axios'

const app = createApp(App)

axios.defaults.baseURL = 'http://localhost:8080/api/blog'
app.config.globalProperties.$http = axios

app.mount('#app')
