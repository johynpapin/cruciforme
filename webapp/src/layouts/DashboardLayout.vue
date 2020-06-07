<template>
  <div>
    <v-navigation-drawer app v-model="drawer" clipped>
      <v-list dense nav>
        <v-list-item link :to="{ name: 'Forms' }">
          <v-list-item-icon>
            <v-icon>mdi-view-dashboard</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>My forms</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link :to="{ name: 'CreateForm' }">
          <v-list-item-icon>
            <v-icon>mdi-plus-circle</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>New form</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template #append>
        <v-divider></v-divider>

        <div class="px-4 py-2 d-flex">
          <v-spacer></v-spacer>

          <v-btn flat icon round>
            <v-icon>mdi-account</v-icon>
          </v-btn>

          <v-btn flat icon round>
            <v-icon>mdi-logout</v-icon>
          </v-btn>
        </div>
      </template>
    </v-navigation-drawer>

    <v-app-bar app clipped-left flat id="app-bar">
      <v-app-bar-nav-icon class="hidden-lg-and-up" @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>Cruciforme</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <router-view/>
      </v-container>
    </v-content>
  </div>
</template>

<script>
export default {
  name: 'HomeLayout',

  data: () => ({
    drawer: undefined
  })
}
</script>

<style>
#app-bar {
  border-bottom-color: rgba(0,0,0,.12) !important;
  border-width: 0 0 thin;
  border-style: solid;
  background-color: #fff !important;
}
</style>
