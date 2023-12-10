<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg =''"></ErrorMsg>
	<div class="navbar">
		<button @click="getMyProfile" class="left-profile-navbar"></button>
		<but
	</div>
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
