import FormsService from '@/services/forms-service'

const state = () => ({
  forms: [],
  createdForm: ''
})

const actions = {
  async get ({ commit }) {
    const forms = await FormsService.getForms()

    commit('setForms', forms)

    return forms
  },

  async create ({ commit }, form) {
    const createdForm = await FormsService.createForm(form)

    commit('setCreatedForm', createdForm)

    return createdForm
  }
}

const mutations = {
  setForms (state, forms) {
    state.forms = forms
  },
  setCreatedForm (state, createdForm) {
    state.createdForm = createdForm
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
