<template>
    <h1>Admin</h1>

    <div style="display: flex; flex-direction: column; text-align: center; height: 100%; width: 100%;">
        <h2 style="margin-bottom: 1rem;">Listagem - <router-link to="/admin/addItem">Adicionar novo item</router-link></h2>

        <table style="margin: 0 auto;" v-if="items">
            <thead>
                <tr>
                    <th>Nome</th>
                    <th>Imagem</th>
                    <th>#</th>
                </tr>
            </thead>

            <tbody>
                <tr v-for="item in items">
                    <td>{{ item.name }}</td>
                    <td>
                        <img :src="`http://localhost:8080/images/${item.img}`" width="150" height="100" alt="">
                    </td>

                    <td>
                        <div class="icons">
                            <i @click="this.$router.push({ name: 'editItem', params: { id: item.id }})" style="margin-right: 1rem;" class='bx bx-edit-alt'></i>
                            <i @click="deleteItem(item.id)" class='bx bx-trash'></i>
                        </div>
                    </td>
                </tr>

                
            </tbody>
        </table>
    </div>
</template>

<script>
import { DELETE_REQUEST, makeRequest } from '../utils/api';

export default{
    data(){
        return {
            items: null,
        };
    },

    methods: {
        async deleteItem(idx) {
            const data = DELETE_REQUEST(`itens/${idx}`, this.$store.state.token);
            const req = await fetch(data.url, data.options);

            if(!req.ok){
                alert("ERRO");
                return;
            }

            this.items = this.items.filter(item => item.id != idx);
        }
    },

    async mounted(){
       let data = await makeRequest("itens", this.$store.state.token);
       this.items=data;
    },
};
</script>

<style scoped>

table {
  border-collapse:separate; 
  border-spacing: 0 1em;
}

tr, td {
    padding: 1rem 3rem;
}

.icons {
    gap: 1rem;
}
.icons i {
    font-size: 1.8rem;
    cursor: pointer;
}
</style>