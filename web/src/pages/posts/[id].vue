<script lang="ts">
  import { usePost } from '@/loaders/post'
  import { usePostComments } from '@/loaders/postComments'
  // eslint-disable-next-line unicorn/prefer-export-from
  export { usePost, usePostComments }
</script>

<script setup lang="ts">
  // eslint-disable-next-line import/first
  import type { Comment } from '@/models/posts.model'
  // eslint-disable-next-line import/first
  import { getInitials } from '@/models/users.model'

  const {
    data: post,
    isLoading: isLoadingPost,
    error: errorPost,
  } = usePost()

  const {
    data: csrfPostCommentsPageAndUsers,
    isLoading: isLoadingComments,
    error: errorComments,
    reload: reloadComments,
  } = usePostComments()

  function getUserInitials (userId: number): string {
    const user = csrfPostCommentsPageAndUsers.value.users[userId]
    return getInitials(user)
  }

  function onCommentSaved (_: Comment) {
    commentSavedSnackbar.value = true
    reloadComments()
  }

  const commentSavedSnackbar = ref(false)
</script>

<template>
  <div class="card">
    <v-skeleton-loader
      v-if="isLoadingPost"
      class="mx-auto border"
      max-width="300"
      type="image, article"
    />
    <v-card v-else-if="errorPost" color="red-lighten-2" variant="tonal">Error</v-card>
    <v-card v-else class="card">
      <v-card-title>{{ post.title }}</v-card-title>
      <v-card-text>{{ post.content }}</v-card-text>
    </v-card>

    <br>

    <v-skeleton-loader v-if="isLoadingComments" type="card" />
    <template v-else>
      <AddComment class="card" :csrf="csrfPostCommentsPageAndUsers.csrf" :post="post" @comment:saved="onCommentSaved" />

      <p>This post has {{ csrfPostCommentsPageAndUsers.commentsPage?.data.length }} comment(s)</p>

      <v-card
        v-for="comment in csrfPostCommentsPageAndUsers.commentsPage?.data ?? []"
        :key="comment.id"
        class="card"
        :subtitle="comment.postedAt"
      >
        <template #prepend>
          <v-avatar
            color="brown"
            size="large"
          >
            <span class="text-h5">{{ getUserInitials(comment.userId) }}</span>
          </v-avatar>
        </template>
        <template #title>
          <span class="font-weight-black">{{ csrfPostCommentsPageAndUsers.users[comment.userId].username }}</span>
        </template>

        <!-- eslint-disable-next-line vue/no-v-text-v-html-on-component -->
        <v-card-text class="pt-4" v-html="comment.comment" />
      </v-card>
    </template>
  </div>
  <v-snackbar
    v-model="commentSavedSnackbar"
    color="primary"
    timeout="1000"
  >
    Comment saved
  </v-snackbar>
</template>

<style scoped lang="sass">
.card
  margin: 12px
</style>
