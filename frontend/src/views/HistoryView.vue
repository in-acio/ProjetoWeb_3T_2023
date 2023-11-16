<template>
    <h1>Historico</h1>
</template>

<script>
import { useToast } from 'vue-toastification';
import { GET_REQUEST } from '../utils/api';

const toast = useToast();

export default{
    data(){
        return {
            votes: null,
        };
    },

    async mounted(){
        try {
            const data = GET_REQUEST('vote', this.$store.state.token);
            const req = await fetch(data.url, data.options);
            const json = await req.json();

            if(!req.ok){
                throw new Error("Erro na requisição");
            }

            console.log(json);
            this.votes = json;
        } catch(err) {
            toast.error(err)
        }
        
    },
};
</script>