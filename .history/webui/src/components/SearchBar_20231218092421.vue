<script>
import SimpleProfileEntry from './SimpleProfileEntry.vue';
export default{
	emits: ['error-occurred'],
	components: {
		SimpleProfileEntry,
	},
	props:{},
	data(){
		return{
			searchQuery: '',
			userList:[].
			searchLimit: 10,
			searchOffset: 0,
			busy: false,
			dataAvailable: true,
		}
	},
	
}
</script>
