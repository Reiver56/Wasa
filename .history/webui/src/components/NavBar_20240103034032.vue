<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg =''"></ErrorMsg>
	<div class="navbar">
		<img src="../assets/logo/gallery.png" alt="logo" width="50" height="50">
		<SearchBar @error-occurred="$emit('error-occurred', value)" />
		<button @click="getMyProfile" class="left-profile-navbar">
			<!--Profile-->
			<font-awesome-icon icon="fa-solid fa-user" />
		</button>
		

	</div>
</template>

<script>
import SearchBar from './SearchBar.vue';
export default{
	components:{
		SearchBar,
	},
	emits: ['error-occurred'],
	data(){
		return{
			errorMsg:'',
		}

	},
	methods:{
		async getMyProfile(){

			this.$router.push(`/users/${localStorage.UserID}`);


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
    width: 100%;
    padding: 0;
    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);
    position: sticky;
    padding: 1em;
    margin-bottom: 1em;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    z-index: 3;
}
.left-profile-navbar {
    background-color: rgb(255, 255, 255, 0.4);
    height: 1.5em;
    width: 1.5em;
    border-radius: 10em;
    padding: 0;
    margin: 0 0 0 0.5em;
    float: left;
    outline: none;
    border: none;
    font-size: 2em;
    line-height: 0;
    color: rgb(0, 0, 0, 0.8)

}
.left-profile-navbar:hover{
	cursor: pointer;
	background-color: rgb(255, 255, 255, 0.8);
}
.home-button{
	background-color: rgb(255, 255, 255, 0.4);
    height: 1.5em;
    width: 1.5em;
    border-radius: 10em;
    padding: 0;
    margin: 0 0 0 0.5em;
    float: left;
    outline: none;
    border: none;
    font-size: 2em;
    line-height: 0;
    color: rgb(0, 0, 0, 0.8)
}

.home-button:hover{
	cursor: pointer;
	background-color: rgb(255, 255, 255, 0.8);
}
.upload-post-button{
	background-color: rgb(255, 255, 255, 0.4);
	height: 1.5em;
	width: 1.5em;
	border-radius: 10em;
	padding: 0;
	margin: 0 0 0 0.5em;
	float: left;
	outline: none;
	border: none;
	font-size: 2em;
	line-height: 0;
	color: rgb(0, 0, 0, 0.8)
}

.upload-post-button:hover{
	cursor: pointer;
	background-color: rgb(255, 255, 255, 0.8);
}
</style>
