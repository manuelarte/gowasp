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

  const setUser = (newUser: User | null) => {
    user.value = newUser
  }

  const logout = async () => {
    await backendClient.logout()
    user.value = null
  }

  return { user, setUser, logout }
})
