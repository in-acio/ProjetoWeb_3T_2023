<template>
    <h1>Adicionar item</h1>

    <h1></h1>

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