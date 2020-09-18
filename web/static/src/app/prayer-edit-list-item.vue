<template>
<div :class="{'prayer-item': true, 'prayer-item-is-edit': isEdit}" style="">
    <fieldset>
        <label for="prayer">Gebetanliegen - #{{ id }}</label>
        <textarea placeholder="Hi CJ …" id="prayer" v-model="prayer" :disabled="!isEdit"></textarea>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Öffentlich</label>
            <label class="switch">
                <input type="checkbox" id="confirmField" v-model="isPublic" :disabled="!isEdit">
                <span class="slider round"></span>
            </label>
        </div>
        <div class="float-right">
            <label class="label-inline" for="confirmField">Bestätigt</label>
            <label class="switch">
                <input type="checkbox" id="confirmField" v-model="approved"  :disabled="!isEdit">
                <span class="slider round"></span>
            </label>
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
                    // get and forward the key from the user to the database backend
                    let queryString = window.location.search;
                    let params = new URLSearchParams(queryString);
                    let key = params.get("key");
                    //TODO: make function for multiple uses

                    this.$ajax.put('/api?key='+key, {
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