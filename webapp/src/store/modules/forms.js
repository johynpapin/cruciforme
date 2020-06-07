import FormsService from '@/services/forms-service'

const state = () => ({
  forms: [],
  createdFormId: ''
})

const actions = {
  async get ({ commit }) {
    const forms = await FormsService.getForms()
    commit('setForms', forms)
  },

  async create ({ commit }) {
    const form = await FormsService.createForm()
    commit('setCreatedFormId', form.id)
  }
}

const mutations = {
  setForms (state, forms) {
    state.forms = forms
  },
  setCreatedFormId (state, createdFormId) {
    state.createdFormId = createdFormId
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
