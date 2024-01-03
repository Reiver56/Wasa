<script>
export default{
	props: ['identifier', 'nickname'],
	methods:{
		clickedUser(){
			this.$emit('clikedUser', this.idetifier)
		}
	},
}
</script>
<temple
