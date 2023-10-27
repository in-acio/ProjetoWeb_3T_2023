import { createRouter, createWebHistory } from 'vue-router'

import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';
import PlayView from '../views/PlayView.vue';
import RankingView from '../views/RankingView.vue';
import ProfileView from '../views/ProfileView.vue';
import HistoryView from '../views/HistoryView.vue';
import AdminPanel from '../views/AdminPanel.vue';
import AddItemView from '../views/AddItemView.vue';
import EditItemView from '../views/EditItemView.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/play',
      name: 'play',
      component: PlayView,
    },
    {
      path: '/rankings',
      name: 'rankings',
      component: RankingView,
    },
    {
      path: '/profile',
      name: 'profile',
      component: ProfileView,
    },
    {
      path: '/history',
      name: 'history',
      component: HistoryView,
    },
    {
      path: '/admin',
      children: [
        {
          path: '',
          component: AdminPanel,
        },
        {
          path: 'addItem',
          name: 'addItem',
          component: AddItemView,
        },
        {
          path: 'editItem/:id',
          name: 'editItem',
          component: EditItemView,
        }
      ],
    },
  ]
})

/*
router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)){
    if(store.state.token){
      next();
      return;
    }
    next('/');
  } else {
    next();
  }
})
*/

export default router
