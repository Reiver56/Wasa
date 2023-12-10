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
	background-color: #3b9ee4;
    height: auto;
    width: auto;
    padding: 0;
    border-radius: 10em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: sticky;
    top: 1em;

    padding: 1em;
    margin-bottom: 1em;

    display: flex;
    flex-direction: row;
    justify-content: space-between;
    z-index: 3;
}
.home-button{
	background-color: #3b9ee4;
	height: 2.5em;
	width: 2.5em;
	border-radius: 10em;
	border: none;
	box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);
	color: white;
	font-size: 1.5em;
	margin: 0;
	padding: 0;
	display: flex;
	justify-content: center;
	align-items: center;
	cursor: pointer;
}
</style>
