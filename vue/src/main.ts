import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import router from './router'
// import store from './store'
import 'virtual:windi.css'

import 'ant-design-vue/es/message/style/css'

createApp(App)
    .use(router)
    // .use(store)
    .mount('#app')
