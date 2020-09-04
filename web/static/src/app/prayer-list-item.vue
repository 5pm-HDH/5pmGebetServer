<template>
    <blockquote>
    <p><em>{{ prayer }}</em> - {{ date }}</p>
    </blockquote>

</template>

<script>
    export default{
        props: ['data'],
        data(){
            return {
                prayer: (this.data.hasOwnProperty('prayer') ? this.data.prayer : null),
                date: (this.data.hasOwnProperty('date') ? this.data.prayer : null),
                is_public: (this.data.hasOwnProperty('is_public') ? this.data.prayer : null),
                approved: (this.data.hasOwnProperty('approved') ? this.data.prayer : null),
            };
        },
        mounted(){

        },
        methods: {
            
        }
      }
</script>