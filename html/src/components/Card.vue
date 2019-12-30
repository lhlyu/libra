<template>
    <div class="libra-card" @click="toHandler">
        <div class="libra-card-icon">
            <img :src="value">
        </div>
        <div class="libra-card-title">
            <slot></slot>
        </div>
    </div>
</template>

<script>

  import util from '@/util/util';

  export default {
    name: "Card",
    props:{
      img:{
        type: String,
        default: null
      },
      to: {
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
      },
      toHandler(){
        if (this.to) {
          window.open(this.to,"_blank")
        }
      }
    }
  };
</script>

<style scoped lang="scss">
    .libra-card{
        background: rgba(#ffffff,0);
        height: 90px;
        min-width: 100%;
        border-radius: 10px;
        transition-delay: 0.2s;
        transition: all 0.4s linear;
        overflow: hidden;
        display: flex;
        flex-direction: row;
        cursor: pointer;
        &:hover{
            background: rgba(#f7f1e3,0.8);
            .libra-card-icon {
                img {
                    border-radius: 10px;
                }
            }
        }
        .libra-card-icon{
            padding: 10px;
            img{
                height: 70px;
                width: 70px;
                border-radius: 50%;
                transition: all 0.4s linear;
            }
        }
        .libra-card-title{
            flex: 1;
            font-size: 18px;
            text-align: center;
            line-height: 90px;
            color: #20bf6b;
        }
    }
</style>
