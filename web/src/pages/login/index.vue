<template>
  <v-card
    class="mx-auto pa-12 pb-8"
    elevation="8"
    max-width="448"
    rounded="lg"
  >

    <v-form ref="form" v-model="valid" fast-fail @submit.prevent>
      <div class="text-subtitle-1 text-medium-emphasis">Account</div>

      <v-text-field
        v-model="username"
        density="compact"
        placeholder="Username"
        prepend-inner-icon="mdi-account-outline"
        required
        :rules="usernameRules"
        variant="outlined"
      />

      <div class="text-subtitle-1 text-medium-emphasis">Password</div>

      <v-text-field
        v-model="password"
        :append-inner-icon="!visible ? 'mdi-eye-off' : 'mdi-eye'"
        density="compact"
        placeholder="Enter your password"
        prepend-inner-icon="mdi-lock-outline"
        required
        :type="visible ? 'text' : 'password'"
        variant="outlined"
        @click:append-inner="visible = !visible"
      />

      <v-btn
        v-if="form"
        block
        class="mb-8"
        color="blue"
        :disabled="!valid || loadingNewUser"
        :loading="loadingExistingUser"
        size="large"
        type="submit"
        variant="tonal"
        @click="login"
      >
        Log In
      </v-btn>

      <p class="text-body-2 text-center">Or</p>
      <v-divider />
      <br>

      <v-btn
        v-if="form"
        block
        class="mb-8"
        color="blue"
        :disabled="!valid || loadingExistingUser"
        :loading="loadingNewUser"
        size="large"
        type="submit"
        variant="tonal"
        @click="signup"
      >
        Create new account
      </v-btn>
    </v-form>

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

<script setup lang="ts">
  import router from '@/router'
  import { backendClient, useUserStore } from '@/stores/app'

  const valid = ref(false)
  const visible = ref(false)
  const loadingExistingUser = ref(false)
  const loadingNewUser = ref(false)
  const err = ref<string | null>(null)

  const form = ref()
  const username = ref('')
  const usernameRules = [
    (value: string | null) => {
      if ((value?.length ?? 0) < 3) return 'Username name must be at least 3 characters.'
      if ((value?.length ?? 0) > 21) return 'Username name must be at most 20 characters.'
      return true
    },
  ]
  const password = ref('')

  const userStore = useUserStore()

  const loading = computed(() => {
    return loadingExistingUser.value || loadingNewUser.value
  })

  async function login () {
    err.value = null
    loadingExistingUser.value = true
    backendClient.login(username.value, password.value)
      .then(user => {
        userStore.setUser(user)
        router.push('/')
      })
      .catch(error => err.value = error.message)
      .finally(
        () => loadingExistingUser.value = false,
      )
  }

  async function signup () {
    err.value = null
    loadingNewUser.value = true
    backendClient.signup(username.value, password.value)
      .then(user => {
        userStore.setUser(user)
        router.push('/')
      })
      .catch(error => err.value = error.message)
      .finally(
        () => loadingNewUser.value = false,
      )
  }
</script>

<style scoped lang="sass">
</style>
