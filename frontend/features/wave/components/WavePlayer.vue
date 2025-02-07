<script lang="ts" setup>
import {
  WavePlayerStatus,
} from '@/features/wave'

const playerStatus = useState<WavePlayerStatus>('playerStatus', () => WavePlayerStatus.EMPTY)

const mediaStream = ref<MediaStream>()
const mediaRecorder = ref<MediaRecorder>()
const mediaRecorderTimer = ref<ReturnType<typeof setTimeout> | null>(null)
const mediaPlayer = ref<HTMLAudioElement>()
const mediaPlayerChunks = ref<Blob[]>([])
const recordedBlob = ref<Blob>()
const responseAudioBlob = ref<Blob>()
const responseAudioURL = ref<string>()
const isResponseError = ref(false)

async function createMediaRecorder() {
  mediaStream.value = await navigator.mediaDevices.getUserMedia({
    audio: true,
  })

  mediaRecorder.value = new MediaRecorder(mediaStream.value)

  mediaRecorder.value.ondataavailable = onDataAvailable
  mediaRecorder.value.onstop = onStopRecording
}

function onDataAvailable(event: BlobEvent) {
  mediaPlayerChunks.value.push(event.data)
}

async function onStopRecording() {
  recordedBlob.value = new Blob(mediaPlayerChunks.value, {
    type: 'audio/wav',
  })

  await uploadAudio()

  if (mediaRecorderTimer.value) {
    clearTimeout(mediaRecorderTimer.value)
    mediaRecorderTimer.value = null
  }
}

function stopMediaStreamTracks() {
  if (!mediaStream.value) {
    return
  }

  mediaStream.value.getTracks().forEach(track => track.stop())
}

function stopMediaPlayer() {
  if (!mediaPlayer.value) {
    return
  }

  mediaPlayer.value.pause()
  mediaPlayer.value.currentTime = 0
}

async function uploadAudio() {
  if (!recordedBlob.value) {
    return
  }

  const formData = new FormData()

  /** */

  formData.append('audio', recordedBlob.value, 'recording.wav')

  playerStatus.value = WavePlayerStatus.LOADING

  try {
    const data = await $fetch<Blob>('https://fartlm.com/fart', {
      method: 'POST',
      body: formData,
    })

    if (data) {
      responseAudioBlob.value = new Blob([data])

      if (responseAudioBlob.value) {
        responseAudioURL.value = URL.createObjectURL(responseAudioBlob.value)

        if (mediaPlayer.value) {
          mediaPlayer.value.src = responseAudioURL.value
        }
      }

      playerStatus.value = WavePlayerStatus.SAVED
    }
  }
  catch {
    isResponseError.value = true

    playerStatus.value = WavePlayerStatus.EMPTY
  }
}

async function onStart() {
  isResponseError.value = false

  await createMediaRecorder()

  if (mediaRecorder.value) {
    mediaRecorder.value.start()

    /** */
    mediaRecorderTimer.value = setTimeout(onStop, 10000)

    playerStatus.value = WavePlayerStatus.RECORDING
  }
}

function onStop() {
  if (mediaRecorder.value && mediaRecorder.value.state === 'recording') {
    playerStatus.value = WavePlayerStatus.SAVED
    mediaRecorder.value.stop()
    stopMediaStreamTracks()
  }
}

function onPlay() {
  if (mediaPlayer.value) {
    mediaPlayer.value.play()

    mediaPlayer.value.addEventListener('ended', onStopPlay)

    playerStatus.value = WavePlayerStatus.PLAYING
  }
}

function onStopPlay() {
  if (mediaPlayer.value) {
    stopMediaPlayer()

    mediaPlayer.value.removeEventListener('ended', onStopPlay)

    playerStatus.value = WavePlayerStatus.SAVED
  }
}

function onDelete() {
  stopMediaPlayer()

  mediaRecorder.value = undefined

  mediaPlayerChunks.value = []

  recordedBlob.value = undefined
  responseAudioBlob.value = undefined

  playerStatus.value = WavePlayerStatus.EMPTY
}
</script>

<template>
  <section class="wave-player">
    <audio ref="mediaPlayer" class="wave-player-audio" />

    <div class="wave-player__title">
      Speak to Fart
    </div>

    <ul class="wave-player__steps">
      <li>Click Start and allow microphone access</li>
      <li>Speak naturally into your device for up to <b>10 seconds</b></li>
      <li>Listen as your voice becomes an instant masterpiece</li>
    </ul>

    <template v-if="isResponseError">
      <div class="wave-player__errors">
        An error has occurred. Please try again later.
      </div>
    </template>

    <div class="wave-player__actions">
      <TransitionGroup name="fade">
        <template v-if="WavePlayerStatus.EMPTY === playerStatus">
          <button class="player-button _composite" @click="onStart">
            <div class="player-buttons-wrapper">
              Record

              <div class="player-action-button">
                <Icon size="24" mode="svg" name="local:mic" />
              </div>
            </div>
          </button>
        </template>

        <template v-if="WavePlayerStatus.RECORDING === playerStatus">
          <div class="player-buttons-wrapper">
            <button class="player-button _composite" @click="onStop">
              Stop

              <div class="player-action-button _red">
                <Icon size="24" mode="svg" name="local:stop" />
              </div>
            </button>
          </div>
        </template>

        <template v-if="WavePlayerStatus.LOADING === playerStatus">
          <div class="player-buttons-wrapper">
            <button class="player-button _solo" disabled>
              <Icon size="24" mode="svg" name="local:spinner" />
            </button>
          </div>
        </template>

        <template v-if="[WavePlayerStatus.SAVED, WavePlayerStatus.PLAYING].includes(playerStatus)">
          <div class="player-buttons-wrapper">
            <button v-if="WavePlayerStatus.SAVED === playerStatus" class="player-button _composite" @click="onPlay">
              Play

              <div class="player-action-button">
                <Icon size="26" mode="svg" name="local:play" />
              </div>
            </button>

            <button v-if="WavePlayerStatus.PLAYING === playerStatus" class="player-button _composite" @click="onStopPlay">
              Stop

              <div class="player-action-button">
                <Icon size="24" mode="svg" name="local:stop" />
              </div>
            </button>

            <template v-if="responseAudioBlob">
              <a class="player-button _solo" :href="responseAudioURL" :download="`fart-sound-${Date.now()}.wav`">
                <Icon size="24" mode="svg" name="local:save" />
              </a>
            </template>

            <button class="player-button _solo _red" @click="onDelete">
              <Icon size="24" mode="svg" name="local:delete" />
            </button>
          </div>
        </template>
      </TransitionGroup>
    </div>
  </section>
</template>

<style lang="scss" scoped>
.wave-player {
  @apply flex;
  @apply flex-col;
  @apply gap-8;
  @apply items-center;
  @apply justify-center;
}

.wave-player__title {
  @apply text-5xl;
  @apply text-white;

  @screen max-large-desktop {
    @apply text-3xl;
  }

  @screen max-tablet {
    @apply text-3xl;
  }
}

.wave-player__steps {
  @apply text-white;
  @apply text-center;
  @apply list-decimal;
  @apply list-inside;
  @apply space-y-2;

  @screen max-tablet {
    @apply text-xs;
  }
}

.wave-player__actions {
  @apply relative;
  @apply mt-8;

  .player-buttons-wrapper {
    @apply flex;
    @apply items-center;
    @apply justify-center;
    @apply gap-4;

    @screen max-desktop {
      @apply gap-2;
    }

    @screen max-laptop {
      @apply gap-2;
    }
  }

  .player-button {
    @apply flex;
    @apply items-center;
    @apply justify-center;

    &._composite {
      @apply gap-4;
      @apply border;
      @apply rounded-button-medium;
      @apply py-2;
      @apply pl-8;
      @apply pr-6;
      @apply text-lg;
      @apply text-white;
      @apply duration-200;

      &:hover {
        @apply bg-white;
        @apply bg-opacity-10;
      }
    }

    &._solo {
      @apply w-16;
      @apply h-16;
      @apply rounded-full;
      @apply text-black;
      @apply bg-secondary;
      @apply duration-200;

      &:hover {
        @apply bg-secondary-hover;
      }

      @screen max-desktop {
        @apply w-14;
        @apply h-14;
      }

      &._red {
        @apply text-white;
        @apply bg-red-500;

        &:hover {
          @apply bg-red-600;
        }
      }
    }
  }

  .player-action-button {
    @apply flex;
    @apply items-center;
    @apply justify-center;
    @apply w-16;
    @apply h-16;
    @apply rounded-full;
    @apply text-black;
    @apply bg-secondary;
    @apply duration-200;

    &:hover {
      @apply bg-secondary-hover;
    }

    @screen max-desktop {
      @apply w-14;
      @apply h-14;
    }

    &._red {
      @apply text-white;
      @apply bg-red-500;

      &:hover {
        @apply bg-red-600;
      }
    }
  }

  > * {
    @apply absolute;
    @apply top-0;
    @apply left-0;
    @apply -translate-x-1/2;
    @apply -translate-y-1/2;
  }
}

.wave-player__errors {
  @apply px-8;
  @apply py-4;
  @apply border;
  @apply border-red-500;
  @apply text-red-600;
  @apply bg-red-500;
  @apply bg-opacity-20;
  @apply rounded-xl;
}
</style>
