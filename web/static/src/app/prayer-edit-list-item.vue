<template>
    <fieldset style="border: 1px solid black; border-radius: 5px; box-shadow: 0 3px 3px rgba(0,0,0,0.2); padding: 4rem;">
        <h3>#{{ id }}</h3>
        <label for="prayer">Gebet</label>
        <textarea placeholder="Hi CJ …" id="prayer" v-model="prayer" :disabled="!isEdit"></textarea>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Öffentlich</label>
            <input type="checkbox" id="confirmField" v-model="isPublic" :disabled="!isEdit">
        </div>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Bestätigt</label>
            <input type="checkbox" id="confirmField" v-model="approved"  :disabled="!isEdit">
        </div>
        
        <input class="button-primary" :value="(isEdit ? 'Speichern' : 'Bearbeiten')" v-on:click="clickButton">
    </fieldset>

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
                    console.log("SAVE CHANGE");
                }else{
                    console.log("CHANGE TO EDIT MODE!");
                }


                this.isEdit = !this.isEdit;
            }
        }
      }
</script>