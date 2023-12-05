<template>

</template>

<script>
export default{
	components:{
	},
	data(){
		return{
			errorMsg:'',
		}

	},
	methods:{
		getMyProfile(){
			this.$router.push('/profile/');
		},
	},

	mounted(){

	},
}
</script>
