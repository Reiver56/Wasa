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
			this.photoURL = _API_URL_ + "users"+this.+"/photos/" + this.photo.id + "/image""";
		}
	}
}


</script>
