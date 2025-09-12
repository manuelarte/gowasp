<script setup lang="ts">
  const rules = [(v: string) => v.length <= 1000 || 'Max 1000 characters']
</script>

<template>
  <v-card>
    <v-card-title>Add comment</v-card-title>
    <v-card-text>
      <v-textarea
        clearable
        counter
        label="Write your comment"
        :rules="rules"
      />
    </v-card-text>
    <v-card-actions>
      <v-btn>Save</v-btn>
    </v-card-actions>
  </v-card>
</template>

<style scoped lang="sass">

</style>
