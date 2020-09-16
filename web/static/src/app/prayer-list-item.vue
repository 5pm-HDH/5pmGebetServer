<template>
    <blockquote :class="{'prayer-item': true, 'show-prayer': isMounted}">
        <p><em>{{ prayer }}</em></p>
    </blockquote>
</template>
<script>
    export default{
        props: ['data'],
        data(){
            return {
                prayer: (this.data.hasOwnProperty('prayer') ? this.data.prayer : null),
                isMounted: false,
            };
        },
        mounted(){
            setTimeout( () => {
                this.isMounted = true;
            }, 500);
        },
        methods: {
            
        }
      }
</script>