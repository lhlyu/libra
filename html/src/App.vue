<template>
  <div id="app">
    <BG img="https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/6.jpg"></BG>
    <router-view />
  </div>
</template>

<script>

  import BG from '@/components/BG';

  export default {
    name: "Index",
    components:{
      BG
    }
  };
</script>

<style lang="scss">
  * {
    margin: 0;
    padding: 0;
  }
  body {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;

  }

</style>
