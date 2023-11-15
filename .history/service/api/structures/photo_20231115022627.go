package structures
import
(
	"time"
	
)
type Photo struct{
	Id_photo string      `json:"id"`
	User_ID  User_ID     `json:"user_id"`
	Likes 	 []Like 	 `json: "likes"`
	Comments []Comment   `json: "comments"`
	time 	 time.Time   `json: "date"`
}