import type { User } from '@/models/users.model.ts'
// Utilities
import { defineStore } from 'pinia'
import { type ApiClient, HttpClient } from '@/api/backend.client'

export const backendClient: ApiClient = new HttpClient()

export const useAppStore = defineStore('app', {
  state: () => ({
    //
  }),
})

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const login = async (username: string, password: string): Promise<User> => {
    return backendClient.login(username, password)
      .then(u => {
        user.value = u
        return u
      })
  }
  const logout = async () => {
    await backendClient.logout()
    user.value = null
  }

  const signup = async (username: string, password: string): Promise<User> => {
    return backendClient.signup(username, password)
      .then(u => {
        user.value = u
        return u
      })
  }

  return { user, login, logout, signup }
})
