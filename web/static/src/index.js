import Vue from 'vue';
import ajax from './ajax';

Vue.prototype.$ajax = ajax;

import example from './app/example.vue';
import prayerCreateForm from './app/prayer-create-form.vue';
import prayerList from './app/prayer-list.vue';
import prayerListItem from './app/prayer-list-item.vue';


Vue.component('example', example);
Vue.component('prayer-create-form', prayerCreateForm);
Vue.component('prayer-list', prayerList);
Vue.component('prayer-list-item', prayerListItem);


var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!'
    }
});