<template>
    <blockquote>
    <p><em>{{ prayer }}</em></p>
    </blockquote>

</template>

<script>
    export default{
        props: ['data'],
        data(){
            return {
                prayer: (this.data.hasOwnProperty('prayer') ? this.data.prayer : null),
                date: (this.data.hasOwnProperty('date') ? this.data.date : null),
                is_public: (this.data.hasOwnProperty('is_public') ? this.data.is_public : null),
                approved: (this.data.hasOwnProperty('approved') ? this.data.approved : null),
            };
        },
        mounted(){

        },
        methods: {
            
        }
      }
</script>