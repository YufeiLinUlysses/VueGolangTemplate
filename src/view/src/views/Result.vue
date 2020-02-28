<template>
    <div>
        <h1>Hi {{person.firstname}}</h1>
        <tabs :person="person"/>
        <router-view />
    </div>
</template>

<script>
import tabs from '../components/tabs.vue'
export default {
  // Set up a child component
  components: {
    tabs
  },
  // Set up data for the component
  data: function() {
    return {
      person: Object
    }
  },
  mounted() {
    // Get the parameter in the url
    this.person = this.$route.params.results
    console.log(typeof this.person)
  }
}
</script>
