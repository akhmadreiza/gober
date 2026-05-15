<!-- src/App.vue -->
<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
*, *::before, *::after { box-sizing: border-box; }
body { margin: 0; padding: 0; }
</style>