<script>
export default{
	props: ['modal_id', 'likes'],
}
</script>

<template>
	<div class="modal my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">Likes</h5>
					<button type="button" class="btn close-btn" data-bs-dismiss="modal" aria-label="Close"><font-awesome-icon :icon="['far', 'times-circle']" /></button>
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
.close-btn{
	position: absolute;
	top: 0px;
	right: 0px;
	left: 390px;

	background-color: transparent;
	border: none;
	color: black;
	font-size: 30px;
	cursor: pointer;
}
.close-btn:hover{
	color: #3b9ee4;
}
</style>
