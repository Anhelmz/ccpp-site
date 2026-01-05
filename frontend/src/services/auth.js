import api from './api'

export async function login(email, password) {
  return api.post('/auth/login', { email, password })
}

export async function fetchCurrentUser() {
  return api.get('/auth/me')
}

export function clearAuth() {
  localStorage.removeItem('authToken')
  localStorage.removeItem('auth_token')
  localStorage.removeItem('authUser')
}

