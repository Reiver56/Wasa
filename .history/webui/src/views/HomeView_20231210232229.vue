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
				isLoading: false,
			}
	},
	methods: {

		getMyStream: async function(){
			try{
				this.isLoading = true;
				let response = await this.$axios.get('users'+ localStorage.getItem('token') + '/stream');
				if (response.data.length != 0){
					this.photos = response.data;
				}

			}

	},
	mounted() {
	}
}
</script>



<style>
</style>
