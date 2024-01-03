<script>
export default{
	data(){
		return{
			liked: false,
			likes: [],
			comments: [],
			photoURL: "",
		}
	},

	methods: {
		loadPhoto(){
			this.photoURL = `/users/${localStorage.UserID}/photos/${this.photo.id}`;
			console.log(this.photoURL);
		},

		likePhoto: async function(){
			try{
				if (!this.liked){

					let response = await axios.put(`/users/${localStorage.UserID}/photos/${this.photo.id}/likes/` , { headers: { 'Authorization': `${localStorage.token}` } } );
					this.liked = true;
					this.likes.push(response.data);
				}else{
					let response = await axios.delete(`/users/${localStorage.UserID}/photos/${this.photo.id}/likes/` , { headers: { 'Authorization': `${localStorage.token}` } } );
					this.liked = false;
					this.likes.pop();
				}
			}catch(err){
				console.log(err);
			}
		},

		addComment: function(comment){
			this.comments.push(comment);
		},

		removeComment: function(comment){
			this.comments.splice(this.comments.indexOf(comment), 1);
		},
		photoOwnerClick: function(){
			this.$router.push(`/users/${this.photo.user.id}`);

		},

		deletePhoto: async function(){
			try{
				let response = await axios.delete(`/users/${localStorage.UserID}/photos/${this.photo.id}`);
				this.$emit("deletePhoto", this.photo.id);

			}catch(err){
				console.log(err);
			}
		},
	},
	async mounted(){
		await this.loadPhoto();
		
	}
}


</script>
<template>

</template>
