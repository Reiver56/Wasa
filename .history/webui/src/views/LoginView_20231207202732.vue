<template>
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
		doLogin(){
			this.isLoading = true;
			try{
				if(this.nicknameValidation.test(this.nickname)){
					
				}else{
					this.errorMsg = 'Nickname must be 3-16 characters long and contain only letters and numbers';
				}
			}
		}
	}
}
</script>

<style>
</style>



