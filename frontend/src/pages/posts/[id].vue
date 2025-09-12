<script lang="ts">
  import { usePost } from '@/loaders/post'
  import { usePostComments } from '@/loaders/postComments'
  export { usePost, usePostComments }
</script>
<script setup lang="ts">

  const post = ref({ title: 'My title', content: 'my content' })

  const {
    data: postManuel,
  } = usePost()

  const {
    data: postCommentsPage,
    isLoading, // a boolean indicating if the loader is fetching data
    error, // an error object if the loader failed
    reload, // a function to refetch the data without navigating
  } = usePostComments()
</script>

<template>
  <v-skeleton-loader
    v-if="isLoading"
    class="mx-auto border"
    max-width="300"
    type="image, article"
  />
  <v-card v-else-if="error" color="red-lighten-2" variant="tonal">Error</v-card>
  <v-card v-else>
    <v-card-title>{{ post.title }}</v-card-title>
    <v-card-text>{{ post.content }}</v-card-text>
  </v-card>
</template>

<style scoped lang="sass">

</style>
