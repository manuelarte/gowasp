import { backendClient, useUserStore } from '@/stores/app'

export function useSession () {
  const userStore = useUserStore()

  async function restoreSession () {
    if (userStore.user) {
      return
    }
    userStore.user = await backendClient.getSession()
  }

  return { restoreSession }
}
