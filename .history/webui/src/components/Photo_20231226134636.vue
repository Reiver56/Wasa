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
			Likes: [],
		}
	},
	props: ["owner", "likes", "comments", "upload_date", "photo_id", "isOwner"],
	methods: {
		loadPhoto(){
			this.photoURL = __API_URL__ + "/users/"+this.owner+"/photos/" + this.photo_id;
			console.log(this.photoURL);
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
			if (this.isOwner){
				return;
			}
			try{
				if (this.liked){
					await this.$axios.delete(`users/${this.$route.params.id}/photos/${this.photo_id}/likes`, {
						headers: {
							Authorization: `Bearer ${localStorage.token}`
						}
					});
					this.liked = false;
					this.$emit('update-like', this.photo_id, false);
				}else{
					await this.$axios.put(`users/${this.$route.params.id}/photos/${this.photo_id}/likes`, {}, {
						headers: {
							Authorization: `Bearer ${localStorage.token}`
						}
					});
					this.liked = true;
					this.$emit('update-like', this.photo_id, true);
				}
			}catch(err){
				this.$emit('error-occurred', err);
			}
		},

		removeCommentFromList(value){
			this.Comments = this.Comments.filter(item => item.comment_id != value);
		},
		addCommentToList(value){
			this.Comments.push(value);
		},

		},


	async mounted(){
		await this.loadPhoto();
		if (this.likes != null){
			this.Likes = this.likes;
			this.liked = this.Likes.some(obj=> obj.user_id == localStorage.userID);
		}
		if (this.comments != null){
			this.Comments = this.comments;
		}

	},
}
</script>
<template>
	<div class="photo-container">
		<LikeModel :model_id="'like_model'+photo_id"
			:likes="Likes" />
		<CommentModel :model_id="'comment_model'+photo_id"
			:comments="Comments"
			:photo_owner="owner"
			:photo_id="photo_id"
			@comment-deleted="removeCommentFromList"
			@postComment="addCommentToList"
			/>
		<div class="d-flex flex-row justify-content-center">
			<div class="user my-user">
				<div class="d-flex justify-content-end">
					<button v-if="isOwner" class="btn my-btn-danger" @click="deletePhoto">Delete</button>
				</div>
				<div class="d-flex justify-content-center photo-background-color">
					<img :src="photoURL" class="photo" alt="Photo">
				</div>
				<div class="user-body">
					<div class="container">
						<div class="d-flex flex-row justify">

						</div>
					</div>
				</div>

			</div>
		</div>

	</div>
</template>
