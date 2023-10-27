<template>
    <h1>Editar item</h1>

    <h1>{{ this.$route.params.id }}</h1>

    <form action="">
        <div>
            <label for="name">Nome</label>
            <input id="name" type="text" placeholder="Nome do item">
        </div>

        <div>
            <label for="img">Imagem</label>
            <input id="img" type="file">
        </div>

        <button>Salvar</button>
    </form>
</template>

<script>
export default {
    data(){
    },

    mounted(){

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