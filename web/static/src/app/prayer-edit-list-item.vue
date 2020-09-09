<template>
<div style="border: 1px solid black; border-radius: 5px; box-shadow: 0 3px 3px rgba(0,0,0,0.2); padding: 4rem; margin-bottom: 2rem;">
    <fieldset>
        <label for="prayer">Gebetanliegen - #{{ id }}</label>
        <textarea placeholder="Hi CJ …" id="prayer" v-model="prayer" :disabled="!isEdit"></textarea>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Öffentlich</label>
            <input type="checkbox" id="confirmField" v-model="isPublic" :disabled="!isEdit">
        </div>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Bestätigt</label>
            <input type="checkbox" id="confirmField" v-model="approved"  :disabled="!isEdit">
        </div>     
    </fieldset>
    <a class="button" v-on:click="clickButton">{{ (isEdit ? 'Speichern' : 'Bearbeiten') }}</a>
</div>
</template>

<script>
    export default{
        props: ['data'],
        data(){
            return {
                id: (this.data.hasOwnProperty('id') ? this.data.id : null),
                prayer: (this.data.hasOwnProperty('prayer') ? this.data.prayer : null),
                date: (this.data.hasOwnProperty('date') ? this.data.date : null),
                isPublic: (this.data.hasOwnProperty('is_public') ? this.data.is_public : null),
                approved: (this.data.hasOwnProperty('approved') ? this.data.approved : null),

                isEdit: false
            };
        },
        mounted(){

        },
        methods: {
            clickButton(){
                if(this.isEdit){
                    this.$ajax.put('/api?key=tWyV2KiZ1YFfqEUiBYg4g8sK3ot72nihkK9AMMZb', {
                        id: this.id,
                        prayer: this.prayer,
                        is_public: this.isPublic,
                        approved: this.approved,
                    }).then( response => {
                        console.log("RESPONSE VON ÄNDERUNG!", response);

                        this.$parent.$emit('reload-prayer-list', {});
                    });
                }else{
                    console.log("CHANGE TO EDIT MODE!");
                }


                this.isEdit = !this.isEdit;
            }
        }
      }
</script>