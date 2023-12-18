<template>
	<div class="search-bar-container">
		<input placeholder="Search for users" autoComplete="none" class="search-bar" v-model="searchQuery" @focusout="exitList" maxlength="16" spellcheck="false"/>

		<div class="search-results" v-show="userList.length">
			<SimpleProfileEntry v-for="user in userList" :key="user.userID" :data="user" @exit-from-list="exitList" />
		</div>
	</div>
</template>

<script>
import SimpleProfileEntry from './SimpleProfileEntry.vue';
export default{
	emits: ['error-occurred'],
	components: {
		SimpleProfileEntry,
	},
	props:{},
	data(){
		return{
			searchQuery: '',
			userList:[],
			busy: false,
			dataAvailable: true,
		}
	},
	methods: {
		async search(){
			console.log(this.searchQuery);
			if (this.searchQuery.length == 0){
				this.userList = [];
				return;
			}
			try {
				let response = await this.$axios.get(`/users?search=${this.searchQuery}`, { headers: { 'Authorization': `${localStorage.token}` } })
				if (response.data.length == 0){
					this.dataAvailable = false;
					return;
				}
				console.log(response.data);
				this.userList = [];
				this.userList.push(...response.data);
			}catch (error){
				this.$emit('error-occurred', error.response.data.message);
			}

		},
		exitList(){
			setTimeout(() => {
				this.userList = [];
				this.searchQuery = '';
				this.dataAvailable = true;
			}, 500);
		},

	},
	mounted(){
		this.search();
	},
	watch: {
		searchQuery(){
			this.search();
		}
	}
}
</script>

<style>
.search-bar-container{
	background-color:aliceblue;
	height: 3em;
	width: 12em;
	border-radius: 10em;

	float: right;
}
.search-bar{
	background-color: transparent;
	height: 2em;
	width: 10em;
	border-radius: 10em;
	padding: 1em;
	box-shadow: none;
	outline: none;
	border: none;
}
.search-results{
	background-color: white;
	height: auto;
	max-height: 20em;

	overflow: auto;

	width: ;
}

</style>
