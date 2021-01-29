<template>
    <nav class="navbar navbar-expand-sm fixed-top navbar-light bg-light">
      <div class="container">
        <a class="navbar-brand" href="/">{{ brand }}</a>
        <button class="navbar-toggler"
                type="button"
                data-toggle="collapse"
                data-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent"
                aria-expanded="false"
                aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item">
              <span class="nav-link">
                <router-link to="/">Home</router-link>
              </span>
            </li>
            <li>
              <span class="nav-link">
                <router-link to="/about">About</router-link>
              </span>
            </li>
            <li>
              <span class="nav-link">
                <router-link to="/rules">Rules</router-link>
              </span>
            </li>
          </ul>
        </div>
      </div>
    </nav>
</template>

<script>
export default {
  name: 'Navbar',
  props: {
    brand: String,
  },
};
</script>

<style scoped>
</style>
