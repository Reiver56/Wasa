<template>
	<div>

		<NavBar @show-upload-form="showUploadPhoto = true" @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />

		<div class="row">
			<Photo
				v-for="post in posts"
				:key="i"

			/>

		</div>
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
