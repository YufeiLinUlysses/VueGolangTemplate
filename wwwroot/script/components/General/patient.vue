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
        data: function(){
            return {
                image:"./Pictures/avatar_icon.jpg"
            }
        }
    })
</script>

<style scoped>
    .redtime {
    background-color:red;
    }
    .welcomeText{
        text-align: center;
        font-size: medium;
    }
    
</style>