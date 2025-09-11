import type { User } from '@/models/users.model.ts'
// Utilities
import { defineStore } from 'pinia'
import { type ApiClient, HttpClient } from '@/services/backend.client'

const backendClient: ApiClient = new HttpClient('http://localhost:8083')

export const useAppStore = defineStore('app', {
  state: () => ({
    //
  }),
})

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const login = async (username: string, password: string) => {
    return backendClient.login(username, password).then(u => user.value = u)
  }
  const logout = async () => {
    await backendClient.logout()
    user.value = null
  }

  return { user, login, logout }
})
