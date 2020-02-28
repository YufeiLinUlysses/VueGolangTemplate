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
  components: {
    tabs
  },
  data: function() {
    return {
      person: Object
    }
  },
  mounted() {
    this.person = this.$route.params.results
    console.log(typeof this.person)
  }
}
</script>
