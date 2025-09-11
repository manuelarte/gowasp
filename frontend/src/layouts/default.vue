<template>
  <v-toolbar>
    <v-toolbar-title text="GOwasp" />
    <template #append>
      <div class="d-flex ga-1">
        <v-btn v-if="!user" icon="mdi-login" to="/login" />
        <v-btn v-else icon="mdi-logout" @click="userStore.logout()" />
      </div>
    </template>
  </v-toolbar>
  <v-main>
    <router-view />
  </v-main>

  <AppFooter />
</template>

<script lang="ts" setup>
  import { useUserStore } from '@/stores/app'

  const userStore = useUserStore()

  const user = ref(userStore.user)
</script>
