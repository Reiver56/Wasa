<template>
	<div class="edit-post-container">
		<div class="edit-post-container-header">
			<span class="edit-post-container-header-title">Edit Post</span>
			<span class="edit-post-container-header-close" @click="$emit('exit-form')">
				<i class="fas fa-times"></i>
			</span>
		</div>
		<div class="edit-post-container-body">
			<div class="edit-post-container-body-image">
				<Cropper
					:image64="image64"
					:editorType="editorType"
					@save="saveEdit"
					:cropperProps="cropperProps"
				/>
			</div>
		</div>
		</div>
</template>




<script>
import { RectangleStencil, CircleStencil, Cropper } from "vue-advanced-cropper";
import { markRaw } from "vue";
import "vue-advanced-cropper/dist/style.css";

export default {
	emits: ['exit-form', 'save-form'],
	components:{
		Cropper,
	},
	props: {
		image64: {
			type: String,
			required: true,
		},
		editorType: {
			type: String,
			required: true,
			validator(value) {
				return ['profile', 'post'].includes(value);
			},

		},
	},
	data() {
		return {
			imageData: {
				imageFile: null,
			},

			cropperProps: {
				stencilSize: {
					width: 200,
					height: 200,
				},

			},
		};
	},
	methods: {
		saveEdit() {
			const image = canvas.toDataURL("image/png")

			// convert image in blob
			const byteString = atob(image.split(',')[1]);
			const mimeString = image.split(',')[0].split(':')[1].split(';')[0];
			const ab = new ArrayBuffer(byteString.length);
			const ia = new Uint8Array(ab);
			for (let i = 0; i < byteString.length; i++) {
				ia[i] = byteString.charCodeAt(i);
			}
			const blob = new Blob([ab], {type: mimeString});

			// create object file
			const file = new File([blob], "image.png", {type: mimeString});
			this.imageData["imageFile"] = file;
			this.$emit('save-form', this.imageData);


		},
	},
	mounted () {
	},


}
</script>


<style>
.cropper {

    height: auto;
    min-height: 30vh;
    max-height: 90vh;

    width: auto;
    min-width: 30vw;
    max-width: 90vw;
}

.cropper-container {
    background-color: #fff;

    height: auto;
    min-height: 30vh;
    max-height: 90vh;

    width: auto;
    min-width: 30vw;
    max-width: 90vw;

    border-radius: 0.5em;

    box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);

    position: fixed;

    padding: 2em;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.caption-form-container {
    height: 5em;
    width: auto;

    margin: 1em 0 0 0;

    display: flex;
    flex-direction: column;

}

.label-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.caption-title {
    font-size: 1em;
    font-weight: 600;
}

.caption-form {
    width: 100%;

    padding: 0.5em;

    border: 0.1em solid rgb(0, 0, 0, 0.9);

    border-radius: 0.5em;

    resize: none;

    font-size: 1em;
    font-weight: 500;
}

.caption-text-counter {
    color: rgb(0, 0, 0, 0.5);

    font-size: 0.8em;
    font-weight: 400;

    float: right;
}

.button-container {
    height: auto;
    width: auto;

    margin-top: 1em;

    display: flex;
    flex-direction: row;
    justify-content: space-around;
}

.button-container button {

    height: 2.5em;
    width: 9em;

    padding: 0;
    border: 0.1em solid rgb(0, 0, 0, 0.9);

    border-radius: 0.5em;

    font-size: 1em;
    font-weight: 600;
}

.button-container button:hover {
    box-shadow: inset 0 0 0 10em rgb(0, 0, 0, 0.1);
}

.save-button {
    background-color: #03C988;
}

.cancel-button {
    background-color: transparent;
}
</style>
