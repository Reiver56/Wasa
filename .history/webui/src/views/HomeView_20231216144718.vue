<template>
	<div>

		<NavBar @show-upload-form="showUploadPhoto = true" @error-occurred="(value) => { errorMsg = value }"/>
			<div class = "row">
				<input id="fileUploader" type="file" class="profile-file-upload" @change="uploadPhoto" accept=".jpg, .png">
			<label  class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> Add </label>
			</div>

		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<div v-if="posts.length === 0" class = "row"><h1 class="d-flex">No Content yet, follow someone</h1></div>
	</div>
</template>



<script>
import NavBar from '@/components/NavBar.vue';

export default {
	components: {
		NavBar,
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
		async uploadPhoto(){
			let photoInput = document.getElementById('file-input');
			let photo = photoInput.files[0];
			let formData = new FormData();
			formData.append('photo', photo);
			try{
				let response = await this.$axios.post(`users/${localStorage.userID}/photos`, formData, { headers: { 'Authorization': `${localStorage.token}` } });
				this.posts.unshift(response.data);
				this.showUploadPhoto = false;
			}catch(e){
				this.errorMsg = e.toString();
			}

		},

		getMyStream: async function(){
			try{
				this.isLoading = true;
				let response = await this.$axios.get(`users/${localStorage.userID}/mystream`, { headers: { 'Authorization': `${localStorage.token}` } });
				if (response.data.length != 0){
					this.photos = response.data;
				}
				this.posts.push(...response.data);

			}catch(e){
				this.errorMsg = e.toString();
			}

	},
	async mounted() {
		this.getMyStream();

		document.addEventListener('scroll', this.trackScrolling);
	},
}
}
</script>



<style>
</style>
