import { createApp, provide, h } from 'vue'
import { createPinia } from 'pinia'

import { DefaultApolloClient } from '@vue/apollo-composable'
import { createApolloClient } from '@/apollo/apollo-config.js'

import App from './App.vue'
import router from './router'

import './assets/main.css'

const pinia = createPinia();

const app = createApp({
  setup() {
    provide(DefaultApolloClient, createApolloClient())
  },

  render: () => h(App),
})

app.use(pinia)
app.use(router)

app.mount('#app')
