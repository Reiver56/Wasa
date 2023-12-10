<template>
	<div>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<div class="home-container">
			<div class="home-header">
				<div class="home-header-title">My Stream</div>
				<div class="home-header-button-container">
					<button class="home-header-button" @click="showPhotos = !showPhotos">Show Photos</button>
				</div>
			</div>
			<div class="home-content">
				<div class="home-content-container">
					<div class="home-content-post-container" v-for="post in posts" :key="post.postID">
						<div class="home-content-post-header">
							<div class="home-content-post-header-nickname">{{post.nickname}}</div>
							<div class="home-content-post-header-timestamp">{{post.timestamp}}</div>
						</div>
						<div class="home-content-post-body">
							<div class="home-content-post-body-text">{{post.text}}</div>
							<div class="home-content-post-body-image-container">
								<img class="home-content-post-body-image" v-if="post.image" :src="post.image" />
							</div>
						</div>
						<div class="home-content-post-footer">
							<div class="home-content-post-footer-likes">{{post.likesCount}} likes</div>
							<div class="home-content-post-footer-like-button-container">
								<LikeButton :postID="post.postID" :likes="post.likes" @update-like="updateLike"/>
							</div>
							<div class="home-content-post-footer-delete-button-container">
								<button class="home-content-post-footer-delete-button" v-if="post.userID == localStorage.userID" @click="deletePost(post.postID)">Delete</button>
							</div>
						</div>
					</div>
				</div>
				<div class="home-content-no-data" v-if="!dataAvailable">No data available</div>
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
