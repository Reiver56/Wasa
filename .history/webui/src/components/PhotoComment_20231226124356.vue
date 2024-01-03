<script>
export default{
	props: ['photo_id','nickname','comment_id','photo_owner','content','author'],
	data(){
		return{
			user: '',
		}
	},
	methods:{
		async deleteComment(){
			try{
				await this.$axios.delete(`users/${this.photo_owner}/photos/${this.photo_id}/comments/${this.comment_id}`, {
					headers: {
						Authorization: `Bearer ${localStorage.token}`
					}
				});
			}catch(err){
				this.$emit('error-occurred', err);
			}
		},
	},
	mounted(){
		this.user = localStorage.token;
	},
}
</script>
<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col-10">
				<h5>{{nickname}} @{{author}}</h5>
			</div>
			<div class="col-2">
				<button v-if="user === author || user === photo_owner" class="btn my-btn-danger" @click="deleteComment">Delete</button>
			</div>
		</div>
		<div class="row">
			<div class="col-12">
				{{content}}
			</div>
		</div>
	</div>
</template>
