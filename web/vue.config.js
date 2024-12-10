const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/static/', // Use relative paths for assets
  devServer: {
    proxy: {
      '/articles/popular': {
        target: 'http://localhost:8080', // Go backend
        changeOrigin: true,
      },
      '/article': {
        target: 'http://localhost:8080', // Go backend
        changeOrigin: true,
      },
    },
  },
});
