<template>

    <div>
        {{ message }}
    </div>

</template>

<script>
    export default{
        data(){
            return {
                message: 'Hello Wooorld!!',
                items: ['one', 'two', 'three'],
            };
        },
        mounted(){
            console.log("MOUNTED EXMPLAE VUE FOR REAL!");

            this.items.forEach(item => {
                console.log(item);
            });
        }
      }
</script>