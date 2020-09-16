<template>
    <div>
        <prayer-list-item v-for="item in sortedPrayerList" :key="item.id" :data="item"></prayer-list-item>
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
            sortedPrayerList: function(){
                return this.prayerList.sort( (prayerA, prayerB) => prayerB.id - prayerA.id);
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

                this.$ajax.get("/api/view?key=" + key, {}).then( response => {
                    this.prayerList = response.data;
                });
            }
        }
      }
</script>