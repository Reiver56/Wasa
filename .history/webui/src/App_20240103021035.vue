<template>
	<main>
		<div class="navbar-header">
			
		</div>
		<RouterView />
	</main>
</template>

<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	components: {
		RouterView,
	},
	data() {
		return {
		}
	},
	methods: {
		goHome() {
			if (this.$route.path == '/home') {
				this.$router.go(0);
				return;
			}
			this.$router.push('/home');
		},
	},
	mounted() {
	},

}
</script>



<style>
</style>
