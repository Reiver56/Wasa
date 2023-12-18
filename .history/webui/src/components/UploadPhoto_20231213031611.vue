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
			

		}
	},
}
</script>
