<template>
	<LoadinSpinner :loading="loading"></LoadinSpinner>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg=''"></ErrorMsg>
	<div class="login-form">
		<div class="logo-d">
			<img class="logo" src="/src/assets/logo/logo_blue.jpg" />
		</div>
		<div class="title-d">
			<h1>Login Page</h1>
		</div>
		<div class="input-d">
			<input :on-submit="doLogin" type="text" name="user_id-form" spellcheck="false" v-model="user_id" maxlength="13" placeholder="Enter id"/>
		</div>
		<div class="input-d">
			<input :on-submit="doLogin" type="text" name="nickname-form" spellcheck="false" v-model="nickname" maxlength="13" placeholder="Enter nickname" />
		</div>
			<div class="button-d">
				<button @click="doLogin"><router-link to ="/./home">Login</router-link></button>
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
			localStorage.nickname = response.data.user.nickname;
			this.$router.push("/home");
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


