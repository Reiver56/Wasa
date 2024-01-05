<script>
export default{
	props: ['modal_id', 'likes'],
}
</script>

<template>
	<div class="modal my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="">
					<h5 class="modal-title ">Likes</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class= "modal-body">
					<UserMini v-for="user in likes"
					:identifier="user.id_user"
					/>

				</div>

		</div>

	</div>
	</div>

</template>

<style>
.my-modal-disp-none{
	display: none;
}
</style>
