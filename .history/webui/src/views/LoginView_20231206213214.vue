<template>
	<div class="login">
		<img src="src/assets//logo_blue.jpg" alt="logo" />
	  <div>
		<form @submit.prevent="submit">
		  <div>
			<label for="nickname">Nickname:</label>
			<input type="text" name="nickname" v-model="form.nickname" />
		  </div>
		  <button type="submit">Submit</button>
		</form>
		<p v-if="showError" id="error">Username or Password is incorrect</p>
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
          this.$router.push("/posts");
          this.showError = false
      } catch (error) {
        this.showError = true
      }
    },
  },
};
</script>

<style scoped>
* {
  box-sizing: border-box;
}
label {
  padding: 12px 12px 12px 0;
  display: inline-block;
}
button[type=submit] {
  background-color: #4CAF50;
  color: white;
  padding: 12px 20px;
  cursor: pointer;
  border-radius:30px;
}
button[type=submit]:hover {
  background-color: #45a049;
}
input {
  margin: 5px;
  box-shadow:0 0 15px 4px rgba(0,0,0,0.06);
  padding:10px;
  border-radius:30px;
}
#error {
  color: red;
}
</style>



