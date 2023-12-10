<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg =''"></ErrorMsg>
	<div class="navbar">
		<button @click="getMyProfile" class="left-profile-navbar"></button>
		<button class="home-button" @click="$router.go(0)"><font-awesome-icon icon="fa-regular fa-square-plus" /></button>
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
.navbar{
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	align-items: center;
	background-color: #1a1a1a;
	height: 50px;
	width: 100%;
	position: fixed;
	top: 0;
	left: 0;
	z-index: 100;
}
</style>
