import axios from 'axios'

const API_URL = process.env.VUE_APP_API_URL + '/forms'

export default {
  async getForms () {
    const result = await axios.get(API_URL)

    return result.data.forms.map(form => ({
      id: form.id
    }))
  },

  async createForm () {
    const result = await axios.post(API_URL)

    return {
      id: result.data.id
    }
  }
}
