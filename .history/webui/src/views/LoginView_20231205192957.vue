<template>

	<LoadinSpinner :loading="loading"></LoadinSpinner>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg=''"></ErrorMsg>
	<div class="login-form">
		<img class="logo" src="/src/assets/logo/logo_blue.jpg" />
		<div class="loginform"></div>
		

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
	width: 500px;
	height: 500px;
}
.login-form {
	display: flex;
	flex-direction: column;
	align-items: center;
}
</style>


