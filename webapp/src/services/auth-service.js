import axios from 'axios'
import jwtDecode from 'jwt-decode'
import { isPast } from 'date-fns'

const API_URL = process.env.VUE_APP_API_URL + '/auth'

let auth = {
  accessToken: '',
  refreshToken: ''
}

function applyAuth ({ accessToken, refreshToken }) {
  auth = {
    accessToken,
    refreshToken
  }

  localStorage.setItem('auth', JSON.stringify({
    accessToken,
    refreshToken
  }))

  axios.defaults.headers.common.Authorization = accessToken

  return auth
}

function clearAuth () {
  auth = {
    accessToken: '',
    refreshtoken: ''
  }

  localStorage.removeItem('auth')

  delete axios.defaults.headers.common.Authorization
}

export default {
  async init () {
    const storedAuth = JSON.parse(localStorage.getItem('auth'))

    if (storedAuth === null) {
      return false
    }

    const claims = jwtDecode(storedAuth.refreshToken)

    if (isPast(claims.exp * 1000)) {
      return false
    }

    applyAuth(storedAuth)

    return true
  },

  async signUp ({ email, password }) {
    await axios.post(API_URL + '/sign-up', {
      email,
      password
    })
  },

  async signIn ({ email, password }) {
    const result = await axios.post(API_URL + '/sign-in', {
      email,
      password
    })

    if (result.data.accessToken) {
      return applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }

    throw new Error('the access token is missing in the server response')
  },

  async refresh () {
    const result = await axios.post(API_URL + '/refresh', {
      refreshToken: auth.refreshToken
    })

    if (result.data.accessToken) {
      return applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }

    throw new Error('the access token is missing in the server response')
  },

  async verify ({ verificationToken, password }) {
    const result = await axios.post(API_URL + '/verify', {
      verificationToken,
      password
    })

    if (result.data.accessToken) {
      return applyAuth({
        accessToken: result.data.accessToken,
        refreshToken: result.data.refreshToken
      })
    }

    throw new Error('the access token is missing in the server response')
  },

  logout: clearAuth
}
