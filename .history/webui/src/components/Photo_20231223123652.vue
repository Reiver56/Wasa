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
			this.photoURL = _API_URL_ + "users"+this.UserID+"/photos/" + this.photo.id;
			console.log(this.photoURL);
		},
		deletePhoto: async function(){
			try{
				let response = await axios.delete(`/users/${localStorage.UserID}/photos/${this.photo.id}`);
				this.$emit("deletePhoto", this.photo.id);

			}catch(err){
				console.log(err);
			}
		},
	}
}


</script>
