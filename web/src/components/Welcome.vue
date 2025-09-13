<script setup lang="ts">
  import type { Page } from '@/models/http.model'
  import type { Post } from '@/models/posts.model'
  import router from '@/router'
  import { backendClient, useUserStore } from '@/stores/app.ts'

  const client = backendClient

  // TODO(manuelarte): this is all over the place
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
  const errorLoadingPostPage = ref<string | null>(null)

  watch(page, () => {
    loadPosts()
  })

  async function loadPosts () {
    loadingPosts.value = true
    client.getPosts(page.value - 1)
      .then(pp => postsPage.value = pp)
      .catch(error => errorLoadingPostPage.value = error)
      .finally(() => loadingPosts.value = false)
  }

  onMounted(() => {
    loadingIntro.value = true
    client.getStaticPost('intro.txt')
      .then(i => intro.value = i)
      .finally(() => loadingIntro.value = false)

    loadPosts()
  })
</script>

<template>
  <h1>Welcome gowasp to Gowasp website</h1>

  <v-alert
    density="compact"
    :text="`This is just for information purposes: Password hash: ${user!.password}`"
    title="Warning"
    type="warning"
    variant="outlined"
  />

  <br>

  <v-progress-linear v-if="loadingIntro" indeterminate />
  <v-card v-else>
    <v-card-title>Intro</v-card-title>
    <v-card-text>
      {{ intro }}
    </v-card-text>
  </v-card>

  <v-skeleton-loader
    v-if="loadingPosts"
    class="mx-auto border"
    type="paragraph"
  />
  <v-card
    v-if="!loadingPosts && postsPage"
    class="mx-auto"
  >
    <v-card-title>Posts</v-card-title>
    <v-list lines="one">
      <v-list-item
        v-for="post in postsPage!.data"
        :key="post.id"
        :value="post"
        @click="router.push(`/posts/${post.id}`)"
      >
        <template #prepend>
          <v-avatar>
            <v-icon color="primary" icon="mdi-post" />
          </v-avatar>
          <v-list-item-title>{{ post.title }}</v-list-item-title>
        </template>
      </v-list-item>
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
