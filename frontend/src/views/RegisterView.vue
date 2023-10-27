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
            const data = USER("", {name: this.name, email: this.email, password: this.password})
            const req = await fetch(data.url, data.options);
            const json = await req.json();

            if(!req.ok){
                toast.error(json.error);
                return;
            }

            
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