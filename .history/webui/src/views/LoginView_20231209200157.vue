
<template>
	<ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"/>
	<div class="login-container">
		<div class="top-login-container"> <span class="login-title">Login</span></div>
		<div class="bottom-login-container">
			<div class="login-input-container">
				<input :on-submit="doLogin" type="text" name="nickname" spellcheck="false" v-model="nickname" maxlength="16">
			</div>
			<div class="login-button-container">
				<button class="login-button" @click="doLogin">Login</button>
			</div>
		</div>
	</div>
	<LoadingSpinner :loading="isLoading"/>
</template>



<style>
</style>



