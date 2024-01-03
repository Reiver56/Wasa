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
			this.photoURL = _API_URL_ + "/photos/" + this.photo.id + "/image";
		}
	}
}


</script>
