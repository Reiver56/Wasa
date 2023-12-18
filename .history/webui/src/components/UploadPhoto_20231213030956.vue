<script>
import EditorPost from '@/components/EditPost.vue';
export default{
	components:{
		EditorPost,
	},
	emits: ['exit-form', 'refresh-data', 'error'],
}
</script>
