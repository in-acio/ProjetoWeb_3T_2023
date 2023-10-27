<template>
    <header>
        <router-link to="/">NOME</router-link>

        <div class="links">
            <router-link to="/play">Jogar</router-link>
            <router-link to="/rankings">Rankings</router-link>
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

    a {
        font-size: 1.2rem;
        text-decoration: none;
        color: white;
        font-weight: 500;
    }
</style>