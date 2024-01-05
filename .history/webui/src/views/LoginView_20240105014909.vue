
<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>
	<div class="background-login">
	<div class="login-container">
		<img src="../assets/logo/logo_blue_1.png" alt="logo" width="500" height="500">
		<div class="top-login-container"> <span class="login-title">Login</span></div>
		<div class="bottom-login-container">
			<div class="login-input-container">
				<input :on-submit="doLogin" type="text" name="nickname-form" spellcheck="false" v-model="nickname" maxlength="16" placeholder="nickname" class="box-nickname">
			</div>
			<div class="login-button-container">
				<button class="btn login-button" @click="doLogin">Login</button>
			</div>
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
.background-login{
	padding: auto;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}
.login-container{
	background-color:aliceblue;
	pa
	border-radius: 4em;
	height: 750px;
	width: 50%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}

.top-login-container{
	height: 50%;
	width: 50%;
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
	color: black;
	text-align: center
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
	background-color: #3b9ee4;
	border-radius: 10em;
	outline: none;
	color: white;
	font-size: 1.5em;
	font-weight: bold;
}

.login-button:hover{
	cursor: pointer;
	background-color: white;
	color: black;
}
.login-input-container input{
	height: 2em;
	width: 10em;
	border-radius: 10em;
	outline: none;
	font-size: 1.5em;
	text-align: center;
	border: 0.09ch solid black;
}
.login-input-container input::placeholder{
	font-size: 1.2em;
	text-align: center;
}
.login-input-container input:focus{
	border: 2px solid #3b9ee4;
}




</style>



