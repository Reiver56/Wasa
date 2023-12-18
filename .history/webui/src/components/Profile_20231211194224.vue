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


			isLoading: false,

		}
	},
}
</script>

<template>


</template>
