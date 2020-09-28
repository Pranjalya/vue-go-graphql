<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <ToDo :username="name" />
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import ToDo from "@/components/ToDo.vue"; // @ is an alias to /src

@Options({
  components: {
    ToDo,
  },
})
export default class Home extends Vue {
  name = "1";
  public getName(): void{
    try {
      const n = prompt("Please enter your name");
      this.name = n == null ? "Unknown" : n;
    } catch(err) {
      this.getName()
    }
  }
  mounted() {
    this.getName();
  }
}
</script>
