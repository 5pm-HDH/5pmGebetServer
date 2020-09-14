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
                  // FIXED && prayer.is_public to also filter that the prayer is public
                    return prayer.approved && prayer.is_public;
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
              // get and forward the key from the user to the database backend
              let queryString = window.location.search;
              let params = new URLSearchParams(queryString);
              let key = params.get("key");
              //TODO: make function for multiple uses

                this.$ajax.get("/api?key=" + key, {}).then( response => {
                    this.prayerList = response.data;
                });
            }
        }
      }
</script>