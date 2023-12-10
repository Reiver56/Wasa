<template>
	<div>
		<ErrorMsg v-if="err">
	</div>
</template>



<script>
export default {
	data: function() {
		return {
			errorMsg: "",
			loading: false,
			posts: [],

			dataAvailable: true,
			isLoading: false,
		}
	},
	methods: {
		getMyStream: async function(){
			this.isLoading = true;
			try{
				const url = 'users/' + localStorage.userID + '/mystream';
				let response = await this.$axios.get(url, { headers: {authorization: localStorage.token}});
				if (response.data == null){
					this.dataAvailable = false;
					this.isLoading = false;
					return;
				}
				this.posts.push(response.data);

			}catch(e){
				this.errorMsg = e.message;
			}
			this.isLoading = false;
		},
	},
	mounted() {

	}
}
</script>



<style>
</style>
