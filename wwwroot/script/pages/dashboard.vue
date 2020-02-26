<template>
<div class="container">
  <div class="row">
    <div class="col-sm">
      <patient 
      :results="patientGeneral"
      />
      <tabs />
    </div>
    <div class="col-sm">
      Data Presented
      <router-view/>
    </div>
    <div class="col-sm">
        <links />
    </div>
  </div>
</div>  
</template>

<script>
import patient from '../components/General/patient.vue'
import tabs from '../components/General/tabs.vue'
import links from '../components/General/links.vue'

export default {
  components: {
      patient,
      tabs,
      links
  },
  props: {
    patientGeneral: Object
  },
  mounted(){
    var index = this.$route.params.id
    this.patientGeneral=this.$route.params.patients[index].gi
  }
}
</script>