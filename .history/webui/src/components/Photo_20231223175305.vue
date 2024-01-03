<script>
export default{
	emits: ['error-occurred', 'update-like'],
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			ownerID: this.postData.ownerID['userID'],
			ownerName: this.postData.ownerID['nickname'],

			// Photo data
			photoID: this.postData.photoID,
			isLikes: this.postData.liked,
			image64: this.postData.image,
			likesCount: this.postData.likesCount,
			commentsCount: this.postData.commentsCount,

			isOwner: false,
			isHoverComment: false,

			profilesList: [],
			textCounter: 0,
			showList: false,
			dataGetter: () => {},
			dataUpdater: () => {},
			additionalData: {},

			errorMsg: "",

		}
	},
	
}
</script>
<template>
	<div class="photo-container">

	</div>
</template>
