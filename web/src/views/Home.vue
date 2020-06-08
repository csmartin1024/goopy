<template>
  <div class="home">
    <img class="image" alt="Vue logo" src="../assets/goop.jpg" />
    <HelloWorld msg="You just got gooped!" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld
  }
};
</script>

<style scoped lang="scss">
.image {
  object-fit: cover;
  height: 300px;
}
</style>
