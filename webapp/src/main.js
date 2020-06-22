import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/vue-meta'
import './plugins/vee-validate'
import './interceptors'
import AuthService from './services/auth-service'

Vue.config.productionTip = false

Vue.directive('focus', {
  inserted (el) {
    el.focus()
  }
})

if (AuthService.init()) {
  store.commit('auth/signIn')
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
