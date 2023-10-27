<template>
    <h1 style="margin-bottom: 1rem;">Perfil</h1>

    <form @submit.prevent="updateProfile">
        <div>
            <label for="name">Nome</label>
            <input id="name" type="text" v-model="name">
        </div>

        <div>
            <label for="email">Email</label>
            <input id="email" type="email" v-model="email" readonly>
        </div>

        <div>
            <label for="oldPass">Senha antiga</label>
            <input v-model="oldPass" id="oldPass" type="password" placeholder="Digite aqui" required>
        </div>

        <div>
            <label for="newPass">Nova senha</label>
            <input v-model="newPass" id="newPass" type="password" placeholder="Digite aqui" required>
        </div>

        <button>Salvar</button>
    </form>
</template>

<script>
import { useToast } from 'vue-toastification';
import { POST_REQUEST } from '../utils/api';

const toast = useToast();

export default {
    data(){
        return {
            name: this.$store.state.username,
            email: this.$store.state.email,
            oldPass: "",
            newPass: "",
        };
    },

    methods: {
        async updateProfile(){
            const data = POST_REQUEST("users", "PUT", this.$store.state.token, {name: this.name, old_password: this.oldPass, new_password: this.newPass });
            const req = await fetch(data.url, data.options);
            const json = await req.json();

            if(!req.ok){
                toast.error("Algo deu errado!");
                return;
            }

            this.$store.commit('changeUsername', json.name);
            toast.success("Dados alterados com sucesso!");

        },
    }
};
</script>

<style scoped>

div {
    display: flex;
    flex-direction: column;
    width: 100%;
    justify-content: center;
    align-items: center;
    margin-bottom: 0.5rem;
}

div label {
    font-size: 1.2rem;
}
</style>