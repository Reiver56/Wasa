<script>
export default{
	props: ['model_id', 'likes'],
}
</script>

<template>
	<div class="model-head model-dialog model" :id="model_id" tabindex="-1" aria-hidden="true">
		<div class=" model-dialog">
			<div class="model-content">
				<div class="model-header">
					<h5 class="model-title">Likes</h5>
					<button type="button" class="btn-close" data-bs-dismiss="model" aria-label="Close"></button>
				</div>
				<div class= "model-body">
					<UserMini v-for="(user, index) in likes"
					:key="index"
					:identifier="user.id"
					:nickname="user.nickname"
					/>

				</div>

		</div>

	</div>
	</div>

</template>
