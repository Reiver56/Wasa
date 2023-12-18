<template>
	<div>
		<NavBar />
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<div class="row">
			<Photo
				v-for="photo in photos"
				:key="photo.photoID"
				:photo="photo"
				:isMyStream="true"
				@delete-photo="deletePhoto"
				/>
		</div>
		<div v-if="photos.length === 0" class = "row">
			<h1 class="d-flex">No Content yet, follow someone</h1>
		</div>
	</div>
</template>



<script>
import NavBar from '@/components/NavBar.vue';
import 
export default {
	components: {
		NavBar,
	},
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
				let response = await this.$axios.get('users'+ localStorage.getItem('token') + '/mystream');
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
