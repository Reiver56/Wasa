<script>
export default{
	emits: ['error-occurred', 'update-like'],
	props: {
		postData: { type: Object, required: true },
	},
	data(){
		return {
			ownerID: this.postData.ownerID[''],

		}
	}
}
</script>
<template>
	<div class="photo-container">

	</div>
</template>
