<script>
export default {
	data(){
		return{
			commentValue:"",
			errorMsg:"",
		}
	},
	props:['modal_id','comments_list','photo_owner','photo_id'],

	methods: {
		async addComment(){
			try{

				console.log(this.comments_list)
				let response = await this.$axios.post("/users/"+ this.photo_owner +"/photos/"+this.photo_id+"/comments",{
					text: this.commentValue
				},{
					headers:{
						'Authorization': `${localStorage.token}`,
						'Content-Type': 'application/json'
					}
				})

				this.$emit('addComment',{
					id: response.data.id,
					id_user: response.data.id_user,
					id_photo: response.data.id_photo,
					text: response.data.text,
				}
				)
				console.log(response.data)
				this.commentValue = ""
				console.log(this.commentValue)

			}catch(e){
				console.log(e.toString())
			}
		},

		deleteComment(value){
			console.log("value "+ value)
			this.$emit('delete-comment',value)
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
					:author="comm.id_user"
					:nickname="comm.nickname"
					:comment_id="comm.id_comment"
					:photo_id="comm.id_photo"
					:content="comm.text"
					:photo_owner="photo_owner"


					@delete-comment="deleteComment"
					/>

                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">

                                <textarea class="form-control" id="exampleFormControlTextarea1"
								placeholder="Add a comment..." rows="1" maxLength="30" v-model="commentValue"></textarea>
                            </div>
                        </div>

                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary"
							@click.prevent="addComment"
							:disabled="commentValue.length < 1 || commentValue.length > 30">
							Send
							</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

</template>

<style>
.my-modal-disp-none{
	display: none;
}
</style>
