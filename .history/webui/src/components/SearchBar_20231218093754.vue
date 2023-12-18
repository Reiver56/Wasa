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
			userList:[],
			busy: false,
			dataAvailable: true,
		}
	},
	methods: {
		async search(){
			if (this.search.length == 0){
				this.userList = [];
				return;
			}
			try {
				const
			}

		},

	},


}
</script>
