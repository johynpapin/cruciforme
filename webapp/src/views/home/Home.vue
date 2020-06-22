<template>
  <div class="text-center pt-10 text-white">
    <h1 class="text-5xl mb-2">Cruci<span class="font-semibold">form</span>e</h1>
    <h2 class="text-2xl font-light mb-10">A form endpoint for static websites</h2>

    <div v-if="!signedIn" class="btns">
      <router-link :to="{ name: 'SignUp' }" tag="button" class="btn btn-primary btn-xl mr-4">Sign up</router-link>
      <router-link :to="{ name: 'SignIn' }" tag="button" class="btn btn-secondary btn-xl">Sign in</router-link>
    </div>

    <div v-else>
      <p class="text-3xl mb-5">Welcome back!</p>

      <div class="btns">
        <router-link :to="{ name: 'Forms' }" tag="button" class="btn btn-primary btn-xl">My forms</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Home',

  computed: {
    ...mapState('auth', {
      signedIn: state => state.signedIn
    })
  }
}
</script>
