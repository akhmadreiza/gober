const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})

// vue.config.js
module.exports = {
  devServer: {
    proxy: {
      '/articles/popular': {
        target: 'http://localhost:8080', // Go backend
        changeOrigin: true,
        // pathRewrite: {
        //   '^/articles/popular': '/articles/popular', // Optional, you can adjust the path if needed
        // },
      },
    },
  },
};
