<template>

    <div v-if="visible">
        <div id="notification-background" v-on:click="closeNotification"></div>
        <div id="notification-body">
            {{ message }}
        </div>
    </div>

</template>

<script>
    export default{
        data(){
            return {
                message: 'Hello Wooorld!!',
                visible: false,
            };
        },
        mounted(){

            this.$eventHub.$on('notification-show', ev => {
                this.message = ev.message;
                this.visible = true;
            });

            this.$eventHub.$on('notification-hide', ev => {
                this.visible = false;
            });
        },
        methods: {
            closeNotification: function(){
                this.visible = false;
            }
        }
      }
</script>