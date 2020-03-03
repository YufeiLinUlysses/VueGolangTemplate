<template>
    <div>
        <h1>Example: 5e5e79d4661ace168d5ab66d</h1>
        <input
        type="text"
        v-model="searchquest.id"
        placeholder="id">
        <button @click.prevent="search()">Search</button>
    </div>
</template>

<script>
import axios from 'axios'

// Set up an axios instance
var webcall = axios.create({
  baseURL: 'http://localhost:12345/',
  timeout: 20000,
  withCredentials: false,
  headers: { 'Content-Type': 'application/json' }
})

export default {
  // Set up data for the component
  data: function() {
    return {
      results: [],
      searchquest: { id: '' }
    }
  },
  methods: {
    // Defin the function search
    search: function() {
      var vm = this
      var quest = vm.searchquest
      var url = 'person/' + vm.searchquest.id
      console.log(quest)
      webcall.get(url)
        .then(function (response) {
          console.log(response.data)
          // Push the app to a second view
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
