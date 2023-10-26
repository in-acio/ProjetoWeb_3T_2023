<template>
        <h2>Login</h2>

        <form action="">
            <input type="email" name="email" placeholder="Email" required>
            <input type="password" name="password" placeholder="Senha" required>

            <button>Entrar</button>
        </form>
</template>

<style scoped>
    h2 {
        margin-bottom: 0.5rem;
    }

    form {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        gap: 1rem;
        width: 100%;
    }

    input {
        outline: none;
        border: 1px solid black;
        width: 20%;
        padding: 0.4rem;
        border-radius: 0.2rem;
    }

    button {
        width: 10%;
        padding: 0.2rem;
    }
</style>