<script setup lang="ts">
  import type { Page } from '@/models/http.model'
  import type { Post } from '@/models/posts.model'
  import router from '@/router'
  import { backendClient, useUserStore } from '@/stores/app.ts'

  const client = backendClient

  const userStore = useUserStore()
  userStore.$subscribe((_, state) => {
    if (!state.user) {
      router.push('/login')
    }
    user.value = state.user
  })

  const user = ref(userStore.user)

  const loadingIntro = ref(true)
  const intro = ref<string | null>(null)

  const page = ref(1)
  const loadingPosts = ref(true)
  const postsPage = ref<Page<Post> | null>(null)
  // TODO const errorLoadingPostPage = ref<string | null>(null)

  onMounted(() => {
    loadingIntro.value = true
    client.getStaticPost('intro.txt')
      .then(i => intro.value = i)
      .finally(() => loadingIntro.value = false)

    loadingPosts.value = true
    client.getPosts(page.value - 1)
      .then(pp => postsPage.value = pp)
      .finally(() => loadingPosts.value = false)
  })
</script>

<template>
  <h1>Welcome gowasp to Gowasp website</h1>

  <v-alert
    density="compact"
    :text="`Warning: This is just for information purposes: Password hash: ${user!.password}`"
    title="Info"
    type="warning"
    variant="outlined"
  />

  <br>

  <h2>Posts</h2>

  <v-progress-linear v-if="loadingIntro" indeterminate />
  <v-card v-else>
    <v-card-title>Intro</v-card-title>
    <v-card-text>
      {{ intro }}
    </v-card-text>
  </v-card>

  <h2>Latest Posts</h2>
  <v-skeleton-loader
    v-if="loadingPosts"
    class="mx-auto border"
    type="paragraph"
  />
  <v-card
    v-if="!loadingPosts && postsPage"
    class="mx-auto"
  >
    <v-list lines="one">
      <v-list-item
        v-for="post in postsPage!.data"
        :key="post.id"
        :title="post.title"
      />
    </v-list>

    <v-pagination
      v-model="page"
      :length="postsPage?._metadata?.totalPages"
      rounded="circle"
    />
  </v-card>
</template>

<style scoped lang="sass">

</style>
