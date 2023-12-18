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

			if (this.)
		}
	}

}
</script>
