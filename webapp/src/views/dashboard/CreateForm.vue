<template>
  <v-row class="flex-column">
    <v-col>
      <h1 class="display-1">Create a new form</h1>
    </v-col>

    <v-col>
      <v-stepper v-model="currentStep" vertical>
        <v-stepper-step :complete="currentStep > 1" step="1" editable>
          Set up the form
        </v-stepper-step>

        <v-stepper-content step="1">
          <v-skeleton-loader :loading="fetchingData" type="article" transition="fade-transition">
            <div>
              <p>The identifier of your form is <kbd>{{ createdFormId }}</kbd>. All you have to do is point your form to <kbd>https://crucifor.me/{{ createdFormId }}</kbd>. Here is an example of a form using it:</p>
              <pre>
    &lt;form action="https://crucifor.me/{{createdFormId}}"&gt;
      &lt;input type="text" name="email"&gt;

      &lt;input type="submit" value="Send"&gt;
    &lt;/form&gt;
              </pre>

              <v-btn color="primary" @click="currentStep = 2">It's done!</v-btn>
            </div>
          </v-skeleton-loader>
        </v-stepper-content>

        <v-stepper-step :complete="currentStep > 2" step="2">
          Send it
        </v-stepper-step>

        <v-stepper-content step="2">
          <p>All you have to do now is send the form. It will then be available in the list of your forms.</p>

          <v-btn color="primary" :to="{ name: 'Forms' }">Show me that list</v-btn>
        </v-stepper-content>
      </v-stepper>
    </v-col>
  </v-row>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'

export default {
  name: 'CreateForm',

  data: () => ({
    fetchingData: true,
    currentStep: 1
  }),

  created () {
    this.fetchData()
  },

  methods: {
    async fetchData () {
      await store.dispatch('forms/create')

      this.fetchingData = false
    }
  },

  computed: {
    ...mapState('forms', {
      createdFormId: state => state.createdFormId
    })
  }
}
</script>
