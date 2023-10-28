import { createStore } from 'vuex'

import { makeRequest } from '../utils/api';
import router from '../router';
import bus from '../utils/Events';

export default createStore({
  state: {
    token: null,
    username: null,
    email: null,
    isAdmin: null,
  },
  getters: {
  },
  mutations: {
    changeToken(state, newToken){
      state.token = newToken;
    },

    changeUsername(state, username){
      state.username = username;
    },

    changeEmail(state, email){
        state.email = email;
    },

    changeIsAdmin(state, isAdmin){
      state.isAdmin = isAdmin;
    },
    
    logout(state){
      state.token = null;
      state.username = null;
      state.email = null;
      state.isAdmin = null;
      document.cookie=`token=; Path=/; Secure; SameSite=Strict`;
    },
    
  },
  actions: {
    async checkAuth({ commit }) {
      const token = document.cookie.replace(/(?:(?:^|.*;\s*)token\s*=\s*([^;]*).*$)|^.*$/, "$1");
      const path = window.location.pathname;

      if (token) {
        try {
          
          const req = await makeRequest('users/session', token);
          commit('changeEmail', req.email);
          commit('changeUsername', req.name);
          commit('changeToken', token);
          commit('changeIsAdmin', req.isAdmin);

          bus.$emit("login", true);
          router.push(path);
        } catch(error){
          commit('logout');
          router.push("/");
        }
      }
    },
  },
  modules: {
  }
})