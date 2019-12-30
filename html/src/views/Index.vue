<template>
    <div class="libra-index">
        <div class="libra-cards">
            <Card v-for="count in 90"
                  :key="count"
                  :img="`http://www.thiswaifudoesnotexist.net/example-${count * 1000}.jpg`"
                  to="http://www.baidu.com">
                测试测试
            </Card>
        </div>
    </div>
</template>

<script>

    import Card from '@/components/Card';

  export default {
    name: "Index",
    components:{
      Card
    }
  };
</script>

<style scoped lang="scss">
    .libra-index{
        position: relative;
        z-index: 10;
        padding: 20px;
        .libra-cards{
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
            grid-row-gap: 40px;
            grid-column-gap: 40px;
            justify-self: center;
            z-index: 1;
            img{
                width: 90px;
                height: 90px;
                border-radius: 50%;
                transition: border-radius 0.3s ease;
                &:hover{
                    border-radius: 10px;
                }
            }
        }
    }
</style>
