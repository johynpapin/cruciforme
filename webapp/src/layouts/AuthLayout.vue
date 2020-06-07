<template>
  <v-content>
    <v-container fluid class="fill-height">
      <v-row no-gutters class="fill-height flex-column">
        <v-col class="shrink">
          <v-btn :to="{ name: 'Home' }" depressed>Go back to Home</v-btn>
        </v-col>

        <v-col class="grow">
          <router-view/>
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
export default {
  name: 'AuthLayout'
}
</script>
