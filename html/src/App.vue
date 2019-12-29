<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style lang="scss">
  * {
    margin: 0;
    padding: 0;
  }
  body {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    /*color: #2c3e50;*/
    color: #ff00ff;
    background: url("https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/3.jpg");
    background-size: cover;
    backdrop-filter: blur(5px);
    &:after{
      content: '';
      position: fixed;
      top: 0;
      left: 0;
      height: 100%;
      width: 100%;
      background: url('assets/bg_cover.png') repeat top left;
      z-index: 0;
    }
  }

</style>
