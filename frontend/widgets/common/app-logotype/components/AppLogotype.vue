<script lang="ts" setup>
const route = useRoute()

/**
 * Определяет маршрут для навигации.
 */
const to = computed(() => {
  if (route.name === 'index') {
    return undefined
  }

  return {
    name: 'index',
  }
})
</script>

<template>
  <template v-if="to">
    <NuxtLink class="app-logotype" :to="to" @click="emits('route:change-page')">
      <Icon class="app-logotype _mobile" size="24" mode="svg" name="local:logotype" />
      <Icon class="app-logotype _desktop" size="32" mode="svg" name="local:logotype-and-name" />
    </NuxtLink>
  </template>

  <template v-else>
    <div class="app-logotype">
      <Icon class="app-logotype _mobile" size="24" mode="svg" name="local:logotype" />
      <Icon class="app-logotype _desktop" size="32" mode="svg" name="local:logotype-and-name" />
    </div>
  </template>
</template>

<style lang="scss" scoped>
.app-logotype {
  &._mobile {
    @apply block;

    @screen min-tablet {
      @apply hidden;
    }
  }

  &._desktop {
    @apply hidden;

    @screen min-tablet {
      @apply block;
    }
  }
}
</style>
