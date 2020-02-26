<template>
<div id="search">
    <input
        type="text"
        aria-label="Sizing example input"
        aria-describedby="inputGroup-sizing-lg"
        v-model="searchquest.id"
        placeholder="id">
    <button
        @click.prevent="verifyInput()">
        Search
    </button>
</div>
</template>

<script>
import axios from 'axios'

var webcall = axios.create({
  baseURL: 'localhost://12345',
  timeout: 20000,
  withCredentials: false,
  headers: { 'Content-Type': 'application/json' }
})

export default {
  data: function () {
    return {
      results: [],
      searchquest: { id: '' }
    }
  },
  methods: {
    search: function () {
      var vm = this
      var quest = vm.searchquest
      webcall.post('searchpatient', quest)
        .then(function (response) {
          console.log(response.data.found)
          if (response.data.found) {
            vm.results = response.data
            // This routes to a next page
            vm.$router.push({ name: 'SearchResult', params: { searchresults: vm.results } })
            vm.$emit('chosenpatient', vm.results)
            console.log(vm.results)
          } else {
            alert('The patient does not exist, please contact person in charge')
          }
        })
        .catch(function (error) {
          console.log(error)
        })
      // This would refresh the models the compnents uses
      vm.searchquest = { id: '' }
    }
  }
}
</script>
