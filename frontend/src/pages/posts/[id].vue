<script lang="ts">
  import { defineBasicLoader } from 'unplugin-vue-router/data-loaders/basic'

  const useUserData = defineBasicLoader('/posts/[id]', async route => {
    // return getUserById(route.params.id)
    return {
      id: route.params.id,
      title: 'My Post',
      content: 'This is my post content',
    }
  })
</script>
<script setup lang="ts">
  const {
    data: post,
    isLoading, // a boolean indicating if the loader is fetching data
    error, // an error object if the loader failed
    reload, // a function to refetch the data without navigating
  } = useUserData()
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
