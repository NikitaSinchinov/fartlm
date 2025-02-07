<script lang="ts" setup>
import {
  AppFooter,
  AppHeader,
} from '@/widgets/layouts'
</script>

<template>
  <div class="app">
    <div class="app-container">
      <header class="app-header">
        <AppHeader />
      </header>
    </div>

    <div class="app-main">
      <slot />
    </div>

    <div class="app-container">
      <footer class="app-footer">
        <AppFooter />
      </footer>
    </div>
  </div>
</template>
