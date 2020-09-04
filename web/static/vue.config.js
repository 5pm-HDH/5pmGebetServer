const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
    // ...
    entry: './src/index.js',
    mode: 'production',
    resolve: {
      alias: {
        'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' for webpack 1
      }
    },
    module: {
        rules: [
            { test: /\.vue$/, use: 'vue-loader'}
        ]
    },
    plugins: [
        // make sure to include the plugin for the magic
        new VueLoaderPlugin()
    ]
  }