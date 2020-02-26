<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import axios from 'axios'
var webcall = axios.create({
        baseURL: 'api',
        timeout: 20000,
        withCredentials: false,
        headers: { 'Content-Type': 'application/json' }
    });
export default {
  
}
</script>
