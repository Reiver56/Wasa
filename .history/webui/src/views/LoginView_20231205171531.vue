<template>
	<img class="logo" src="../assets/logo_blue.png" />
	<h1>Login Page</h1>
	<div class="login-form">
		<input type="text" v-model="user_id" placeholder="Enter user identfier"/>
		<input type="text" v-model="nickname" placeholder="Enter user nickname"/>
		<button @click="login">Login</button>
		<p>
			<router-link to ><
		</p>
	</div>
</template>

<script>
export default {
	name: 'LoginView',
};
</script>

