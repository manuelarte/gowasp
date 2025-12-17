<template>
  <v-toolbar color="primary">
    <v-toolbar-title text="GOwasp App" />
    <template #prepend>
      <v-btn v-tooltip="'Home'" icon="mdi-home" to="/" />
    </template>
    <template #append>
      <div v-if="!user" class="d-flex ga-1">
        <v-btn v-if="!user" v-tooltip="'Login'" icon="mdi-login" to="/login" />
      </div>
      <v-menu v-else>
        <template v-slot:activator="{ props }">
          <v-avatar
            id="avatar"
            color="secondary"
            size="large"
            v-bind="props"
          >
            <span class="text-h5">{{ getInitials(user) }}</span>
          </v-avatar>
        </template>

        <v-list>
          <v-list-item v-if="!user" to="/login" append-icon="mdi-login">
            <v-list-item-title>
              Login
            </v-list-item-title>
          </v-list-item>
          <v-list-item v-else @click="userStore.logout()" append-icon="mdi-logout">
            <v-list-item-title>
              Logout
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
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
  import {getInitials} from "@/models/users.model.ts";

  // TODO(manuelarte): this is all over the place
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
