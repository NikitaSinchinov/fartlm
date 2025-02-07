<script lang="ts" setup>
const isLoading = defineModel()
</script>

<template>
  <Transition name="loader">
    <div v-if="isLoading" class="app-loader">
      <Icon size="36" mode="svg" name="local:spinner" />
    </div>
  </Transition>
</template>

<style lang="scss" scoped>
.app-loader {
  @apply fixed;
  @apply top-0;
  @apply left-0;
  @apply flex;
  @apply items-center;
  @apply justify-center;
  @apply w-full;
  @apply h-full;
  @apply transition-all;
  @apply duration-200;
  @apply text-primary;
  @apply bg-white/70;
  @apply z-50;
}
</style>
