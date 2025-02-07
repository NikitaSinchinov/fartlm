<script setup lang="ts">
import {
  AppLogotype,
} from '@/widgets/common'

/**
 * Меню навигации с элементами и их маршрутами.
 */
const menu = [
  {
    to: 'https://x.com/FartLMtech',
    icon: 'local:x',
  },

  {
    to: '/',
    icon: 'local:dexscreen',
  },

  {
    to: '/',
    icon: 'local:dextools',
  },
]
</script>

<template>
  <div class="app-header">
    <AppLogotype class="app-header-logotype" />

    <div class="app-header-menu">
      <div class="app-header-menu__list">
        <NuxtLink
          v-for="(item, index) in menu"
          :key="index"
          :to="item.to"
          target="_blank"
          class="app-header-menu__item"
        >
          <Icon size="20" mode="svg" :name="item.icon" class="app-header-menu__icon" />
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.app-header {
  @apply flex;
  @apply items-center;
  @apply justify-between;
  @apply py-4;
}
.app-header-logotype {
  @apply text-primary;
}

.app-header-menu__list {
  @apply flex;
  @apply flex-row;
  @apply items-center;
  @apply gap-4;

  @screen max-tablet {
    @apply gap-2;
  }
}

.app-header-menu__item {
  @apply flex;
  @apply items-center;
  @apply justify-center;
  @apply w-12;
  @apply h-12;
  @apply rounded-full;
  @apply border;

  @screen max-tablet {
    @apply w-10;
    @apply h-10;
  }
}

.app-header-menu__icon {
  @apply duration-200;
  @apply cursor-pointer;
  @apply text-accent-300;

  &:hover {
    @apply text-accent;
  }

  &._active {
    @apply text-accent;
  }
}
</style>
