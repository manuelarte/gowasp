<script setup lang="ts">
  import router from '@/router'
  import { useUserStore } from '@/stores/app.ts'

  const visible = ref(false)
  const loading = ref(false)
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

  async function login () {
    err.value = null
    loading.value = true
    userStore.login(username.value, password.value)
      .then(_ => {
        router.push('/')
      })
      .catch(error => err.value = error.message)
      .finally(
        () => loading.value = false,
      )
  }

  async function signup () {
    err.value = null
    loading.value = true
    userStore.signup(username.value, password.value)
      .then(_ => {
        router.push('/')
      })
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

    <v-form ref="form" fast-fail @submit.prevent>
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
        v-if="form && !loading"
        block
        class="mb-8"
        color="blue"
        size="large"
        type="submit"
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

      <p class="text-body-2 text-center">Or</p>
      <v-divider/>
      <br>

      <v-btn
        v-if="form && !loading"
        block
        class="mb-8"
        color="blue"
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

<style scoped lang="sass">
</style>
