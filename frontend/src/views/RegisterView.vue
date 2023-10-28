<template>
    <h2>Registro</h2>

    <form @submit.prevent="register">
        <input type="text" name="name" placeholder="Nome" v-model="name" required>
        <input type="email" name="email" placeholder="Email" v-model="email" required>
        <input type="password" name="password" placeholder="Senha" v-model="password" required>

        <button>Registrar</button>
    </form>

    <p>Já possui uma conta? Clique <router-link to="/">aqui</router-link></p>
</template>

<script>
import { useToast } from 'vue-toastification';
import { USER } from '../utils/api';

const toast = useToast();

export default {
    data(){
        return {
            name: "",
            email: "",
            password: "",
        };
    },

    methods: {
        async register(){
            let data = USER("", {name: this.name, email: this.email, password: this.password})
            let req = await fetch(data.url, data.options);
            let json = await req.json();

            if(!req.ok){
                toast.error(json.error);
                return;
            }

            data = USER("/login", {email: this.email, password: this.password})
            req = await fetch(data.url, data.options);
            json = await req.json();

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