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
			
		}
	}
}
</script>

<style>
</style>



