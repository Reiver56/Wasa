<template>
	 <div class="upload-form-background" @click.self="this.$emit('exit-upload-form')">
        <div class="upload-form-container" v-if="!file64">
            <div class="drag-drop-area-container">
                <button class="drag-drop-area" @click="this.$refs.file.click()">
                    <input type="file" ref="file" accept=".jpg,.jpeg" @change="onChange" hidden />
                    <span class="drag-drop-area-text">
                        Upload your photo here
                    </span>
                    <span class="drag-drop-area-subtext">
                        max size 5MB, only jpg, jpeg
                    </span>
					<div class="upload-icon-container">
						<font-awesome-icon icon="upload" />
					</div>


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
.upload-form-background {
    background-color: rgba(0, 0, 0, 0.5);
    height: 100%;
    width: 100%;
    padding: 0;
    z-index: 4;

    position: fixed;
    top: 0;
    left: 0;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.upload-form-container {
    background-color: #fff;
    height: 40em;
    width: 30em;
    padding: 0;
    border-radius: 0.5em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: fixed;

    padding: 2em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;

}

.drag-drop-area-container {

    height: 100%;
    width: 100%;

    padding: 0;


    cursor: pointer;
}

.drag-drop-area {
    background-color: #c7deeb83;

    height: 100%;
    width: 100%;

    border-radius: 0.5em;
    border: 0.3em dashed rgb(0, 0, 0, 0.5);

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;


    position: relative;

    opacity: 1;
}

.drag-drop-area-text {
    font-size: 1.5em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    text-align: center;
    margin-top: 9em;
}

.drag-drop-area-subtext {
    font-size: 1em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    text-align: center;
    margin-bottom: 15em;
}

.bottom-area {
    height: 5em;
    width: auto;

    padding: 1em;
    display: flex;
    flex-direction: row;
    justify-content: center;

}

.upload-button {
    background-color: #0D92DD;

    height: 2.5em;
    width: 9em;

    padding: 0;

    border-radius: 10em;
    border: none;

    font-size: 1.2em;
    font-weight: 600;
}
.upload-button:hover{
	cursor: pointer;
	background-color: #0D99DD;
}

.upload-icon-container{
	height: 100%;
	width: 100%;
	font-size: 1.5em;
    font-weight: 500;
    color: rgb(0, 0, 0, 0.5);
    text-align: center;
    margin-top: 9em;

}
</style>
