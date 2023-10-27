<template>
  <Header/>

  <div class="content">
    <RouterView />
  </div>
</template>

<script>
import Header from './components/Header.vue';
export default {
  components: {
    Header,
  },

  created(){
    this.$store.dispatch('checkAuth');
  }
};

</script>