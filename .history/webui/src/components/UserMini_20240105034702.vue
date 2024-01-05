<script>
export default{
	props: ['identifier', 'nickname'],
	methods:{
		clickedUser(){
			this.$emit('clikedUser', this.idetifier)
		}
	},
}
</script>
<template>
	<div class="row mb-2 mt-2">
		<div class="user-mini col d-flex justify-content-center">
			<div class="user-body">
				<h5 @click="clickedUser" class="user-mini-title d-flex jusfity-content-center">{{identifier}}</h5>
			</div>
		</div>
	</div>
</template>
<style>
.user{
	color: var(--primary);
}
.user-mini{
	color: black;
	width: 25%;
	border: none;
	border-color: black;
}
.user-mini-title:hover{
	cursor: pointer;
	font-weight: bold;
}
</style>
