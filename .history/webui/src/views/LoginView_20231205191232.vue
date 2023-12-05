<template>
	<img class="logo" src="/src/assets/logo/logo_blue.jpg" />
	
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
};
</script>

<style>
.logo {
	width: 400px;
	height: 400px;
}
</style>


