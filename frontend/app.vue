<script setup lang="ts">
import {
  AppLoader,
} from './widgets/layouts'

/** */

const isLoadingPage = useState('isLoadingPage', () => false)

/** */

const nuxtApp = useNuxtApp()

nuxtApp.hook('page:loading:start', () => {
  isLoadingPage.value = true
})

nuxtApp.hook('page:loading:end', () => {
  isLoadingPage.value = false
})

/** */

useHead({
  meta: [
    {
      name: 'robots',
      content: 'noindex, nofollow',
    },
  ],
})
</script>

<template>
  <div class="app">
    <AppLoader v-model="isLoadingPage" />

    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>
