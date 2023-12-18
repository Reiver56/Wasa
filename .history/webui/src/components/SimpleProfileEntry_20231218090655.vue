<template>
	<div class="profile-entries" @click="goToProfile">
		<

	</div>
</template>



<script>
export default{
	emits: ['exit-from-list', 'error-occurred', 'data-updated'],
	props: {
		data: { type: Object, required: true },
	},
	data(){
		return {
			nickname: this.$props.data.nickname,
			userID: this.$props.data.userID,

		}
	},
	methods: {
		goToProfile(){
			this.$router.push(`/users/${this.userID}`);
			this.$emit('exit-from-list');
		},
	},
}
</script>
