<script lang="ts" setup>
import {
  WavePlayer,
  WaveSymbol,
} from '@/features/wave'
</script>

<template>
  <section id="wave" class="wave">
    <WaveSymbol class="wave__symbol">
      <WavePlayer />
    </WaveSymbol>
  </section>
</template>

<style lang="scss" scoped>
.wave {
  @apply flex;
  @apply items-center;
  @apply justify-center;
  @apply overflow-hidden;
  @apply h-[825px];

  background: linear-gradient(to bottom, #0906AF, #9AC0F1);

  @screen max-extra-desktop {
    @apply h-[725px];
  }

  @screen max-large-desktop {
    @apply h-[625px];
  }

  @screen max-laptop {
    @apply h-[525px];
  }

  @screen max-tablet {
    @apply h-[525px];
  }
}

.wave__symbol {
  @apply w-2/3;
  @apply h-full;

  @screen max-large-desktop {
    @apply w-[900px];
  }

  @screen max-tablet {
    @apply w-[800px];
  }
}
</style>
