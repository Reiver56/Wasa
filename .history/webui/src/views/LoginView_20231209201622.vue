
<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>
	<div class="login-container">
		<div class="top-login-container"> <span class="login-title">Login</span></div>
		<div class="bottom-login-container">
			<div class="login-input-container">
				<input :on-submit="doLogin" type="text" name="nickname" spellcheck="false" v-model="nickname" maxlength="16">
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
		async doLogin() {
			this.isLoading = true;
			try{
				if(this.nicknameValidation.test(this.nickname)) throw 'Invalid nickname';
				if(this.nickname.length < 3 || this.nickname.length > 16) throw new Error('Nickname must be between 3 and 16 characters');
				let response = await this.$axios.post('/session', {nickname: this.nickname.trim(),});
				localStorage.userID = response.data.userID;
				localStorage.nickname = response.data.nickname;
				this.$router.push('/home');
				axios
			}catch(e){
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
</style>



