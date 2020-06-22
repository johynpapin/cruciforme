<template>
  <div>
    <h1 class="text-5xl font-bold mb-8">Create a new form</h1>

    <div v-if="currentStep === 1">

      <validation-observer v-slot="{ invalid, handleSubmit }">
        <form @submit.prevent="handleSubmit(createForm)">
          <h2 class="text-xl mb-5">First, let's give it a name!</h2>
          <text-input name="Name" rules="required" v-model="name"></text-input>

          <h2 class="text-xl mb-5 mt-4">Where should it send the answers?</h2>
          <text-input name="Email" rules="required|email" v-model="email"></text-input>

          <div class="btns">
            <button type="input" class="btn btn-primary">Create the form</button>
          </div>
        </form>
      </validation-observer>
    </div>

    <div v-else-if="currentStep === 2">
      <h2 class="text-2xl font-medium mb-5">Waiting for email validation</h2>

      <p class="text-lg mb-5">We just sent an email to the address of the form. Please follow the instructions in that email to continue.</p>

      <div class="btns">
        <router-link tag="button" class="btn btn-primary" :to="{ name: 'Forms' }">I didn't get anything, I want another one!</router-link>

        <router-link tag="button" class="btn btn-secondary" :to="{ name: 'Forms' }">I'll do it later, show me the form!</router-link>
      </div>
    </div>

    <div v-else-if="currentStep === 3">
      <h2 class="text-2xl font-medium mb-5">Waiting for the first response</h2>

      <p class="text-lg mb-5">All right, the form's ready to use. Here's how you do it:</p>

      <ol id="send-help-steps" class="list-decimal mb-5">
        <li>
          <span>Take an HTML form, for example :</span>
          <pre>&lt;form&gt;
  &lt;input type="text" name="username"&gt;

  &lt;button type="submit"&gt;Send&lt;/button&gt;
&lt;/form&gt;</pre>
        </li>

        <li>
          <span>Set the action of the form to <span class="bg-gray-800 text-white rounded px-2 py-1 select-all">{{ formAction }}</span>.</span>
          <pre>&lt;form action="{{ formAction }}"&gt;</pre>
        </li>

        <li>
          <span>That's it, now just send it!</span>
        </li>
      </ol>

      <div class="btns">
        <router-link tag="button" class="btn btn-primary" :to="{ name: 'Forms' }">I'll do it later, show me the form!</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'
import { ValidationObserver } from 'vee-validate'
import TextInput from '@/components/TextInput.vue'

export default {
  name: 'CreateForm',

  data: () => ({
    currentStep: 1,
    name: '',
    email: ''
  }),

  methods: {
    async createForm () {
      const createdForm = await store.dispatch('forms/create', {
        name: this.name,
        email: this.email
      })

      if (createdForm.confirmed) {
        this.currentStep = 3
      } else {
        this.currentStep = 2
      }
    }
  },

  computed: {
    ...mapState('forms', [
      'createdForm'
    ]),

    formAction () {
      return 'https://crucifor.me/' + this.createdForm.id
    }
  },

  components: {
    ValidationObserver,
    TextInput
  }
}
</script>

<style>
#send-help-steps li pre {
  @apply text-gray-600 py-5;
}
</style>
