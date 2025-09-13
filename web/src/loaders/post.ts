import { defineBasicLoader } from 'unplugin-vue-router/data-loaders/basic'
import { backendClient } from '@/stores/app.ts'

export const usePost = defineBasicLoader('/posts/[id]', to =>
  backendClient.getPost(Number.parseInt(to.params.id)),
)
