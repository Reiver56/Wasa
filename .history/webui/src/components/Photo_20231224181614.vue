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
			this.photoURL = _API_URL + "/users/"+this.owner+"/photos/" + this.photo_id;
		},
		async deletePhoto(){
			try{
				await this.$axios.delete(`users/${this.$route.params.id}/photos/${this.photo_id}`, {
					headers: {
						Authorization: `Bearer ${localStorage.token}`
					}
				});
				this.$emit('photo-deleted', this.photo_id);
			}catch(err){
				this.$emit('error-occurred', err);
			}
		},
		photoOwnerClick(){
			this.$router.push(`/users/${this.owner}`);
		},

		async toggleLike(){
			if ()
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
