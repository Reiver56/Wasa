<template>
	<div class="search-bar">
		<input placeholder="Search for users" autoComplete="none" class="search"/>
		<div class="search-results">
			<SimpleProfileEntry v-for="user in userList" :key="user.userID" :data="user" @exit-from-list="userList = []" @error-occurred="$emit('error-occurred', value)" />
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
			if (this.search.length == 0){
				this.userList = [];
				return;
			}
			try {
				let response = await this.$axios.get('/users', { headers: { 'Authorization': `${localStorage.token}` } })
				console.log(response.data);
			}catch (error){
				this.$emit('error-occurred', error.response.data.message);
			}

		},

	},


}
</script>
