import Vue from 'vue';
import ajax from './ajax';

import example from './app/example.vue';

Vue.component('example', example);

var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!'
    }
});