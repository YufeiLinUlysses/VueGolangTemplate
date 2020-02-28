<template>
  <div>
    <search />
  </div>
</template>

<script>
import search from '../components/search.vue'

export default {
  name: 'home',
  // Set up a child component
  components: {
    search
  },
  // Set up the data for the component
  data() {
    return {
      showPetForm: false,
      formData: {
        name: '',
        age: 0,
        species: null
      }
    }
  }
}
</script>
