<script>
export default{
	props: ['modal_id', 'likes'],
}
</script>

<template>
	<div class="modal-head modal-dialog modal-dialog-centered modal-dialog-scrollable my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class=" modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="model-title">Likes</h5>
					<button type="button" class="btn-close" data-bs-dismiss="model" aria-label="Close"></button>
				</div>
				<div class= "modal-body">
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
