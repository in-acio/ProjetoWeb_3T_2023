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
import store from '../store';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView,
      meta: {
        requiresGuest: true,
      },
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
      meta: {
        requiresGuest: true,
      },
    },
    {
      path: '/play',
      name: 'play',
      component: PlayView,
      meta: {
        requiresAuth: true,
      },
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
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/history',
      name: 'history',
      component: HistoryView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/admin',
      meta: {
        requiresAuth: true,
        requiresAdmin: true,
      },
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

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some((x) => x.meta.requiresAuth);
  const requiresGuest = to.matched.some((x) => x.meta.requiresGuest);
  const requiresAdmin = to.matched.some((x) => x.meta.requiresAdmin);

  const isLoggedin = store.state.token;

  if (requiresAuth && !isLoggedin) {
    next("/");
  } else if (requiresGuest && isLoggedin) {
    next("/play");
  } else if (requiresAdmin && !store.state.isAdmin) {
    next("/");
  } else {
    next();
  }
});

export default router
