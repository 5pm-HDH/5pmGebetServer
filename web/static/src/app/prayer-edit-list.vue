<template>
    <div>
        <div class="prayer-list-not-approved">
            <h3>nicht bestätigt:</h3>
            <prayer-edit-list-item v-for="item in notApprovedPrayerList" :key="item.id" :data="item"></prayer-edit-list-item>
        </div>
        <div class="prayer-list-public">
            <h3>bestätigt, öffentlich:</h3>
            <prayer-edit-list-item v-for="item in approvedPublicPrayerList" :key="item.id" :data="item"></prayer-edit-list-item>
        </div>
        <div class="prayer-list-not-public">
            <h3>bestätigt, nicht-öffentlich:</h3>
            <prayer-edit-list-item v-for="item in approvedNotPublicPrayerList" :key="item.id" :data="item"></prayer-edit-list-item>
        </div>
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
            approvedPublicPrayerList: {
                get(){
                    return this.approvedPrayerList.filter( prayer => {
                        return prayer.is_public;
                    });
                }
            },
            approvedNotPublicPrayerList: {
                get(){
                    return this.approvedPrayerList.filter( prayer => {
                        return !prayer.is_public;
                    });
                }
            },
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