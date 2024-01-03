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
<template>
	<div class="row mb-2 mt-2">
		<div class="col d-flex justify-content-center">
			<>
		</div>

	</div>
</template>
