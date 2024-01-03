<script>
export default {
	data(){
		return{
			commentValue:"",
		}
	},
	props:['modal_id','comments_list','photo_owner','photo_id', 'error_occurred'],

	methods: {
		async addComment(){
			try{
				console.log(this.comments_list)

				let response = await this.$axios.post(`users/${this.photo_owner}/photos/${this.photo_id}/comments`,{
					text: this.commentValue
				},{
					headers:{
						'Authorization': localStorage.token,
						'Content-Type': 'application/json'
					}
				})
				console.log(response.data)

				this.$emit('addComment',{
					comment_id: response.data.comment_id,
					photo_id: this.photo_id,
					user_id: localStorage.token,
					text: this.commentValue,

					})
					console.log(this.comment_id),
				this.commentValue = ""

			}catch(err){
				this.$emit('error-occurred', err);
			}
		},

		eliminateComment(value){
			this.$emit('eliminateComment',value)
		},

		addCommentToList(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
			console.log(newCommentJSON)
		},
	},
}
</script>

<template>
    <div class="modal fade my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable ">
            <div class="modal-content">

                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modal_id">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>

                <div class="modal-body">
                    <PhotoComment v-for="comm in comments_list"
					:author="comm.user_id"
					:nickname="comm.nickname"
					:comment_id="comm.comment_id"
					:photo_id="comm.photo_id"
					:content="comm.comment"
					:photo_owner="photo_owner"


					@eliminateComment="eliminateComment"
					/>

                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">

                                <textarea class="form-control" id="exampleFormControlTextarea1"
								placeholder="Add a comment..." rows="1" maxLength="255" v-model="commentValue"></textarea>
                            </div>
                        </div>

                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary"
							@click.prevent="addComment"
							:disabled="commentValue.length < 1 || commentValue.length > 255">
							Post
							</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

</template>

<style>
.btn-close{
	color: black;
}
.my-modal-disp-none{
	display: none;
}
</style>
