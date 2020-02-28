<template>
  <div>
    <search />
  </div>
</template>

<script>
import search from '../components/search.vue'

export default {
  name: 'home',
  components: {
    search
  },
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
