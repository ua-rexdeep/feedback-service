module.exports = {
   env: {
      node: true,
      jest: true,
   },
   extends: ['airbnb-base', 'plugin:jest/recommended', 'plugin:security/recommended'],
   plugins: ['jest', 'security'],

   parserOptions: {
      requireConfigFile: false,
      babelOptions: {
         babelrc: false,
         configFile: false,
      },
   },
   parser: '@babel/eslint-parser',
   rules: {
      'no-console': 'error',
      'func-names': 'off',
      'no-underscore-dangle': 'off',
      'consistent-return': 'off',
      'jest/expect-expect': 'off',
      'security/detect-object-injection': 'off',
      indent: [2, 3],
      'linebreak-style': ['error', 'windows'],
      camelcase: 'off',
      'object-curly-newline': ['error', {
         consistent: true,
         minProperties: 5,
         multiline: true,
      }],
      'import/no-extraneous-dependencies': ['error', { devDependencies: true }],
   },
};
