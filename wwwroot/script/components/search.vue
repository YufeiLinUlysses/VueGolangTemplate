<template>
    <div class="container">        
        <div class="jumbotron text-center" style="background-color:lightblue">
            <h1 class="display-4">Please Enter the Information of the Patient</h1>
            <div class="input-group input-group-lg justify-content-center">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="inputGroup-sizing-lg">OCVID</span>
                </div>
                <input 
                type="text" 
                aria-label="Sizing example input" 
                aria-describedby="inputGroup-sizing-lg" 
                v-model="searchquest.OCVID" 
                placeholder="V0001">
            </div>
            <font size="12">Or</font>
            <div class="input-group input-group-lg justify-content-center">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="inputGroup-sizing-lg">MR#</span>
                </div>
                <input 
                type="text" 
                aria-label="Sizing example input" 
                aria-describedby="inputGroup-sizing-lg" 
                v-model="searchquest.MRNum" 
                placeholder="M0001">
            </div>
            <font size="12">Or</font>
            <div class="input-group input-group-lg justify-content-center">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="inputGroup-sizing-lg">First Name</span>
                </div>
                <input 
                type="text" 
                aria-label="Sizing example input" 
                aria-describedby="inputGroup-sizing-lg" 
                v-model="searchquest.FName" 
                placeholder="First Name">
            </div>
            <div class="input-group input-group-lg justify-content-center">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="inputGroup-sizing-lg">Last Name</span>
                </div>
                <input 
                type="text" 
                aria-label="Sizing example input" 
                aria-describedby="inputGroup-sizing-lg" 
                v-model="searchquest.LName" 
                placeholder="Last Name">
            </div>
            <div class="input-group input-group-lg justify-content-center">
                <div class="input-group-prepend">
                    <span class="input-group-text" id="inputGroup-sizing-lg">DOB</span>
                </div>
                <input 
                type="text" 
                aria-label="Sizing example input" 
                aria-describedby="inputGroup-sizing-lg" 
                v-model="searchquest.DOB" 
                placeholder="01/01/2020">
            </div>
            <button @click.prevent="verifyInput()"> Search </button>
        </div>
        <!-- The Result of a list of Patient that has the related personal information -->
        <!-- <p>results</p>
        <span v-for="x in searchquest">
            <a href="#" @click="chosenpatient(x)">{{x}}</a><br>
        </span> -->
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

    export default({
        data: function() {
            return{
                results:[],
                searchquest:{OCVID:"",MRNum:"",FName:"",LName:"",DOB:""}
            }
        },
        methods:{
            verifyInput: function(){
                var vm = this
                var quest = vm.searchquest
                const ocvidregex1 = /V\d{4}/
                if (quest.OCVID!="" && quest.OCVID.match(ocvidregex1)==null){
                    alert("Please check your OCVID")
                    return
                }
                const mrnumregex1 = /M\d{4}/
                if (quest.MRNum!="" && quest.MRNum.match(mrnumregex1)==null){
                    alert("Please check your OCVID")
                    return
                }

                var useName1 = (quest.DOB!="")&&(quest.FName==""||quest.LName=="")
                var useName2 = (quest.FName!="")&&(quest.DOB==""||quest.LName=="")
                var useName3 = (quest.LName!="")&&(quest.DOB==""||quest.FName=="")

                if (useName1||useName2||useName3){
                    alert("Need more information")
                    return
                }
                webcall.post("searchpatient",quest)
                    .then(function (response) {
                        console.log(response.data.found)
                        if (response.data.found){
                            vm.results = response.data
                            //This routes to a next page
                            vm.$router.push({ name:'SearchResult', params: {searchresults: vm.results}})
                            vm.$emit('chosenpatient',vm.results)
                            console.log(vm.results)
                        }
                        else{
                            alert("The patient does not exist, please contact person in charge")
                        }
                        
                    })
                    .catch(function (error) {
                        console.log(error);
                    });

                //This would refresh the models the compnents uses 
                vm.searchquest = {OCVID:"",MRNum:"",FName:"",LName:"",DOB:""} 
            },
            chosenpatient: function(patient){
                this.$emit('chosenpatient',patient)
            }
        }
    })
</script>