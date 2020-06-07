<template>
  <v-row class="flex-column">
    <v-col>
      <h1 class="display-1">My forms</h1>
    </v-col>

    <v-col>
      <v-row>
        <v-col v-for="form in forms" :key="form.id">
          <v-card>
            <v-card-title>Hey</v-card-title>
          </v-card>
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script>
import store from '@/store'
import { mapState } from 'vuex'

export default {
  name: 'Forms',

  created () {
    this.fetchData()
  },

  methods: {
    fetchData () {
      store.dispatch('forms/get')
    }
  },

  computed: {
    ...mapState('forms', {
      forms: state => state.forms
    })
  }
}
</script>
