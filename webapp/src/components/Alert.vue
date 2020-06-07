<template>
  <v-snackbar :value="currentAlert !== null" :color="alertColor">
    {{ alertMessage }}
  </v-snackbar>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Alert',

  computed: {
    ...mapState('alert', {
      currentAlert: state => state.currentAlert
    }),

    alertColor () {
      if (this.currentAlert === null) {
        return 'info'
      }

      switch (this.currentAlert.type) {
        case 'success':
          return 'success'
        case 'error':
          return 'error'
        case 'info':
        default:
          return 'info'
      }
    },

    alertMessage () {
      if (this.currentAlert === null) {
        return ''
      }

      return this.currentAlert.message
    }
  }
}
</script>
