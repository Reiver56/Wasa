<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	components: {
		RouterView,
	},
	data(){
		return{
		}
	},
	methods: {
		goHome(){
			if(this.$route.path == "/"){
				this.$router.go(0);return;
		}
		this.$router.push("/");
		},
		mounted(){
		}
}
</script>

<template>
	<main>
		<di class="""">
	</main>
</template>

<style>
</style>
