import axios from 'axios'

function isAccessTokenExpiredError (errorResponse) {
  return errorResponse.status === 401

  // TODO: check response error code
}

function refreshTokenAndRetry (error) {
  return Promise.reject(error)
}

function responseErrorHandler (error) {
  const errorResponse = error.response

  if (isAccessTokenExpiredError(errorResponse)) {
    return refreshTokenAndRetry(error)
  }

  return Promise.reject(error)
}

axios.interceptors.response.use(undefined, responseErrorHandler)
