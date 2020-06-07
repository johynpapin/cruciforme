const ALERT_TIMEOUT_MS = 5000

const state = () => ({
  currentAlert: null,
  timeoutId: null
})

const actions = {
  async close ({ commit, state }) {
    if (state.timeoutId !== null) {
      clearTimeout(state.timeoutId)
      commit('setTimeoutId', null)
    }

    commit('setCurrentAlert', null)
  },

  async create ({ dispatch, commit, state }, alert) {
    if (state.currentAlert) {
      await dispatch('close')
    }

    commit('setCurrentAlert', alert)

    commit('setTimeoutId', setTimeout(() => {
      commit('setTimeoutId', null)
      commit('setCurrentAlert', null)
    }, ALERT_TIMEOUT_MS))
  }
}

const mutations = {
  setCurrentAlert (state, currentAlert) {
    state.currentAlert = currentAlert
  },
  setTimeoutId (state, timeoutId) {
    state.timeoutId = timeoutId
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
