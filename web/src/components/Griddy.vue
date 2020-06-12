<template>
    <div>
        <h2>Test</h2>
        <ul>
            <li class="wrapper" v-for="(item, index) in items" :key="index">
                <div class="letter">{{ item.color }}</div>
                <div class="letter">{{ item.speed }}</div>
            </li>
        </ul>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
@Component
export default class Griddy extends Vue {
    @Prop() private items!: [];
}
</script>

<style scoped lang="scss">
.wrapper {
    display: grid;
    grid-template-columns: 1fr 1fr;
}
.letter {
    background-color: #0069b3;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
    font-size: 70px;
    color: white;
    transition: all 0.3s ease;
}
</style>
