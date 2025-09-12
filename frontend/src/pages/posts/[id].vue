<script lang="ts">
  import { usePost } from '@/loaders/post'
  import { usePostComments } from '@/loaders/postComments'
  // eslint-disable-next-line unicorn/prefer-export-from
  export { usePost, usePostComments }
</script>

<script setup lang="ts">
  // eslint-disable-next-line import/first
  import type { User } from '@/models/users.model'

  const {
    data: post,
    isLoading: isLoadingPost,
    error: errorPost,
  } = usePost()

  const {
    data: postCommentsPage,
    isLoading: isLoadingComments,
    error: errorComments,
    reload: reloadComments,
  } = usePostComments()
</script>

<template>
  <div class="card">
    <v-skeleton-loader
      v-if="isLoadingComments"
      class="mx-auto border"
      max-width="300"
      type="image, article"
    />
    <v-card v-else-if="errorComments" color="red-lighten-2" variant="tonal">Error</v-card>
    <v-card v-else class="card">
      <v-card-title>{{ post.title }}</v-card-title>
      <v-card-text>{{ post.content }}</v-card-text>
    </v-card>

    <br>

    <p>This post has {{ postCommentsPage?.data.length }} comment(s)</p>

    <v-card
      v-for="comment in postCommentsPage?.data ?? []"
      :key="comment.id"
      class="card"
      prepend-avatar="https://cdn.vuetifyjs.com/images/john.jpg"
      :subtitle="comment.postedAt"
    >
      <template #title>
        <span class="font-weight-black">{{ comment.userId }}</span>
      </template>

      <v-card-text class="pt-4">
        {{ comment.comment }}
      </v-card-text>
    </v-card>
  </div>
</template>

<style scoped lang="sass">
.card
  margin: 12px
</style>
