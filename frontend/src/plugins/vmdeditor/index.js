import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
import createTipPlugin from '@kangc/v-md-editor/lib/plugins/tip/index';

// 引入所有语言包
import hljs from 'highlight.js';

VueMarkdownEditor.use(githubTheme, {
    Hljs: hljs,
});
VueMarkdownEditor.use(createTipPlugin());

export {VueMarkdownEditor}