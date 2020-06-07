import Vue from 'vue'
import Vuex from 'vuex'

import auth from './modules/auth'
import forms from './modules/forms'
import alert from './modules/alert'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    auth,
    forms,
    alert
  },
  strict: debug
})
