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
			this.$emit('delete-comment',value)
			location.reload();
		},

		addCommentToParent(newCommentJSON){
			this.$emit('addComment',newCommentJSON)
		},
	},
}
</script>

<template>
    <div class="modal my-modal-disp-none" :id="modal_id" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable ">
            <div class="modal-content">

                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modal_id">Comments</h1>
                    <button type="button" class="btn close-btn" data-bs-dismiss="modal" aria-label="Close"><font-awesome-icon :icon="['far', 'times-circle']" /></button>
                </div>

                <div class="modal-body">
                    <PhotoComment v-for="(comm,index) in comments_list"
					:key="index"
					:author="comm.id_user"
					:comment_id="comm.id"
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

                                <textarea class="form-control" placeholder="Add a comment..." rows="1" maxLength="255" v-model="commentValue"></textarea>
                            </div>
                        </div>

                        <div class="align-items-center">
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
.my-modal-disp-none{
	display: none;
}
.close-btn{
	position: absolute;
	top: 0px;
	right: 0px;
	left: 430px;

	background-color: transparent;
	border: none;
	color: black;
	font-size: 30px;
	cursor: pointer;
}
.close-btn:hover{
	color: #3b9ee4;
}
</style>
