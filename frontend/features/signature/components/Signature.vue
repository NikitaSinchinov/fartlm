<template>
  <div class="signature">
    <div class="app-container">
      <div class="signature__body">
        <div class="signature__title">
          CA:
        </div>

        <div class="signature__value">
          Not live yet
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.signature {
    @apply bg-secondary;
    @apply text-accent;
}

.signature__body {
    @apply flex;
    @apply items-center;
    @apply gap-2;
    @apply py-4;
}

.signature__title {
    @apply font-bold;
    @apply text-xl;
}

.signature__value {
    @apply text-xl;
}
</style>
