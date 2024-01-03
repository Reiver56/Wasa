<script>
export default{
	emits: ['error-occurred', 'update-like'],
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			photoURL: "",
			liked: false,
			Comments: [],
			allLikes: [],
		}
	},
	props: ["owner", "likes", "comments", "upload_date", "photo_id", "isOwner"],
	methods: {
		loadPhoto(){
			this.photoURL = _API_URL + "/p/" + this.photo_id;
		},

	},
	mounted(){
		await this.loadPhoto();

		if (this.ownerID == localStorage.userID){
			this.isOwner = true;
		}
	}
}
</script>
<template>
	<div class="photo-container">

	</div>
</template>
