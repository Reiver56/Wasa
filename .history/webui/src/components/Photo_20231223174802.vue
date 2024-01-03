<script>
export default{
	emits: ['error-occurred', 'update-like'],
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			ownerID: this.postData.ownerID['userID'],
			ownerName: this.postData.ownerID['username'],

		}
	}
}
</script>
<template>
	<div class="photo-container">

	</div>
</template>
