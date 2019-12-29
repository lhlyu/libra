<template>
    <div class="libra-card">
        <slot></slot>
    </div>
</template>

<script>
  export default {
    name: "Card"
  };
</script>

<style scoped lang="scss">
    .libra-card{
        background: rgba(#ffffff,0.8);
        height: 90px;
        min-width: 100%;
        border-radius: 10px;
        transition: all 0.3s linear;
        overflow: hidden;
        &:hover{
            transform: translateY(-5px);
            box-shadow: 0 0 3px rgba(#555555,0.9),8px 8px 8px rgba(#555555,0.6);
        }
    }
</style>
