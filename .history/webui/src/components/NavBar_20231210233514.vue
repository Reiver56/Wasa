<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg =''"></ErrorMsg>
	<div class="navbar">
		<div class="navbar-left">
			<div class="navbar-logo">
				<img src="/static/logo.png" alt="logo" />
			</div>
			<div class="navbar-title">
				<span>Photo Stream</span>
			</div>
		</div><
</template>

<script>
export default{
	components:{
	},
	data(){
		return{
			errorMsg:'',
		}

	},
	methods:{
		getMyProfile(){
			this.$router.push(`/profile/${this.$store.state.user_id}`);
		},
	},

	mounted(){

	},
}
</script>

<style>

</style>
