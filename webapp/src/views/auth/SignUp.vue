<template>
  <v-row align="center" justify="center" class="fill-height">
    <v-col cols="11" sm="8" md="4">
      <v-card class="elevation-12">
        <v-toolbar color="primary" dark flat>
          <v-toolbar-title>Sign up</v-toolbar-title>
        </v-toolbar>

        <v-form @submit.prevent="submit">
          <v-card-text>
            <v-text-field
              label="Email"
              prepend-icon="mdi-account"
              type="email"
              v-model="email"
            ></v-text-field>

            <v-text-field
              label="Password"
              prepend-icon="mdi-lock"
              type="password"
              v-model="password"
            ></v-text-field>

              <v-text-field
              label="Password verification"
              prepend-icon="mdi-lock"
              type="password"
              v-model="passwordVerification"
            ></v-text-field>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn color="primary" type="submit">Sign up</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'SignUp',

  data: () => ({
    email: '',
    password: '',
    passwordVerification: ''
  }),

  methods: {
    submit () {
      // TODO: handle sign up
    }
  }
}
</script>
