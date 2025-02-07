<script lang="ts" setup>
import {
  WavePlayerStatus,
} from '@/features/wave'

const playerStatus = useState<WavePlayerStatus>('playerStatus')

const isRecording = computed(() => {
  return playerStatus.value === WavePlayerStatus.RECORDING
})

const isPlaying = computed(() => {
  return playerStatus.value === WavePlayerStatus.PLAYING
})

const isActiveWave = computed(() => {
  return isRecording.value || isPlaying.value
})
</script>

<template>
  <div class="wave-symbol" :class="{ _fast: isActiveWave, _slow: !isActiveWave }">
    <svg width="1382" height="717" viewBox="0 0 1382 717" fill="none" xmlns="http://www.w3.org/2000/svg">
      <circle class="circle" cx="690.42" cy="358.365" r="358.365" fill="white" fill-opacity="0.3" />
      <path class="right" d="M1072.06 639.38C1072.06 649.101 1078.72 657.084 1086.85 655.994C1179.73 643.534 1252.76 514.769 1252.76 357.759C1252.76 200.749 1179.73 71.9838 1086.85 59.5244C1078.72 58.434 1072.06 66.4168 1072.06 76.1378V639.38Z" fill="white" fill-opacity="0.2" />
      <path class="right-far" d="M1270.95 586.939C1270.95 600.489 1281.73 608.791 1289.83 600.848C1345.25 546.514 1381.31 458.146 1381.31 358.365C1381.31 258.585 1345.25 170.216 1289.83 115.882C1281.73 107.94 1270.95 116.242 1270.95 129.792V586.939Z" fill="white" fill-opacity="0.1" />
      <path class="left" d="M308.035 639.38C308.035 649.101 301.418 657.084 293.344 655.993C201.088 643.534 128.549 514.769 128.549 357.759C128.549 200.749 201.088 71.9837 293.344 59.5243C301.418 58.4338 308.035 66.4167 308.035 76.1377V639.38Z" fill="white" fill-opacity="0.2" />
      <path class="left-far" d="M110.359 586.939C110.359 600.489 99.5829 608.791 91.4813 600.848C36.0584 546.514 -0.000155151 458.146 -0.000155151 358.365C-0.000155151 258.585 36.0584 170.216 91.4813 115.882C99.5829 107.94 110.359 116.242 110.359 129.792V586.939Z" fill="white" fill-opacity="0.1" />
    </svg>

    <div class="wave-symbol__body">
      <slot />
    </div>
  </div>
</template>

<style lang="scss" scoped>
@keyframes moveToCenter {
  0%, 100% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(var(--translate-x));
  }
}

@keyframes circleTransform {
  0%, 100% {
    transform: scale(1);
    transform-origin: center;
  }
  50% {
    transform: scale(0.9);
    transform-origin: center;
  }
}

.wave-symbol {
  @apply relative;
  @apply flex;
  @apply items-center;
  @apply justify-center;

  &._fast {
    .left, .left-far, .right, .right-far {
      animation-duration: 4s;
    }
  }

  &._slow {
    .left, .left-far, .right, .right-far {
      animation-duration: 16s;
    }
  }

  .circle {
    animation-name: moveToCenter;
    animation-iteration-count: infinite;
    animation-timing-function: cubic-bezier(0.5, 0, 0.5, 1);
  }

  .left {
    --translate-x: 50px;
    animation-name: moveToCenter;
    animation-iteration-count: infinite;
    animation-timing-function: cubic-bezier(0.5, 0, 0.5, 1);
  }

  .left-far {
    --translate-x: 100px;
    animation-name: moveToCenter;
    animation-iteration-count: infinite;
    animation-timing-function: cubic-bezier(0.5, 0, 0.5, 1);
  }

  .right {
    --translate-x: -50px;
    animation-name: moveToCenter;
    animation-iteration-count: infinite;
    animation-timing-function: cubic-bezier(0.5, 0, 0.5, 1);
  }

  .right-far {
    --translate-x: -100px;
    animation-name: moveToCenter;
    animation-iteration-count: infinite;
    animation-timing-function: cubic-bezier(0.5, 0, 0.5, 1);
  }
}

.wave-symbol__body {
  position: absolute;
  @apply top-1/2;
  @apply left-1/2;
  @apply -translate-x-1/2;
  @apply -translate-y-1/2;
}
</style>
