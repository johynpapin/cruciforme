<template>
  <div class="min-h-screen flex items-center justify-center">
    <div class="rounded overflow-hidden shadow-lg bg-white py-8 px-10 w-full max-w-2xl">
      <router-view/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AuthLayout'
}
</script>
