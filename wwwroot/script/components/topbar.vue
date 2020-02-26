<template>
    <div :v-if="data.found==true" class="container">
        <div class ="row">
            <div class="col">
                <div>
                    <img v-bind:src="image">
                </div>       
            </div>

            <div class="col">
                <div class="welcomeText">
                    <p>Welcome, {{data.fName}}'s {{num}} time Visit</p>
                    <p>Next appointment: {{nextApp}}</p>
                    <p>Has been Waiting {{time}} mins</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default({
        props:['data'],

        data: function(){
            return {
                image:"./Pictures/svmclogo.jpg",
                num: 1,
                nextApp: "03/01/2020",
                time:"15", 
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