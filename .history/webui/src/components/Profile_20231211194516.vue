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
				let response = await this.$axios.get(`users/${}`)
			}
			catch(e){
				this.errorMsg = e.message;
			}
		}
	}
}
</script>

<template>


</template>
