<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <ToDo msg="Welcome to the Test Center"/>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import ToDo from '@/components/ToDo.vue'; // @ is an alias to /src

@Options({
  components: {
    ToDo,
  },
})
export default class Home extends Vue {}
</script>
