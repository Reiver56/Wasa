<template>
	<img class="logo" src="/src/assets/logo/logo_blue.jpg" />
	<h1>Login Page</h1>
	<div class="login-form">
		<input type="text" v-model="user_id" placeholder="Enter user identfier"/>
		<input type="text" v-model="nickname" placeholder="Enter user nickname"/>
		<button @click="login" router-link to="/">Login</button>
		<p>
			<router-link to="/" >login</router-link>
		</p>
	</div>
</template>

<script>
export default {
	name: 'LoginView',
};
</script>

