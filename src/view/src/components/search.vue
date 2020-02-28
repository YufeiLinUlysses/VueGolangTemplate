<template>
    <div>
        <h1>Example: 5e5535be9f0a1ceabc5a8f4c</h1>
        <input
        type="text"
        v-model="searchquest.id"
        placeholder="id">
        <button @click.prevent="search()">Search</button>
    </div>
</template>

<script>
import axios from 'axios'

var webcall = axios.create({
  baseURL: 'http://localhost:12345/',
  timeout: 20000,
  withCredentials: false,
  headers: { 'Content-Type': 'application/json' }
})

export default {
  data: function() {
    return {
      results: [],
      searchquest: { id: '' }
    }
  },
  methods: {
    search: function() {
      var vm = this
      var quest = vm.searchquest
      var url = 'person/' + vm.searchquest.id
      console.log(quest)
      webcall.get(url)
        .then(function (response) {
          console.log(response.data)
          vm.$router.push({
            name: 'result',
            params: {
              results: response.data
            }
          })
        })
    }
  }
}
</script>
