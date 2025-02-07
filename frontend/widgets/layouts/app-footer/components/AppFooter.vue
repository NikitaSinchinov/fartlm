<script setup lang="ts">
import {
  AppLogotype,
} from '@/widgets/common'
</script>

<template>
  <div class="app-footer">
    <AppLogotype />

    <div class="app-footer__body">
      <p class="app-footer__description">
        The world’s first voice-to-fart neural network.
      </p>

      <p class="app-footer__rights">
        © {{ new Date().getFullYear() }} FartLM. All rights reserved. $FLM Token.
      </p>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.app-footer {
  @apply flex;
  @apply flex-col;
  @apply gap-8;
  @apply py-6;
  @apply text-primary;

  .app-footer__body {
    @apply flex;
    @apply flex-col;
    @apply gap-4;
  }

  .app-footer__description {
    @apply text-accent-950;
  }

  .app-footer__rights {
    @apply text-accent-500;
  }
}
</style>
