<template>
	<div>

		<NavBar  @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />
		<button class="home-button" @click="getMyStream">
			<font-awesome-icon icon="fa-solid fa-home" /></button>

		<div class="row">
			<Photo
				v-for="post in posts"
				:key="index"
				:owner="post"
				:photo_id="post.photo_id"
				:comments="post.comments != null ? post.comments : []"
				:likes="post.likes != null ? post.likes : []"
				:upload_date="post.upload_date"

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
			}
	},
	methods: {
		async getMyStream(){
			this.isLoading = true;
			console.log(localStorage.UserID);
			try{
				this.isLoading = true;
				let response = await this.$axios.get(`users/${localStorage.UserID}/mystream`, { headers: { 'Authorization': `${localStorage.token}` } });
				console.log(response.data);
				if (response.data != null){
					this.post = response.data;
					return;
				}

			}catch(e){
				this.errorMsg = e.toString();
			}
			this.isLoading = false;

	},
	 mounted(){
		this.getMyStream();
		console.log(this.posts);
		document.addEventListener('scroll', this.trackScrolling);
	},
}
}
</script>



<style>
</style>
