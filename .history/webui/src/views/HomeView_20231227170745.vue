<template>
	<div>

		<NavBar  @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />

		<div class="row">
			<Photo
				v-for="post in posts"
				:key="index"
				:owner="post.owner"
				:photo_id="post.photo_id"
				:comments="post.comments != null ? post.comments : []"
				:likes="post.likes != null ? post.likes : []"
				:upload_date="post.upload_date"

			/>

		</div>
		<div v-if="posts.length === 0" class = "row"><h1 class="d-flex">No Content yet, follow someone</h1></div>
	</div>
</template>



<script>
import NavBar from '@/components/NavBar.vue';

export default {
	components: {
		NavBar,
	},
	data: function() {
		return {
				posts: [],
				errorMsg: "",
				isLoading: false,
			}
	},
	methods: {
		: async getMyStream(){
			console.log("getMyStream");
			try{
				this.isLoading = true;
				let response = await this.$axios.get(`users/${localStorage.userID}/mystream`, { headers: { 'Authorization': `${localStorage.token}` } });
				if (response.data.length != 0){
					this.photos = response.data;
				}
				this.posts.push(...response.data);

			}catch(e){
				this.errorMsg = e.toString();
			}

	},
	async mounted(){
		await this.getMyStream();

		document.addEventListener('scroll', this.trackScrolling);
	},
}
}
</script>



<style>
</style>
