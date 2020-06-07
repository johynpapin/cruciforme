<template>
  <v-app>
    <router-view></router-view>

    <alert></alert>
  </v-app>
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
