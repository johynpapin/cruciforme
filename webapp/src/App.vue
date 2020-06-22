<template>
  <div class="min-h-screen bg-indigo-900">
    <router-view></router-view>

    <alert></alert>
  </div>
</template>

<script>
import Alert from './components/Alert.vue'

export default {
  name: 'App',

  components: {
    Alert
  }
}
</script>

<style>
/* purgecss start ignore */
@tailwind base;

@tailwind components;
@import './styles/components.css';
/* purgecss end ignore */

@tailwind utilities;
@import './styles/utilities.css';
</style>
