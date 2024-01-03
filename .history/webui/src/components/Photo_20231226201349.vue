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
						Authorization: `${localStorage.token}`
					}
				});
				this.$emit('photo-deleted', this.photo_id);
			}catch(err){
				this.$emit('error-occurred', err);
			}
		},
		photoOwnerClick: function(){
			this.$router.replace(`/users/${this.owner}`);
		},

		async toggleLike(){
			if (this.isOwner){
				return;
			}
			try{
				if (this.liked){
					await this.$axios.delete(`users/${this.$route.params.id}/photos/${this.photo_id}/likes`, {
						headers: {
							Authorization: `${localStorage.token}`
						}
					});
					this.liked = false;
					this.$emit('update-like', this.photo_id, false);
				}else{
					await this.$axios.put(`users/${this.$route.params.id}/photos/${this.photo_id}/likes`, {}, {
						headers: {
							Authorization: `${localStorage.token}`
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
		console.log(this.owner);
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
					<button v-if="isOwner" class="btn my-btn-danger" @click="deletePhoto">
						<font-awesome-icon icon="fa-regular fa-trash" />
					</button>
				</div>
				<div class="d-flex justify-content-center photo-background-color">
					<img :src="photoURL" class="photo" alt="Photo">
				</div>
				<div class="user-body">
					<div class="container">
						<div class="d-flex flex-row justify-content-end align-items-center mb-2">
							<button class="my-trnsp-btn-user m-0 p-1 me-auto" @click="photoOwnerClick">
								<i>{{owner}}</i>
							</button>
							<button class="my-trnsp-btn-heart">
								<font-awesome-icon icon="fa-regular fa-heart me-1" />
                                <i @click="toggleLike" :class="'me-1 my-heart-color w-100 h-100 fa '+(liked ? 'fa-heart' : 'fa-heart-o') "></i>
                                <i data-bs-toggle="model" :data-bs-target="'#like_model'+photo_id" class="my-comment-color ">
                                    {{Likes.length}}
                                </i>
                            </button>
							<button class="my-trnsp-btn-comment"
							data-bs-toggle="model" :data-bs-target="'#comment_model'+photo_id">
								<font-awesome-icon icon="fa-regular fa-comment me-1" />
                                <i class="my-comment-color t me-1" @click="commentClick"></i>
                                <i class="my-comment-color-2"> {{Comments != null ? Comments.length : 0}}</i>
                            </button>
						</div>
						<div class="d-flex flex-row justify-content-start align-items-center ">
                            <p>{{upload_date}}</p>
                        </div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
<style>

.photo-background-color{
	background-color: rgba(128, 128, 128, 0);
	border-radius: 4em;

}
.photo{
	width: 100%;
	height: auto;
	border-radius: 4em;

}
.my-user{
	width: 27em;
	border-color: black;
	border-width: thin;
	border-radius: 4em;
}
.my-heart-color{
	color: grey;
}
.my-comment-color{
	color: grey;
}
.my-comment-color-2{
	color: var(--color-blue-light);
}
.my-btn-danger:hover{
	border: none;
	color: var(--color-blue-danger);
	transform: scale(1.1);
}
.my-btn-danger{
	border: none;
}
.my-trnsp-btn-heart{
	background-color: transparent;
	border: none;
}
.my-trnsp-btn-heart:hover{
	color: rgb(237, 115, 115)
}
.my-trnsp-btn-comment{
	background-color: transparent;
	border: none;
}
.my-trnsp-btn-comment:hover{
	color: rgb(129, 199, 212);
}
.my-trnsp-btn{
	background-color: transparent;
	border: none;
}

.my-trnsp-btn-user{
	background-color: transparent;
	border: none;
}


</style>
