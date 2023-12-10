<template>
	<div>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />

	</div>
</template>



<script>
export default {
	data: function() {
		return {
				photos: [],
				errorMsg: "",
			}
	},
	methods: {

		getMyStream: async function(){
			

	},
	mounted() {
	}
}
</script>



<style>
</style>
