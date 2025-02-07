<template>
  <section class="hero">
    <div class="hero-action">
      Make Your Speech Fartastic
    </div>

    <div class="hero-description">
      Revolutionary AI that converts any voice input into a symphony of perfectly timed, pitch-matched sound effects while maintaining the original rhythm and melody
    </div>
  </section>
</template>

<style lang="scss" scoped>
.hero {
  @apply flex;
  @apply items-end;
  @apply justify-between;
  @apply gap-8;
  @apply pt-6;
  @apply pb-12;

  @screen max-laptop {
    @apply flex-col;
    @apply items-start;
    @apply gap-6;
    @apply py-8;
  }

  @screen max-tablet {
    @apply py-6;
  }
}

.hero__body {
  @apply flex;
  @apply flex-col;
}

.hero-action {
  @apply max-w-[900px];
  @apply text-7xl;
  @apply leading-20;
  @apply text-accent;
  @apply font-medium;

  @screen max-tablet {
    @apply text-6xl;
    @apply leading-18;
  }

  @screen max-phone {
    @apply text-5xl;
    @apply leading-16;
  }

  @screen max-small-phone {
    @apply text-4xl;
    @apply leading-14;
  }
}

.hero-description {
  @apply max-w-[450px];
  @apply text-accent-400;
}
</style>
