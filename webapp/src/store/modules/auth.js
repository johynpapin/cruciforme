import AuthService from '@/services/auth-service'

const state = () => ({
  signedIn: false
})

const actions = {
  async signIn ({ commit }, user) {
    try {
      await AuthService.signIn(user)
      commit('signin')
    } catch (err) {
      commit('logout')
      throw err
    }
  },

  async signUp ({ commit }, user) {
    try {
      await AuthService.signUp(user)
      commit('signin')
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
  signin (state) {
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
