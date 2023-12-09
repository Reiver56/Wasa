<template>
	<div class="login">
		<img src="src/assets/logo/logo_blue.jpg" alt="logo" />
	  <div>
		<form @submit.prevent="submit">
		  <div>
			<label for="nickname">Nickname:</label>
			<input type="text" name="nickname" v-model="form.nickname" />
		  </div>
		  <button type="submit">Submit</button>
		</form>
		<p v-if="showError" id="error">Nickname is incorrect</p>
	  </div>
	</div>
  </template>


<script>

</script>

<style>
</style>



