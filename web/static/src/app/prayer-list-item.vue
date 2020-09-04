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
                date: null,
                is_public: null,
                is_approved: null,
            };
        },
        mounted(){

        },
        methods: {
            
        }
      }
</script>