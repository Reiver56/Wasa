<template>
	<div class="edit-post-container">
		<cropper class="cropper" :image="image64" :auto-zoom="true" ref="cropper" :stencil-component="cropperProps.stencilComponent" :stencil-size="cropperProps.stencilSize" image-restriction="stencil" :stencil-props="{ aspectRatio : 1/ 1}"/>
		<div class="edit-post-container-header">
			<span class="edit-post-container-header-title">Edit Post</span>
			<span class="edit-post-container-header-close" @click="$emit('exit-form')">
				<i class="fas fa-times"></i>
			</span>
		</div>
		<div class="button-container">
            <button class="cancel-button" @click="$emit('exit-upload-form')">Cancel</button>
            <button class="save-button" @click="saveEdit">Save and Upload</button>
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


.edit-post-container {
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

.edit-post-container-header {
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
