<script setup lang="ts">
import type {
  NuxtError,
} from '#app'
import {
  NuxtLayout,
} from '#components'

defineProps({
  error: Object as () => NuxtError,
})
</script>

<template>
  <NuxtLayout name="error">
    <div class="error-page">
      <div class="error">
        <h1 class="error__title">
          {{ error.statusCode }} ERROR
        </h1>

        <NuxtLink class="error__link" to="/">
          Go back home
        </NuxtLink>
      </div>
    </div>
  </NuxtLayout>
</template>

<style lang="scss" scoped>
.error-page {
  @apply flex;
  @apply flex-col;
  @apply items-center;
  @apply justify-center;
  @apply grow;
  @apply shrink;
  @apply bg-accent-100;
  @apply text-accent;
}

.error {
  @apply flex;
  @apply flex-col;
  @apply items-center;
  @apply justify-center;
  @apply gap-10;
  @apply grow;
  @apply shrink;
  @apply w-full;
  @apply h-full;
}

.error__title {
  @apply text-8xl;

  @screen max-laptop {
    @apply text-7xl;
  }

  @screen max-tablet {
    @apply text-6xl;
  }

  @screen max-phone {
    @apply text-5xl;
  }
}

.error__link {
  @apply app-button;
  @apply bg-secondary;
  @apply text-accent;
}
</style>
