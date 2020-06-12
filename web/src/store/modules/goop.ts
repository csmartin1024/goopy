import {
    VuexModule,
    Module,
    MutationAction,
    Action,
} from 'vuex-module-decorators';
import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import VueAxios from 'vue-axios';

Vue.use(Vuex);
Vue.use(VueAxios, axios);

// Vue.axios.defaults.baseURL = 'http://localhost:8000/api';

@Module({ namespaced: true })
class Goop extends VuexModule {
    public color = '';
    public opacity = '';
    public speed = '';
    public texture = '';
    public viscosity = '';
    public goops = [];

    @MutationAction({ mutate: ['color'] })
    async updateColor(newColor: string) {
        console.log(newColor);
        return { color: newColor };
    }

    @MutationAction({ mutate: ['opacity'] })
    async updateOpacity(newOpacity: string) {
        return { opacity: newOpacity };
    }

    @MutationAction({ mutate: ['speed'] })
    async updateSpeed(newSpeed: number) {
        return { speed: newSpeed };
    }

    @MutationAction({ mutate: ['texture'] })
    async updateTexture(newTexture: string) {
        return { texture: newTexture };
    }

    @MutationAction({ mutate: ['viscosity'] })
    async updateViscosity(newViscosity: string) {
        return { viscosity: newViscosity };
    }

    @MutationAction({ mutate: ['goops'] })
    async loadGoops() {
        const result = await Vue.axios.get('/api/goops');
        console.log(result);
        return { goops: result.data };
    }
}
export default Goop;
