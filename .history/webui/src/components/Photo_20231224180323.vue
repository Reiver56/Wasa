<script>
export default{
	emits: ['error-occurred', 'update-like'],
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			photoURL: "",
			liked: false,
			Comments: [],
			allLikes: [],
		}
	},
	methods: {

	},
	mounted(){
		if (this.ownerID == localStorage.userID){
			this.isOwner = true;
		}
	}
}
</script>
<template>
	<div class="photo-container">

	</div>
</template>
