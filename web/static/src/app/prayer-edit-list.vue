<template>
    <div>
        <h3>Nicht bestätigt:</h3>
        <prayer-edit-list-item v-for="item in notApprovedPrayerList" :key="item.id" :data="item"></prayer-edit-list-item>
        <h3>Bestätigt:</h3>
        <prayer-edit-list-item v-for="item in approvedPrayerList" :key="item.id" :data="item"></prayer-edit-list-item>
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
            approvedPrayerList: {
                get(){
                    return this.sortedPrayerList.filter( prayer => {
                        return prayer.approved;
                    });
                }
            },
            notApprovedPrayerList: {
                get(){
                    return this.sortedPrayerList.filter( prayer => {
                        return !prayer.approved;
                    });
                }
            },
            sortedPrayerList: {
                get(){
                    return this.prayerList.sort( (prayerA, prayerB) => {
                        return parseInt(prayerA.id) - parseInt(prayerB.id);
                    });
                }
            }
        },
        mounted(){
            this.loadPrayerList();

            setInterval(this.loadPrayerList, 1000);

            this.$on('reload-prayer-list', ev => {
                this.loadPrayerList();
            });
        },
        methods: {
            loadPrayerList(){

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