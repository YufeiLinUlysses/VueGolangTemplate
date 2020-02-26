import axios from "axios";

var webcall = axios.create({
    baseURL: 'api',
    timeout: 20000,
    withCredentials: false,
    headers: { 'Content-Type': 'application/json' }
});

export default {
    state: {},
    getters: {},
    mutations: {},
    actions: {}
}