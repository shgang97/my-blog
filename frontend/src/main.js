import {createApp} from 'vue';
import App from './App.vue';
import axios from 'axios';

// 配置请求根路径
axios.defaults.baseURL = 'http://localhost:8081';


const app = createApp(App);
// 将 axios 作为全局的自定义属性，每个组件可以在内部直接访问
app.config.globalProperties.axios = axios;
app.mount('#app');
