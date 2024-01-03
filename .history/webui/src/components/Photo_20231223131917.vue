<script>
export default{
	data(){
		return{
			liked: false,
			likes: [],
			comments: [],
			photoURL: "",
		}
	},

	methods: {
		loadPhoto(){
			this.photoURL = `/users/${localStorage.UserID}/photos/${this.photo.id}`;
			console.log(this.photoURL);
		},

		likePhoto: async function(){
			try{
				if (!this.liked){

					let response = await axios.put(`/users/${localStorage.UserID}/photos/${this.photo.id}/likes/` , { headers: { 'Authorization': `${localStorage.token}` } } );
					this.liked = true;
					this.likes.push(response.data);
				}else{
					let response = await axios.delete(`/users/${localStorage.UserID}/photos/${this.photo.id}/likes/` , { headers: { 'Authorization': `${localStorage.token}` } } );
					this.liked = false;
					this.likes.pop();
				}
			}catch(err){
				console.log(err);
			}
		},

		addComment: function(comment){
			this.comments.push(comment);
		},

		removeComment: function(comment){
			this.comments.splice(this.comments.indexOf(comment), 1);
		},
		photoOwnerClick: function(){
			this.$router.push(`/users/${this.photo.user.id}`);

		},

		deletePhoto: async function(){
			try{
				let response = await axios.delete(`/users/${localStorage.UserID}/photos/${this.photo.id}`);
				this.$emit("deletePhoto", this.photo.id);

			}catch(err){
				console.log(err);
			}
		},
	},
	async mounted(){
		await this.loadPhoto();
		if ( this.likes != null){
			this.likes = this.likes;
		}
		if ( this.likes != null){
			this.likes = this.likes.some(obj => obj.user.id === localStorage.UserID);
		}
		if ( this.comments != null){
			this.comments = this.comments;
		}
	}
}


</script>
<template>
	<div class="photo-container">
		<div class="photo-header">
			<div class="photo-user-icon" @click="photoOwnerClick">
				{{ photo.user.nickname[0] }}
			</div>
			<div class="photo-nickname" @click="photoOwnerClick">
				{{ photo.user.nickname }}
			</div>
			<div class="photo-delete" v-if="photo.user.id == localStorage.UserID" @click="deletePhoto">
				<i class="fas fa-trash-alt"></i>
			</div>
		</div>
		<div class="photo-image-container">
			<img :src="photoURL" class="photo-image">
		</div>
		<div class="photo-footer">
			<div class="photo-likes">
				<div class="photo-like-icon" @click="likePhoto">
					<i class="fas fa-heart" v-if="liked"></i>
					<i class="far fa-heart" v-else></i>
				</div>
				<div class="photo-like-count">
					{{ likes.length }}
				</div>
			</div>
			<div class="photo-comments">
				<div class="photo-comment" v-for="comment in comments">
					<div class="photo-comment-nickname" @click="comment.user.id == localStorage.UserID ? $router.push(`/users/${localStorage.UserID}`) : $router.push(`/users/${comment.user.id}`)">
						{{ comment.user.nickname }}
					</div>
					<div class="photo-comment-text">
						{{ comment.text }}
					</div>
					<div class="photo-comment-delete" v-if="comment.user.id == localStorage.UserID" @click="removeComment(comment)">
						<i class="fas fa-trash-alt"></i>
					</div>
				</div>
			</div>
			<div class="photo-comment-input">
				<PhotoCommentInput :photoID="photo.id" @addComment="addComment"/>
			</div>
		</div>
	</div>
</template>
