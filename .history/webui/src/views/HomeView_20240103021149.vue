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
NavBar{
	background-color: #3b9ee4;
	height: auto;
	width: auto;
	padding: 0;
	border-radius: 10em;
	box-shadow: 0 0.2em 0 0 rgba(176, 179, 184, 0.50);
	position: sticky;
	top: 1em;
	padding: 1em;
	margin-bottom: 1em;
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	z-index: 3;
}
</style>
