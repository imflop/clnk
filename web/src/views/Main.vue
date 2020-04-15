<template>
  <div class="main">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <ShortenForm/>
  </div>
</template>

<script>
// @ is an alias to /src
import ShortenForm from '@/components/ShortenForm.vue';

export default {
  name: 'Main',
  components: {
    ShortenForm,
  },
};
</script>
