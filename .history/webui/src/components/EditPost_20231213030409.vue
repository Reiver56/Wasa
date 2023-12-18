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
import {RectagleStencil, CircleStencil, Cropper} from 'vue-advanced-spinner';
import { markRaw } from 'vue';

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


