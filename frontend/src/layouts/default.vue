<template>
  <v-toolbar color="primary">
    <v-toolbar-title text="GOwasp" />
    <template #append>
      <div class="d-flex ga-1">
        <v-btn v-if="!user" v-tooltip="'Login'" icon="mdi-login" to="/login" />
        <v-btn v-else v-tooltip="'Logout'" icon="mdi-logout" @click="userStore.logout()" />
      </div>
    </template>
  </v-toolbar>
  <v-main>
    <router-view />
  </v-main>

  <AppFooter />
</template>

<script lang="ts" setup>
  import router from '@/router'
  import { useUserStore } from '@/stores/app'

  const userStore = useUserStore()
  userStore.$subscribe((_, state) => {
    if (!state.user) {
      router.push('/login')
    }
    user.value = state.user
  })

  const user = ref(userStore.user)

  onMounted(() => {
    if (!user.value) {
      router.push('/login')
    }
  })
</script>
