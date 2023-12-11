<script>
import NavBar from '@/components/NavBar.vue';
export default {
	props: {
		data :  { type: Object, required: true },
	},
	data(){
		return{
			nickname: this.$props.data['nickname'],
			userID: this.$props.data['id'],
		}
	}
}
</script>

<template>
	<NavBar />
</template>
