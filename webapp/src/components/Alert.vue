<template>
  <div class="border px-4 py-3 rounded fixed inset-x-0 bottom-0 z-30" :class="alertClasses" role="alert" v-if="currentAlert !== null">
    <span class="block sm:inline">{{ alertMessage }}</span>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Alert',

  computed: {
    ...mapState('alert', {
      currentAlert: state => state.currentAlert
    }),

    alertClasses () {
      if (this.currentAlert === null) {
        return []
      }

      switch (this.currentAlert.type) {
        case 'success':
          return ['bg-green-100', 'border-green-400', 'text-green-700']
        case 'error':
          return ['bg-red-100', 'border-red-400', 'text-red-700']
        case 'info':
        default:
          return ['bg-blue-100', 'border-blue-400', 'text-blue-700']
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
