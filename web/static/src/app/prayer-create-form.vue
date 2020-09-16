<template>
    <form>
    <fieldset>
        <label for="commentField">Gebet</label>
        <textarea placeholder="Gebetsanliegen" id="commentField" v-model="prayer"></textarea>
        
        <div class="float-left">
            <input type="checkbox" id="confirmField" v-model="isPublic">
            <label class="label-inline" for="confirmField">Öffentlich beten</label>
        </div>
        <br>
    </fieldset>
    <a class="button-primary button" v-on:click="sendForm">Senden</a>
    </form>

</template>

<script>
    export default{
        data(){
            return {
                prayer: null,
                isPublic: true
            };
        },
        mounted(){

        },
        methods: {
            sendForm: function(){
                if(this.prayer == null || this.prayer == ""){
                    this.$eventHub.$emit('notification-show', {
                        message: "Bitte Gebetsanliegen ausfüllen."
                    });
                    return;
                }


                this.$ajax.post('/api', {
                    prayer: this.prayer,
                    is_public: this.isPublic,
                }).then( response => {
                    
                    this.$eventHub.$emit('notification-show', {
                        message: "Gebetsanliegen wurde abgesendet."
                    });

                    this.prayer = null;
                    this.isPublic = true;

                }).error( error => {

                    this.$eventHub.$emit('notification-show', {
                        message: "Fehler beim Absenden."
                    });

                });
            }
        }
      }
</script>