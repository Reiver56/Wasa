<template>
	<img class="logo" src="/src/assets/logo/logo_blue.jpg" />
	<LoadinSpinner :loading="loading"></LoadinSpinner>
	<ErrorMsg></ErrorMsg>
	<h1>Login Page</h1>
	<div class="login-form">
		<input type="text" v-model="user_id" placeholder="Enter user id"/>
		<input type="text" v-model="nickname" placeholder="Enter user nickname"/>
		<button @click="dologin">
			Login
			<router-link to="/" ></router-link>
		</button>

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
.logo {
	width: 400px;
	height: 400px;
}
</style>


