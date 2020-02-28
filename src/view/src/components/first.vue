<template>
    <h1>{{firstname}}</h1>
</template>

<script>
export default {
  data: function() {
    return {
      firstname: ''
    }
  },
  mounted() {
    this.firstname = this.$route.params.firstn
  }
}
</script>
