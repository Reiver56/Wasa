<template>
	<div>

		<NavBar @show-upload-form="showUploadPhoto = true" @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<div v-if="posts.length === 0" class = "row"><h1 class="d-flex">No Content yet, follow someone</h1></div>
		<div class="row">
			<Photo v-for="(photo, index) in posts" :key="index" :photo_id="photo.photo_id" :owner="photo.user_id" :nickname="photo.nickname" :likes="photo.likes" :comments="photo.comments" :isOwner="photo.isOwner" :photo_url="photo.photo_url" :description="photo.description" :created_at="photo.created_at" :updated_at="photo.updated_at" :isLiked="photo.isLiked" @error-occurred="(value) => { errorMsg = value }" />
		</div>
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
