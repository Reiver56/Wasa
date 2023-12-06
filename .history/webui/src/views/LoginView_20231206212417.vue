<template>
	<div class="login">
	  <div>
		<form @submit.prevent="submit">
		  <div>
			<label for="username">Username:</label>
			<input type="text" name="username" v-model="form.username" />
		  </div>
		  <div>
			<label for="password">Password:</label>
			<input type="password" name="password" v-model="form.password" />
		  </div>
		  <button type="submit">Submit</button>
		</form>
		<p v-if="showError" id="error">Username or Password is incorrect</p>
	  </div>
	</div>
  </template>


<script>
import { mapActions } from "vuex";
export default {
  name: "Login",
  components: {},
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      showError: false
    };
  },
  methods: {
    ...mapActions(["LogIn"]),
    async submit() {
      const User = new FormData();
      User.append("username", this.form.username);
      User.append("password", this.form.password);
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

<style>
body {
	background-color: whitesmoke;
}
.logo {
	width: 500px;
	height: 500px;

}

.login-form {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content:  space-between;

}

.title-d {
	font-size: 2.5em;
	line-height: 1.5em;
	font-weight: bold;
	margin-top: 0.5em;
}
.input-d {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content:  space-between;
	margin-top: 0.5em;
}
.input-d input {
	width: 200px;
	font-size: 1.5em;
	line-height: 1.5em;
	text-align: center;
	border: 1px solid black;
	border-radius: 5px;
}
.button-d {
	margin-top: 0.5em;
}
.button-d button {
	background-color: white;
	width: 200px;
	font-size: 1.5em;
	line-height: 1.5em;
	text-align: center;
	border: 1px solid black;
	border-radius: 5px;
}
.button-d button:hover {
	background-color: #0D92DD;
	color: white;
}
</style>


