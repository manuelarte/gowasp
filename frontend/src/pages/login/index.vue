<script setup lang="ts">
  import { useUserStore } from '@/stores/app.ts'

  const visible = ref(false)
  const loading = ref(false)
  const err = ref<string | null>(null)

  const userStore = useUserStore()

  async function login () {
    err.value = null
    loading.value = true
    userStore.login('username', 'password')
      .catch(error => err.value = error.message)
      .finally(
        () => loading.value = false,
      )
  }
</script>

<template>
  <v-card
    class="mx-auto pa-12 pb-8"
    elevation="8"
    max-width="448"
    rounded="lg"
  >
    <div class="text-subtitle-1 text-medium-emphasis">Account</div>

    <v-text-field
      density="compact"
      placeholder="Username"
      prepend-inner-icon="mdi-account-outline"
      variant="outlined"
    />

    <div class="text-subtitle-1 text-medium-emphasis">Password</div>

    <v-text-field
      :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
      density="compact"
      placeholder="Enter your password"
      prepend-inner-icon="mdi-lock-outline"
      :type="visible ? 'text' : 'password'"
      variant="outlined"
      @click:append-inner="visible = !visible"
    />

    <v-btn
      v-if="!loading"
      block
      class="mb-8"
      color="blue"
      size="large"
      variant="tonal"
      @click="login"
    >
      Log In
    </v-btn>
    <v-progress-circular
      v-else
      class="mb-8 align-center"
      color="primary"
      indeterminate
    />

    <v-card
      v-if="err"
      class="mb-12"
      color="red-lighten-2"
      variant="tonal"
    >
      <v-card-text class="text-medium-emphasis text-caption">
        {{ err! }}
      </v-card-text>
    </v-card>
  </v-card>
</template>

<style scoped lang="sass">
</style>
