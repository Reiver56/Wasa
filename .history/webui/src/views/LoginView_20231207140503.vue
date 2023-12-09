<template>
	<div class="login">
		<img src="src/assets/logo/logo_blue.jpg" alt="logo" />
	  <div>
		<form @submit.prevent="submit">
		  <div>
			<label for="nickname">Nickname:</label>
			<input type="text" name="nickname" v-model="form.nickname" />
		  </div>
		  <button type="submit">Submit</button>
		</form>
		<p v-if="showError" id="error">Nickname is incorrect</p>
	  </div>
	</div>
  </template>


<script>

export default {
  name: "Login",
  components: {},
  data() {
    return {
      form: {
        nickname: "",
      },
      showError: false
    };
  },
  methods: {

    async submit() {
      const User = new FormData();
      User.append("username", this.form.nickname);
      try {
          await this.LogIn(User);
          this.$router.push("/home");
          this.showError = false
      } catch (error) {
        this.showError = true
      }
    },
  },
};
</script>

<style>

</style>



