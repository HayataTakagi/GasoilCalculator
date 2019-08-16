import Vue from 'vue'
import Router from 'vue-router'

import HelloWorld from '../components/pages/HelloWorld'
import ShowFee from '../components/pages/ShowFee'

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: '/hello',
      component: HelloWorld
    },
    {
      path: '/fee',
      component: ShowFee
    }
  ]
});

export default router
