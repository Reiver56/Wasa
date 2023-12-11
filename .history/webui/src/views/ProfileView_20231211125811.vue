<template>
 <NavBar />
</template>

<script>

export default {
	
	data: function(){
		return{
			errorMsg: "",
			userExist: false,
			isLoading: false,

			nickname: "",
			followStatus: false,
			currentIsBanned: false,

			followersCount: 0,
			followingCount: 0,
			photosCount: 0,

			photos: [],
			followers: [],
			following: [],
		}
	},
	watch: {
		Path(newid, oldid){

		}
	},
}
</script>
