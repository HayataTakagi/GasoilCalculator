import Vue from 'vue'
import Router from 'vue-router'

import HelloWorld from '../components/pages/HelloWorld'
import ShowArticle from './../components/pages/ShowArticle'

Vue.use(Router);

let router = new Router({
    routes: [
        {
            path: '/hello',
            component: HelloWorld
        },
        {
            path: '/articles',
            component: ShowArticle
        }
    ]
});

export default router
