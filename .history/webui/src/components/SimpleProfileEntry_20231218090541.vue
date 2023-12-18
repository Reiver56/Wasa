<script>
export default{
	emits: ['exit-from-list', 'error-occurred', 'data-updated'],
	props: {
		data: { type: Object, required: true },
	},
	data(){
		return {
			nickname: this.$props.data.nickname,
			userID: this.$props.data.userID,

		}
	},
	methods: {
		goToProfile(){
			this.$router.push(`/users/${this.userID}`);
			
		},

	},
}
</script>
