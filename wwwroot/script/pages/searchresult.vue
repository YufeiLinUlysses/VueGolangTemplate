<template>
  <div class="container text-center">
    <h1 class="display-4">Search Result</h1>
    <table class="table">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">First Name</th>
          <th scope="col">Last Name</th>
          <th scope="col">DOB</th>
          <th scope="col">OCVID</th>
          <th scope="col">MR#</th>
        </tr>
      </thead>
      
      <tbody 
      v-for="(pat,index) in patients"
      :key="index"
      >
        <tr>
          <th scope="row">{{index}}</th>
          <router-link :to="{name:'Dashboard',params:{patients:patients, id:index}}">
            {{pat.gi.fName}}
          </router-link>
          <td>{{pat.gi.lName}}</td>
          <td>{{pat.gi.dob}}</td>
          <td>{{pat.gi.ocvid}}</td>
          <td>{{pat.gi.mrNum}}</td>
        </tr>
      </tbody>
    </table>
  </div>
  
</template>

<script>
export default {
  props: {
    patients: Array
  },
  mounted(){
    this.patients=this.$route.params.searchresults.patients
  }
}
</script>