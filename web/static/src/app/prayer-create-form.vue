<template>
    <form>
    <fieldset>
        <label for="commentField" class="textarea-title">Gebet</label>
        <textarea placeholder="" class="textarea" id="commentField" v-model="prayer"></textarea>
        
        <div class="public-option-group">
            <label class="switch">
                <input type="checkbox" id="confirmField" v-model="isPublic">
                <span class="slider round"></span>
            </label>
            <label class="label-inline public-option-text" for="confirmField">Öffentlich beten</label>
        </div>
        <br>
    </fieldset>
    <div class="button-senden-group">
        <a class="button-primary button button-senden" v-on:click="sendForm">Senden</a>
    </div>
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