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
		function: async function(){
			const response = await axios.get(_API_URL_ + "users"+this.UserID+"/photos/" + this.photo.id);
			this.likes = response.data.likes;
			this.comments = response.data.comments;
		},
	}
}


</script>
