import axios from 'axios'
import AuthService from './services/auth-service'

function isAccessTokenExpiredError (errorResponse) {
  if (errorResponse.status !== 401) {
    return false
  }

  if (!errorResponse.data.error) {
    return false
  }

  return errorResponse.data.error.code === 'auth-access-token-expired'
}

async function refreshTokenAndRetry (error) {
  const { accessToken } = await AuthService.refresh()

  const config = error.config
  config.headers.Authorization = accessToken

  return axios.request(config)
}

async function responseErrorHandler (error) {
  const errorResponse = error.response

  if (isAccessTokenExpiredError(errorResponse)) {
    return refreshTokenAndRetry(error)
  }

  throw error
}

axios.interceptors.response.use(undefined, responseErrorHandler)
