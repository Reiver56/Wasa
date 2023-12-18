<template>
	<div class="profile-entries" @click="goToProfile">
		<div class="profile-user-icon"><font-awesome-icon icon="fa-regular fa-user" /></div>
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
			console
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
.profile-user-icon{
	background-color: rgb(255, 255, 255, 0.4);
    height: 1.5em;
    width: 1.5em;
    border-radius: 10em;
    padding: 0;
    margin: 0 0 0 0.5em;
    float: left;
    outline: none;
    border: none;
    font-size: 2em;
    line-height: 0;
    color: rgb(0, 0, 0, 0.8)
}
.profile-nickname{
	font-size: 1em;
	font-weight: 600;
}
</style>

