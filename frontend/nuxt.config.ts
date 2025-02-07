// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',

  devtools: {
    enabled: true,
  },

  app: {
    head: {
      title: 'FartLM | AI-Powered Voice Transformation Platform',
      meta: [
        {
          name: 'keywords',
          content: 'FartLM, voice transformation, AI audio converter, blockchain audio, sound processing, voice modulation, neural network audio, speech transformation, $FLM token, audio AI technology',
        },
        {
          name: 'description',
          content: 'Transform your voice with advanced AI technology. Real-time audio processing meets blockchain innovation. Experience the future of sound transformation',
        },
        {
          name: 'viewport',
          content: 'width=device-width, initial-scale=1',
        },
        {
          name: 'apple-mobile-web-app-title',
          content: 'FartLM',
        },
        {
          property: 'og:title',
          content: 'FartLM | AI-Powered Voice Transformation Platform',
        },
        {
          property: 'og:description',
          content: 'Transform your voice with advanced AI technology. Real-time audio processing meets blockchain innovation. Experience the future of sound transformation',
        },
        {
          property: 'og:image',
          content: '/og-image.png',
        },
        {
          property: 'og:url',
          content: 'https://fartlm.com/',
        },
        {
          property: 'og:type',
          content: 'website',
        },
        {
          property: 'language',
          content: 'en',
        },
      ],
      link: [
        {
          rel: 'icon',
          type: 'image/png',
          href: '/favicon-96x96.png',
          sizes: '96x96',
        },
        {
          rel: 'icon',
          type: 'image/svg+xml',
          href: '/favicon.svg',
        },
        {
          rel: 'shortcut icon',
          href: '/favicon.ico',
        },
        {
          rel: 'apple-touch-icon',
          sizes: '180x180',
          href: '/apple-touch-icon.png',
        },
        {
          rel: 'manifest',
          href: '/site.webmanifest',
        },
      ],
    },
  },

  modules: [
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxtjs/sitemap',
    '@nuxtjs/robots',
    '@vueuse/nuxt',
    'nuxt-og-image',
  ],

  css: ['@/assets/styles/index.scss'],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  site: {
    name: 'Fart LM',
    url: 'https://fartlm.com/',
  },

  robots: {
    enabled: true,
  },

  icon: {
    customCollections: [
      {
        prefix: 'local',
        dir: './assets/icons',
      },
    ],
  },
})
