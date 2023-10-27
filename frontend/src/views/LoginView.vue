
<template>
        <h2>Login</h2>

        <form action="">
            <input type="email" name="email" placeholder="Email" required>
            <input type="password" name="password" placeholder="Senha" required>

            <button>Entrar</button>
        </form>

        <p>Ainda não possui uma conta? Clique <router-link to="/register">aqui</router-link></p>
</template>

<style scoped>
    h2 {
        margin-bottom: 0.5rem;
    }

    p {
        margin-top: 1rem;
    }
</style>