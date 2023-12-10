
<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>
	<div class="login-container">
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
					nickname: this.nickname
				});
				localS
				this.$router.push('/home');

			} catch(e){
				this.errorMsg = e.message;
			};
			this.isLoading = false;
		}
	},

	mounted() {
	},
}
</script>

<style>
</style>



