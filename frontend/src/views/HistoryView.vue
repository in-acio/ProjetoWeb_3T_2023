<template>
    <h1>Historico</h1>

    <div style="display: flex; flex-direction: column; text-align: center; width: 100%; height: 100%;">
        <table style="margin: 0 auto;" v-if="votes">
                <thead>
                    <tr>
                        <th>Nome</th>
                        <th>Imagem</th>
                        <th>#</th>
                    </tr>
                </thead>

                <tbody>
                    <tr v-for="vote in votes">
                        <td>{{ vote.name }}</td>
                        <td>
                            <img :src="`http://localhost:8080/images/${vote.img}`" width="150" height="100" alt="">
                        </td>

                        <td>
                            <div class="buttons">
                                <button @click="vote.Value = 0; changeVote(vote.ItemID, 0)" :class="{'btnNoC': vote.Value==0 }" class="btnNo">
                                    <i :class="{'whiteText': vote.Value==0 }" class='bx bx-x'></i>
                                </button>

                                <button @click="vote.Value = 1; changeVote(vote.ItemID, 1)" :class="{'btnMehC': vote.Value==1 }" class="btnMeh">
                                    <i :class="{'whiteText': vote.Value==1 }" class='bx bx-meh'></i>
                                </button>

                                <button @click="vote.Value = 2; changeVote(vote.ItemID, 2)" :class="{'btnYesC': vote.Value==2 }" class="btnYes">
                                    <i :class="{'whiteText': vote.Value==2 }" class='bx bx-check'></i>
                                </button>
                            </div>  
                        </td>
                    </tr>

                    
                </tbody>
            </table>
        </div>
</template>

<script>
import { useToast } from 'vue-toastification';
import { GET_REQUEST, POST_REQUEST } from '../utils/api';

const toast = useToast();

export default{
    data(){
        return {
            votes: null,
        };
    },

    methods: {
        async changeVote(id, value){
            try {
                const data = POST_REQUEST('vote', 'POST', this.$store.state.token, { item_id: id, value: value });
                const req = await fetch(data.url, data.options);

                if(!req.ok){
                    throw new Error("Erro na requisição");
                }
            } catch(err) {
                toast.error(err)
            }
        }
    },

    async mounted(){
        try {
            const data = GET_REQUEST('vote', this.$store.state.token);
            const req = await fetch(data.url, data.options);
            const json = await req.json();

            if(!req.ok){
                throw new Error("Erro na requisição");
            }

            this.votes = json;
        } catch(err) {
            toast.error(err)
        }
        
    },
};
</script>

<style scoped>
table {
  border-collapse:separate; 
  border-spacing: 0 1em;
}

tr, td {
    padding: 1rem 3rem;
}

.whiteText {
    color: white;
}

.buttons button {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    border: 2px solid;
    background-color: transparent;
    transition: 0.2s all ease-in;
}

.buttons button i {
    font-size: 1.5rem;
}

.btnYes {
    color: #5cb85c;
    border-color: #5cb85c;
}

.btnYesC{ background-color: #5cb85c !important; border-color: #5cb85c; }

.btnMeh {
    color: #4cc9f0;
    border-color: #4cc9f0;
    margin: 0 0.5rem;
}

.btnMehC { background-color: #4cc9f0 !important; border-color: #4cc9f0; }

.btnNo {
    color: #ed4337;
    border-color: #ed4337;
}
.btnNoC { background-color: #ed4337 !important; border-color: #ed4337;}

</style>