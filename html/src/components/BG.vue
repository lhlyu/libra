<template>
    <div class="libra-bg">
        <img class="libra-img" :src="value" v-if="value">
    </div>
</template>

<script>

  import util from '@/util/util';

  export default {
    name: "BG",
    props:{
      img:{
        type: String,
        default: null
      }
    },
    data(){
      return {
        value: null
      }
    },
    async mounted(){
      this.loadImg()
    },
    methods:{
      loadImg(){
        if(this.img){
          let _this = this
          util.loadImageAsync(this.img).then(v =>{
            _this.value = v.src
          })
        }
      }
    }
  };
</script>

<style scoped lang="scss">
    .libra-bg{
        position: relative;
        &:after{
            content: '';
            position: fixed;
            top: 0;
            left: 0;
            height: 100%;
            width: 100%;
            background: rgba(#000000,0.5);
            /*background: url('../assets/bg_cover.png') repeat top left;*/
            /*backdrop-filter: blur(5px);*/
            z-index: 0;
        }
        .libra-img{
            position: fixed;
            height: 100%;
            width: 100%;
            object-fit: cover;

        }
    }

</style>
