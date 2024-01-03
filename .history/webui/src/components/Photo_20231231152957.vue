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
	props: ["owner", "likes", "comments", "date", "photo_id", "isOwner"],
	methods: {
		loadPhoto(){
			console.log(this)
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
				console.log("qui"+this.liked);
				if (this.liked){
					await this.$axios.delete(`users/${this.owner}/photos/${this.photo_id}/likes/${localStorage.UserID}`, {
						headers: {
							Authorization: `${localStorage.token}`
						}
					});
					this.Likes.pop();
					this.liked = false;
					this.$emit('update-like', this.photo_id, false);
				}else{
					await this.$axios.put(`users/${this.owner}/photos/${this.photo_id}/likes/${localStorage.UserID}`,
					 		this.Likes.push({
							user_id: localStorage.token,
						}),
						{headers: {
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
		addCommentToList(comment){
			this.Comments.push(comment);
		},

		},

	async mounted(){
		console.log(this.isOwner);
		await this.loadPhoto();
		if (this.likes != null){
			this.Likes = this.likes;
		}
		if (this.likes != null){
			this.liked = this.Likes.some(obj=> obj.id_user === localStorage.token);
			console.log(this.liked);
		}

		if (this.comments != null){
			this.Comments = this.comments;
		}

	},
}
</script>
<template>
	<div class="photo-container">
		<LikeModal :modal_id="'like_modal'+photo_id"
			:likes="Likes" />
		<CommentModal :modal_id="'comment_modal'+photo_id"
			:comments_list="Comments"
			:photo_owner="owner"
			:photo_id="photo_id"

			@comment-deleted="removeCommentFromList"
			@postComment="addCommentToList"
			/>
		<div class="d-flex flex-row justify-content-center">
			<div class="user my-user">
				<div class="d-flex justify-content-end">
					<button v-if="isOwner" class="btn my-btn-danger" @click="deletePhoto">
						<font-awesome-icon icon="fa-solid fa-trash" />
					</button>
				</div>
				<div class="d-flex justify-content-center photo-background-color">
					<img :src="photoURL" class="photo" alt="Photo">
				</div>
				<div class="user-body">
					<div class="container">
						<div class="d-flex flex-row justify-content-end align-items-center mb-2">
							<button class="my-trnsp-btn-user m-0 p-1 me-auto" @click="photoOwnerClick">
								<i class="owner-button">{{owner}}</i>
							</button>
							<button class="my-trnsp-btn-heart m-0 p-1 d-flex justify-content-center align-items-center " @click="toggleLike" >
								<i class="my-heart" @click="toggleLike"><font-awesome-icon icon="fa-regular fa-heart me-1" /></i>
                            </button>
							<button data-bs-toggle="modal" :data-bs-target="'#like_modal'+photo_id" class="counter-of-like-color m-0 p-1 ">{{Likes.length}}</button>
							<button class="my-trnsp-btn m-0 p-1 d-flex justify-content-center align-items-center"
							@click="commentClick" data-bs-toggle="modal" :data-bs-target="'#comment_modal'+photo_id">
								<font-awesome-icon icon="fa-regular fa-comment me-1" />
								<div class="comments-counter-container">{{Comments != null ? Comments.length : 0}}</div>
                            </button>
						</div>
						<div class="d-flex flex-row justify-content-start align-items-center ">
                            <p>{{date}}</p>
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
.my-heart{
	color: black;
	width: 12px;
}
.my-heart:hover{
	color: red;
	transform: scale(1.1);
}
.counter-of-like-color{
	color: black;
	background-color: transparent;
	border: none;
	width: 18px;
}

.comments-counter-container{
	background-color: white;
	border-radius: 50%;
	width: 1.5em;
	height: 1.5em;
	display: flex;
	justify-content: center;
	align-items: center;
}
.my-comment-color{
	color: black;
}
.my-comment-color:hover{
	color: rgb(129, 199, 212);
	transform: scale(1.1);
}

.owner-button{
	color: black;
	font-size: 1.2em;
}
.owner-button:hover{
	transform: scale(1.1);
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
