<template>
	<div>

		<NavBar  @error-occurred="(value) => { errorMsg = value }"/>
		<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
		<LoadingSpinner :loading="isLoading" />

		<div class="row">
			<Photo
				v-for="post in posts"
				:owner="post.user_id"
				:photo_id="post.id"
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
	data() {
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
				}
				else{
					this.posts = [];
				}
				this.posts.push(...response.data);
				console.log(this.posts);

			}catch(e){
				console.log(e.toString());
			}
			this.isLoading = false;

		},
	},
	 mounted(){
		if (!localStorage.token){
			this.$router.replace('/login');
		}
		this.getMyStream();
		console.log(this.posts);
		document.addEventListener('scroll', this.trackScrolling);
	},

}
</script>



<style>

</style>
