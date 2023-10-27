<template>
        <h2>Login</h2>

        <form @submit.prevent="login">
            <input type="email" name="email" placeholder="Email" v-model="email" required>
            <input type="password" name="password" placeholder="Senha" v-model="password" required>

            <button>Entrar</button>
        </form>

        <p>Ainda não possui uma conta? Clique <router-link to="/register">aqui</router-link></p>
</template>

<script>
import { useToast } from 'vue-toastification';
import { USER } from '../utils/api';

const toast = useToast();

export default {
    inject: ['$bus'],

    data(){
        return {
            email: "",
            password: "",
        };
    },
    
    methods: {
        async login(){
            const data = USER("/login", {email: this.email, password: this.password})
            const req = await fetch(data.url, data.options);
            const json = await req.json();

            if(!req.ok){
                toast.error(json.error);
                return;
            }
            
            this.$store.commit('changeToken', json.token);
            this.$store.commit('changeEmail', json.email);
            this.$store.commit('changeUsername', json.name);
            this.$store.commit('changeIsAdmin', json.isAdmin);
            document.cookie=`token=${json.token}; Path=/; Secure; SameSite=Strict`;

            this.$bus.$emit("login", true);
            this.$router.push('/play');
        },
    }
};
</script>

<style scoped>
    h2 {
        margin-bottom: 0.5rem;
    }

    p {
        margin-top: 1rem;
    }
</style>