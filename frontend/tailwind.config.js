/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './features/**/*.{vue,js,ts}',
    './layouts/**/*.{vue,js,ts}',
    './pages/**/*.{vue,js,ts}',
    './widgets/**/*.{vue,js,ts}',
  ],

  theme: {
    screens: {
      'min-small-phone': {
        min: '480px',
      },

      'max-small-phone': {
        max: '480px',
      },

      'min-phone': {
        min: '640px',
      },

      'max-phone': {
        max: '640px',
      },

      /** */

      'min-tablet': {
        min: '768px',
      },

      'max-tablet': {
        max: '768px',
      },

      /** */

      'min-laptop': {
        min: '1024px',
      },

      'max-laptop': {
        max: '1024px',
      },

      /** */

      'min-desktop': {
        min: '1280px',
      },

      'max-desktop': {
        max: '1280px',
      },

      /** */

      'min-large-desktop': {
        min: '1536px',
      },

      'max-large-desktop': {
        max: '1536px',
      },

      /** */

      'min-extra-desktop': {
        min: '1920px',
      },

      'max-extra-desktop': {
        max: '1920px',
      },
    },

    extend: {
      colors: {
        primary: {
          50: '#ecf5ff',
          100: '#ddecff',
          200: '#c2dcff',
          300: '#9dc3ff',
          400: '#76a0ff',
          500: '#557eff',
          600: '#304ff4',
          700: '#2a42d8',
          800: '#253aae',
          900: '#263789',
          950: '#161e50',
          DEFAULT: '#304FF4',
          hover: '#2a42d8',
        },

        secondary: {
          50: '#ffffe4',
          100: '#fdffc4',
          200: '#f9ff90',
          300: '#efff51',
          400: '#dffe08',
          500: '#c3e500',
          600: '#97b700',
          700: '#728b00',
          800: '#5a6d07',
          900: '#4b5c0b',
          950: '#273400',
          DEFAULT: '#DFFE08',
          hover: '#c3e500',
        },

        accent: {
          50: '#F5F6F6',
          100: '#E4E7E9',
          200: '#CCD2D5',
          300: '#A9B2B7',
          400: '#7E8B92',
          500: '#637077',
          600: '#555F65',
          700: '#495055',
          800: '#40454A',
          900: '#393C40',
          950: '#0C0D0E',
          DEFAULT: '#0C0D0E',
          hover: '#393C40',
        },

        tertiary: {
          50: '#F6F7F8',
          100: '#ECEDEF',
          200: '#DCDEE1',
          300: '#C3C7CD',
          400: '#A6ABB4',
          500: '#9094A1',
          600: '#7F8391',
          700: '#727483',
          800: '#60626D',
          900: '#4F5159',
          950: '#333438',
          DEFAULT: '#ECEDEF',
          hover: '#DCDEE1',
        },
      },

      borderRadius: {
        'button-small': '24px',
        'button-medium': '32px',
        'button-large': '36px',

        'hamburger-button-small': '8px',
        'hamburger-button-medium': '12px',
        'hamburger-button-large': '16px',
      },

      fontFamily: {
        sans: '"Manrope", ui-sans-serif, system-ui, -apple-system, "Segoe UI", sans-serif',
      },

      fontSize: {
        phone: '14px',
        tablet: '15px',
        laptop: '16px',
      },

      lineHeight: {
        10: '2.5rem',
        12: '3rem',
        14: '3.5rem',
        16: '4rem',
        18: '4.5rem',
        20: '5rem',
        22: '6rem',
        24: '7rem',
        26: '8rem',
        28: '9rem',
        30: '10rem',
      },
    },
  },
}
