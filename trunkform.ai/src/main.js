import { createApp } from 'vue'
import index from './app.vue'
import router from './router'
import './style.css'

const app = createApp(index)

app.use(router)
app.mount('#app')
