<template>
  <validation-provider :name="name" :rules="rules" v-slot="{ errors, classes }" slim>
    <div class="input-group" :class="classes">
      <label>
        <span>{{ name }}</span>
        <input class="input" :type="type" v-model="currentValue">
      </label>

      <span class="input-error" v-if="errors.length > 0">{{ errors[0] }}</span>
    </div>
  </validation-provider>
</template>

<script>
import { ValidationProvider } from 'vee-validate'

export default {
  name: 'TextInput',

  props: {
    value: {
      type: String,
      default: ''
    },
    rules: {
      type: [String, Object],
      default: ''
    },
    name: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: 'text'
    }
  },

  data: () => ({
    currentValue: ''
  }),

  watch: {
    currentValue (value) {
      this.$emit('input', value)
    }
  },

  components: {
    ValidationProvider
  }
}
</script>
