<script setup lang="ts">
import { RouterView } from 'vue-router'
import {ref} from "vue";
import type {User} from "@/models/user.ts";

const user = ref<User | null>(null)
</script>

<template>
  <div class="navbar bg-base-100 shadow-sm">
    <div class="flex-1">
      <RouterLink class="btn btn-ghost text-xl" to="/">GOwasp Application</RouterLink>
    </div>
    <div class="flex gap-2">
      <div v-if="user != null" class="dropdown dropdown-end">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
          <div class="w-10 rounded-full">
            <img
              alt="Tailwind CSS Navbar component"
              src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
          </div>
        </div>
        <ul
          tabindex="0"
          class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
          <li><a>Logout</a></li>
        </ul>
      </div>
      <div v-else role="button" class="btn btn-ghost btn-circle avatar">
        <div class="w-10 rounded-full">
          <RouterLink to="/login">
            <img
              alt="Login link"
              src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
          </RouterLink>
        </div>
      </div>
    </div>
  </div>

  <RouterView />
</template>

<style scoped>
</style>
