<script>
export default{
	emits: ['exit-from-list', 'error-occurred', 'data-updated'],
	props: {
		data: { type: Object, required: true },
	}
}
</script>
