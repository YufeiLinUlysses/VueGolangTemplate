<template>
    <div>
        <search @chosenpatient="chosenpatient"></search>
    </div>
</template>
<script>
    import Vue from 'vue'
    import axios from 'axios'
    import moment from 'moment';
    import { preference } from 'vue-preferences'
    import topbar from './components/topbar.vue'
    import search from './components/search.vue'

    var webcall = axios.create({
        baseURL: 'api',
        timeout: 20000,
        withCredentials: false,
        headers: { 'Content-Type': 'application/json' }
    });

    export default {
        props: [],
        components: {
            topbar,
            search
        },
        data() {
            return {
                curpatient: {}
            }
        },
        methods: {
            chosenpatient(patient) {
                this.curpatient = patient
            }
        },
        computed: {
            isdarkmode: preference('isdarkmodeenabled', { defaultValue: false })

        },
        watch: {
        },
        created() {
        },
        mounted() {
            console.log("Main View Mounted!");
        }
    }

</script>
<style>
    .menucustom2 {
        padding: 0px;
    }
</style>
