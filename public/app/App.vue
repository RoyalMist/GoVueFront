<template>
    <span id="hello">{{who}}</span>
</template>

<script>
    export default {
        name: "App",
        data() {
            return {
                who: 'Go & Vue'
            }
        }
    }
</script>

<style scoped lang="stylus">
    #hello
        color cadetblue
</style>
