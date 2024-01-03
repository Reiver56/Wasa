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
		<div class="model-dialog model-dialog-centered">
			<div class="model-header">
				<div class="model-header">
					<h1 class="model-title" :id="model_id">Comments</h1>
					<button type="button" class="btn-close" data-bs-dismiss="model" aria-label="Close"></button>
				</div>
				<div class="model-body">
					<PhotoComment v-for="(comment,index) in comments"
					:key="index"
					:author="comment.user_id"
					:nickname="comment.nickname"
					:comment_id="comment.comment_id"
					:photo_id="comment.photo_id"
					:content="comment.content"
					:photo_owner="photo_owner"

					@comment-deleted="deleteComment"
					/>
				</div>
				<div class="model-footer d-flex justify-content-center w-100">
					<div class="row w-100">
						<div class="col-10">
							<div class="mb-3 me-auto">
								<textarea class="form-control" id="exampleFormTextArea" rows="1" placeholder="Add a comment..." maxlength="30" v-model="commentText"></textarea>
							</div>
						</div>
						<div class="col-2 d-flex align-items-center">
							<button type="button" class="btn btn-primary" @click.prevent="postComment" :disabled="commentText.length < 1 || commentText.length > 30">Post</button>
						</div>
					</div>

				</div>
			</div>
		</div>
	</div>
</template>
<style>
.my
</style>
