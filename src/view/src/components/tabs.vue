<template>
<div id='tabs'>
    <button type="button" @click.prevent="first()">{{person.firstname}}</button>
    <button type="button" @click.prevent="last()">{{person.lastname}}</button>
</div>
</template>

<script>
export default {
  props: ['person'],
  methods: {
    first: function() {
      this.$router.push({
        name: 'First',
        params: {
          firstn: this.person.firstname
        }
      })
    },
    last: function() {
      this.$router.push({
        name: 'Last',
        params: {
          lastn: this.person.lastname
        }
      })
    }
  },
  mounted() {
    console.log(this.person.firstname)
  }
}
</script>
