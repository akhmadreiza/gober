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
/* Global style for responsive images */
.article-body img {
  max-width: 100% !important;
  height: auto !important;
  display: block;
  margin: 10px 0;
  object-fit: contain; /* Ensure the image fits within the container */
}
</style>