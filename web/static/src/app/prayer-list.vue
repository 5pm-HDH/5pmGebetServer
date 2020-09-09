<template>
    <div>
        <prayer-list-item v-for="item in latestPrayerList" :key="item.id" :data="item"></prayer-list-item>
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
        computed: {
            approvedPrayerList: function(){
                return this.prayerList.filter( prayer => {
                    return prayer.approved;
                });
            },
            latestPrayerList: function(){
                return this.approvedPrayerList.slice(Math.max(this.approvedPrayerList.length - 5, 0))
            }
        },
        mounted(){
            this.loadPrayerList();

            setInterval(this.loadPrayerList, 500);
        },
        methods: {
            loadPrayerList()
            {
                this.$ajax.get("/api?key=tWyV2KiZ1YFfqEUiBYg4g8sK3ot72nihkK9AMMZb", {}).then( response => {
                    this.prayerList = response.data;
                });
            }
        }
      }
</script>