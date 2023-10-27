import { createStore } from 'vuex'

export default createStore({
  state: {
    token: null,
    username: null,
    email: null,
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
    
    logout(state){
      state.token = null;
      state.username = null;
      state.email = null;
      document.cookie=`token=; Path=/; Secure; SameSite=Strict`;
    },
  },
  actions: {
  },
  modules: {
  }
})