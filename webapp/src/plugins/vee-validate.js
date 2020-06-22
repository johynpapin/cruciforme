import { configure, extend } from 'vee-validate'
import { required, email } from 'vee-validate/dist/rules'

configure({
  classes: {
    valid: 'is-valid',
    invalid: 'is-invalid'
  }
})

extend('required', {
  ...required,
  message: 'This field is required'
})

extend('email', email)
