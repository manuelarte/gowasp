<script setup lang="ts">
  import type { Comment, Post } from '@/models/posts.model'
  import { backendClient } from '@/stores/app'

  const props = defineProps({
    post: {
      type: Object as PropType<Post>,
      required: true,
    },
  })

  const emits = defineEmits<{
    // Event to notify the parent component that the comment has been saved.
    'comment:saved': [Comment]
  }>()

  function submit () {
    const newComment = { comment: commentContent.value, userId: 1 }
    isSaving.value = true
    backendClient.postPostComment(props.post.id, newComment)
      .then(c => {
        emits('comment:saved', c)
      })
      .finally(() => isSaving.value = false)
  }

  const rules = [(v: string | null) => (v?.length ?? 0) <= 1000 || 'Max 1000 characters']

  const valid = ref(false)
  const isSaving = ref(false)
  const commentContent = ref('')
</script>

<template>
  <v-card title="My comment">
    <template #prepend>
      <v-avatar
        color="brown"
        size="large"
      >
        <span class="text-h5">MD</span>
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
      <v-card-actions>
        <v-btn :disabled="!valid || isSaving" type="submit">Save</v-btn>
      </v-card-actions>
    </v-form>
  </v-card>
</template>

<style scoped lang="sass">

</style>
