<script>
export default{
	props: ['modal_id', 'comments', 'photo_owner', 'photo_id'],
	data(){
		return{
			commentText: "",
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
	<div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable">
			<div class="modal-header">
				<div class="modal-header">
					<h1 class="modal-title fs-5" :id="modal_id">Comments</h1>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<PhotoComment v-for="comment in comments"
					:author="comment.user_id"
					:nickname="comment.nickname"
					:comment_id="comment.comment_id"
					:photo_id="comment.photo_id"
					:content="comment.content"
					:photo_owner="photo_owner"

					@comment-deleted="deleteComment"
					/>
				</div>

				<div class="modal-footer d-flex justify-content-center w-100">

				</div>
			</div>
		</div>
	</div>
</template>
<style>
.my-modal-disp-none{
	background-color:#ffffff;
	box-sizing: ;
	display: none;

}
.modal-title{
	color: #000000;
}
</style>
