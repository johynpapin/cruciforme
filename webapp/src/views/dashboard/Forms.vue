<template>
  <div>
    <h1 class="text-5xl font-bold mb-8">My forms</h1>

    <div class="flex flex-wrap">
      <div class="w-1/3 p-2" v-for="form in forms" :key="form.id">
        <div class="card p-2">
          <h2 class="text-lg font-medium text-gray-800">{{ form.name }}</h2>
        </div>
      </div>
    </div>
  </div>
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
