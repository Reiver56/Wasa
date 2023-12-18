<template>
	 <div class="upload-form-background" @click.self="this.$emit('exit-upload-form')">
        <div class="upload-form-container" v-if="!file64">
            <div class="drag-drop-area-container">
                <button class="drag-drop-area" @click="this.$refs.file.click()">
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                    <span class="drag-drop-area-text">
                        Drop your photo here
                    </span>
                    <span class="drag-drop-area-subtext">
                        max size 5MB, only jpg, jpeg
                    </span>

                </button>
            </div>
            <div class="bottom-area">
                <button @click="this.$refs.file.click()" class="upload-button">Choose File
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                </button>

            </div>
        </div>

        <EditorPost :image64="file64" :editorType="this.$props.photoType" v-if="file64"
            @exit-upload-form="this.$emit('exit-upload-form')" @save-upload-form="saveData" />
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
