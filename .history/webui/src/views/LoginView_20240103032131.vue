
<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>
	<div class="login-container">
		<img src="../assets//logo.png" alt="logo" width="200" height="200">
		<div class="top-login-container"> <span class="login-title">Login</span></div>
		<div class="bottom-login-container">
			<div class="login-input-container">
				<input :on-submit="doLogin" type="text" name="nickname-form" spellcheck="false" v-model="nickname" maxlength="16">
			</div>
			<div class="login-button-container">
				<button class="login-button" @click="doLogin">Login</button>
			</div>
		</div>
	</div>
	<LoadingSpinner :loading="isLoading"/>
</template>

<script>
export default{
	data() {
		return{
			nickname: "",
			errorMsg: "",
			isLoading: false,
			nicknameValidation: new RegExp('^[a-zA-Z0-9]{3,16}$'),

		}
	},
	methods: {
		doLogin: async function() {
			this.isLoading = true;
			try{

				let response = await this.$axios.post('/session', {
					nickname: this.nickname,
				});
				console.log(response.data);
				localStorage.UserID = response.data.user.id;
				localStorage.nickname = response.data.user.nickname;
				localStorage.token = response.data.token;

				this.$router.replace('/home');

			} catch(e){
				this.errorMsg = e.message;
			};
			this.isLoading = false;
		}
	},

	mounted() {
		localStorage.clear();
	},
}
</script>

<style>
.login-container{
	background-color:aliceblue;
	height: 100vh;
	width: 100vw;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.top-login-container{
	height: 50%;
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.bottom-login-container{
	height: 50%;
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.login-title{
	font-size: 3em;
	font-weight: bold;
	color: #3b9ee4;
}
.login-input-container{
	height: 50%;
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.login-button-container{
	height: 50%;
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.login-button{
	height: 2em;
	width: 10em;
	background-color: white;
	border-radius: 10em;
	outline: none;
	color: black;
	font-size: 1.5em;
	font-weight: bold;
}

.login-button:hover{
	cursor: pointer;
	background-color: #3b9ee4;
	color: white;
	transform: scale(1.1);
}
.login-input-container input{
	height: 2em;
	width: 10em;
	border-radius: 10em;

	outline: none;
	font-size: 1.5em;
	text-align: center;
}



</style>



