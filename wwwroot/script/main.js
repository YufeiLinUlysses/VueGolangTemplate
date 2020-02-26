//import 'bootstrap/dist/js/bootstrap.js'
//import 'font-awesome/css/font-awesome.css'
// import { parsePhoneNumberFromString } from 'libphonenumber-js'
import Vue from 'vue'
import axios from 'axios'
import VueLocalStorage from 'vue-ls';
import VueRouter from 'vue-router'
import moment from 'moment'
import 'animate.css/animate.min.css'
import 'bootstrap/dist/css/bootstrap.css'
import '@fortawesome/fontawesome-free/css/all.css'
import VuePreferences from 'vue-preferences';
import 'lodash'

import App from './App.vue'
import sr from './pages/searchresult.vue'
import dashboard from './pages/dashboard.vue'
import financial from './components/Financial/financialgeneral.vue'
import care from './components/CareCoordination/ccgeneral.vue'
import quality from './components/QualityMeasure/qmgeneral.vue'
import social from './components/SocialDeterminant/sdgeneral.vue'
import health from './components/HealthInfo/higeneral.vue'

Vue.use(VuePreferences);
Vue.use(VueRouter);

Vue.use(VueLocalStorage, {
    namespace: 'vuejs__'
});
Vue.directive('focus', {
    inserted: function (el) {
        el.focus();
    },
    update: function (el) {
        Vue.nextTick(function () {
            el.focus();
        })
    }
})
var webcall = axios.create({
    baseURL: 'api',
    timeout: 6000,
    withCredentials:false, 
    headers: { 'Content-Type': 'application/json' }

});




Vue.filter('timeago', function (value) {
    if (value) {
        return moment(String(value)).fromNow(false);
    }
});




Vue.filter('formatElipsi', function (value) {
    if (value) {
        if (value.length > 25)
            return value.substring(0, 25) + '...';
        else
            return value;


    }
});

webcall.interceptors.request.use(function(config) {
 
    return config;
}, function(error) {

    return Promise.reject(error);
});
//webcall.interceptors.response.use(function(response) {

//    return response;
//}, function (error) {
//    alert("fail");

//    if (error.toString().indexOf("403") !== -1) {   
//               
//                console.log(response.data);        

//    }
//    return Promise.reject(error);
//    });



var routes = [   
    { path: '/', component: App },
    { path: '/searchresult/:searchresults', name: 'SearchResult', component: sr, props: true },
    { path: '/dashboard/:patients/:id', 
      name:'Dashboard', 
      component: dashboard, 
      props:true,
      children:[
          { path: '/financial', name: 'Financial', component: financial},
          { path: '/care', name: 'CareCoor', component: care },
          { path: '/social', name: 'SocialDeter', component: social }, 
          { path: '/quality', name: 'QualityMeasure', component: quality },
          { path: '/health', name: 'HealthInfo', component: health }
      ]
    }
];

var router = new VueRouter({
    routes: routes
});

new Vue({
    el: '#app',
    router: router,
    data: { api: webcall }
    //render: h => h(App)
})

