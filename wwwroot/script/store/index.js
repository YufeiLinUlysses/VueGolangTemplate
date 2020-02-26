import Vue from "vue";
import Vuex from "vuex";

import patient from "./modules/patient";

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        patient
    }
});