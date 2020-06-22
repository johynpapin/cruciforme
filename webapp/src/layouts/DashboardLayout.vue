<template>
  <div class="flex">
    <sidebar></sidebar>

    <div class="bg-gray-100 flex-grow flex flex-col">
      <navbar></navbar>

      <div class="flex-grow px-20">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from '@/components/Sidebar.vue'
import Navbar from '@/components/Navbar.vue'

export default {
  name: 'HomeLayout',

  data: () => ({
    drawer: undefined
  }),

  components: {
    Sidebar,
    Navbar
  }
}
</script>
