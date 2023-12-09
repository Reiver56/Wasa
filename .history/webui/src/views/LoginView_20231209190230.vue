<template>
	<div class="login-container">
		<div class="top-login-container"> <span class="login-title">Log in</span></div>
		<div class="bottom-login-container">
			<div class="login-input-container">
				<input type="text" v-model="nickname" placeholder="Nickname" class="login-input" />
			</div>
			<div class="login-button-container">
				<button class="login-button" @click="doLogin">Log in</button>
			</div>
			<div class="login-error-container">
				<span class="login-error">{{errorMsg}}</span>
			</div>
		</div>
	</div>
	<LoadingSpinner :loadi
</template>

<script>
export default{
	data(){
		return{
			nickname: '',
			errorMsg: '',
			isLoading: false,
			nicknameValidation: new RegExp('^[a-zA-Z0-9]{3,16}$'),
		}
	},
	methods:{
		async doLogin(){
			this.isLoading = true;
			try{
				if(this.nicknameValidation.test(this.nickname)) throw new Error('Invalid nickname');
				if(this.nickname.length < 3 || this.nickname.length > 16) throw new Error('Nickname must be between 3 and 16 characters');
				let response = await this.$axios.post('/session', {nickname: this.nickname});
				localStorage.userID = response.data.userID;
				localStorage.nickname = response.data.nickname;
				this.$router.push('/home');

			}catch(e){
				this.errorMsg = e.message;
			};
			this.isLoading = false;
		}


	},
	mounted(){
			localStorage.clear();
		}
}
</script>

<style>
</style>



