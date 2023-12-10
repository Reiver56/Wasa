<template>
	<div>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />

		<div class="home-container">
			<div class="home-header">
				<div class="home-header-title">
					<span class="home-title">My Stream</span>
				</div>
				<div class="home-header-button">
					<button class="home-button" @click="showPhotos = !showPhotos">
						{{ showPhotos ? 'Hide Photos' : 'Show Photos' }}
					</button>
				</div>
			</div>
			<div class="home-content">
				<div class="home-content-container">
					<div class="home-content-item" v-for="post in posts" :key="post.postID">
						<Post :post="post" :showPhotos="showPhotos" @update-like="updateLike" @delete-post="deletePost" />
					</div>
				</div>
			</div>

	</div>
</template>



<script>
export default {
	data: function() {
		return {
			errorMsg: "",
			loading: false,
			posts: [],

			showPhotos: false,

			busy: false,
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
		updateLike(data) {
			this.posts.forEach(post =>{
				if (post.postID == data.postID){
					post.likes = data.likes;
					post.likesCount++;
				}
			})

		},
		deletePost: async function(postID){
			const index = this.posts.findIndex(post => post.postID == postID && post.userID == localStorage.userID);
			try{
				await this.$axios.delete('users/${localStorage.userID}/photos/${photoID}', { headers: {authorization: localStorage.token}});
				this.posts.splice(index, 1);
			}catch(e){
				this.errorMsg = e.message;
			}
		},
	},
	mounted() {
		this.getMyStream();
		document.addEventListener('scroll', this.trackScrolling);

	}
}
</script>



<style>
</style>
