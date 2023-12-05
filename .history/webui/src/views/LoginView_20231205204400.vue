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
			<input type="text" v-model="user_id" placeholder="Enter user id" spellcheck="false" maxlength="13"/>
		</div>
		<div class="input-d">
			<input type="text" v-model="nickname" placeholder="Enter user nickname" spellcheck="false" maxlength="13"/>
		</div>
		<div class="button-d">
			<button @click="dologin">Login</button>
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
	}
};
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
	background-color: white;
	box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.2);
	padding: 20px;
	width: 25em;
	
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content:  space-between;

}
</style>


