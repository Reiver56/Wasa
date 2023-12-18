<script>
export default {
	data(){
		return{
			errorMsg: "",

			userID: "",
			nickname: "",
			followersCount: 0,
			followingCount: 0,
			photosCount: 0,

			isFollowed: false,

			isOwner: false,

			followButton: "Follow",

			textCounter: 0,
			profiles: [],
			header: "",

			posts: [],
			showPosts: false,
			postViewData: {},

			dataGet : () => { },
			showList: false,

			isLoading: false,

		}
	},
	methods: {
		getProfile: async function() {
			this.isLoading = true;
			try{
				let response = await this.$axios.get(`users/${this.userID}`, { headers : { 'Authorization': `${localStorage.token}` } });
				this.userID = response.data.user.user_id;
				this.nickname = response.data.user.nickname;
				this.followersCount = response.data.user.followers.length;
				this.followingCount = response.data.user.following.length;
				this.photosCount = response.data.user.photos.length;
				this.isFollowed = response.data.isFollowed;
				this.isOwner = localStorage.UserID == this.userID;
				this.followButton = this.isFollowed ? "Unfollow" : "Follow";

			}
			catch(e){
				this.errorMsg = e.message;
			}
			this.isLoading = false;
		},


		async getPhotos(){
			this.isLoading = true;
			try{
				let response = await this.$axios.get(`users/${this.userID}/photos`, { headers : { 'Authorization': `${localStorage.token}` } });
				this.posts = response.data;
				this.showPosts = true;
			}
			catch(e){
				this.errorMsg = e.message;
			}
			this.isLoading = false;
		},
		editingNickname() {
			if (this.isOwner){
				document.querySelectorAll(".top-body-profile")
			}
		},
	},
	mounted() {
		this.getProfile();
	}

}
</script>

<template>
	<LoadingSpinner :loading="isLoading" />
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>

	<div class = "top-container-profile">

	</div>


</template>
