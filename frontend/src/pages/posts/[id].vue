<script lang="ts">
  import { usePost } from '@/loaders/post'
  import { usePostComments } from '@/loaders/postComments'
  // eslint-disable-next-line unicorn/prefer-export-from
  export { usePost, usePostComments }
</script>

<script setup lang="ts">

  const {
    data: post,
    isLoading: isLoadingPost,
    error: errorPost,
  } = usePost()

  const {
    data: postCommentsPageAndUsers,
    isLoading: isLoadingComments,
    error: errorComments,
    reload: reloadComments,
  } = usePostComments()

  function getUsername (userId: number): string {
    return postCommentsPageAndUsers.value.users[userId].username
  }

  function onCommentSaved (_: Comment): void {
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
      <AddComment class="card" :post="post" @comment:saved="onCommentSaved($event)" />

      <p>This post has {{ postCommentsPageAndUsers.commentsPage?.data.length }} comment(s)</p>

      <v-card
        v-for="comment in postCommentsPageAndUsers.commentsPage?.data ?? []"
        :key="comment.id"
        class="card"
        :subtitle="comment.postedAt"
      >
        <template #prepend>
          <v-avatar
            color="brown"
            size="large"
          >
            <span class="text-h5">{{ getUsername(comment.userId).substring(0, 2).toUpperCase() }}</span>
          </v-avatar>
        </template>
        <template #title>
          <span class="font-weight-black">{{ getUsername(comment.userId) }}</span>
        </template>

        <v-card-text class="pt-4">
          {{ comment.comment }}
        </v-card-text>
      </v-card>
    </template>
  </div>
  <v-snackbar
    v-model="commentSavedSnackbar"
    timeout="1000"
  >
    Comment saved
  </v-snackbar>
</template>

<style scoped lang="sass">
.card
  margin: 12px
</style>
