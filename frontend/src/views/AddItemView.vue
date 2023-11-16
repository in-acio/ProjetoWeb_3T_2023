<template>
    <h1>Adicionar item</h1>

    <h1></h1>

    <form @submit.prevent="saveItem">
        <div>
            <label for="name">Nome</label>
            <input v-model="name" required id="name" type="text" name="name" placeholder="Nome do item">
        </div>

        <div>
            <label for="img">Imagem</label>
            <input ref="file" id="img" type="file" name="file" required accept="image/*">
        </div>

        <button>Salvar</button>
    </form>
</template>

<script>
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
    data(){
        return {
            name: "",
            file: ""
        };
    },
    methods: {
        async saveItem() {
            let dataForm = new FormData();
            for (let file of this.$refs.file.files) {
                dataForm.append(`file`, file);
            }
            dataForm.append("name", this.name);

            try {
                let res = await fetch(`http://localhost:8080/itens`, {
                    method: 'POST',
                    body: dataForm,
                    headers: {
                        'Authorization': `Bearer ${this.$store.state.token}`
                    },
                });
                
                if(res.status == 201) {
                    toast.success("Item adicionado!");
                    this.name = "";
                } else {
                    const json = await res.json();
                    toast.error(json.error);
                }
            } catch(err) {
                toast.error("Erro!");
            }
        },
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