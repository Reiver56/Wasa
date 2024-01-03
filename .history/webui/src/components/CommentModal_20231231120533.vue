<script>
export default {
	data(){
		return{
			commentValue:"",
		}
	},
	props:['modal_id','comments_list','photo_owner','photo_id'],

	methods: {
		async addComment(){
			try{
				// Comment post: /users/:id/photos/:photo_id/comments
				let response = await this.$axios.post("/users/"+ this.photo_owner +"/photos/"+this.photo_id+"/comments",{
					comment: this.commentValue
				},{
					headers:{
						'Authorization': localStorage.token,
						'Content-Type': 'application/json'
					}
				})
				console.log(response.data)

				this.$emit('addComment',{
					comment_id: response.data.comment_id,
					id_photo: this.photo_id,
					id_user: localStorage.token,
					text: this.commentValue
					,}
				)
				this.commentValue = ""

			}catch(e){
				console.log(e.toString())
			}
		},

		eliminateCommentToParent(value){
			this.$emit('eliminateComment',value)
		},

		addCommentToParent(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
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
                    <PhotoComment v-for="(comm,index) in comments_list"
					:key="index"
					:author="comm.user_id"
					:nickname="comm.nickname"
					:comment_id="comm.comment_id"
					:photo_id="comm.photo_id"
					:content="comm.comment"
					:photo_owner="photo_owner"


					@eliminateComment="eliminateCommentToParent"
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
