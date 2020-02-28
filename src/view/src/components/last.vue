<template>
    <h1>{{lastname}}</h1>
</template>

<script>
export default {
  data: function() {
    return {
      lastname: ''
    }
  },
  mounted() {
    this.lastname = this.$route.params.lastn
  }
}
</script>
