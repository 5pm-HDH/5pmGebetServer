import Vue from 'vue';
import ajax from './ajax';

Vue.prototype.$ajax = ajax;

import example from './app/example.vue';
import prayerCreateForm from './app/prayer-create-form.vue';

Vue.component('example', example);
Vue.component('prayer-create-form', prayerCreateForm);

var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!'
    }
});