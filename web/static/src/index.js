import Vue from 'vue';
import ajax from './ajax';

// PROTOTYPE
Vue.prototype.$ajax = ajax;
Vue.prototype.$eventHub = new Vue();

// COMPONENTS
import example from './app/example.vue';
import notification from './app/notification.vue';
import prayerCreateForm from './app/prayer-create-form.vue';
import prayerList from './app/prayer-list.vue';
import prayerListItem from './app/prayer-list-item.vue';
import prayerEditList from './app/prayer-edit-list.vue';
import prayerEditListItem from './app/prayer-edit-list-item.vue';


Vue.component('example', example);
Vue.component('notification', notification);
Vue.component('prayer-create-form', prayerCreateForm);
Vue.component('prayer-list', prayerList);
Vue.component('prayer-list-item', prayerListItem);
Vue.component('prayer-edit-list', prayerEditList);
Vue.component('prayer-edit-list-item', prayerEditListItem);

var app = new Vue({
    el: '#app',
    data: {
      message: 'Hello Vue!'
    }
});