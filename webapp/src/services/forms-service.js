import axios from 'axios'

const API_URL = process.env.VUE_APP_API_URL + '/forms'

function mapForm (form) {
  return {
    id: form.id,
    name: form.name,
    email: form.email,
    confirmed: form.confirmed
  }
}

export default {
  async getForms () {
    const result = await axios.get(API_URL)

    return result.data.forms.map(mapForm)
  },

  async createForm (form) {
    const result = await axios.post(API_URL, form)

    return mapForm(result.data.form)
  }
}
