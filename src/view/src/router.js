import Vue from 'vue'
import Router from 'vue-router'

import Home from './views/Home.vue'
import Result from './views/Result.vue'
import First from './components/first.vue'
import Last from './components/last.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/result/:results',
      name: 'result',
      component: Result,
      props: true,
      children: [
        {
          path: '/first/:firstn',
          name: 'First',
          component: First,
          props: true
        },
        {
          path: '/last/:lastn',
          name: 'Last',
          component: Last,
          props: true
        }
      ]
    }
  ]
})
