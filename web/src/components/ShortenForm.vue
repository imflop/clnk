<template>
  <div class="shorten-form">
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <div class="input-group">
          <input type="text"
                 class="form-control form-control-lg"
                 id="textInput"
                 aria-describedby="textHelp"
                 v-model="original_url">
          <div class="input-group-append">
            <button class="btn btn-outline-secondary" type="submit" id="button-addon1">
              Shorten
            </button>
          </div>
        </div>
        <small id="textHelp" class="form-text text-muted">
            We'll never share your email with anyone else.
        </small>
      </div>
    </form>
  </div>
</template>

<script>
import HTTP from '../api/requester';

export default {
  data() {
    return { original_url: '' };
  },
  name: 'ShortenForm',
  props: {
    msg: String,
  },
  methods: {
    handleSubmit() {
      HTTP.post('/', {
        original_url: this.original_url,
      }).then((response) => {
        this.original_url = response.data.short_url;
        console.log(response.data.short_url);
      }).catch((e) => { console.error(e); });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
