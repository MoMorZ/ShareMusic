import Init from '../pages/Init.vue'
import Vue from 'vue'
import Router from 'vue-router'
import Play from '../pages/Play.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            component: Init
        },
        {
            path: '/player',
            component: Play
        }
    ]
})