<template>
  <div id="app">
    <Navbar brand="t9.pm"/>
    <main role="main" class="container flex-shrink-0">
      <router-view/>
    </main>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';

export default {
  components: {
    Navbar,
  },
};
</script>

<style>
main.container { padding: 100px 15px 0;}
.container { max-width: 680px; }
</style>
