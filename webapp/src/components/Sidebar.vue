<template>
  <nav id="sidebar" class="h-screen z-10 pt-5 flex-none">
    <div class="px-8 mb-10">
      <h1 class="text-3xl tracking-tighter font-bold text-white mb-4">Cruciforme</h1>

      <router-link tag="button" :to="{ name: 'CreateForm' }" class="btn rounded-tl-none rounded-full bg-white w-full inline-flex items-center justify-center" id="new-form-btn">
        <i class="material-icons mr-1">add</i>
        <span>New Form</span>
      </router-link>
    </div>

    <ul>
      <router-link tag="li" :to="{ name: 'Forms' }" class="w-full cursor-pointer"><i class="material-icons">dashboard</i> <span>My forms</span></router-link>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'Sidebar'
}
</script>

<style>
#sidebar {
  width: 18rem;
  box-shadow: 5px 0 10px 0 rgba(0,0,0,.15) !important;
}

#new-form-btn {
  color: #7754f8;
}

#sidebar ul li {
  @apply pl-8 h-10 font-semibold text-lg inline-flex items-center border-l-4 text-white duration-75;
  transition-property: border-width;
}

#sidebar ul li:hover {
  @apply border-l-8;
}

#sidebar ul li i {
  @apply mr-3;
}
</style>
