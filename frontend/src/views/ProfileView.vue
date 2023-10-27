<template>
    <h1 style="margin-bottom: 1rem;">Perfil</h1>

    <form action="">
        <div>
            <label for="name">Nome</label>
            <input id="name" type="text" value="Nome do usuario">
        </div>

        <div>
            <label for="email">Email</label>
            <input id="email" type="email" value="email@emai.com" readonly>
        </div>

        <div>
            <label for="oldPass">Senha antiga</label>
            <input id="oldPass" type="password" placeholder="Digite aqui">
        </div>

        <div>
            <label for="newPass">Nova senha</label>
            <input id="newPass" type="password" placeholder="Digite aqui">
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