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
		<div class>

		</div>

	</div>
</template>
