import AuthService from '@/services/auth-service'

const state = () => ({
  signedIn: false
})

const actions = {
  async signIn ({ commit }, user) {
    try {
      await AuthService.signIn(user)
      commit('signIn')
    } catch (err) {
      commit('logout')
      throw err
    }
  },

  async signUp ({ commit }, user) {
    await AuthService.signUp(user)
  },

  async verify ({ commit }, request) {
    try {
      await AuthService.verify(request)
      commit('signIn')
    } catch (err) {
      commit('logout')
      throw err
    }
  },

  logout ({ commit }) {
    AuthService.logout()

    commit('logout')
  }
}

const mutations = {
  signIn (state) {
    state.signedIn = true
  },
  logout (state) {
    state.signedIn = false
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
