<template>
	<div>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<uploadPhoto v-if = "showPhotos" :photoType="'post'" @exit-upload-form="showPhotos = false"
		 @refresh="$router.go(0)" @error-occurred="(value) => { errorMsg = value }" >

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
