<script>
export default{
	props: ['photo_id','nickname','comment_id','photo_owner','content','author'],
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
	
}
</script>
