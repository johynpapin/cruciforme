<template>
  <validation-observer v-slot="{ handleSubmit }">
    <form @submit.prevent="handleSubmit(verify)">
      <h1 class="text-3xl mb-3 font-semibold text-gray-700">Account verification</h1>

      <p class="mb-6">Simply enter your password to complete the verification and sign in.</p>

      <text-input name="Password" type="password" rules="required" v-model="password"></text-input>

      <div class="btns">
        <button class="btn btn-primary" type="submit">Verify my account</button>
      </div>
    </form>
  </validation-observer>
</template>

<script>
import store from '@/store'
import { ValidationObserver } from 'vee-validate'
import TextInput from '@/components/TextInput.vue'

export default {
  name: 'Verify',

  props: {
    verificationToken: String
  },

  data: () => ({
    password: ''
  }),

  methods: {
    async verify () {
      await store.dispatch('auth/verify', {
        verificationToken: this.verificationToken,
        password: this.password
      })
      await store.dispatch('alert/create', {
        type: 'success',
        message: 'Your account has been verified.'
      })

      this.$router.push({
        name: 'Forms'
      })
    }
  },

  components: {
    ValidationObserver,
    TextInput
  }
}
</script>
