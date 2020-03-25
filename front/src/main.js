import Vue from 'vue'
import Main from './Main.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
import router from './routers'
import axios from 'axios'
import Vuex from 'vuex'
import store from './store/index'

Vue.use(ElementUI)
Vue.use(Vuex)

Vue.config.productionTip = false
Vue.prototype.$axios = axios


new Vue({
  router,
  store,
  render: h => h(Main),
}).$mount('#app')
