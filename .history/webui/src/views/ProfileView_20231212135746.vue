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
			listType: "",

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
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outline = "auto";
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outlineColor = "#3b9ee4";
			}
		},

		async saveChangeNickname() {
			if (this.isOwner){
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outline = "none";
				if (this.nickname == "" | this.nickname.length <3){
					this.nickname = localStorage.nickname;
					return;
				}
				this.isLoading = true;

				try{
					let response = await this.$axios.put(`users/${this.userID}`, { nickname: this.nickname }, { headers : { 'Authorization': `${localStorage.token}` } });
					localStorage.nickname = this.nickname;
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}
		},

		getFollowers() {
			this.header = "Followers";
			this.listType ="simple";
			this.showList = true;
			this.dataGet = async (usersArray, dataAvailable) => {
				try{
					let response = await this.$axios.get(`users/${this.userID}/followers`, { headers : { 'Authorization': `${localStorage.token}` } });
					usersArray = response.data;
					dataAvailable = true;
					usersArray.push(...response.data);
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}

		},

		getFollowing() {
			this.header = "Followers";
			this.listType ="simple";
			this.showList = true;
			this.dataGet = async (usersArray, dataAvailable) => {
				try{
					let response = await this.$axios.get(`users/${this.userID}/following`, { headers : { 'Authorization': `${localStorage.token}` } });
					usersArray = response.data;
					dataAvailable = true;
					usersArray.push(...response.data);
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}

		},
		freeLists(){
			this.showList = false;
			this.usersArray = [];
			this.header = "";
		},
		async follow(){
			if (this.isFollowed){
				try{
					let _ 
				}
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
		<div class="top-body-profile-container">

		</div>
	</div>


</template>
