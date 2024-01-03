<template>
	<LoadingSpinner :loading="isLoading" />
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>

	<div class = "top-container-profile">
		<div class="top-body-profile-container">
			<button class="home-button" @click="$router.push('/home')">
				<font-awesome-icon icon="fa-solid fa-home" />
			</button>
			<button class="option-button" @click="options=true" @focusout="closeOption">
				<font-awesome-icon icon="fa-solid fa-ellipsis-v" />
			</button>
			<div v-if="options && isOwner" class="option-container-menu">
				<div class="option-menu-item" @click="getBans">
					<font-awesome-icon icon="fa-solid fa-ban" />
					<span>Banned Users</span>
				</div>
				<div class="option-menu-item" @click="doLogout">
					<font-awesome-icon icon="fa-solid fa-sign-out" />
					<span>Logout</span>
				</div>
			</div>
			<div v-else-if="options" class="option-container-menu">
				<div class="option-menu-item" @click="banUser">
					<font-awesome-icon icon="fa-solid fa-ban" />
					<span>Ban User</span>
				</div>
			</div>
		</div>
		<input :readonly="!isOwner" v-model="nickname" @input="editingNickname" @focusout="saveChangeNickname" class="top-body-profile-user-nickname" maxlength="16" spellcheck="false">
		<div class="top-body-profile-user-info">
			<div class="top-body-profile-user-info-item" @click="getFollowers">
				<span class="top-body-profile-user-info-item-number">{{followersCount}}</span>
				<span class="top-body-profile-user-info-item-text">Followers</span>
			</div>
			<div class="top-body-profile-user-info-item" @click="getFollowing">
				<span class="top-body-profile-user-info-item-number">{{followingCount}}</span>
				<span class="top-body-profile-user-info-item-text">Following</span>
			</div>
			<div class="top-body-profile-user-info-item">
				<span class="top-body-profile-user-info-item-number">{{photosCount}}</span>
				<span class="top-body-profile-user-info-item-text">Photos</span>
			</div>
		</div>
		<div class="top-body-profile-user-follow">
			<button v-if="!isOwner" class="top-body-profile-user-follow-button" @click="follow">{{followButton}}</button>
		</div>

		<ProfilesList v-if="showList" :dataGet="dataGet" :header="header" :listType="listType" @exit-list="freeLists" />


		<div class = "row">
				<div class="col-12 d-flex justify-content-center">
					<input id="fileUploader" type="file" class="profile-file-upload" @change="uploadPhoto" accept=".jpg, .png">
			<label v-if="isOwner" class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> New Photo<font-awesome-icon icon="image" /> </label>
		</div>
		<div class="row ">
                    <div class="col-3"></div>
                    <div class="col-6">
                        <hr class="border border-dark">
                    </div>
                    <div class="col-3"></div>
                </div>

		</div>
		<div class = "posts-container">
			<span v-if ="(posts.length == 0 )" class="no-posts">No posts</span>
			<Photo v-for="post in posts"
			:key="post.id"

			:owner="this.userID"
			:photo_id="post.id"
			:likes="post.likes"
			:comments="post.comments"
			:upload_date="post.upload_date"
			:isOwner="isOwner"


			@removePhoto="deletePhoto"
			/>


		</div>
</div>


</template>

<script>


export default {
	components:{

	},

	data(){
		return{
			errorMsg: "",

			userID: "",
			nickname: "",
			followersCount: 0,
			followerList : [],
			followingCount: 0,
			followingList : [],
			photosCount: 0,

			isFollowed: false,

			isOwner: false,

			// Upload Photo
			showUploadPhoto: false,

			followButton: "Follow",

			dataAvailable: true,

			textCounter: 0,
			profiles: [],
			header: "",
			listType: "",

			posts: [],
			showPosts: false,
			postViewData: {},

			dataGet : () => { },
			showList: false,

			options: false,

			isLoading: false,

		}
	},
	methods: {
		async uploadPhoto(){
			let photoInput = document.getElementById('fileUploader');
			const photo = photoInput.files[0];
			const reader = new FileReader();

			console.log(photo);

			reader.readAsArrayBuffer(photo);
			reader.onload = async () => {

				let response = await this.$axios.post(`users/${this.$route.params.id}/photos`,
				reader.result,
				 {
					headers: {
						'Authorization': `${localStorage.token}`,
						'Content-Type': photo.type,
					},
				});
				console.log(response);
				this.posts.unshift(response.data);
				this.photosCount++;

			}

		},




		getProfile: async function() {
			this.isLoading = true;
			try{

				let response = await this.$axios.get(`users/${this.$route.params.id}`, { headers : { 'Authorization': `${localStorage.token}` } });
				console.log(response.data);
				this.userID = response.data.user_id;
				this.nickname = response.data.nickname;
				this.followersCount = response.data.followers != null ? response.data.followers.length : 0;
				this.followingCount = response.data.following != null ? response.data.following.length : 0;
				this.photosCount = response.data.photos != null ? response.data.photos.length : 0;
				this.posts = response.data.photos != null ? response.data.photos : [];
				this.isFollowed = response.data.isFollowed;
				this.isOwner = localStorage.UserID == this.userID;
				this.followButton = this.isFollowed ? "Unfollow" : "Follow";
				console.log(this.posts)

			}
			catch(e){
				this.errorMsg = e.message;
			}
			this.isLoading = false;
		},



		editingNickname() {
			if (this.isOwner){
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outline = "auto";
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outlineColor = "#3b9ee4";
			}
		},
		// funzione per il cambio del nickname
		async saveChangeNickname() {
			if (this.isOwner){
				document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outline = "none";
				if (this.nickname == "" | this.nickname.length <3){
					this.nickname = localStorage.nickname;
					return;
				}
				this.isLoading = true;

				try{
					let response = await this.$axios.put(`users/${this.userID}`, { nickname: this.nickname }, { headers : { 'Authorization': `${localStorage.token}` } });
					localStorage.nickname = this.nickname;
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}
		},

		getFollowers() {
			this.header = "Followers";
			this.listType ="simple";
			this.showList = true;
			this.dataGet = async (usersArray, dataAvailable) => {
				try{
					let response = await this.$axios.get(`users/${this.userID}/followers`, { headers : { 'Authorization': `${localStorage.token}` } });
					usersArray = response.data;
					dataAvailable = true;
					usersArray.push(...response.data);
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}

		},
		closeOption() {
			setTimeout(() => {
				this.options = false;
			}, 100);

		},

		getFollowing() {
			this.header = "Followers";
			this.listType ="simple";
			this.showList = true;
			this.dataGet = async (usersArray, dataAvailable) => {
				try{
					let response = await this.$axios.get(`users/${this.userID}/following`, { headers : { 'Authorization': `${localStorage.token}` } });
					usersArray = response.data;
					dataAvailable = true;
					usersArray.push(...response.data);
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}

		},
		freeLists(){
			this.showList = false;
			this.usersArray = [];
			this.header = "";
		},


		updateProfile(){
			this.getProfile();
		},
		async follow(){
			if (this.isFollowed){
				try{

					let response = await this.$axios.delete(`users/${localStorage.UserID}/followers/${this.userID}`, { headers : { 'Authorization': `${localStorage.token}` } });
					this.isFollowed = false;
					this.followButton = "Follow";
					this.followersCount--;

				}catch(e){
					this.errorMsg = e.message;
				}
			}else{
				try{
					let response = await this.$axios.put(`users/${localStorage.UserID}/followers/${this.userID}`, {},{ headers : { 'Authorization': `${localStorage.token}` } });
					this.isFollowed = true;
					this.followButton = "Unfollow";
					this.followersCount++;
				}catch(e){
					this.errorMsg = e.message;
				}
			}

		},
		openPhoto(data){
			this.showPosts = true;
			this.postViewData = data;
		},
		doLogout(){
			localStorage.clear();
			this.$router.push('/login');
		},
		getBans() {
			this.header = "Banned Users";
			this.listType ="simple";
			this.showList = true;
			this.dataGet = async (usersArray, dataAvailable) => {
				try{
					let response = await this.$axios.get(`users/${this.user_id}/bans`, { headers : { 'Authorization': `${localStorage.token}` } });
					usersArray = response.data;
					dataAvailable = true;
					usersArray.push(...response.data);
				}
				catch(e){
					this.errorMsg = e.message;
				}
			}

		},
		async banUser(){
			try{
				let response = await this.$axios.put(`users/${this.user_id}/ban`, { headers : { 'Authorization': `${localStorage.token}` } });
				this.$router.push('/home');
			}catch(e){
				this.errorMsg = e.message;
			}
			this.options = false;
		},

		deletePhoto(photo_id){
			this.posts = this.posts.filter(item => item.photo_id != photo_id);
			this.photosCount--;
		},







	},

	beforeMount() {
		if (!localStorage.token){
			this.$router.push('/login');
		}
		if (localStorage.userID == this.$route.params.userID){
			this.isOwner = true;
		}
	},

	mounted() {
		this.getProfile();

		if (this.isOwner){
			document.querySelectorAll(".top-body-profile-user-nickname")[0].style.outline = "none";
		}
		document.addEventListener('scroll', this.handleScroll);
	},

	beforeRouteUpdate(to, from) {
		this.posts = [];
		this.dataAvailable = true;

		this.getProfile();


	},

}
</script>

<style>
.top-container-profile{
	display: flex;
	flex-direction: column;
	align-items: center;
	width: 100%;
	height: 100%;
}

.profile-file-upload{
	display: none;
}
.my-btn-add-photo{
	background-color: #3b9ee4;
	color: white;
	border-radius: 5px;
	padding: 5px;
	cursor: pointer;
}
</style>

