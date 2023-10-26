<template>
    <header>
        <h2>NOME</h2>

        <div class="links">
            <p>Jogar</p>
            <p>Rankings</p>
            <p>Login</p>
        </div>
    </header>
</template>

<style scoped>
    header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
    }

    h2 {
        font-weight: 300;
    }

    header .links {
        display: flex;
        gap: 1rem;
        font-weight: 300;
    }
</style>