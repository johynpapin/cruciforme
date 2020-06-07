import axios from 'axios'

const API_URL = process.env.VUE_APP_API_URL + '/auth'

const auth = {
  accessToken: '',
  refreshToken: ''
}

function applyAuth ({ accessToken, refreshToken }) {
  auth.accessToken = accessToken
  auth.refreshToken = refreshToken

  localStorage.setItem('auth', JSON.stringify({
    accessToken,
    refreshToken
  }))

  axios.defaults.headers.common.Authorization = accessToken
}

export default {
  async signUp ({ email, password }) {
    const result = await axios.post(API_URL + '/sign-up', {
      email,
      password
    })

    if (result.data.accessToken) {
      applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }
  },

  async signIn ({ email, password }) {
    const result = await axios.post(API_URL + '/sign-in', {
      email,
      password
    })

    if (result.data.accessToken) {
      applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }
  },

  async refresh () {
    const result = await axios.post(API_URL + '/refresh', {
      refreshToken: auth.refreshToken
    })

    if (result.data.accessToken) {
      applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }
  },

  logout () {
    auth.accessToken = ''
    auth.refreshToken = ''

    localStorage.removeItem('auth')

    delete axios.defaults.headers.common.Authorization
  }
}
