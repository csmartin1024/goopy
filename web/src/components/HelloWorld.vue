<template>
    <div>
        <div class="hello">
            <h1>{{ msg }}</h1>
        </div>
        <div class="details" style="display:none;">
            <h6>Color</h6>
            <input :value="color" @input="updateColor($event.target.value)" />

            <h6>Opacity</h6>
            <input :value="opacity" @input="updateOpacity($event.target.value)" />

            <h6>Speed</h6>
            <input type="number" :value="speed" @input="updateSpeed($event.target.value)" />

            <h6>Texture</h6>
            <input :value="texture" @input="updateTexture($event.target.value)" />

            <h6>Viscosity</h6>
            <input :value="viscosity" @input="updateViscosity($event.target.value)" />
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { namespace } from 'vuex-class';
const Goop = namespace('goop');

@Component
export default class HelloWorld extends Vue {
    @Prop() private msg!: string;
    @Goop.State
    public color!: string;

    @Goop.State
    public opacity!: string;

    @Goop.State
    public speed!: number;

    @Goop.State
    public texture!: string;

    @Goop.State
    public viscosity!: string;

    @Goop.Action
    public updateColor!: (newColor: string) => string;

    @Goop.Action
    public updateOpacity!: (newOpacity: string) => string;

    @Goop.Action
    public updateSpeed!: (newSpeed: number) => number;

    @Goop.Action
    public updateTexture!: (newTexture: string) => string;

    @Goop.Action
    public updateViscosity!: (newViscosity: string) => string;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
    margin: 40px 0 0;
}
ul {
    list-style-type: none;
    padding: 0;
}
li {
    display: inline-block;
    margin: 0 10px;
}
a {
    color: #42b983;
}
</style>
