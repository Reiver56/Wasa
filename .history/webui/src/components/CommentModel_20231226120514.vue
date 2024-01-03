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
				let response = await this.$axios.post(`users/${this.$route.params.id}`, { headers : { 'Authorization': `${localStorage.token}` } })
			}
		}
	},

}
</script>
