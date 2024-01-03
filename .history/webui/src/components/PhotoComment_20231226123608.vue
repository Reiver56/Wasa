<script>
export default{
	props: ['photo_id','nickname','comment_id','photo_owner','content','author'],
	methods:{
		async deleteComment(){
			try{
				a

			}catch(err){
				this.$emit('error-occurred', err);
			}

		},
	},
}
</script>
