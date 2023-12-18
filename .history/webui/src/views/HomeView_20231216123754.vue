<template>
	<div>
		<UploadPhoto v-if="showUploadPhoto" :photoType="'post'" @exit-form="showUploadPhoto = false" @refresh-data="$router.go(0)" @error-occurred="(value) => { errorMsg = value}" />
		<NavBar @show-upload-form="showUploadPhoto=true" @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<div v-if="posts.length === 0" class = "row"><h1 class="d-flex">No Content yet, follow someone</h1></div>
	</div>
</template>



<script>
import NavBar from '@/components/NavBar.vue';
import UploadPhoto from '@/components/UploadPhoto.vue';
export default {
	components: {
		NavBar,
		UploadPhoto,
	},
	data: function() {
		return {
				posts: [],
				errorMsg: "",
				isLoading: false,

				showUploadPhoto: false,
			}
	},
	methods: {

		getMyStream: async function(){
			try{
				this.isLoading = true;
				let response = await this.$axios.get(`users/${localStorage.userID}/mystream`);
				if (response.data.length != 0){
					this.photos = response.data;
				}

			}catch(e){
				this.errorMsg = e.toString();
			}

	},
	async mounted() {
		await this.getMyStream();
	},
}
}
</script>



<style>
</style>
