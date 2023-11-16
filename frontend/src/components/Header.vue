<template>
    <header>
        <router-link to="/">NOME</router-link>

        <div class="links">
            <router-link to="/play">Jogar</router-link>
            <router-link to="/rankings">Rankings</router-link>
            <router-link v-if="userLoggedIn" to="/history">Histórico</router-link>
            <router-link v-if="userLoggedIn" to="/profile">Perfil</router-link>
            <router-link v-if="userLoggedIn && isAdmin" to="/admin">Admin</router-link>
            <a style="cursor: pointer;" v-if="userLoggedIn" @click.prevent="logout">Logout</a>
        </div>
    </header>
</template>

<script>

export default {
    inject: ['$bus'],

   data(){
        return {
            userLoggedIn: false,
            isAdmin: false,
        };
   },

   methods: {
    logout(){
        this.$store.commit('logout');
        window.location = "/";
    },
   },

   created(){
        this.$bus.$on('login', (v) => {
            this.userLoggedIn=true;
            this.isAdmin=v;
        });
   },
};
</script>

<style scoped>
    header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
    }

    h2 {
        font-weight: 300;
    }

    header .links {
        display: flex;
        gap: 1rem;
        font-weight: 300;
    }

    a {
        font-size: 1.2rem;
        text-decoration: none;
        color: white;
        font-weight: 500;
    }
</style>