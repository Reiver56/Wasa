<script>
import EditorPost from '@/components/EditPost.vue';
export default{
	components:{
		EditorPost,
	},
	emits: ['exit-form', 'refresh-data', 'error'],
	props: {
		postData: {
			type: String,
			required: true,
			validator(value) {
				return ['post', 'profile'].includes(value);
			},
		},
	},
	data() {
		return {
			file64: "",
			errorMsg: "",
		}
	},
	methods: {
		onChange() {
			const file = this.$refs.file.files[0];
			// check if file is jpg and jpeg, that are the only supported formats
			if (file.type != "image/jpeg" && file.type != "image/jpg") {
				this.errorMsg = "Only jpg and jpeg files are supported";
				return;
			}

			// check if file is less than 5MB
			if (file.size > 5242880) {
				this.errorMsg = "File is too big, max size is 5MB";
				return;
			}

			// convert file to base64
			const reader = new FileReader();
			reader.readAsDataURL(file);
			reader.onload = () => {
				this.file64 = reader.result;
			};

		},
		saveData(data) {
			if (this.file.size == 'post') {
				this.createPost(data);
			}

		},

		async createPost(postData) {
			const formData = new FormData();
			formData.append('image', postData.imageFile);

			try{}
		}


	},
}
</script>
