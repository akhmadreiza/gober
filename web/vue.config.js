const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/static/', // Use relative paths for assets
  devServer: {
    proxy: {
      '/articles/popular': {
        target: process.env.VUE_APP_GOBER_API_URL, // Go backend
        changeOrigin: true,
      },
      '/article': {
        target: process.env.VUE_APP_GOBER_API_URL, // Go backend
        changeOrigin: true,
      },
    },
  },
});
