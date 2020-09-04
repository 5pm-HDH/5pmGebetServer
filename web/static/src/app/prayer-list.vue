<template>
    <div>
        <prayer-list-item v-for="item in prayerList" :key="item.id" :data="item"></prayer-list-item>
    </div>
</template>

<script>
    export default{
        data(){
            return {
                prayerList: [
                    {id: 1, prayer: 'Ich bin krank!', is_public: false, is_approved: true},
                    {id: 2, prayer: 'Ich bin kränker!', is_public: false, is_approved: true},
                    {id: 3, prayer: 'Ich bin cool', is_public: false, is_approved: true},
                    {id: 4, prayer: 'Ich bin traurig', is_public: false, is_approved: true},
                    {id: 5, prayer: 'Ich bin betrübt', is_public: false, is_approved: true},
                ],
            };
        },
        mounted(){

        },
        methods: {
            
        }
      }
</script>