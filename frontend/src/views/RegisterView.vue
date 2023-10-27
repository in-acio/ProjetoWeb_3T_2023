<template>
    <h2>Registro</h2>

    <form action="">
        <input type="text" name="name" placeholder="Nome" required>
        <input type="email" name="email" placeholder="Email" required>
        <input type="password" name="password" placeholder="Senha" required>

        <button>Registrar</button>
    </form>

    <p>Já possui uma conta? Clique <router-link to="/">aqui</router-link></p>
</template>

<style scoped>
h2 {
    margin-bottom: 0.5rem;
}

p {
    margin-top: 1rem;
}
</style>