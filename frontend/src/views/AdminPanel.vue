<template>
    <h1>Admin</h1>

    <div style="display: flex; flex-direction: column; height: 100%; width: 100%; text-align: center; padding: 2rem;">
        <h2 style="margin-bottom: 1rem;">Listagem - <router-link to="/admin/addItem">Adicionar novo item</router-link></h2>

        <table>
            <thead>
                <tr>
                    <th>Nome</th>
                    <th>Imagem</th>
                    <th>#</th>
                </tr>
            </thead>

            <tbody>
                <tr v-for="i in 20">
                    <td>Nome do item</td>
                    <td>
                        <img src="https://prefeitura.rio/wp-content/uploads/2021/12/WhatsApp-Image-2021-12-16-at-16.42.03.jpeg" width="100" height="100" alt="">
                    </td>

                    <td>
                        <div class="icons">
                            <i @click="this.$router.push({ name: 'editItem', params: { id: '123' }})" style="margin-right: 1rem;" class='bx bx-edit-alt'></i>
                            <i class='bx bx-trash' ></i>
                        </div>
                    </td>
                </tr>

                
            </tbody>
        </table>
    </div>
</template>

<style scoped>

table {
  border-collapse:separate; 
  border-spacing: 0 1em;
}

.icons {
    gap: 1rem;
}
.icons i {
    font-size: 1.8rem;
    cursor: pointer;
}
</style>