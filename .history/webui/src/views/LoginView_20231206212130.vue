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
export default {
	name: 'LoginView',
	data(){
		return{
			nickname:'',
			user_id:'',
			loading:false,
			errorMsg:'',
		}
	},
	methods: {
		async doLogin(){
			this.loading = true;
			try{
				if (this.nickname == '' || this.user_id == ''){
					this.errorMsg = 'Please enter nickname and id';
					return;
				}
				if(this.nickname == ''){
				this.errorMsg = 'Please enter nickname';
				return;
			}
			if(this.user_id == ''){
				this.errorMsg = 'Please enter id';
				return;
			}
			let response = await this.$axios.post('/login', this.user_id.trim(), this.nickname.trim());
			this.$router.push('/home');
			localStorage.user_id = response.data.user.user_id;
			localStorage.nickname = response.data.user.nickname
			}catch(e){
			this.errorMsg = e.response.data;
		};
		this.loading = false;




	}
	},
	mounted(){
		localStorage.clear();
	},



}


</script>

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


