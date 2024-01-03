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
				let response = await this.$axios.post(`users/${this.photo_owner}/photos/${this.photo_id}/comments`,{ user} ,{
					headers : { 'Authorization': `${localStorage.token}` }
				});
			}catch(err){
				this.$emit('error-occurred', err);
			}
		}
	},

}
</script>
