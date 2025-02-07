<script setup lang="ts">
import {
  Hero,
} from '@/features/hero'

import {
  Signature,
} from '@/features/signature'

import {
  Wave,
} from '@/features/wave'
</script>

<template>
  <div class="page">
    <div class="app-container">
      <Hero class="hero" />
    </div>

    <Signature class="signature" />

    <Wave class="wave" />
  </div>
</template>
