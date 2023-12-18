<template>
	<div class="profile-entries" @click="goToProfile">
		<span class="profile-nickname">{{ nickname }}</span>

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

<style>
.profile-entries{
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	align-items: center;
	padding: 10px;
	border-bottom: 1px solid #e0e0e0;
}
.profile-nickname{
	font-size: 1em;
	font-weight: 500;
}
</style>

