<template>
        <div v-for="(item, idx) in arr">
            <div v-motion-pop v-if="currentId == idx">
                <img :src="item.img" alt="">

                <div class="buttons">
                    <button class="btnNo">
                        <i class='bx bx-x'></i>
                    </button>

                    <button class="btnMeh">
                        <i class='bx bx-meh'></i>
                    </button>

                    <button @click="currentId++" class="btnYes">
                        <i class='bx bx-check'></i>
                    </button>
                </div>  
            </div>
        </div>

        <div class="done" v-if="currentId>=arr.length">
            <h2>Você já votou em todos os itens!</h2>
            <p>Para alterar algum voto, acesse o <router-link to="/history">histórico</router-link></p>
        </div>
</template>

<script>
export default {
    data() {
        return {
            currentId: 0,
            arr: [
                {
                    img: "https://cbissn.ibict.br/images/phocagallery/galeria2/thumbs/phoca_thumb_l_image03_grd.png",
                },
                {
                    img: "https://proex.ufpa.br/images/galeria_em_artigos/image04_grd.png",
                },
                {
                    img: "https://cbissn.ibict.br/images/phocagallery/galeria2/thumbs/phoca_thumb_l_image03_grd.png",
                },
                {
                    img: "https://cbissn.ibict.br/images/phocagallery/galeria2/thumbs/phoca_thumb_l_image03_grd.png",
                },
            ],
        };
    },
};
</script>

<style scoped>
.buttons {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 1rem;
}

.buttons button {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    border: 2px solid;
    background-color: transparent;
    transition: 0.2s all ease-in;
}

.buttons button:hover {
    color: white;
}

.btnYes {
    color: #5cb85c;
    border-color: #5cb85c;
}

.btnYes:hover { background-color: #5cb85c; border-color: #5cb85c; }

.btnMeh {
    color: #4cc9f0;
    border-color: #4cc9f0;
}

.btnMeh:hover { background-color: #4cc9f0; border-color: #4cc9f0; }

.btnNo {
    color: #ed4337;
    border-color: #ed4337;
}
.btnNo:hover { background-color: #ed4337; border-color: #ed4337; }


i {
    font-size: 2rem;
}

.done {
    text-align: center;
}

.done h2 {
    margin-bottom: 1rem;
}

.done a {
    color: white;
}


</style>