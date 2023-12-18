<template>
	<div class="upload-photo-container">
		<div class="upload-photo-container-header">
			<span class="upload-photo-container-header-title">Upload Photo</span>
			<span class="upload-photo-container-header-exit" @click="$emit('exit-form')">
				<i class="fas fa-times"></i>
			</span>
		</div>
		<div class="upload-photo-container-body">
			<div class="upload-photo-container-body-image">
				<Cropper
					:image64="file64"
					:editorType="postData"
					@save="saveData"
					:cropperProps="cropperProps"
				/>
			</div>
			<div class="upload-photo-container-body-input">
				<input type="file" ref="file" @change="onChange" class="upload-photo-container-body-input-file">
			</div>
		</div>
		<div class="upload-photo-container-footer">
			<span class="upload-photo-container-footer-error">{{errorMsg}}</span>
		</div>
	</div>
</template>


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

			try{
				let response = await this.$axios.post(`/users/${localStorage.userID}/photos`, formData, { headers : { 'Authorization': `${localStorage.token}` , 'content-type': 'multipart/form-data'} });
				this.$emit('refresh-data');
				setTimeout(() => {
					this.$emit('exit-form');
				}, 1000);

			}catch(err){
				this.$emit('error', err);
			}
		}


	},
	watch: {
		errorMsg() {
			this.$emit('error', this.errorMsg);
			console.log(this.errorMsg);
		},
	},
}
</script>

<style>
</style>
