<template>
    <div class="container text-center">
        <img v-bind:src="image" height="150" width="150">
        <p>{{results.fName}} {{results.lName}}, {{results.age}}</p>
        <p>{{results.gender}}</p>
        <p>{{results.dob}}</p>
        <p>Risk Score: {{results.riskScore}}</p>
        <p>{{results.riskLevel}}</p>
        <p>{{results.contact}}</p>
        <p>{{results.ocvid}}</p>
        <p>{{results.mrNum}}</p>
    </div>
</template>

<script>
export default({
    props:{
        results: Object,
    },
    mounted(){
        var index = this.$route.params.id
        this.patientGeneral=this.$route.params.patients[index].gi
    }
})
</script>
