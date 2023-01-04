import {createApp} from 'vue';
import App from './App.vue';
import './index.css';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import router from './router/index.js';
import '../public/style/iconfont/iconfont.css';
import store from './store';
import axios from 'axios';
import {VueMarkdownEditor}  from './plugins/vmdeditor/index';
import dateformat from "./dateformat";

const app = createApp(App);

axios.defaults.baseURL = 'http://shgang.cn:8080/api/blog';
app.config.globalProperties.$http = axios;
app.config.globalProperties.$dataFormat = dateformat;

app.use(ElementPlus);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

app.use(router);
app.use(store);
app.use(VueMarkdownEditor);
app.mount('#app');
