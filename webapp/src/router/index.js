import Vue from 'vue'
import VueRouter from 'vue-router'

import HomeLayout from '../layouts/HomeLayout.vue'
import AuthLayout from '../layouts/AuthLayout.vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'

import Home from '../views/home/Home.vue'
import SignIn from '../views/auth/SignIn.vue'
import SignUp from '../views/auth/SignUp.vue'
import AfterSignUp from '../views/auth/AfterSignUp.vue'
import Verify from '../views/auth/Verify.vue'
import CreateForm from '../views/dashboard/CreateForm.vue'
import Forms from '../views/dashboard/Forms.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: HomeLayout,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home
      }
    ]
  },
  {
    path: '/auth/',
    component: AuthLayout,
    children: [
      {
        path: 'sign-in',
        name: 'SignIn',
        component: SignIn
      },
      {
        path: 'sign-up',
        name: 'SignUp',
        component: SignUp
      },
      {
        path: 'after-sign-up',
        name: 'AfterSignUp',
        component: AfterSignUp
      },
      {
        path: 'verify',
        name: 'Verify',
        component: Verify,
        props: (route) => ({
          verificationToken: route.query.token
        })
      }
    ]
  },
  {
    path: '/dashboard/',
    component: DashboardLayout,
    children: [
      {
        path: 'create-form',
        name: 'CreateForm',
        component: CreateForm
      },
      {
        path: 'forms',
        name: 'Forms',
        component: Forms
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
