import { defineBasicLoader } from 'unplugin-vue-router/data-loaders/basic'
import { backendClient } from '@/stores/app.ts'

export const usePostComments = defineBasicLoader('/posts/[id]', to =>
  backendClient.getPostComments(Number.parseInt(to.params.id)),
)
