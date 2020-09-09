<template>
    <form>
    <fieldset>
        <label for="commentField">Gebet</label>
        <textarea placeholder="Gebetsanliegen" id="commentField" v-model="prayer"></textarea>
        
        <div class="float-left">
            <input type="checkbox" id="confirmField" v-model="isPublic">
            <label class="label-inline" for="confirmField">Öffentlich (anonym) beten</label>
        </div>
        <input class="button-primary" value="Senden" v-on:click="dump">
    </fieldset>
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
            dump: function(){
                this.$ajax.post('/api', {
                    prayer: this.prayer,
                    is_public: this.isPublic,
                }).then( response => {
                    console.log("RESPONSE", response);
                });
            }
        }
      }
</script>