<template>
  <div class="w-full flex pt-5 pr-5">
    <div class="flex-grow"></div>

    <div>
      <button @click="logout" class="btn btn-secondary">Log out</button>
    </div>
  </div>
</template>

<script>
import store from '@/store'

export default {
  name: 'Navbar',

  methods: {
    async logout () {
      await store.dispatch('auth/logout')
      this.$router.push({
        name: 'Home'
      })
    }
  }
}
</script>
