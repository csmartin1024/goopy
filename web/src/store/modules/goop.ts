import { VuexModule, Module, MutationAction } from 'vuex-module-decorators';
@Module({ namespaced: true })
class Goop extends VuexModule {
    public color = '';
    public opacity = '';
    public speed = '';
    public texture = '';
    public viscosity = '';

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
}
export default Goop;
