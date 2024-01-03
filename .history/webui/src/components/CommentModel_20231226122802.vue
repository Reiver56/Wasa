<script>
export default{
	props: ['model_id', 'comments', 'photo_owner', 'photo_id'],
	data(){
		return{
			commentText: '',
		}
	},
	methods: {
		async postComment(){
			try{
				let response = await this.$axios.post(`users/${this.photo_owner}/photos/${this.photo_id}/comments`,{ comment: this.commentText} ,{
					headers : { 'Authorization': `${localStorage.token}` }
				});
				this.$emit('comment-posted', response.data);
			}catch(err){
				this.$emit('error-occurred', err);
			}
		},
		deleteComment(value){
			this.$emit('comment-deleted', value);

		},
		addCommentToList(value){
			this.$emit('postComment', value);
		},
	},

}
</script>
<template>
	<div class="model fade my-model-disp-none" :id="model_id" tabindex="-1" aria-hidden="true">
		<div class=" model-dialog">
			<div class="model-content">
				<div class="model-header">
					<h5 class="model-title">Comments</h5>
					<button type="button" class="btn-close" data-bs-dismiss="model" aria-label="Close"></button>
				</div>
				<div class= "model-body">
					<div class="row">
						<div class="col-12">
							<div class="input-group mb-3">
								<input type="text" class="form-control" placeholder="Comment" v-model="commentText">
								<button class="btn btn-outline-secondary" type="button" id="button-addon2" @click="postComment">Post</button>
							</div>
						</div>
					</div>
					<Comment v-for="(comment, index) in comments"
					:key="index"
					:comment_id="comment.comment_id"
					:comment_text="comment.comment_text"
					:comment_owner="comment.comment_owner"
					:photo_owner="photo_owner"
					:photo_id="photo_id"
					:deleteComment="deleteComment"
					:addCommentToList="addCommentToList"
					/>
				</div>
			</div>

	</div>
</template>
