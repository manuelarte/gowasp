import type { User } from '@/models/users.model.ts'
import { defineBasicLoader } from 'unplugin-vue-router/data-loaders/basic'
import { backendClient } from '@/stores/app'

export const usePostComments = defineBasicLoader('/posts/[id]', async to => {
  const commentsPage = await backendClient.getPostComments(Number.parseInt(to.params.id))
  const users: Record<number, User> = {}
  for (const i in commentsPage.data) {
    const comment = commentsPage.data[i]
    users[comment.userId] = await backendClient.getUser(comment.userId)
  }

  return { commentsPage, users }
})
