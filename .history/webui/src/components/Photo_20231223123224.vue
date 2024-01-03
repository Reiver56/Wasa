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
		
	}
}


</script>
