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
					<h1 class="model-title" :id>Comments</h1>

				</div>
			</div>
		</div>
	</div>
</template>
