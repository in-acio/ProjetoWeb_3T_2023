<template>
    <h1>Editar item</h1>

    <form @submit.prevent="editItem" v-if="item">
        <div>
            <label for="name">Nome</label>
            <input id="name" type="text" v-model="item.name">
        </div>

        <div>
            <label for="img">Imagem</label>
            <img width="250" height="200" :src="`http://localhost:8080/images/${item.img}`" alt="">
            <input id="img" type="file" ref="file">
        </div>

        <button>Salvar</button>
    </form>
</template>

<script>
import { useToast } from 'vue-toastification';
import { makeRequest } from '../utils/api';

const toast = useToast();

export default {
    data(){
        return {
            item: null,
        };
    },

    methods: {
        async editItem(){
            let dataForm = new FormData();
            for (let file of this.$refs.file.files) {
                dataForm.append(`file`, file);
            }
            dataForm.append("name", this.item.name);

            let res = await fetch(`http://localhost:8080/itens/${this.$route.params.id}`, {
                method: 'PUT',
                body: dataForm,
                headers: {
                    'Authorization': `Bearer ${this.$store.state.token}`
                },
            });
            
            if(res.status == 200) {
                toast.success("Item atualizado!");
            } else {
                toast.error("Algo deu errado!");
            }
        },
    },

    async mounted(){
        try {
            const data = await makeRequest(`itens/${this.$route.params.id}`, this.$store.state.token);
            this.item = data;
        } catch(err) {
            window.location.href="/";
        }
    },
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