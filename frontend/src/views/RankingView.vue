<template>

    <div class="rankingContainer">
        <div v-for="(rank,idx) in ranking" class="rankingBox">
            <h1>{{ title[idx] }}</h1>
            <button @click="reorder(idx)">Inverter</button>

            <table v-if="ranking">
                <thead>
                    <tr>
                        <th>Lugar</th>
                        <th>Nome</th>
                        <th>Votos</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(item,idx) in rank">
                        <td>{{idx+1}}</td>
                        <td>{{item.name}}</td>
                        <td>{{item.vote_count}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

</template>

<script>
import { GET_REQUEST } from '../utils/api';

export default{
    data(){
        return {
            ranking: null,
            title: ['Menos queridos', 'Mais neutros', 'Mais queridos'],
        };
    },

    async mounted(){
        const data = GET_REQUEST("itens/ranking", "");
        const req = await fetch(data.url, data.options);
        const json = await req.json();

        this.ranking = json;
    },

    methods: {
        reorder(idx){
            this.ranking[idx] = this.ranking[idx].reverse();
        },
    },
};
</script>

<style scoped>
    .rankingContainer {
        display: flex;
        justify-content: space-around;
        align-items: flex-start;
        flex-wrap: wrap;
        height: 100%;
        width: 100%;
        margin-top: 5rem;
        padding: 2rem;
    }

    .rankingBox {
        text-align: center;
        margin-bottom: 2rem;
    }

    button {
        width: 50%;
        padding: 0.2rem;
    }

    table {
        text-align: center;
        width: 100%;
        font-size: 1.2rem;
        border-collapse: collapse;
        padding: 1rem;
        padding: 2rem;
        border-radius: 1rem;
    }

    tbody tr:nth-child(1) {
        background-color: #ffd700;
        color: black;
        font-weight: bold;
    }

    tbody tr:nth-child(2) {
        background-color: #c0c0c0;
        color: black;
        font-weight: bold;
    }

    tbody tr:nth-child(3) {
        background-color: #cd7f32;
        color: black;
        font-weight: bold;
    }

    tr {
        margin: 1rem;
    }

    @media (max-width: 720px) {
        .rankingContainer {
            display: block;
        }
    }
</style>