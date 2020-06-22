<template>
  <validation-observer v-slot="{ handleSubmit }">
    <form @submit.prevent="handleSubmit(signUp)">
      <h1 class="text-3xl mb-6 font-semibold text-gray-700">Sign up</h1>

      <text-input name="Email" type="email" rules="required|email" v-model="email"></text-input>
      <text-input name="Password" type="password" rules="required" v-model="password"></text-input>
      <text-input name="Password verification" type="password" rules="required"></text-input>

      <div class="btns">
        <button class="btn btn-primary" type="submit">Sign up</button>
      </div>
    </form>
  </validation-observer>
</template>

<script>
import store from '@/store'
import { ValidationObserver } from 'vee-validate'
import TextInput from '@/components/TextInput.vue'

export default {
  name: 'SignUp',

  data: () => ({
    email: '',
    password: '',
    passwordVerification: ''
  }),

  methods: {
    async signUp () {
      const user = {
        email: this.email,
        password: this.password
      }

      await store.dispatch('auth/signUp', user)

      this.$router.push({
        name: 'AfterSignUp'
      })
    }
  },

  components: {
    ValidationObserver,
    TextInput
  }
}
</script>
