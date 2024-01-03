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
		deleteCo
	},

}
</script>
