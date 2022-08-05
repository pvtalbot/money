import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user.js'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/app',
      component: () => import('@/views/AppLayout.vue'),
      children: [
        {
          path: 'home',
          name: 'home',
          component: () => import('@/views/HomeView.vue'),
        },
        {
          path: 'graphs',
          name: 'graphs',
          component: () => import('@/views/GraphsView.vue'),
        }
      ],
    }
  ]
})


router.beforeEach((to, from, next) => {
  const userStore = useUserStore();
  if (to.name !== 'login' && !userStore.userLoggedIn) next({name: 'login'});
  else next();
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore();
  if (to.name === 'login' && userStore.userLoggedIn) next({name: 'home'})
  else next();
})

export default router
