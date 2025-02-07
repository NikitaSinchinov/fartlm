import antfu from '@antfu/eslint-config'

export default antfu(
  {
    vue: true,
    typescript: true,
  },

  {
    rules: {
      'indent': [
        'error',
        2,
        {
          ArrayExpression: 'first',
          ObjectExpression: 'first',
          ImportDeclaration: 'first',
          flatTernaryExpressions: false,
        },
      ],

      'comma-style': [
        'error',
        'last',
      ],

      'object-property-newline': [
        'error',
        {
          allowMultiplePropertiesPerLine: false,
        },
      ],

      'object-curly-spacing': [
        'error',
        'always',
      ],

      'object-curly-newline': [
        'error',
        {
          ObjectExpression: {
            multiline: true,
            consistent: true,
            minProperties: 1,
          },
          ObjectPattern: {
            multiline: true,
            consistent: true,
            minProperties: 1,
          },
          ImportDeclaration: {
            multiline: true,
            consistent: true,
            minProperties: 1,
          },
          ExportDeclaration: {
            multiline: true,
            consistent: true,
            minProperties: 1,
          },
        },
      ],

      'array-element-newline': [
        'error',
        {
          ArrayExpression: 'always',
          ArrayPattern: {
            multiline: true,
          },
        },
      ],

      'array-bracket-spacing': [
        'error',
        'never',
      ],

      'array-bracket-newline': [
        'error',
        {
          multiline: true,
          minItems: null,
        },
      ],

      'function-paren-newline': [
        'error',
        'multiline',
      ],

      'template-curly-spacing': [
        'error',
        'never',
      ],

      'vue/custom-event-name-casing': 'off',
    },
  },
)
