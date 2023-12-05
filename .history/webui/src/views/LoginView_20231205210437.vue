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
			<div class="button-d">
			<button @click="dologin">Login</button>
			</div>
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
	width: 200px;
	font-size: 1.5em;
	line-height: 1.5em;
	text-align: center;
	border: 1px solid black;
	border-radius: 5px;
}

</style>


