<template>
  <v-content>
    <v-row>
      <v-spacer/>

      <v-col class="shrink">
        <v-btn :to="{ name: 'SignIn' }" depressed>Sign in</v-btn>
      </v-col>

      <v-col class="shrink">
        <v-btn :to="{ name: 'SignUp' }" depressed color="primary">Sign up</v-btn>
      </v-col>
    </v-row>

    <v-container fluid>
      <router-view/>
    </v-container>
  </v-content>
</template>

<script>
export default {
  name: 'HomeLayout'
}
</script>
