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
                //    {id: 1, prayer: 'Ich bin krank!', is_public: false, is_approved: true},
                ],
            };
        },
        mounted(){
            this.$ajax.get("/api?key=tWyV2KiZ1YFfqEUiBYg4g8sK3ot72nihkK9AMMZb", {}).then( response => {
                this.prayerList = response.data;
            });
        },
        methods: {
            
        }
      }
</script>