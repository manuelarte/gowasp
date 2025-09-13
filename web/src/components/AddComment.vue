<script setup lang="ts">
  import type { Comment, Post } from '@/models/posts.model'
  import { getInitials } from '@/models/users.model.ts'
  import router from '@/router'
  import { backendClient, useUserStore } from '@/stores/app'

  const props = defineProps({
    post: {
      type: Object as PropType<Post>,
      required: true,
    },
    csrf: {
      type: String,
      required: true,
      validator: (v: string, _) => v.length > 0,
    },
  })

  const emits = defineEmits<{
    // Event to notify the parent component that the comment has been saved.
    'comment:saved': [Comment]
  }>()

  function submit () {
    const newComment = { comment: commentContent.value, userId: user.value!.id, csrf: props.csrf }
    isSaving.value = true
    errorSaving.value = null
    backendClient.postPostComment(props.post.id, newComment)
      .then(c => {
        errorSaving.value = null
        emits('comment:saved', c)
      })
      .catch(error => errorSaving.value = error.message)
      .finally(() => isSaving.value = false)
  }

  // TODO(manuelarte): this is all over the place
  const userStore = useUserStore()
  userStore.$subscribe((_, state) => {
    if (!state.user) {
      router.push('/login')
    }
    user.value = state.user!
  })
  const user = ref(userStore.user!)

  const rules = [(v: string | null) => (v?.length ?? 0) <= 1000 || 'Max 1000 characters']

  const valid = ref(false)
  const isSaving = ref(false)
  const errorSaving = ref<string | null>(null)
  const commentContent = ref('')
</script>

<template>
  <v-card title="My comment">
    <template #prepend>
      <v-avatar
        color="brown"
        size="large"
      >
        <span class="text-h5">{{ getInitials(user) }}</span>
      </v-avatar>
    </template>
    <v-form v-model="valid" :type="submit" @submit.prevent="submit">
      <v-card-text>
        <v-textarea
          v-model="commentContent"
          clearable
          counter
          label="Write your comment"
          :rules="rules"
        />
      </v-card-text>

      <v-card
        v-if="errorSaving"
        class="mb-12"
        color="red-lighten-2"
        variant="tonal"
      >
        <v-card-text class="text-medium-emphasis text-caption">
          {{ errorSaving! }}
        </v-card-text>
      </v-card>

      <v-card-actions>
        <v-btn :disabled="!valid || isSaving" type="submit">Save</v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<style scoped lang="sass">

</style>
