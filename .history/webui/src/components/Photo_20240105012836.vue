<script>
export default{
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			photoURL: "",
			liked: false,
			Comments: [],
			Likes: [],
			errorMsg: "",
		}
	},
	props: ["owner", "likes", "comments", "date", "photo_id", "isOwner"],
	methods: {
		loadPhoto(){

			this.photoURL = __API_URL__ + "/users/"+this.owner+"/photos/" + this.photo_id;
			//console.log(this.photoURL);
		},
		async deletePhoto(){
			try{
				await this.$axios.delete(`users/${this.$route.params.id}/photos/${this.photo_id}`, {
					headers: {
						Authorization: `${localStorage.token}`
					}
				});
				location.reload();
			}catch(err){
				this.errorMsg = err.message;
			}
		},
		photoOwnerClick: function(){
			console.log(this.Comments);
			this.$router.push(`/users/${this.owner}`);
		},

		async toggleLike(){
			if (this.isOwner){
				return;
			}
			try{

				if (this.liked){
					await this.$axios.delete(`users/${this.owner}/photos/${this.photo_id}/likes/${localStorage.UserID}`, {
						headers: {
							'Authorization': `${localStorage.token}`,
						}
					});
					//console.log("qui "+this.liked);
					this.Likes.pop();
					this.liked = false;

				}else{
					console.log(localStorage.token);
					await this.$axios.put(`users/${this.owner}/photos/${this.photo_id}/likes/${localStorage.UserID}`, {}, {
						headers: {
							Authorization: `${localStorage.token}`
						}
					});
					//console.log("que "+this.liked);
					this.Likes.push({id_user: localStorage.UserID});
					this.liked = true;
				}
			}catch(err){
				this.errorMsg = err.message;
			}
		},

		removeCommentFromList(value){
			console.log("value "+ value);
			this.Comments = this.Comments.filter(item=> item.comment_id !== value)
		},

		addCommentToList(comment){
			console.log( "comment "+ comment);
			this.Comments.push(comment)
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

				@eliminateComment="removeCommentFromList"
				@addComment="addCommentToList"
		/>
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
								<i class="btn owner-button">{{owner}}</i>
							</button>
							<button class="btn my-trnsp-btn-heart m-0 p-1 d-flex justify-content-center align-items-center " @click="toggleLike" >
								<i class="my-heart"><font-awesome-icon icon="fa-regular fa-heart me-1" /></i>
                            </button>
							<button data-bs-toggle="modal" :data-bs-target="'#like_modal'+photo_id" class="counter-of-like-color m-0 p-1 ">{{Likes.length}}</button>
							<button class="btn my-trnsp-btn-comment m-0 p-1 d-flex justify-content-center align-items-center"
							@click="commentClick" data-bs-toggle="modal" :data-bs-target="'#comment_modal'+photo_id">
								<font-awesome-icon icon="fa-regular fa-comment me-1" />

                            </button>
							<div class="comments-counter-container">{{Comments != null ? Comments.length : 0}}</div>
						</div>
						<div class="d-flex flex-row justify-content-start align-items-center ">
                            <p>{{ date }}</p>
                        </div>
					</div>
				</div>
			</div>
		</div>
</template>
<style>
.photo-container{
	width: 100%;
	height: auto;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
}

.photo-background-color{
	background-color: rgba(128, 128, 128, 0);
	border-radius: 4em;

}
.photo{
	width: 100%;
	height: auto;
	border-radius: 4em;

}
.us
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
}

.owner-button{
	color: black;
	font-size: 1.2em;
}
.owner-button:hover{
	color: rgb(129, 199, 212);
}

.my-btn-danger:hover{
	border: none;
	color: var(--color-blue-danger);

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
.my-trnsp-btn-comment:hover{
	color: rgb(129, 199, 212);
}
.comments-counter-container{
	color: black;
}
.comments-counter-container{
	width: 10px;
	font-size: 16px;
}



</style>
