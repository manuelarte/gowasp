<script setup lang="ts">
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

  onMounted(() => {
    loadingIntro.value = true
    client.post('intro.txt')
      .then(i => intro.value = i)
      .finally(() => loadingIntro.value = false)
  })
</script>

<template>
  <h1>Welcome gowasp to Gowasp website</h1>

  <v-alert
    density="compact"
    :text="`Warning: This is just for information purposes: Password hash: ${user.password}`"
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
  TODO
</template>

<style scoped lang="sass">

</style>
